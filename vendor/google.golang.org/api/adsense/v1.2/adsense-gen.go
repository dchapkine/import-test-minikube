// Package adsense provides access to the AdSense Management API.
//
// See https://developers.google.com/adsense/management/
//
// Usage example:
//
//   import "google.golang.org/api/adsense/v1.2"
//   ...
//   adsenseService, err := adsense.New(oauthHttpClient)
package adsense // import "google.golang.org/api/adsense/v1.2"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "adsense:v1.2"
const apiName = "adsense"
const apiVersion = "v1.2"
const basePath = "https://www.googleapis.com/adsense/v1.2/"

// OAuth2 scopes used by this API.
const (
	// View and manage your AdSense data
	AdsenseScope = "https://www.googleapis.com/auth/adsense"

	// View your AdSense data
	AdsenseReadonlyScope = "https://www.googleapis.com/auth/adsense.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Accounts = NewAccountsService(s)
	s.Adclients = NewAdclientsService(s)
	s.Adunits = NewAdunitsService(s)
	s.Customchannels = NewCustomchannelsService(s)
	s.Reports = NewReportsService(s)
	s.Savedadstyles = NewSavedadstylesService(s)
	s.Urlchannels = NewUrlchannelsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Accounts *AccountsService

	Adclients *AdclientsService

	Adunits *AdunitsService

	Customchannels *CustomchannelsService

	Reports *ReportsService

	Savedadstyles *SavedadstylesService

	Urlchannels *UrlchannelsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewAccountsService(s *Service) *AccountsService {
	rs := &AccountsService{s: s}
	rs.Adclients = NewAccountsAdclientsService(s)
	rs.Adunits = NewAccountsAdunitsService(s)
	rs.Customchannels = NewAccountsCustomchannelsService(s)
	rs.Reports = NewAccountsReportsService(s)
	rs.Savedadstyles = NewAccountsSavedadstylesService(s)
	rs.Urlchannels = NewAccountsUrlchannelsService(s)
	return rs
}

type AccountsService struct {
	s *Service

	Adclients *AccountsAdclientsService

	Adunits *AccountsAdunitsService

	Customchannels *AccountsCustomchannelsService

	Reports *AccountsReportsService

	Savedadstyles *AccountsSavedadstylesService

	Urlchannels *AccountsUrlchannelsService
}

func NewAccountsAdclientsService(s *Service) *AccountsAdclientsService {
	rs := &AccountsAdclientsService{s: s}
	return rs
}

type AccountsAdclientsService struct {
	s *Service
}

func NewAccountsAdunitsService(s *Service) *AccountsAdunitsService {
	rs := &AccountsAdunitsService{s: s}
	rs.Customchannels = NewAccountsAdunitsCustomchannelsService(s)
	return rs
}

type AccountsAdunitsService struct {
	s *Service

	Customchannels *AccountsAdunitsCustomchannelsService
}

func NewAccountsAdunitsCustomchannelsService(s *Service) *AccountsAdunitsCustomchannelsService {
	rs := &AccountsAdunitsCustomchannelsService{s: s}
	return rs
}

type AccountsAdunitsCustomchannelsService struct {
	s *Service
}

func NewAccountsCustomchannelsService(s *Service) *AccountsCustomchannelsService {
	rs := &AccountsCustomchannelsService{s: s}
	rs.Adunits = NewAccountsCustomchannelsAdunitsService(s)
	return rs
}

type AccountsCustomchannelsService struct {
	s *Service

	Adunits *AccountsCustomchannelsAdunitsService
}

func NewAccountsCustomchannelsAdunitsService(s *Service) *AccountsCustomchannelsAdunitsService {
	rs := &AccountsCustomchannelsAdunitsService{s: s}
	return rs
}

type AccountsCustomchannelsAdunitsService struct {
	s *Service
}

func NewAccountsReportsService(s *Service) *AccountsReportsService {
	rs := &AccountsReportsService{s: s}
	rs.Saved = NewAccountsReportsSavedService(s)
	return rs
}

type AccountsReportsService struct {
	s *Service

	Saved *AccountsReportsSavedService
}

func NewAccountsReportsSavedService(s *Service) *AccountsReportsSavedService {
	rs := &AccountsReportsSavedService{s: s}
	return rs
}

type AccountsReportsSavedService struct {
	s *Service
}

func NewAccountsSavedadstylesService(s *Service) *AccountsSavedadstylesService {
	rs := &AccountsSavedadstylesService{s: s}
	return rs
}

type AccountsSavedadstylesService struct {
	s *Service
}

func NewAccountsUrlchannelsService(s *Service) *AccountsUrlchannelsService {
	rs := &AccountsUrlchannelsService{s: s}
	return rs
}

type AccountsUrlchannelsService struct {
	s *Service
}

func NewAdclientsService(s *Service) *AdclientsService {
	rs := &AdclientsService{s: s}
	return rs
}

type AdclientsService struct {
	s *Service
}

func NewAdunitsService(s *Service) *AdunitsService {
	rs := &AdunitsService{s: s}
	rs.Customchannels = NewAdunitsCustomchannelsService(s)
	return rs
}

type AdunitsService struct {
	s *Service

	Customchannels *AdunitsCustomchannelsService
}

func NewAdunitsCustomchannelsService(s *Service) *AdunitsCustomchannelsService {
	rs := &AdunitsCustomchannelsService{s: s}
	return rs
}

type AdunitsCustomchannelsService struct {
	s *Service
}

func NewCustomchannelsService(s *Service) *CustomchannelsService {
	rs := &CustomchannelsService{s: s}
	rs.Adunits = NewCustomchannelsAdunitsService(s)
	return rs
}

type CustomchannelsService struct {
	s *Service

	Adunits *CustomchannelsAdunitsService
}

func NewCustomchannelsAdunitsService(s *Service) *CustomchannelsAdunitsService {
	rs := &CustomchannelsAdunitsService{s: s}
	return rs
}

type CustomchannelsAdunitsService struct {
	s *Service
}

func NewReportsService(s *Service) *ReportsService {
	rs := &ReportsService{s: s}
	rs.Saved = NewReportsSavedService(s)
	return rs
}

type ReportsService struct {
	s *Service

	Saved *ReportsSavedService
}

func NewReportsSavedService(s *Service) *ReportsSavedService {
	rs := &ReportsSavedService{s: s}
	return rs
}

type ReportsSavedService struct {
	s *Service
}

func NewSavedadstylesService(s *Service) *SavedadstylesService {
	rs := &SavedadstylesService{s: s}
	return rs
}

type SavedadstylesService struct {
	s *Service
}

func NewUrlchannelsService(s *Service) *UrlchannelsService {
	rs := &UrlchannelsService{s: s}
	return rs
}

type UrlchannelsService struct {
	s *Service
}

type Account struct {
	// Id: Unique identifier of this account.
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is, in this case adsense#account.
	Kind string `json:"kind,omitempty"`

	// Name: Name of this account.
	Name string `json:"name,omitempty"`

	// Premium: Whether this account is premium.
	Premium bool `json:"premium,omitempty"`

	// SubAccounts: Sub accounts of the this account.
	SubAccounts []*Account `json:"subAccounts,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Account) MarshalJSON() ([]byte, error) {
	type noMethod Account
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type Accounts struct {
	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Items: The accounts returned in this list response.
	Items []*Account `json:"items,omitempty"`

	// Kind: Kind of list this is, in this case adsense#accounts.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token used to page through accounts. To
	// retrieve the next page of results, set the next request's "pageToken"
	// value to this.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Etag") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Accounts) MarshalJSON() ([]byte, error) {
	type noMethod Accounts
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type AdClient struct {
	// ArcOptIn: Whether this ad client is opted in to ARC.
	ArcOptIn bool `json:"arcOptIn,omitempty"`

	// Id: Unique identifier of this ad client.
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is, in this case adsense#adClient.
	Kind string `json:"kind,omitempty"`

	// ProductCode: This ad client's product code, which corresponds to the
	// PRODUCT_CODE report dimension.
	ProductCode string `json:"productCode,omitempty"`

	// SupportsReporting: Whether this ad client supports being reported on.
	SupportsReporting bool `json:"supportsReporting,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ArcOptIn") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AdClient) MarshalJSON() ([]byte, error) {
	type noMethod AdClient
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type AdClients struct {
	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Items: The ad clients returned in this list response.
	Items []*AdClient `json:"items,omitempty"`

	// Kind: Kind of list this is, in this case adsense#adClients.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token used to page through ad clients. To
	// retrieve the next page of results, set the next request's "pageToken"
	// value to this.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Etag") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AdClients) MarshalJSON() ([]byte, error) {
	type noMethod AdClients
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type AdStyle struct {
	// Colors: The colors which are included in the style. These are
	// represented as six hexadecimal characters, similar to HTML color
	// codes, but without the leading hash.
	Colors *AdStyleColors `json:"colors,omitempty"`

	// Corners: The style of the corners in the ad.
	Corners string `json:"corners,omitempty"`

	// Font: The font which is included in the style.
	Font *AdStyleFont `json:"font,omitempty"`

	// Kind: Kind this is, in this case adsense#adStyle.
	Kind string `json:"kind,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Colors") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AdStyle) MarshalJSON() ([]byte, error) {
	type noMethod AdStyle
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AdStyleColors: The colors which are included in the style. These are
// represented as six hexadecimal characters, similar to HTML color
// codes, but without the leading hash.
type AdStyleColors struct {
	// Background: The color of the ad background.
	Background string `json:"background,omitempty"`

	// Border: The color of the ad border.
	Border string `json:"border,omitempty"`

	// Text: The color of the ad text.
	Text string `json:"text,omitempty"`

	// Title: The color of the ad title.
	Title string `json:"title,omitempty"`

	// Url: The color of the ad url.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Background") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AdStyleColors) MarshalJSON() ([]byte, error) {
	type noMethod AdStyleColors
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AdStyleFont: The font which is included in the style.
type AdStyleFont struct {
	// Family: The family of the font.
	Family string `json:"family,omitempty"`

	// Size: The size of the font.
	Size string `json:"size,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Family") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AdStyleFont) MarshalJSON() ([]byte, error) {
	type noMethod AdStyleFont
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type AdUnit struct {
	// Code: Identity code of this ad unit, not necessarily unique across ad
	// clients.
	Code string `json:"code,omitempty"`

	// ContentAdsSettings: Settings specific to content ads (AFC) and
	// highend mobile content ads (AFMC).
	ContentAdsSettings *AdUnitContentAdsSettings `json:"contentAdsSettings,omitempty"`

	// CustomStyle: Custom style information specific to this ad unit.
	CustomStyle *AdStyle `json:"customStyle,omitempty"`

	// FeedAdsSettings: Settings specific to feed ads (AFF).
	FeedAdsSettings *AdUnitFeedAdsSettings `json:"feedAdsSettings,omitempty"`

	// Id: Unique identifier of this ad unit. This should be considered an
	// opaque identifier; it is not safe to rely on it being in any
	// particular format.
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is, in this case adsense#adUnit.
	Kind string `json:"kind,omitempty"`

	// MobileContentAdsSettings: Settings specific to WAP mobile content ads
	// (AFMC).
	MobileContentAdsSettings *AdUnitMobileContentAdsSettings `json:"mobileContentAdsSettings,omitempty"`

	// Name: Name of this ad unit.
	Name string `json:"name,omitempty"`

	// SavedStyleId: ID of the saved ad style which holds this ad unit's
	// style information.
	SavedStyleId string `json:"savedStyleId,omitempty"`

	// Status: Status of this ad unit. Possible values are:
	// NEW: Indicates that the ad unit was created within the last seven
	// days and does not yet have any activity associated with it.
	//
	// ACTIVE: Indicates that there has been activity on this ad unit in the
	// last seven days.
	//
	// INACTIVE: Indicates that there has been no activity on this ad unit
	// in the last seven days.
	Status string `json:"status,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AdUnit) MarshalJSON() ([]byte, error) {
	type noMethod AdUnit
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AdUnitContentAdsSettings: Settings specific to content ads (AFC) and
// highend mobile content ads (AFMC).
type AdUnitContentAdsSettings struct {
	// BackupOption: The backup option to be used in instances where no ad
	// is available.
	BackupOption *AdUnitContentAdsSettingsBackupOption `json:"backupOption,omitempty"`

	// Size: Size of this ad unit.
	Size string `json:"size,omitempty"`

	// Type: Type of this ad unit.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BackupOption") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AdUnitContentAdsSettings) MarshalJSON() ([]byte, error) {
	type noMethod AdUnitContentAdsSettings
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AdUnitContentAdsSettingsBackupOption: The backup option to be used in
// instances where no ad is available.
type AdUnitContentAdsSettingsBackupOption struct {
	// Color: Color to use when type is set to COLOR.
	Color string `json:"color,omitempty"`

	// Type: Type of the backup option. Possible values are BLANK, COLOR and
	// URL.
	Type string `json:"type,omitempty"`

	// Url: URL to use when type is set to URL.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Color") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AdUnitContentAdsSettingsBackupOption) MarshalJSON() ([]byte, error) {
	type noMethod AdUnitContentAdsSettingsBackupOption
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AdUnitFeedAdsSettings: Settings specific to feed ads (AFF).
type AdUnitFeedAdsSettings struct {
	// AdPosition: The position of the ads relative to the feed entries.
	AdPosition string `json:"adPosition,omitempty"`

	// Frequency: The frequency at which ads should appear in the feed (i.e.
	// every N entries).
	Frequency int64 `json:"frequency,omitempty"`

	// MinimumWordCount: The minimum length an entry should be in order to
	// have attached ads.
	MinimumWordCount int64 `json:"minimumWordCount,omitempty"`

	// Type: The type of ads which should appear.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AdPosition") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AdUnitFeedAdsSettings) MarshalJSON() ([]byte, error) {
	type noMethod AdUnitFeedAdsSettings
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AdUnitMobileContentAdsSettings: Settings specific to WAP mobile
// content ads (AFMC).
type AdUnitMobileContentAdsSettings struct {
	// MarkupLanguage: The markup language to use for this ad unit.
	MarkupLanguage string `json:"markupLanguage,omitempty"`

	// ScriptingLanguage: The scripting language to use for this ad unit.
	ScriptingLanguage string `json:"scriptingLanguage,omitempty"`

	// Size: Size of this ad unit.
	Size string `json:"size,omitempty"`

	// Type: Type of this ad unit.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MarkupLanguage") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AdUnitMobileContentAdsSettings) MarshalJSON() ([]byte, error) {
	type noMethod AdUnitMobileContentAdsSettings
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type AdUnits struct {
	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Items: The ad units returned in this list response.
	Items []*AdUnit `json:"items,omitempty"`

	// Kind: Kind of list this is, in this case adsense#adUnits.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token used to page through ad units. To
	// retrieve the next page of results, set the next request's "pageToken"
	// value to this.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Etag") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AdUnits) MarshalJSON() ([]byte, error) {
	type noMethod AdUnits
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type AdsenseReportsGenerateResponse struct {
	// Averages: The averages of the report. This is the same length as any
	// other row in the report; cells corresponding to dimension columns are
	// empty.
	Averages []string `json:"averages,omitempty"`

	// Headers: The header information of the columns requested in the
	// report. This is a list of headers; one for each dimension in the
	// request, followed by one for each metric in the request.
	Headers []*AdsenseReportsGenerateResponseHeaders `json:"headers,omitempty"`

	// Kind: Kind this is, in this case adsense#report.
	Kind string `json:"kind,omitempty"`

	// Rows: The output rows of the report. Each row is a list of cells; one
	// for each dimension in the request, followed by one for each metric in
	// the request. The dimension cells contain strings, and the metric
	// cells contain numbers.
	Rows [][]string `json:"rows,omitempty"`

	// TotalMatchedRows: The total number of rows matched by the report
	// request. Fewer rows may be returned in the response due to being
	// limited by the row count requested or the report row limit.
	TotalMatchedRows int64 `json:"totalMatchedRows,omitempty,string"`

	// Totals: The totals of the report. This is the same length as any
	// other row in the report; cells corresponding to dimension columns are
	// empty.
	Totals []string `json:"totals,omitempty"`

	// Warnings: Any warnings associated with generation of the report.
	Warnings []string `json:"warnings,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Averages") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AdsenseReportsGenerateResponse) MarshalJSON() ([]byte, error) {
	type noMethod AdsenseReportsGenerateResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type AdsenseReportsGenerateResponseHeaders struct {
	// Currency: The currency of this column. Only present if the header
	// type is METRIC_CURRENCY.
	Currency string `json:"currency,omitempty"`

	// Name: The name of the header.
	Name string `json:"name,omitempty"`

	// Type: The type of the header; one of DIMENSION, METRIC_TALLY,
	// METRIC_RATIO, or METRIC_CURRENCY.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Currency") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AdsenseReportsGenerateResponseHeaders) MarshalJSON() ([]byte, error) {
	type noMethod AdsenseReportsGenerateResponseHeaders
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type CustomChannel struct {
	// Code: Code of this custom channel, not necessarily unique across ad
	// clients.
	Code string `json:"code,omitempty"`

	// Id: Unique identifier of this custom channel. This should be
	// considered an opaque identifier; it is not safe to rely on it being
	// in any particular format.
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is, in this case adsense#customChannel.
	Kind string `json:"kind,omitempty"`

	// Name: Name of this custom channel.
	Name string `json:"name,omitempty"`

	// TargetingInfo: The targeting information of this custom channel, if
	// activated.
	TargetingInfo *CustomChannelTargetingInfo `json:"targetingInfo,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CustomChannel) MarshalJSON() ([]byte, error) {
	type noMethod CustomChannel
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// CustomChannelTargetingInfo: The targeting information of this custom
// channel, if activated.
type CustomChannelTargetingInfo struct {
	// AdsAppearOn: The name used to describe this channel externally.
	AdsAppearOn string `json:"adsAppearOn,omitempty"`

	// Description: The external description of the channel.
	Description string `json:"description,omitempty"`

	// Location: The locations in which ads appear. (Only valid for content
	// and mobile content ads). Acceptable values for content ads are:
	// TOP_LEFT, TOP_CENTER, TOP_RIGHT, MIDDLE_LEFT, MIDDLE_CENTER,
	// MIDDLE_RIGHT, BOTTOM_LEFT, BOTTOM_CENTER, BOTTOM_RIGHT,
	// MULTIPLE_LOCATIONS. Acceptable values for mobile content ads are:
	// TOP, MIDDLE, BOTTOM, MULTIPLE_LOCATIONS.
	Location string `json:"location,omitempty"`

	// SiteLanguage: The language of the sites ads will be displayed on.
	SiteLanguage string `json:"siteLanguage,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AdsAppearOn") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CustomChannelTargetingInfo) MarshalJSON() ([]byte, error) {
	type noMethod CustomChannelTargetingInfo
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type CustomChannels struct {
	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Items: The custom channels returned in this list response.
	Items []*CustomChannel `json:"items,omitempty"`

	// Kind: Kind of list this is, in this case adsense#customChannels.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token used to page through custom
	// channels. To retrieve the next page of results, set the next
	// request's "pageToken" value to this.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Etag") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CustomChannels) MarshalJSON() ([]byte, error) {
	type noMethod CustomChannels
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type SavedAdStyle struct {
	// AdStyle: The AdStyle itself.
	AdStyle *AdStyle `json:"adStyle,omitempty"`

	// Id: Unique identifier of this saved ad style. This should be
	// considered an opaque identifier; it is not safe to rely on it being
	// in any particular format.
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is, in this case adsense#savedAdStyle.
	Kind string `json:"kind,omitempty"`

	// Name: The user selected name of this SavedAdStyle.
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AdStyle") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *SavedAdStyle) MarshalJSON() ([]byte, error) {
	type noMethod SavedAdStyle
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type SavedAdStyles struct {
	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Items: The saved ad styles returned in this list response.
	Items []*SavedAdStyle `json:"items,omitempty"`

	// Kind: Kind of list this is, in this case adsense#savedAdStyles.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token used to page through ad units. To
	// retrieve the next page of results, set the next request's "pageToken"
	// value to this.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Etag") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *SavedAdStyles) MarshalJSON() ([]byte, error) {
	type noMethod SavedAdStyles
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type SavedReport struct {
	// Id: Unique identifier of this saved report.
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is, in this case adsense#savedReport.
	Kind string `json:"kind,omitempty"`

	// Name: This saved report's name.
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *SavedReport) MarshalJSON() ([]byte, error) {
	type noMethod SavedReport
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type SavedReports struct {
	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Items: The saved reports returned in this list response.
	Items []*SavedReport `json:"items,omitempty"`

	// Kind: Kind of list this is, in this case adsense#savedReports.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token used to page through saved reports.
	// To retrieve the next page of results, set the next request's
	// "pageToken" value to this.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Etag") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *SavedReports) MarshalJSON() ([]byte, error) {
	type noMethod SavedReports
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type UrlChannel struct {
	// Id: Unique identifier of this URL channel. This should be considered
	// an opaque identifier; it is not safe to rely on it being in any
	// particular format.
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is, in this case adsense#urlChannel.
	Kind string `json:"kind,omitempty"`

	// UrlPattern: URL Pattern of this URL channel. Does not include
	// "http://" or "https://". Example: www.example.com/home
	UrlPattern string `json:"urlPattern,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *UrlChannel) MarshalJSON() ([]byte, error) {
	type noMethod UrlChannel
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type UrlChannels struct {
	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Items: The URL channels returned in this list response.
	Items []*UrlChannel `json:"items,omitempty"`

	// Kind: Kind of list this is, in this case adsense#urlChannels.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token used to page through URL channels.
	// To retrieve the next page of results, set the next request's
	// "pageToken" value to this.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Etag") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *UrlChannels) MarshalJSON() ([]byte, error) {
	type noMethod UrlChannels
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// method id "adsense.accounts.get":

type AccountsGetCall struct {
	s            *Service
	accountId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Get information about the selected AdSense account.
func (r *AccountsService) Get(accountId string) *AccountsGetCall {
	c := &AccountsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *AccountsGetCall) QuotaUser(quotaUser string) *AccountsGetCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// Tree sets the optional parameter "tree": Whether the tree of sub
// accounts should be returned.
func (c *AccountsGetCall) Tree(tree bool) *AccountsGetCall {
	c.urlParams_.Set("tree", fmt.Sprint(tree))
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *AccountsGetCall) UserIP(userIP string) *AccountsGetCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsGetCall) Fields(s ...googleapi.Field) *AccountsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsGetCall) IfNoneMatch(entityTag string) *AccountsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsGetCall) Context(ctx context.Context) *AccountsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *AccountsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId": c.accountId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.accounts.get" call.
// Exactly one of *Account or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Account.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AccountsGetCall) Do() (*Account, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Account{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get information about the selected AdSense account.",
	//   "httpMethod": "GET",
	//   "id": "adsense.accounts.get",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account to get information about.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tree": {
	//       "description": "Whether the tree of sub accounts should be returned.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "accounts/{accountId}",
	//   "response": {
	//     "$ref": "Account"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.accounts.list":

type AccountsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: List all accounts available to this AdSense account.
func (r *AccountsService) List() *AccountsListCall {
	c := &AccountsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of accounts to include in the response, used for paging.
func (c *AccountsListCall) MaxResults(maxResults int64) *AccountsListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through accounts. To retrieve the next page, set
// this parameter to the value of "nextPageToken" from the previous
// response.
func (c *AccountsListCall) PageToken(pageToken string) *AccountsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *AccountsListCall) QuotaUser(quotaUser string) *AccountsListCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *AccountsListCall) UserIP(userIP string) *AccountsListCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsListCall) Fields(s ...googleapi.Field) *AccountsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsListCall) IfNoneMatch(entityTag string) *AccountsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsListCall) Context(ctx context.Context) *AccountsListCall {
	c.ctx_ = ctx
	return c
}

func (c *AccountsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.accounts.list" call.
// Exactly one of *Accounts or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Accounts.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccountsListCall) Do() (*Accounts, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Accounts{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all accounts available to this AdSense account.",
	//   "httpMethod": "GET",
	//   "id": "adsense.accounts.list",
	//   "parameters": {
	//     "maxResults": {
	//       "description": "The maximum number of accounts to include in the response, used for paging.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through accounts. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts",
	//   "response": {
	//     "$ref": "Accounts"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.accounts.adclients.list":

type AccountsAdclientsListCall struct {
	s            *Service
	accountId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: List all ad clients in the specified account.
func (r *AccountsAdclientsService) List(accountId string) *AccountsAdclientsListCall {
	c := &AccountsAdclientsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of ad clients to include in the response, used for paging.
func (c *AccountsAdclientsListCall) MaxResults(maxResults int64) *AccountsAdclientsListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through ad clients. To retrieve the next page,
// set this parameter to the value of "nextPageToken" from the previous
// response.
func (c *AccountsAdclientsListCall) PageToken(pageToken string) *AccountsAdclientsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *AccountsAdclientsListCall) QuotaUser(quotaUser string) *AccountsAdclientsListCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *AccountsAdclientsListCall) UserIP(userIP string) *AccountsAdclientsListCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsAdclientsListCall) Fields(s ...googleapi.Field) *AccountsAdclientsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsAdclientsListCall) IfNoneMatch(entityTag string) *AccountsAdclientsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsAdclientsListCall) Context(ctx context.Context) *AccountsAdclientsListCall {
	c.ctx_ = ctx
	return c
}

func (c *AccountsAdclientsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/adclients")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId": c.accountId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.accounts.adclients.list" call.
// Exactly one of *AdClients or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *AdClients.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccountsAdclientsListCall) Do() (*AdClients, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AdClients{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all ad clients in the specified account.",
	//   "httpMethod": "GET",
	//   "id": "adsense.accounts.adclients.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account for which to list ad clients.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of ad clients to include in the response, used for paging.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through ad clients. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/adclients",
	//   "response": {
	//     "$ref": "AdClients"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.accounts.adunits.get":

type AccountsAdunitsGetCall struct {
	s            *Service
	accountId    string
	adClientId   string
	adUnitId     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Gets the specified ad unit in the specified ad client for the
// specified account.
func (r *AccountsAdunitsService) Get(accountId string, adClientId string, adUnitId string) *AccountsAdunitsGetCall {
	c := &AccountsAdunitsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.adClientId = adClientId
	c.adUnitId = adUnitId
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *AccountsAdunitsGetCall) QuotaUser(quotaUser string) *AccountsAdunitsGetCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *AccountsAdunitsGetCall) UserIP(userIP string) *AccountsAdunitsGetCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsAdunitsGetCall) Fields(s ...googleapi.Field) *AccountsAdunitsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsAdunitsGetCall) IfNoneMatch(entityTag string) *AccountsAdunitsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsAdunitsGetCall) Context(ctx context.Context) *AccountsAdunitsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *AccountsAdunitsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/adclients/{adClientId}/adunits/{adUnitId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":  c.accountId,
		"adClientId": c.adClientId,
		"adUnitId":   c.adUnitId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.accounts.adunits.get" call.
// Exactly one of *AdUnit or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *AdUnit.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AccountsAdunitsGetCall) Do() (*AdUnit, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AdUnit{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the specified ad unit in the specified ad client for the specified account.",
	//   "httpMethod": "GET",
	//   "id": "adsense.accounts.adunits.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "adClientId",
	//     "adUnitId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account to which the ad client belongs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adClientId": {
	//       "description": "Ad client for which to get the ad unit.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adUnitId": {
	//       "description": "Ad unit to retrieve.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/adclients/{adClientId}/adunits/{adUnitId}",
	//   "response": {
	//     "$ref": "AdUnit"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.accounts.adunits.list":

type AccountsAdunitsListCall struct {
	s            *Service
	accountId    string
	adClientId   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: List all ad units in the specified ad client for the specified
// account.
func (r *AccountsAdunitsService) List(accountId string, adClientId string) *AccountsAdunitsListCall {
	c := &AccountsAdunitsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.adClientId = adClientId
	return c
}

// IncludeInactive sets the optional parameter "includeInactive":
// Whether to include inactive ad units. Default: true.
func (c *AccountsAdunitsListCall) IncludeInactive(includeInactive bool) *AccountsAdunitsListCall {
	c.urlParams_.Set("includeInactive", fmt.Sprint(includeInactive))
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of ad units to include in the response, used for paging.
func (c *AccountsAdunitsListCall) MaxResults(maxResults int64) *AccountsAdunitsListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through ad units. To retrieve the next page, set
// this parameter to the value of "nextPageToken" from the previous
// response.
func (c *AccountsAdunitsListCall) PageToken(pageToken string) *AccountsAdunitsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *AccountsAdunitsListCall) QuotaUser(quotaUser string) *AccountsAdunitsListCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *AccountsAdunitsListCall) UserIP(userIP string) *AccountsAdunitsListCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsAdunitsListCall) Fields(s ...googleapi.Field) *AccountsAdunitsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsAdunitsListCall) IfNoneMatch(entityTag string) *AccountsAdunitsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsAdunitsListCall) Context(ctx context.Context) *AccountsAdunitsListCall {
	c.ctx_ = ctx
	return c
}

func (c *AccountsAdunitsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/adclients/{adClientId}/adunits")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":  c.accountId,
		"adClientId": c.adClientId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.accounts.adunits.list" call.
// Exactly one of *AdUnits or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *AdUnits.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AccountsAdunitsListCall) Do() (*AdUnits, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AdUnits{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all ad units in the specified ad client for the specified account.",
	//   "httpMethod": "GET",
	//   "id": "adsense.accounts.adunits.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "adClientId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account to which the ad client belongs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adClientId": {
	//       "description": "Ad client for which to list ad units.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "includeInactive": {
	//       "description": "Whether to include inactive ad units. Default: true.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of ad units to include in the response, used for paging.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through ad units. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/adclients/{adClientId}/adunits",
	//   "response": {
	//     "$ref": "AdUnits"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.accounts.adunits.customchannels.list":

type AccountsAdunitsCustomchannelsListCall struct {
	s            *Service
	accountId    string
	adClientId   string
	adUnitId     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: List all custom channels which the specified ad unit belongs
// to.
func (r *AccountsAdunitsCustomchannelsService) List(accountId string, adClientId string, adUnitId string) *AccountsAdunitsCustomchannelsListCall {
	c := &AccountsAdunitsCustomchannelsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.adClientId = adClientId
	c.adUnitId = adUnitId
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of custom channels to include in the response, used for
// paging.
func (c *AccountsAdunitsCustomchannelsListCall) MaxResults(maxResults int64) *AccountsAdunitsCustomchannelsListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through custom channels. To retrieve the next
// page, set this parameter to the value of "nextPageToken" from the
// previous response.
func (c *AccountsAdunitsCustomchannelsListCall) PageToken(pageToken string) *AccountsAdunitsCustomchannelsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *AccountsAdunitsCustomchannelsListCall) QuotaUser(quotaUser string) *AccountsAdunitsCustomchannelsListCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *AccountsAdunitsCustomchannelsListCall) UserIP(userIP string) *AccountsAdunitsCustomchannelsListCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsAdunitsCustomchannelsListCall) Fields(s ...googleapi.Field) *AccountsAdunitsCustomchannelsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsAdunitsCustomchannelsListCall) IfNoneMatch(entityTag string) *AccountsAdunitsCustomchannelsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsAdunitsCustomchannelsListCall) Context(ctx context.Context) *AccountsAdunitsCustomchannelsListCall {
	c.ctx_ = ctx
	return c
}

func (c *AccountsAdunitsCustomchannelsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/adclients/{adClientId}/adunits/{adUnitId}/customchannels")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":  c.accountId,
		"adClientId": c.adClientId,
		"adUnitId":   c.adUnitId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.accounts.adunits.customchannels.list" call.
// Exactly one of *CustomChannels or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *CustomChannels.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccountsAdunitsCustomchannelsListCall) Do() (*CustomChannels, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &CustomChannels{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all custom channels which the specified ad unit belongs to.",
	//   "httpMethod": "GET",
	//   "id": "adsense.accounts.adunits.customchannels.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "adClientId",
	//     "adUnitId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account to which the ad client belongs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adClientId": {
	//       "description": "Ad client which contains the ad unit.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adUnitId": {
	//       "description": "Ad unit for which to list custom channels.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of custom channels to include in the response, used for paging.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through custom channels. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/adclients/{adClientId}/adunits/{adUnitId}/customchannels",
	//   "response": {
	//     "$ref": "CustomChannels"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.accounts.customchannels.get":

type AccountsCustomchannelsGetCall struct {
	s               *Service
	accountId       string
	adClientId      string
	customChannelId string
	urlParams_      gensupport.URLParams
	ifNoneMatch_    string
	ctx_            context.Context
}

// Get: Get the specified custom channel from the specified ad client
// for the specified account.
func (r *AccountsCustomchannelsService) Get(accountId string, adClientId string, customChannelId string) *AccountsCustomchannelsGetCall {
	c := &AccountsCustomchannelsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.adClientId = adClientId
	c.customChannelId = customChannelId
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *AccountsCustomchannelsGetCall) QuotaUser(quotaUser string) *AccountsCustomchannelsGetCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *AccountsCustomchannelsGetCall) UserIP(userIP string) *AccountsCustomchannelsGetCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsCustomchannelsGetCall) Fields(s ...googleapi.Field) *AccountsCustomchannelsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsCustomchannelsGetCall) IfNoneMatch(entityTag string) *AccountsCustomchannelsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsCustomchannelsGetCall) Context(ctx context.Context) *AccountsCustomchannelsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *AccountsCustomchannelsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/adclients/{adClientId}/customchannels/{customChannelId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":       c.accountId,
		"adClientId":      c.adClientId,
		"customChannelId": c.customChannelId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.accounts.customchannels.get" call.
// Exactly one of *CustomChannel or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *CustomChannel.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccountsCustomchannelsGetCall) Do() (*CustomChannel, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &CustomChannel{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get the specified custom channel from the specified ad client for the specified account.",
	//   "httpMethod": "GET",
	//   "id": "adsense.accounts.customchannels.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "adClientId",
	//     "customChannelId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account to which the ad client belongs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adClientId": {
	//       "description": "Ad client which contains the custom channel.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "customChannelId": {
	//       "description": "Custom channel to retrieve.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/adclients/{adClientId}/customchannels/{customChannelId}",
	//   "response": {
	//     "$ref": "CustomChannel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.accounts.customchannels.list":

type AccountsCustomchannelsListCall struct {
	s            *Service
	accountId    string
	adClientId   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: List all custom channels in the specified ad client for the
// specified account.
func (r *AccountsCustomchannelsService) List(accountId string, adClientId string) *AccountsCustomchannelsListCall {
	c := &AccountsCustomchannelsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.adClientId = adClientId
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of custom channels to include in the response, used for
// paging.
func (c *AccountsCustomchannelsListCall) MaxResults(maxResults int64) *AccountsCustomchannelsListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through custom channels. To retrieve the next
// page, set this parameter to the value of "nextPageToken" from the
// previous response.
func (c *AccountsCustomchannelsListCall) PageToken(pageToken string) *AccountsCustomchannelsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *AccountsCustomchannelsListCall) QuotaUser(quotaUser string) *AccountsCustomchannelsListCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *AccountsCustomchannelsListCall) UserIP(userIP string) *AccountsCustomchannelsListCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsCustomchannelsListCall) Fields(s ...googleapi.Field) *AccountsCustomchannelsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsCustomchannelsListCall) IfNoneMatch(entityTag string) *AccountsCustomchannelsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsCustomchannelsListCall) Context(ctx context.Context) *AccountsCustomchannelsListCall {
	c.ctx_ = ctx
	return c
}

func (c *AccountsCustomchannelsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/adclients/{adClientId}/customchannels")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":  c.accountId,
		"adClientId": c.adClientId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.accounts.customchannels.list" call.
// Exactly one of *CustomChannels or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *CustomChannels.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccountsCustomchannelsListCall) Do() (*CustomChannels, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &CustomChannels{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all custom channels in the specified ad client for the specified account.",
	//   "httpMethod": "GET",
	//   "id": "adsense.accounts.customchannels.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "adClientId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account to which the ad client belongs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adClientId": {
	//       "description": "Ad client for which to list custom channels.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of custom channels to include in the response, used for paging.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through custom channels. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/adclients/{adClientId}/customchannels",
	//   "response": {
	//     "$ref": "CustomChannels"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.accounts.customchannels.adunits.list":

type AccountsCustomchannelsAdunitsListCall struct {
	s               *Service
	accountId       string
	adClientId      string
	customChannelId string
	urlParams_      gensupport.URLParams
	ifNoneMatch_    string
	ctx_            context.Context
}

// List: List all ad units in the specified custom channel.
func (r *AccountsCustomchannelsAdunitsService) List(accountId string, adClientId string, customChannelId string) *AccountsCustomchannelsAdunitsListCall {
	c := &AccountsCustomchannelsAdunitsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.adClientId = adClientId
	c.customChannelId = customChannelId
	return c
}

// IncludeInactive sets the optional parameter "includeInactive":
// Whether to include inactive ad units. Default: true.
func (c *AccountsCustomchannelsAdunitsListCall) IncludeInactive(includeInactive bool) *AccountsCustomchannelsAdunitsListCall {
	c.urlParams_.Set("includeInactive", fmt.Sprint(includeInactive))
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of ad units to include in the response, used for paging.
func (c *AccountsCustomchannelsAdunitsListCall) MaxResults(maxResults int64) *AccountsCustomchannelsAdunitsListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through ad units. To retrieve the next page, set
// this parameter to the value of "nextPageToken" from the previous
// response.
func (c *AccountsCustomchannelsAdunitsListCall) PageToken(pageToken string) *AccountsCustomchannelsAdunitsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *AccountsCustomchannelsAdunitsListCall) QuotaUser(quotaUser string) *AccountsCustomchannelsAdunitsListCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *AccountsCustomchannelsAdunitsListCall) UserIP(userIP string) *AccountsCustomchannelsAdunitsListCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsCustomchannelsAdunitsListCall) Fields(s ...googleapi.Field) *AccountsCustomchannelsAdunitsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsCustomchannelsAdunitsListCall) IfNoneMatch(entityTag string) *AccountsCustomchannelsAdunitsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsCustomchannelsAdunitsListCall) Context(ctx context.Context) *AccountsCustomchannelsAdunitsListCall {
	c.ctx_ = ctx
	return c
}

func (c *AccountsCustomchannelsAdunitsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/adclients/{adClientId}/customchannels/{customChannelId}/adunits")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":       c.accountId,
		"adClientId":      c.adClientId,
		"customChannelId": c.customChannelId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.accounts.customchannels.adunits.list" call.
// Exactly one of *AdUnits or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *AdUnits.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AccountsCustomchannelsAdunitsListCall) Do() (*AdUnits, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AdUnits{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all ad units in the specified custom channel.",
	//   "httpMethod": "GET",
	//   "id": "adsense.accounts.customchannels.adunits.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "adClientId",
	//     "customChannelId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account to which the ad client belongs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adClientId": {
	//       "description": "Ad client which contains the custom channel.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "customChannelId": {
	//       "description": "Custom channel for which to list ad units.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "includeInactive": {
	//       "description": "Whether to include inactive ad units. Default: true.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of ad units to include in the response, used for paging.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through ad units. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/adclients/{adClientId}/customchannels/{customChannelId}/adunits",
	//   "response": {
	//     "$ref": "AdUnits"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.accounts.reports.generate":

type AccountsReportsGenerateCall struct {
	s            *Service
	accountId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Generate: Generate an AdSense report based on the report request sent
// in the query parameters. Returns the result as JSON; to retrieve
// output in CSV format specify "alt=csv" as a query parameter.
func (r *AccountsReportsService) Generate(accountId string, startDate string, endDate string) *AccountsReportsGenerateCall {
	c := &AccountsReportsGenerateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.urlParams_.Set("startDate", startDate)
	c.urlParams_.Set("endDate", endDate)
	return c
}

// Currency sets the optional parameter "currency": Optional currency to
// use when reporting on monetary metrics. Defaults to the account's
// currency if not set.
func (c *AccountsReportsGenerateCall) Currency(currency string) *AccountsReportsGenerateCall {
	c.urlParams_.Set("currency", currency)
	return c
}

// Dimension sets the optional parameter "dimension": Dimensions to base
// the report on.
func (c *AccountsReportsGenerateCall) Dimension(dimension ...string) *AccountsReportsGenerateCall {
	c.urlParams_.SetMulti("dimension", append([]string{}, dimension...))
	return c
}

// Filter sets the optional parameter "filter": Filters to be run on the
// report.
func (c *AccountsReportsGenerateCall) Filter(filter ...string) *AccountsReportsGenerateCall {
	c.urlParams_.SetMulti("filter", append([]string{}, filter...))
	return c
}

// Locale sets the optional parameter "locale": Optional locale to use
// for translating report output to a local language. Defaults to
// "en_US" if not specified.
func (c *AccountsReportsGenerateCall) Locale(locale string) *AccountsReportsGenerateCall {
	c.urlParams_.Set("locale", locale)
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of rows of report data to return.
func (c *AccountsReportsGenerateCall) MaxResults(maxResults int64) *AccountsReportsGenerateCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// Metric sets the optional parameter "metric": Numeric columns to
// include in the report.
func (c *AccountsReportsGenerateCall) Metric(metric ...string) *AccountsReportsGenerateCall {
	c.urlParams_.SetMulti("metric", append([]string{}, metric...))
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *AccountsReportsGenerateCall) QuotaUser(quotaUser string) *AccountsReportsGenerateCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// Sort sets the optional parameter "sort": The name of a dimension or
// metric to sort the resulting report on, optionally prefixed with "+"
// to sort ascending or "-" to sort descending. If no prefix is
// specified, the column is sorted ascending.
func (c *AccountsReportsGenerateCall) Sort(sort ...string) *AccountsReportsGenerateCall {
	c.urlParams_.SetMulti("sort", append([]string{}, sort...))
	return c
}

// StartIndex sets the optional parameter "startIndex": Index of the
// first row of report data to return.
func (c *AccountsReportsGenerateCall) StartIndex(startIndex int64) *AccountsReportsGenerateCall {
	c.urlParams_.Set("startIndex", fmt.Sprint(startIndex))
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *AccountsReportsGenerateCall) UserIP(userIP string) *AccountsReportsGenerateCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsReportsGenerateCall) Fields(s ...googleapi.Field) *AccountsReportsGenerateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsReportsGenerateCall) IfNoneMatch(entityTag string) *AccountsReportsGenerateCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do and Download
// methods. Any pending HTTP request will be aborted if the provided
// context is canceled.
func (c *AccountsReportsGenerateCall) Context(ctx context.Context) *AccountsReportsGenerateCall {
	c.ctx_ = ctx
	return c
}

func (c *AccountsReportsGenerateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/reports")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId": c.accountId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Download fetches the API endpoint's "media" value, instead of the normal
// API response value. If the returned error is nil, the Response is guaranteed to
// have a 2xx status code. Callers must close the Response.Body as usual.
func (c *AccountsReportsGenerateCall) Download() (*http.Response, error) {
	res, err := c.doRequest("media")
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckMediaResponse(res); err != nil {
		res.Body.Close()
		return nil, err
	}
	return res, nil
}

// Do executes the "adsense.accounts.reports.generate" call.
// Exactly one of *AdsenseReportsGenerateResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *AdsenseReportsGenerateResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccountsReportsGenerateCall) Do() (*AdsenseReportsGenerateResponse, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AdsenseReportsGenerateResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Generate an AdSense report based on the report request sent in the query parameters. Returns the result as JSON; to retrieve output in CSV format specify \"alt=csv\" as a query parameter.",
	//   "httpMethod": "GET",
	//   "id": "adsense.accounts.reports.generate",
	//   "parameterOrder": [
	//     "accountId",
	//     "startDate",
	//     "endDate"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account upon which to report.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "currency": {
	//       "description": "Optional currency to use when reporting on monetary metrics. Defaults to the account's currency if not set.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z]+",
	//       "type": "string"
	//     },
	//     "dimension": {
	//       "description": "Dimensions to base the report on.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z_]+",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "endDate": {
	//       "description": "End of the date range to report on in \"YYYY-MM-DD\" format, inclusive.",
	//       "location": "query",
	//       "pattern": "\\d{4}-\\d{2}-\\d{2}|(today|startOfMonth|startOfYear)(([\\-\\+]\\d+[dwmy]){0,3}?)",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "Filters to be run on the report.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z_]+(==|=@).+",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "locale": {
	//       "description": "Optional locale to use for translating report output to a local language. Defaults to \"en_US\" if not specified.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z_]+",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of rows of report data to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "50000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "metric": {
	//       "description": "Numeric columns to include in the report.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z_]+",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "sort": {
	//       "description": "The name of a dimension or metric to sort the resulting report on, optionally prefixed with \"+\" to sort ascending or \"-\" to sort descending. If no prefix is specified, the column is sorted ascending.",
	//       "location": "query",
	//       "pattern": "(\\+|-)?[a-zA-Z_]+",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "startDate": {
	//       "description": "Start of the date range to report on in \"YYYY-MM-DD\" format, inclusive.",
	//       "location": "query",
	//       "pattern": "\\d{4}-\\d{2}-\\d{2}|(today|startOfMonth|startOfYear)(([\\-\\+]\\d+[dwmy]){0,3}?)",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startIndex": {
	//       "description": "Index of the first row of report data to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "5000",
	//       "minimum": "0",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "accounts/{accountId}/reports",
	//   "response": {
	//     "$ref": "AdsenseReportsGenerateResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ],
	//   "supportsMediaDownload": true
	// }

}

// method id "adsense.accounts.reports.saved.generate":

type AccountsReportsSavedGenerateCall struct {
	s             *Service
	accountId     string
	savedReportId string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
}

// Generate: Generate an AdSense report based on the saved report ID
// sent in the query parameters.
func (r *AccountsReportsSavedService) Generate(accountId string, savedReportId string) *AccountsReportsSavedGenerateCall {
	c := &AccountsReportsSavedGenerateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.savedReportId = savedReportId
	return c
}

// Locale sets the optional parameter "locale": Optional locale to use
// for translating report output to a local language. Defaults to
// "en_US" if not specified.
func (c *AccountsReportsSavedGenerateCall) Locale(locale string) *AccountsReportsSavedGenerateCall {
	c.urlParams_.Set("locale", locale)
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of rows of report data to return.
func (c *AccountsReportsSavedGenerateCall) MaxResults(maxResults int64) *AccountsReportsSavedGenerateCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *AccountsReportsSavedGenerateCall) QuotaUser(quotaUser string) *AccountsReportsSavedGenerateCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// StartIndex sets the optional parameter "startIndex": Index of the
// first row of report data to return.
func (c *AccountsReportsSavedGenerateCall) StartIndex(startIndex int64) *AccountsReportsSavedGenerateCall {
	c.urlParams_.Set("startIndex", fmt.Sprint(startIndex))
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *AccountsReportsSavedGenerateCall) UserIP(userIP string) *AccountsReportsSavedGenerateCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsReportsSavedGenerateCall) Fields(s ...googleapi.Field) *AccountsReportsSavedGenerateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsReportsSavedGenerateCall) IfNoneMatch(entityTag string) *AccountsReportsSavedGenerateCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsReportsSavedGenerateCall) Context(ctx context.Context) *AccountsReportsSavedGenerateCall {
	c.ctx_ = ctx
	return c
}

func (c *AccountsReportsSavedGenerateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/reports/{savedReportId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":     c.accountId,
		"savedReportId": c.savedReportId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.accounts.reports.saved.generate" call.
// Exactly one of *AdsenseReportsGenerateResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *AdsenseReportsGenerateResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccountsReportsSavedGenerateCall) Do() (*AdsenseReportsGenerateResponse, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AdsenseReportsGenerateResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Generate an AdSense report based on the saved report ID sent in the query parameters.",
	//   "httpMethod": "GET",
	//   "id": "adsense.accounts.reports.saved.generate",
	//   "parameterOrder": [
	//     "accountId",
	//     "savedReportId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account to which the saved reports belong.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "locale": {
	//       "description": "Optional locale to use for translating report output to a local language. Defaults to \"en_US\" if not specified.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z_]+",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of rows of report data to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "50000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "savedReportId": {
	//       "description": "The saved report to retrieve.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startIndex": {
	//       "description": "Index of the first row of report data to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "5000",
	//       "minimum": "0",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "accounts/{accountId}/reports/{savedReportId}",
	//   "response": {
	//     "$ref": "AdsenseReportsGenerateResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.accounts.reports.saved.list":

type AccountsReportsSavedListCall struct {
	s            *Service
	accountId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: List all saved reports in the specified AdSense account.
func (r *AccountsReportsSavedService) List(accountId string) *AccountsReportsSavedListCall {
	c := &AccountsReportsSavedListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of saved reports to include in the response, used for paging.
func (c *AccountsReportsSavedListCall) MaxResults(maxResults int64) *AccountsReportsSavedListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through saved reports. To retrieve the next page,
// set this parameter to the value of "nextPageToken" from the previous
// response.
func (c *AccountsReportsSavedListCall) PageToken(pageToken string) *AccountsReportsSavedListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *AccountsReportsSavedListCall) QuotaUser(quotaUser string) *AccountsReportsSavedListCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *AccountsReportsSavedListCall) UserIP(userIP string) *AccountsReportsSavedListCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsReportsSavedListCall) Fields(s ...googleapi.Field) *AccountsReportsSavedListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsReportsSavedListCall) IfNoneMatch(entityTag string) *AccountsReportsSavedListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsReportsSavedListCall) Context(ctx context.Context) *AccountsReportsSavedListCall {
	c.ctx_ = ctx
	return c
}

func (c *AccountsReportsSavedListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/reports/saved")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId": c.accountId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.accounts.reports.saved.list" call.
// Exactly one of *SavedReports or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *SavedReports.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccountsReportsSavedListCall) Do() (*SavedReports, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &SavedReports{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all saved reports in the specified AdSense account.",
	//   "httpMethod": "GET",
	//   "id": "adsense.accounts.reports.saved.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account to which the saved reports belong.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of saved reports to include in the response, used for paging.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through saved reports. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/reports/saved",
	//   "response": {
	//     "$ref": "SavedReports"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.accounts.savedadstyles.get":

type AccountsSavedadstylesGetCall struct {
	s              *Service
	accountId      string
	savedAdStyleId string
	urlParams_     gensupport.URLParams
	ifNoneMatch_   string
	ctx_           context.Context
}

// Get: List a specific saved ad style for the specified account.
func (r *AccountsSavedadstylesService) Get(accountId string, savedAdStyleId string) *AccountsSavedadstylesGetCall {
	c := &AccountsSavedadstylesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.savedAdStyleId = savedAdStyleId
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *AccountsSavedadstylesGetCall) QuotaUser(quotaUser string) *AccountsSavedadstylesGetCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *AccountsSavedadstylesGetCall) UserIP(userIP string) *AccountsSavedadstylesGetCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsSavedadstylesGetCall) Fields(s ...googleapi.Field) *AccountsSavedadstylesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsSavedadstylesGetCall) IfNoneMatch(entityTag string) *AccountsSavedadstylesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsSavedadstylesGetCall) Context(ctx context.Context) *AccountsSavedadstylesGetCall {
	c.ctx_ = ctx
	return c
}

func (c *AccountsSavedadstylesGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/savedadstyles/{savedAdStyleId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":      c.accountId,
		"savedAdStyleId": c.savedAdStyleId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.accounts.savedadstyles.get" call.
// Exactly one of *SavedAdStyle or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *SavedAdStyle.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccountsSavedadstylesGetCall) Do() (*SavedAdStyle, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &SavedAdStyle{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List a specific saved ad style for the specified account.",
	//   "httpMethod": "GET",
	//   "id": "adsense.accounts.savedadstyles.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "savedAdStyleId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account for which to get the saved ad style.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "savedAdStyleId": {
	//       "description": "Saved ad style to retrieve.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/savedadstyles/{savedAdStyleId}",
	//   "response": {
	//     "$ref": "SavedAdStyle"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.accounts.savedadstyles.list":

type AccountsSavedadstylesListCall struct {
	s            *Service
	accountId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: List all saved ad styles in the specified account.
func (r *AccountsSavedadstylesService) List(accountId string) *AccountsSavedadstylesListCall {
	c := &AccountsSavedadstylesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of saved ad styles to include in the response, used for
// paging.
func (c *AccountsSavedadstylesListCall) MaxResults(maxResults int64) *AccountsSavedadstylesListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through saved ad styles. To retrieve the next
// page, set this parameter to the value of "nextPageToken" from the
// previous response.
func (c *AccountsSavedadstylesListCall) PageToken(pageToken string) *AccountsSavedadstylesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *AccountsSavedadstylesListCall) QuotaUser(quotaUser string) *AccountsSavedadstylesListCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *AccountsSavedadstylesListCall) UserIP(userIP string) *AccountsSavedadstylesListCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsSavedadstylesListCall) Fields(s ...googleapi.Field) *AccountsSavedadstylesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsSavedadstylesListCall) IfNoneMatch(entityTag string) *AccountsSavedadstylesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsSavedadstylesListCall) Context(ctx context.Context) *AccountsSavedadstylesListCall {
	c.ctx_ = ctx
	return c
}

func (c *AccountsSavedadstylesListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/savedadstyles")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId": c.accountId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.accounts.savedadstyles.list" call.
// Exactly one of *SavedAdStyles or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *SavedAdStyles.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccountsSavedadstylesListCall) Do() (*SavedAdStyles, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &SavedAdStyles{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all saved ad styles in the specified account.",
	//   "httpMethod": "GET",
	//   "id": "adsense.accounts.savedadstyles.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account for which to list saved ad styles.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of saved ad styles to include in the response, used for paging.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through saved ad styles. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/savedadstyles",
	//   "response": {
	//     "$ref": "SavedAdStyles"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.accounts.urlchannels.list":

type AccountsUrlchannelsListCall struct {
	s            *Service
	accountId    string
	adClientId   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: List all URL channels in the specified ad client for the
// specified account.
func (r *AccountsUrlchannelsService) List(accountId string, adClientId string) *AccountsUrlchannelsListCall {
	c := &AccountsUrlchannelsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.accountId = accountId
	c.adClientId = adClientId
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of URL channels to include in the response, used for paging.
func (c *AccountsUrlchannelsListCall) MaxResults(maxResults int64) *AccountsUrlchannelsListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through URL channels. To retrieve the next page,
// set this parameter to the value of "nextPageToken" from the previous
// response.
func (c *AccountsUrlchannelsListCall) PageToken(pageToken string) *AccountsUrlchannelsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *AccountsUrlchannelsListCall) QuotaUser(quotaUser string) *AccountsUrlchannelsListCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *AccountsUrlchannelsListCall) UserIP(userIP string) *AccountsUrlchannelsListCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsUrlchannelsListCall) Fields(s ...googleapi.Field) *AccountsUrlchannelsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsUrlchannelsListCall) IfNoneMatch(entityTag string) *AccountsUrlchannelsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsUrlchannelsListCall) Context(ctx context.Context) *AccountsUrlchannelsListCall {
	c.ctx_ = ctx
	return c
}

func (c *AccountsUrlchannelsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "accounts/{accountId}/adclients/{adClientId}/urlchannels")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"accountId":  c.accountId,
		"adClientId": c.adClientId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.accounts.urlchannels.list" call.
// Exactly one of *UrlChannels or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *UrlChannels.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccountsUrlchannelsListCall) Do() (*UrlChannels, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &UrlChannels{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all URL channels in the specified ad client for the specified account.",
	//   "httpMethod": "GET",
	//   "id": "adsense.accounts.urlchannels.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "adClientId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account to which the ad client belongs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adClientId": {
	//       "description": "Ad client for which to list URL channels.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of URL channels to include in the response, used for paging.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through URL channels. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/adclients/{adClientId}/urlchannels",
	//   "response": {
	//     "$ref": "UrlChannels"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.adclients.list":

type AdclientsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: List all ad clients in this AdSense account.
func (r *AdclientsService) List() *AdclientsListCall {
	c := &AdclientsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of ad clients to include in the response, used for paging.
func (c *AdclientsListCall) MaxResults(maxResults int64) *AdclientsListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through ad clients. To retrieve the next page,
// set this parameter to the value of "nextPageToken" from the previous
// response.
func (c *AdclientsListCall) PageToken(pageToken string) *AdclientsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *AdclientsListCall) QuotaUser(quotaUser string) *AdclientsListCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *AdclientsListCall) UserIP(userIP string) *AdclientsListCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdclientsListCall) Fields(s ...googleapi.Field) *AdclientsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AdclientsListCall) IfNoneMatch(entityTag string) *AdclientsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AdclientsListCall) Context(ctx context.Context) *AdclientsListCall {
	c.ctx_ = ctx
	return c
}

func (c *AdclientsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "adclients")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.adclients.list" call.
// Exactly one of *AdClients or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *AdClients.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AdclientsListCall) Do() (*AdClients, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AdClients{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all ad clients in this AdSense account.",
	//   "httpMethod": "GET",
	//   "id": "adsense.adclients.list",
	//   "parameters": {
	//     "maxResults": {
	//       "description": "The maximum number of ad clients to include in the response, used for paging.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through ad clients. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "adclients",
	//   "response": {
	//     "$ref": "AdClients"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.adunits.get":

type AdunitsGetCall struct {
	s            *Service
	adClientId   string
	adUnitId     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Gets the specified ad unit in the specified ad client.
func (r *AdunitsService) Get(adClientId string, adUnitId string) *AdunitsGetCall {
	c := &AdunitsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.adClientId = adClientId
	c.adUnitId = adUnitId
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *AdunitsGetCall) QuotaUser(quotaUser string) *AdunitsGetCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *AdunitsGetCall) UserIP(userIP string) *AdunitsGetCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdunitsGetCall) Fields(s ...googleapi.Field) *AdunitsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AdunitsGetCall) IfNoneMatch(entityTag string) *AdunitsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AdunitsGetCall) Context(ctx context.Context) *AdunitsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *AdunitsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "adclients/{adClientId}/adunits/{adUnitId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"adClientId": c.adClientId,
		"adUnitId":   c.adUnitId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.adunits.get" call.
// Exactly one of *AdUnit or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *AdUnit.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AdunitsGetCall) Do() (*AdUnit, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AdUnit{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the specified ad unit in the specified ad client.",
	//   "httpMethod": "GET",
	//   "id": "adsense.adunits.get",
	//   "parameterOrder": [
	//     "adClientId",
	//     "adUnitId"
	//   ],
	//   "parameters": {
	//     "adClientId": {
	//       "description": "Ad client for which to get the ad unit.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adUnitId": {
	//       "description": "Ad unit to retrieve.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "adclients/{adClientId}/adunits/{adUnitId}",
	//   "response": {
	//     "$ref": "AdUnit"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.adunits.list":

type AdunitsListCall struct {
	s            *Service
	adClientId   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: List all ad units in the specified ad client for this AdSense
// account.
func (r *AdunitsService) List(adClientId string) *AdunitsListCall {
	c := &AdunitsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.adClientId = adClientId
	return c
}

// IncludeInactive sets the optional parameter "includeInactive":
// Whether to include inactive ad units. Default: true.
func (c *AdunitsListCall) IncludeInactive(includeInactive bool) *AdunitsListCall {
	c.urlParams_.Set("includeInactive", fmt.Sprint(includeInactive))
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of ad units to include in the response, used for paging.
func (c *AdunitsListCall) MaxResults(maxResults int64) *AdunitsListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through ad units. To retrieve the next page, set
// this parameter to the value of "nextPageToken" from the previous
// response.
func (c *AdunitsListCall) PageToken(pageToken string) *AdunitsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *AdunitsListCall) QuotaUser(quotaUser string) *AdunitsListCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *AdunitsListCall) UserIP(userIP string) *AdunitsListCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdunitsListCall) Fields(s ...googleapi.Field) *AdunitsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AdunitsListCall) IfNoneMatch(entityTag string) *AdunitsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AdunitsListCall) Context(ctx context.Context) *AdunitsListCall {
	c.ctx_ = ctx
	return c
}

func (c *AdunitsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "adclients/{adClientId}/adunits")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"adClientId": c.adClientId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.adunits.list" call.
// Exactly one of *AdUnits or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *AdUnits.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AdunitsListCall) Do() (*AdUnits, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AdUnits{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all ad units in the specified ad client for this AdSense account.",
	//   "httpMethod": "GET",
	//   "id": "adsense.adunits.list",
	//   "parameterOrder": [
	//     "adClientId"
	//   ],
	//   "parameters": {
	//     "adClientId": {
	//       "description": "Ad client for which to list ad units.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "includeInactive": {
	//       "description": "Whether to include inactive ad units. Default: true.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of ad units to include in the response, used for paging.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through ad units. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "adclients/{adClientId}/adunits",
	//   "response": {
	//     "$ref": "AdUnits"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.adunits.customchannels.list":

type AdunitsCustomchannelsListCall struct {
	s            *Service
	adClientId   string
	adUnitId     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: List all custom channels which the specified ad unit belongs
// to.
func (r *AdunitsCustomchannelsService) List(adClientId string, adUnitId string) *AdunitsCustomchannelsListCall {
	c := &AdunitsCustomchannelsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.adClientId = adClientId
	c.adUnitId = adUnitId
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of custom channels to include in the response, used for
// paging.
func (c *AdunitsCustomchannelsListCall) MaxResults(maxResults int64) *AdunitsCustomchannelsListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through custom channels. To retrieve the next
// page, set this parameter to the value of "nextPageToken" from the
// previous response.
func (c *AdunitsCustomchannelsListCall) PageToken(pageToken string) *AdunitsCustomchannelsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *AdunitsCustomchannelsListCall) QuotaUser(quotaUser string) *AdunitsCustomchannelsListCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *AdunitsCustomchannelsListCall) UserIP(userIP string) *AdunitsCustomchannelsListCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdunitsCustomchannelsListCall) Fields(s ...googleapi.Field) *AdunitsCustomchannelsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AdunitsCustomchannelsListCall) IfNoneMatch(entityTag string) *AdunitsCustomchannelsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AdunitsCustomchannelsListCall) Context(ctx context.Context) *AdunitsCustomchannelsListCall {
	c.ctx_ = ctx
	return c
}

func (c *AdunitsCustomchannelsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "adclients/{adClientId}/adunits/{adUnitId}/customchannels")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"adClientId": c.adClientId,
		"adUnitId":   c.adUnitId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.adunits.customchannels.list" call.
// Exactly one of *CustomChannels or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *CustomChannels.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AdunitsCustomchannelsListCall) Do() (*CustomChannels, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &CustomChannels{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all custom channels which the specified ad unit belongs to.",
	//   "httpMethod": "GET",
	//   "id": "adsense.adunits.customchannels.list",
	//   "parameterOrder": [
	//     "adClientId",
	//     "adUnitId"
	//   ],
	//   "parameters": {
	//     "adClientId": {
	//       "description": "Ad client which contains the ad unit.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adUnitId": {
	//       "description": "Ad unit for which to list custom channels.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of custom channels to include in the response, used for paging.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through custom channels. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "adclients/{adClientId}/adunits/{adUnitId}/customchannels",
	//   "response": {
	//     "$ref": "CustomChannels"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.customchannels.get":

type CustomchannelsGetCall struct {
	s               *Service
	adClientId      string
	customChannelId string
	urlParams_      gensupport.URLParams
	ifNoneMatch_    string
	ctx_            context.Context
}

// Get: Get the specified custom channel from the specified ad client.
func (r *CustomchannelsService) Get(adClientId string, customChannelId string) *CustomchannelsGetCall {
	c := &CustomchannelsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.adClientId = adClientId
	c.customChannelId = customChannelId
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *CustomchannelsGetCall) QuotaUser(quotaUser string) *CustomchannelsGetCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *CustomchannelsGetCall) UserIP(userIP string) *CustomchannelsGetCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomchannelsGetCall) Fields(s ...googleapi.Field) *CustomchannelsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *CustomchannelsGetCall) IfNoneMatch(entityTag string) *CustomchannelsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CustomchannelsGetCall) Context(ctx context.Context) *CustomchannelsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *CustomchannelsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "adclients/{adClientId}/customchannels/{customChannelId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"adClientId":      c.adClientId,
		"customChannelId": c.customChannelId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.customchannels.get" call.
// Exactly one of *CustomChannel or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *CustomChannel.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *CustomchannelsGetCall) Do() (*CustomChannel, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &CustomChannel{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get the specified custom channel from the specified ad client.",
	//   "httpMethod": "GET",
	//   "id": "adsense.customchannels.get",
	//   "parameterOrder": [
	//     "adClientId",
	//     "customChannelId"
	//   ],
	//   "parameters": {
	//     "adClientId": {
	//       "description": "Ad client which contains the custom channel.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "customChannelId": {
	//       "description": "Custom channel to retrieve.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "adclients/{adClientId}/customchannels/{customChannelId}",
	//   "response": {
	//     "$ref": "CustomChannel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.customchannels.list":

type CustomchannelsListCall struct {
	s            *Service
	adClientId   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: List all custom channels in the specified ad client for this
// AdSense account.
func (r *CustomchannelsService) List(adClientId string) *CustomchannelsListCall {
	c := &CustomchannelsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.adClientId = adClientId
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of custom channels to include in the response, used for
// paging.
func (c *CustomchannelsListCall) MaxResults(maxResults int64) *CustomchannelsListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through custom channels. To retrieve the next
// page, set this parameter to the value of "nextPageToken" from the
// previous response.
func (c *CustomchannelsListCall) PageToken(pageToken string) *CustomchannelsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *CustomchannelsListCall) QuotaUser(quotaUser string) *CustomchannelsListCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *CustomchannelsListCall) UserIP(userIP string) *CustomchannelsListCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomchannelsListCall) Fields(s ...googleapi.Field) *CustomchannelsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *CustomchannelsListCall) IfNoneMatch(entityTag string) *CustomchannelsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CustomchannelsListCall) Context(ctx context.Context) *CustomchannelsListCall {
	c.ctx_ = ctx
	return c
}

func (c *CustomchannelsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "adclients/{adClientId}/customchannels")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"adClientId": c.adClientId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.customchannels.list" call.
// Exactly one of *CustomChannels or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *CustomChannels.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *CustomchannelsListCall) Do() (*CustomChannels, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &CustomChannels{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all custom channels in the specified ad client for this AdSense account.",
	//   "httpMethod": "GET",
	//   "id": "adsense.customchannels.list",
	//   "parameterOrder": [
	//     "adClientId"
	//   ],
	//   "parameters": {
	//     "adClientId": {
	//       "description": "Ad client for which to list custom channels.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of custom channels to include in the response, used for paging.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through custom channels. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "adclients/{adClientId}/customchannels",
	//   "response": {
	//     "$ref": "CustomChannels"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.customchannels.adunits.list":

type CustomchannelsAdunitsListCall struct {
	s               *Service
	adClientId      string
	customChannelId string
	urlParams_      gensupport.URLParams
	ifNoneMatch_    string
	ctx_            context.Context
}

// List: List all ad units in the specified custom channel.
func (r *CustomchannelsAdunitsService) List(adClientId string, customChannelId string) *CustomchannelsAdunitsListCall {
	c := &CustomchannelsAdunitsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.adClientId = adClientId
	c.customChannelId = customChannelId
	return c
}

// IncludeInactive sets the optional parameter "includeInactive":
// Whether to include inactive ad units. Default: true.
func (c *CustomchannelsAdunitsListCall) IncludeInactive(includeInactive bool) *CustomchannelsAdunitsListCall {
	c.urlParams_.Set("includeInactive", fmt.Sprint(includeInactive))
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of ad units to include in the response, used for paging.
func (c *CustomchannelsAdunitsListCall) MaxResults(maxResults int64) *CustomchannelsAdunitsListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through ad units. To retrieve the next page, set
// this parameter to the value of "nextPageToken" from the previous
// response.
func (c *CustomchannelsAdunitsListCall) PageToken(pageToken string) *CustomchannelsAdunitsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *CustomchannelsAdunitsListCall) QuotaUser(quotaUser string) *CustomchannelsAdunitsListCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *CustomchannelsAdunitsListCall) UserIP(userIP string) *CustomchannelsAdunitsListCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomchannelsAdunitsListCall) Fields(s ...googleapi.Field) *CustomchannelsAdunitsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *CustomchannelsAdunitsListCall) IfNoneMatch(entityTag string) *CustomchannelsAdunitsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CustomchannelsAdunitsListCall) Context(ctx context.Context) *CustomchannelsAdunitsListCall {
	c.ctx_ = ctx
	return c
}

func (c *CustomchannelsAdunitsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "adclients/{adClientId}/customchannels/{customChannelId}/adunits")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"adClientId":      c.adClientId,
		"customChannelId": c.customChannelId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.customchannels.adunits.list" call.
// Exactly one of *AdUnits or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *AdUnits.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *CustomchannelsAdunitsListCall) Do() (*AdUnits, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AdUnits{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all ad units in the specified custom channel.",
	//   "httpMethod": "GET",
	//   "id": "adsense.customchannels.adunits.list",
	//   "parameterOrder": [
	//     "adClientId",
	//     "customChannelId"
	//   ],
	//   "parameters": {
	//     "adClientId": {
	//       "description": "Ad client which contains the custom channel.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "customChannelId": {
	//       "description": "Custom channel for which to list ad units.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "includeInactive": {
	//       "description": "Whether to include inactive ad units. Default: true.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of ad units to include in the response, used for paging.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through ad units. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "adclients/{adClientId}/customchannels/{customChannelId}/adunits",
	//   "response": {
	//     "$ref": "AdUnits"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.reports.generate":

type ReportsGenerateCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Generate: Generate an AdSense report based on the report request sent
// in the query parameters. Returns the result as JSON; to retrieve
// output in CSV format specify "alt=csv" as a query parameter.
func (r *ReportsService) Generate(startDate string, endDate string) *ReportsGenerateCall {
	c := &ReportsGenerateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.urlParams_.Set("startDate", startDate)
	c.urlParams_.Set("endDate", endDate)
	return c
}

// AccountId sets the optional parameter "accountId": Accounts upon
// which to report.
func (c *ReportsGenerateCall) AccountId(accountId ...string) *ReportsGenerateCall {
	c.urlParams_.SetMulti("accountId", append([]string{}, accountId...))
	return c
}

// Currency sets the optional parameter "currency": Optional currency to
// use when reporting on monetary metrics. Defaults to the account's
// currency if not set.
func (c *ReportsGenerateCall) Currency(currency string) *ReportsGenerateCall {
	c.urlParams_.Set("currency", currency)
	return c
}

// Dimension sets the optional parameter "dimension": Dimensions to base
// the report on.
func (c *ReportsGenerateCall) Dimension(dimension ...string) *ReportsGenerateCall {
	c.urlParams_.SetMulti("dimension", append([]string{}, dimension...))
	return c
}

// Filter sets the optional parameter "filter": Filters to be run on the
// report.
func (c *ReportsGenerateCall) Filter(filter ...string) *ReportsGenerateCall {
	c.urlParams_.SetMulti("filter", append([]string{}, filter...))
	return c
}

// Locale sets the optional parameter "locale": Optional locale to use
// for translating report output to a local language. Defaults to
// "en_US" if not specified.
func (c *ReportsGenerateCall) Locale(locale string) *ReportsGenerateCall {
	c.urlParams_.Set("locale", locale)
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of rows of report data to return.
func (c *ReportsGenerateCall) MaxResults(maxResults int64) *ReportsGenerateCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// Metric sets the optional parameter "metric": Numeric columns to
// include in the report.
func (c *ReportsGenerateCall) Metric(metric ...string) *ReportsGenerateCall {
	c.urlParams_.SetMulti("metric", append([]string{}, metric...))
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *ReportsGenerateCall) QuotaUser(quotaUser string) *ReportsGenerateCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// Sort sets the optional parameter "sort": The name of a dimension or
// metric to sort the resulting report on, optionally prefixed with "+"
// to sort ascending or "-" to sort descending. If no prefix is
// specified, the column is sorted ascending.
func (c *ReportsGenerateCall) Sort(sort ...string) *ReportsGenerateCall {
	c.urlParams_.SetMulti("sort", append([]string{}, sort...))
	return c
}

// StartIndex sets the optional parameter "startIndex": Index of the
// first row of report data to return.
func (c *ReportsGenerateCall) StartIndex(startIndex int64) *ReportsGenerateCall {
	c.urlParams_.Set("startIndex", fmt.Sprint(startIndex))
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *ReportsGenerateCall) UserIP(userIP string) *ReportsGenerateCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReportsGenerateCall) Fields(s ...googleapi.Field) *ReportsGenerateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ReportsGenerateCall) IfNoneMatch(entityTag string) *ReportsGenerateCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do and Download
// methods. Any pending HTTP request will be aborted if the provided
// context is canceled.
func (c *ReportsGenerateCall) Context(ctx context.Context) *ReportsGenerateCall {
	c.ctx_ = ctx
	return c
}

func (c *ReportsGenerateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "reports")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Download fetches the API endpoint's "media" value, instead of the normal
// API response value. If the returned error is nil, the Response is guaranteed to
// have a 2xx status code. Callers must close the Response.Body as usual.
func (c *ReportsGenerateCall) Download() (*http.Response, error) {
	res, err := c.doRequest("media")
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckMediaResponse(res); err != nil {
		res.Body.Close()
		return nil, err
	}
	return res, nil
}

// Do executes the "adsense.reports.generate" call.
// Exactly one of *AdsenseReportsGenerateResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *AdsenseReportsGenerateResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ReportsGenerateCall) Do() (*AdsenseReportsGenerateResponse, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AdsenseReportsGenerateResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Generate an AdSense report based on the report request sent in the query parameters. Returns the result as JSON; to retrieve output in CSV format specify \"alt=csv\" as a query parameter.",
	//   "httpMethod": "GET",
	//   "id": "adsense.reports.generate",
	//   "parameterOrder": [
	//     "startDate",
	//     "endDate"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Accounts upon which to report.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "currency": {
	//       "description": "Optional currency to use when reporting on monetary metrics. Defaults to the account's currency if not set.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z]+",
	//       "type": "string"
	//     },
	//     "dimension": {
	//       "description": "Dimensions to base the report on.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z_]+",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "endDate": {
	//       "description": "End of the date range to report on in \"YYYY-MM-DD\" format, inclusive.",
	//       "location": "query",
	//       "pattern": "\\d{4}-\\d{2}-\\d{2}|(today|startOfMonth|startOfYear)(([\\-\\+]\\d+[dwmy]){0,3}?)",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "Filters to be run on the report.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z_]+(==|=@).+",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "locale": {
	//       "description": "Optional locale to use for translating report output to a local language. Defaults to \"en_US\" if not specified.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z_]+",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of rows of report data to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "50000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "metric": {
	//       "description": "Numeric columns to include in the report.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z_]+",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "sort": {
	//       "description": "The name of a dimension or metric to sort the resulting report on, optionally prefixed with \"+\" to sort ascending or \"-\" to sort descending. If no prefix is specified, the column is sorted ascending.",
	//       "location": "query",
	//       "pattern": "(\\+|-)?[a-zA-Z_]+",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "startDate": {
	//       "description": "Start of the date range to report on in \"YYYY-MM-DD\" format, inclusive.",
	//       "location": "query",
	//       "pattern": "\\d{4}-\\d{2}-\\d{2}|(today|startOfMonth|startOfYear)(([\\-\\+]\\d+[dwmy]){0,3}?)",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startIndex": {
	//       "description": "Index of the first row of report data to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "5000",
	//       "minimum": "0",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "reports",
	//   "response": {
	//     "$ref": "AdsenseReportsGenerateResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ],
	//   "supportsMediaDownload": true
	// }

}

// method id "adsense.reports.saved.generate":

type ReportsSavedGenerateCall struct {
	s             *Service
	savedReportId string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
}

// Generate: Generate an AdSense report based on the saved report ID
// sent in the query parameters.
func (r *ReportsSavedService) Generate(savedReportId string) *ReportsSavedGenerateCall {
	c := &ReportsSavedGenerateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.savedReportId = savedReportId
	return c
}

// Locale sets the optional parameter "locale": Optional locale to use
// for translating report output to a local language. Defaults to
// "en_US" if not specified.
func (c *ReportsSavedGenerateCall) Locale(locale string) *ReportsSavedGenerateCall {
	c.urlParams_.Set("locale", locale)
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of rows of report data to return.
func (c *ReportsSavedGenerateCall) MaxResults(maxResults int64) *ReportsSavedGenerateCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *ReportsSavedGenerateCall) QuotaUser(quotaUser string) *ReportsSavedGenerateCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// StartIndex sets the optional parameter "startIndex": Index of the
// first row of report data to return.
func (c *ReportsSavedGenerateCall) StartIndex(startIndex int64) *ReportsSavedGenerateCall {
	c.urlParams_.Set("startIndex", fmt.Sprint(startIndex))
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *ReportsSavedGenerateCall) UserIP(userIP string) *ReportsSavedGenerateCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReportsSavedGenerateCall) Fields(s ...googleapi.Field) *ReportsSavedGenerateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ReportsSavedGenerateCall) IfNoneMatch(entityTag string) *ReportsSavedGenerateCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ReportsSavedGenerateCall) Context(ctx context.Context) *ReportsSavedGenerateCall {
	c.ctx_ = ctx
	return c
}

func (c *ReportsSavedGenerateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "reports/{savedReportId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"savedReportId": c.savedReportId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.reports.saved.generate" call.
// Exactly one of *AdsenseReportsGenerateResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *AdsenseReportsGenerateResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ReportsSavedGenerateCall) Do() (*AdsenseReportsGenerateResponse, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AdsenseReportsGenerateResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Generate an AdSense report based on the saved report ID sent in the query parameters.",
	//   "httpMethod": "GET",
	//   "id": "adsense.reports.saved.generate",
	//   "parameterOrder": [
	//     "savedReportId"
	//   ],
	//   "parameters": {
	//     "locale": {
	//       "description": "Optional locale to use for translating report output to a local language. Defaults to \"en_US\" if not specified.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z_]+",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of rows of report data to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "50000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "savedReportId": {
	//       "description": "The saved report to retrieve.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startIndex": {
	//       "description": "Index of the first row of report data to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "5000",
	//       "minimum": "0",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "reports/{savedReportId}",
	//   "response": {
	//     "$ref": "AdsenseReportsGenerateResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.reports.saved.list":

type ReportsSavedListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: List all saved reports in this AdSense account.
func (r *ReportsSavedService) List() *ReportsSavedListCall {
	c := &ReportsSavedListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of saved reports to include in the response, used for paging.
func (c *ReportsSavedListCall) MaxResults(maxResults int64) *ReportsSavedListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through saved reports. To retrieve the next page,
// set this parameter to the value of "nextPageToken" from the previous
// response.
func (c *ReportsSavedListCall) PageToken(pageToken string) *ReportsSavedListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *ReportsSavedListCall) QuotaUser(quotaUser string) *ReportsSavedListCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *ReportsSavedListCall) UserIP(userIP string) *ReportsSavedListCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReportsSavedListCall) Fields(s ...googleapi.Field) *ReportsSavedListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ReportsSavedListCall) IfNoneMatch(entityTag string) *ReportsSavedListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ReportsSavedListCall) Context(ctx context.Context) *ReportsSavedListCall {
	c.ctx_ = ctx
	return c
}

func (c *ReportsSavedListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "reports/saved")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.reports.saved.list" call.
// Exactly one of *SavedReports or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *SavedReports.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ReportsSavedListCall) Do() (*SavedReports, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &SavedReports{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all saved reports in this AdSense account.",
	//   "httpMethod": "GET",
	//   "id": "adsense.reports.saved.list",
	//   "parameters": {
	//     "maxResults": {
	//       "description": "The maximum number of saved reports to include in the response, used for paging.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through saved reports. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "reports/saved",
	//   "response": {
	//     "$ref": "SavedReports"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.savedadstyles.get":

type SavedadstylesGetCall struct {
	s              *Service
	savedAdStyleId string
	urlParams_     gensupport.URLParams
	ifNoneMatch_   string
	ctx_           context.Context
}

// Get: Get a specific saved ad style from the user's account.
func (r *SavedadstylesService) Get(savedAdStyleId string) *SavedadstylesGetCall {
	c := &SavedadstylesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.savedAdStyleId = savedAdStyleId
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *SavedadstylesGetCall) QuotaUser(quotaUser string) *SavedadstylesGetCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *SavedadstylesGetCall) UserIP(userIP string) *SavedadstylesGetCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SavedadstylesGetCall) Fields(s ...googleapi.Field) *SavedadstylesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *SavedadstylesGetCall) IfNoneMatch(entityTag string) *SavedadstylesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SavedadstylesGetCall) Context(ctx context.Context) *SavedadstylesGetCall {
	c.ctx_ = ctx
	return c
}

func (c *SavedadstylesGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "savedadstyles/{savedAdStyleId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"savedAdStyleId": c.savedAdStyleId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.savedadstyles.get" call.
// Exactly one of *SavedAdStyle or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *SavedAdStyle.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SavedadstylesGetCall) Do() (*SavedAdStyle, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &SavedAdStyle{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get a specific saved ad style from the user's account.",
	//   "httpMethod": "GET",
	//   "id": "adsense.savedadstyles.get",
	//   "parameterOrder": [
	//     "savedAdStyleId"
	//   ],
	//   "parameters": {
	//     "savedAdStyleId": {
	//       "description": "Saved ad style to retrieve.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "savedadstyles/{savedAdStyleId}",
	//   "response": {
	//     "$ref": "SavedAdStyle"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.savedadstyles.list":

type SavedadstylesListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: List all saved ad styles in the user's account.
func (r *SavedadstylesService) List() *SavedadstylesListCall {
	c := &SavedadstylesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of saved ad styles to include in the response, used for
// paging.
func (c *SavedadstylesListCall) MaxResults(maxResults int64) *SavedadstylesListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through saved ad styles. To retrieve the next
// page, set this parameter to the value of "nextPageToken" from the
// previous response.
func (c *SavedadstylesListCall) PageToken(pageToken string) *SavedadstylesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *SavedadstylesListCall) QuotaUser(quotaUser string) *SavedadstylesListCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *SavedadstylesListCall) UserIP(userIP string) *SavedadstylesListCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SavedadstylesListCall) Fields(s ...googleapi.Field) *SavedadstylesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *SavedadstylesListCall) IfNoneMatch(entityTag string) *SavedadstylesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SavedadstylesListCall) Context(ctx context.Context) *SavedadstylesListCall {
	c.ctx_ = ctx
	return c
}

func (c *SavedadstylesListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "savedadstyles")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.savedadstyles.list" call.
// Exactly one of *SavedAdStyles or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *SavedAdStyles.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *SavedadstylesListCall) Do() (*SavedAdStyles, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &SavedAdStyles{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all saved ad styles in the user's account.",
	//   "httpMethod": "GET",
	//   "id": "adsense.savedadstyles.list",
	//   "parameters": {
	//     "maxResults": {
	//       "description": "The maximum number of saved ad styles to include in the response, used for paging.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through saved ad styles. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "savedadstyles",
	//   "response": {
	//     "$ref": "SavedAdStyles"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}

// method id "adsense.urlchannels.list":

type UrlchannelsListCall struct {
	s            *Service
	adClientId   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: List all URL channels in the specified ad client for this
// AdSense account.
func (r *UrlchannelsService) List(adClientId string) *UrlchannelsListCall {
	c := &UrlchannelsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.adClientId = adClientId
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of URL channels to include in the response, used for paging.
func (c *UrlchannelsListCall) MaxResults(maxResults int64) *UrlchannelsListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through URL channels. To retrieve the next page,
// set this parameter to the value of "nextPageToken" from the previous
// response.
func (c *UrlchannelsListCall) PageToken(pageToken string) *UrlchannelsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// QuotaUser sets the optional parameter "quotaUser": Available to use
// for quota purposes for server-side applications. Can be any arbitrary
// string assigned to a user, but should not exceed 40 characters.
// Overrides userIp if both are provided.
func (c *UrlchannelsListCall) QuotaUser(quotaUser string) *UrlchannelsListCall {
	c.urlParams_.Set("quotaUser", quotaUser)
	return c
}

// UserIP sets the optional parameter "userIp": IP address of the site
// where the request originates. Use this if you want to enforce
// per-user limits.
func (c *UrlchannelsListCall) UserIP(userIP string) *UrlchannelsListCall {
	c.urlParams_.Set("userIp", userIP)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlchannelsListCall) Fields(s ...googleapi.Field) *UrlchannelsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *UrlchannelsListCall) IfNoneMatch(entityTag string) *UrlchannelsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UrlchannelsListCall) Context(ctx context.Context) *UrlchannelsListCall {
	c.ctx_ = ctx
	return c
}

func (c *UrlchannelsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "adclients/{adClientId}/urlchannels")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"adClientId": c.adClientId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "adsense.urlchannels.list" call.
// Exactly one of *UrlChannels or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *UrlChannels.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *UrlchannelsListCall) Do() (*UrlChannels, error) {
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &UrlChannels{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all URL channels in the specified ad client for this AdSense account.",
	//   "httpMethod": "GET",
	//   "id": "adsense.urlchannels.list",
	//   "parameterOrder": [
	//     "adClientId"
	//   ],
	//   "parameters": {
	//     "adClientId": {
	//       "description": "Ad client for which to list URL channels.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of URL channels to include in the response, used for paging.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through URL channels. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "adclients/{adClientId}/urlchannels",
	//   "response": {
	//     "$ref": "UrlChannels"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsense",
	//     "https://www.googleapis.com/auth/adsense.readonly"
	//   ]
	// }

}
