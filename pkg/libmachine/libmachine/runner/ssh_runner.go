/*
Copyright 2023 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package runner

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/kballard/go-shellquote"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"golang.org/x/sync/errgroup"
	"k8s.io/klog/v2"
	"k8s.io/minikube/pkg/libmachine/libmachine/utils"
	"k8s.io/minikube/pkg/minikube/assets"
	"k8s.io/minikube/pkg/util/retry"
)

// SSHRunner runs commands through SSH.
//
// It implements the CommandRunner interface.
type SSHRunner struct {
	Ip      string
	KeyPath string
	UsrName string
	Port    int

	C *ssh.Client
	S *ssh.Session
}

type sshReadableFile struct {
	length      int
	sourcePath  string
	permissions string
	sess        *ssh.Session
	modTime     time.Time
	reader      io.Reader
}

// GetLength returns length of file
func (s *sshReadableFile) GetLength() int {
	return s.length
}

// GetSourcePath returns asset name
func (s *sshReadableFile) GetSourcePath() string {
	return s.sourcePath
}

// GetPermissions returns permissions
func (s *sshReadableFile) GetPermissions() string {
	return s.permissions
}

func (s *sshReadableFile) GetModTime() (time.Time, error) {
	return s.modTime, nil
}

func (s *sshReadableFile) Read(p []byte) (int, error) {
	if s.GetLength() == 0 {
		return 0, fmt.Errorf("attempted read from a 0 length asset")
	}
	return s.reader.Read(p)
}

func (s *sshReadableFile) Seek(_ int64, _ int) (int64, error) {
	return 0, fmt.Errorf("Seek is not implemented for sshReadableFile")
}

func (s *sshReadableFile) Close() error {
	return s.sess.Close()
}

// NewSSHRunner returns a new SSHRunner that will run commands
// through the ssh.Client provided.
func NewSSHRunner(ip, keyPath, usrName string, port int) *SSHRunner {
	return &SSHRunner{
		C:       nil,
		Ip:      ip,
		KeyPath: keyPath,
		UsrName: usrName,
		Port:    port,
	}
}

// client returns an ssh client (uses retry underneath)
func (s *SSHRunner) client() (*ssh.Client, error) {
	if s.C != nil {
		return s.C, nil
	}

	c, err := utils.NewSSHClient(s.Ip, s.KeyPath, s.UsrName, s.Port)
	if err != nil {
		return nil, errors.Wrap(err, "new client")
	}
	s.C = c
	return s.C, nil
}

// session returns an ssh session, retrying if necessary
func (s *SSHRunner) session() (*ssh.Session, error) {
	var sess *ssh.Session
	getSession := func() (err error) {
		client, err := s.client()
		if err != nil {
			return errors.Wrap(err, "new client")
		}

		sess, err = client.NewSession()
		if err != nil {
			klog.Warningf("session error, resetting client: %v", err)
			s.C = nil
			return err
		}
		return nil
	}

	if err := retry.Expo(getSession, 250*time.Millisecond, 2*time.Second); err != nil {
		return nil, err
	}

	return sess, nil
}

// RemoveFile runs a command to delete a file on the remote.
func (s *SSHRunner) RemoveFile(f assets.CopyableFile) error {
	dst := path.Join(f.GetTargetDir(), f.GetTargetName())
	klog.Infof("rm: %s", dst)

	sess, err := s.session()
	if err != nil {
		return errors.Wrap(err, "getting ssh session")
	}

	defer sess.Close()
	return sess.Run(fmt.Sprintf("sudo rm %s", dst))
}

// teeSSH runs an SSH command, streaming stdout, stderr to logs
func teeSSH(s *ssh.Session, cmd string, outB io.Writer, errB io.Writer) error {
	outPipe, err := s.StdoutPipe()
	if err != nil {
		return errors.Wrap(err, "stdout")
	}

	errPipe, err := s.StderrPipe()
	if err != nil {
		return errors.Wrap(err, "stderr")
	}
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		if err := teePrefix(ErrPrefix, errPipe, errB, klog.V(8).Infof); err != nil {
			klog.Errorf("tee stderr: %v", err)
		}
		wg.Done()
	}()
	go func() {
		if err := teePrefix(OutPrefix, outPipe, outB, klog.V(8).Infof); err != nil {
			klog.Errorf("tee stdout: %v", err)
		}
		wg.Done()
	}()
	err = s.Run(cmd)
	wg.Wait()
	return err
}

// RunCmd implements the Command Runner interface to run a exec.Cmd object
func (s *SSHRunner) RunCmd(cmd *exec.Cmd) (*RunResult, error) {
	var err error // x7NOTE: cut this out

	if cmd.Stdin != nil {
		return nil, fmt.Errorf("SSHRunner does not support stdin - you could be the first to add it")
	}

	rr := &RunResult{Args: cmd.Args}
	klog.Infof("Run: %v", rr.Command())

	defer func() {
		klog.Infof("x7DBG = returning err: %v", err)
	}()

	var outb, errb io.Writer
	start := time.Now()

	if cmd.Stdout == nil {
		var so bytes.Buffer
		outb = io.MultiWriter(&so, &rr.Stdout)
	} else {
		outb = io.MultiWriter(cmd.Stdout, &rr.Stdout)
	}

	if cmd.Stderr == nil {
		var se bytes.Buffer
		errb = io.MultiWriter(&se, &rr.Stderr)
	} else {
		errb = io.MultiWriter(cmd.Stderr, &rr.Stderr)
	}

	sess, err := s.session()
	if err != nil {
		return rr, errors.Wrap(err, "NewSession")
	}

	defer func() {
		if err := sess.Close(); err != nil {
			if err != io.EOF {
				klog.Errorf("session close: %v", err)
			}
		}
	}()

	err = teeSSH(sess, shellquote.Join(cmd.Args...), outb, errb)
	elapsed := time.Since(start)

	if exitError, ok := err.(*exec.ExitError); ok {
		rr.ExitCode = exitError.ExitCode()
	}
	// Decrease log spam
	if elapsed > (1 * time.Second) {
		klog.Infof("Completed: %s: (%s)", rr.Command(), elapsed)
	}
	if err == nil {
		return rr, nil
	}

	return rr, fmt.Errorf("%s: %v\nstdout:\n%s\nstderr:\n%s", rr.Command(), err, rr.Stdout.String(), rr.Stderr.String())
}

// teeSSHStart starts a non-blocking SSH command, streaming stdout, stderr to logs
func teeSSHStart(s *ssh.Session, cmd string, outB io.Writer, errB io.Writer, wg *sync.WaitGroup) error {
	outPipe, err := s.StdoutPipe()
	if err != nil {
		return errors.Wrap(err, "stdout")
	}

	errPipe, err := s.StderrPipe()
	if err != nil {
		return errors.Wrap(err, "stderr")
	}

	go func() {
		if err := teePrefix(ErrPrefix, errPipe, errB, klog.V(8).Infof); err != nil {
			klog.Errorf("tee stderr: %v", err)
		}
		wg.Done()
	}()
	go func() {
		if err := teePrefix(OutPrefix, outPipe, outB, klog.V(8).Infof); err != nil {
			klog.Errorf("tee stdout: %v", err)
		}
		wg.Done()
	}()

	return s.Start(cmd)
}

// StartCmd implements the Command Runner interface to start a exec.Cmd object
func (s *SSHRunner) StartCmd(cmd *exec.Cmd) (*StartedCmd, error) {
	if cmd.Stdin != nil {
		return nil, fmt.Errorf("SSHRunner does not support stdin - you could be the first to add it")
	}

	if s.S != nil {
		return nil, fmt.Errorf("another SSH command has been started and is currently running")
	}

	var wg sync.WaitGroup
	wg.Add(2)
	rr := &RunResult{Args: cmd.Args}
	sc := &StartedCmd{cmd: cmd, rr: rr, wg: &wg}
	klog.Infof("Start: %v", rr.Command())

	var outb, errb io.Writer

	if cmd.Stdout == nil {
		var so bytes.Buffer
		outb = io.MultiWriter(&so, &rr.Stdout)
	} else {
		outb = io.MultiWriter(cmd.Stdout, &rr.Stdout)
	}

	if cmd.Stderr == nil {
		var se bytes.Buffer
		errb = io.MultiWriter(&se, &rr.Stderr)
	} else {
		errb = io.MultiWriter(cmd.Stderr, &rr.Stderr)
	}

	sess, err := s.session()
	if err != nil {
		return sc, errors.Wrap(err, "NewSession")
	}

	s.S = sess

	err = teeSSHStart(s.S, shellquote.Join(cmd.Args...), outb, errb, &wg)

	return sc, err
}

// WaitCmd implements the Command Runner interface to wait until a started exec.Cmd object finishes
func (s *SSHRunner) WaitCmd(sc *StartedCmd) (*RunResult, error) {
	if s.S == nil {
		return nil, fmt.Errorf("there is no SSH command started")
	}

	rr := sc.rr

	err := s.S.Wait()
	if exitError, ok := err.(*exec.ExitError); ok {
		rr.ExitCode = exitError.ExitCode()
	}

	sc.wg.Wait()

	if err := s.S.Close(); err != io.EOF {
		klog.Errorf("session close: %v", err)
	}

	s.S = nil

	if err == nil {
		return rr, nil
	}

	return rr, fmt.Errorf("%s: %v\nstdout:\n%s\nstderr:\n%s", rr.Command(), err, rr.Stdout.String(), rr.Stderr.String())
}

// Copy copies a file to the remote over SSH.
func (s *SSHRunner) CopyFile(f assets.CopyableFile) error {
	dst := path.Join(path.Join(f.GetTargetDir(), f.GetTargetName()))

	// For small files, don't bother risking being wrong for no performance benefit
	if f.GetLength() > 2048 {
		exists, err := fileExists(s, f, dst)
		if err != nil {
			klog.Infof("existence check for %s: %v", dst, err)
		}

		if exists {
			klog.Infof("copy: skipping %s (exists)", dst)
			return nil
		}
	}

	src := f.GetSourcePath()
	klog.Infof("scp %s --> %s (%d bytes)", src, dst, f.GetLength())
	if f.GetLength() == 0 {
		klog.Warningf("0 byte asset: %+v", f)
	}

	sess, err := s.session()
	if err != nil {
		return errors.Wrap(err, "NewSession")
	}
	defer func() {
		if err := sess.Close(); err != nil {
			if err != io.EOF {
				klog.Errorf("session close: %v", err)
			}
		}
	}()

	w, err := sess.StdinPipe()
	if err != nil {
		return errors.Wrap(err, "StdinPipe")
	}
	// The scpcmd below *should not* return until all data is copied and the
	// StdinPipe is closed. But let's use errgroup to make it explicit.
	var g errgroup.Group
	var copied int64

	g.Go(func() error {
		defer w.Close()
		header := fmt.Sprintf("C%s %d %s\n", f.GetPermissions(), f.GetLength(), f.GetTargetName())
		fmt.Fprint(w, header)
		if f.GetLength() == 0 {
			klog.Warningf("asked to copy a 0 byte asset: %+v", f)
			fmt.Fprint(w, "\x00")
			return nil
		}

		copied, err = io.Copy(w, f)
		if err != nil {
			return errors.Wrap(err, "io.Copy")
		}
		if copied != int64(f.GetLength()) {
			return fmt.Errorf("%s: expected to copy %d bytes, but copied %d instead", f.GetTargetName(), f.GetLength(), copied)
		}
		fmt.Fprint(w, "\x00")
		return nil
	})

	scp := fmt.Sprintf("sudo test -d %s && sudo scp -t %s", f.GetTargetDir(), f.GetTargetDir())
	mtime, err := f.GetModTime()
	if err != nil {
		klog.Infof("error getting modtime for %s: %v", dst, err)
	} else if mtime != (time.Time{}) {
		scp += fmt.Sprintf(" && sudo touch -d \"%s\" %s", mtime.Format(layout), dst)
	}
	out, err := sess.CombinedOutput(scp)
	if err != nil {
		return fmt.Errorf("%s: %s\noutput: %s", scp, err, out)
	}
	return g.Wait()
}

// CopyFileFrom copies a file from the remote over SSH.
func (s *SSHRunner) CopyFileFrom(f assets.CopyableFile) error {
	dst := path.Join(path.Join(f.GetTargetDir(), f.GetTargetName()))

	sess, err := s.session()
	if err != nil {
		return errors.Wrap(err, "NewSession")
	}
	defer func() {
		if err := sess.Close(); err != nil {
			if err != io.EOF {
				klog.Errorf("session close: %v", err)
			}
		}
	}()

	cmd := exec.Command("stat", "-c", "%s", dst)
	rr, err := s.RunCmd(cmd)
	if err != nil {
		return fmt.Errorf("%s: %v", cmd, err)
	}
	length, err := strconv.Atoi(strings.TrimSuffix(rr.Stdout.String(), "\n"))
	if err != nil {
		return err
	}
	src := f.GetSourcePath()
	klog.Infof("scp %s --> %s (%d bytes)", dst, src, length)
	f.SetLength(length)

	r, err := sess.StdoutPipe()
	if err != nil {
		return errors.Wrap(err, "StdoutPipe")
	}
	w, err := sess.StdinPipe()
	if err != nil {
		return errors.Wrap(err, "StdinPipe")
	}
	// The scpcmd below *should not* return until all data is copied and the
	// StdinPipe is closed. But let's use errgroup to make it explicit.
	var g errgroup.Group
	var copied int64

	g.Go(func() error {
		defer w.Close()
		br := bufio.NewReader(r)
		fmt.Fprint(w, "\x00")
		b, err := br.ReadBytes('\n')
		if err != nil {
			return errors.Wrap(err, "ReadBytes")
		}
		if b[0] != 'C' {
			return fmt.Errorf("unexpected: %v", b)
		}
		fmt.Fprint(w, "\x00")

		copied = 0
		for copied < int64(length) {
			n, err := io.CopyN(f, br, int64(length))
			if err != nil {
				return errors.Wrap(err, "io.CopyN")
			}
			copied += n
		}
		fmt.Fprint(w, "\x00")
		err = sess.Wait()
		if err != nil {
			return err
		}
		return nil
	})

	scp := fmt.Sprintf("sudo scp -f %s", f.GetTargetPath())
	err = sess.Start(scp)
	if err != nil {
		return fmt.Errorf("%s: %s", scp, err)
	}
	return g.Wait()
}

// ReadableFile returns assets.ReadableFile for the sourcePath (via `stat` command)
func (s *SSHRunner) ReadableFile(sourcePath string) (assets.ReadableFile, error) {
	klog.V(4).Infof("NewsshReadableFile: %s -> %s", sourcePath)

	if !strings.HasPrefix(sourcePath, "/") {
		return nil, fmt.Errorf("sourcePath must be an absolute Path. Relative Path is not allowed")
	}

	// get file size and modtime of the destination
	rr, err := s.RunCmd(exec.Command("stat", "-c", "%#a %s %y", sourcePath))
	if err != nil {
		return nil, err
	}

	stdout := strings.TrimSpace(rr.Stdout.String())
	outputs := strings.SplitN(stdout, " ", 3)

	permission := outputs[0]
	size, err := strconv.Atoi(outputs[1])
	if err != nil {
		return nil, err
	}

	modTime, err := time.Parse(layout, outputs[2])
	if err != nil {
		return nil, err
	}

	sess, err := s.session()
	if err != nil {
		return nil, errors.Wrap(err, "NewSession")
	}

	r, err := sess.StdoutPipe()
	if err != nil {
		return nil, errors.Wrap(err, "StdOutPipe")
	}

	cmd := fmt.Sprintf("cat %s", sourcePath)
	if err := sess.Start(cmd); err != nil {
		return nil, err
	}

	return &sshReadableFile{
		length:      size,
		sourcePath:  sourcePath,
		permissions: permission,
		reader:      r,
		modTime:     modTime,
		sess:        sess,
	}, nil
}
