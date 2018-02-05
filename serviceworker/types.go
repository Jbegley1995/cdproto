package serviceworker

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/jbegley1995/cdproto/target"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Registration serviceWorker registration.
type Registration struct {
	RegistrationID string `json:"registrationId"`
	ScopeURL       string `json:"scopeURL"`
	IsDeleted      bool   `json:"isDeleted"`
}

// VersionRunningStatus [no description].
type VersionRunningStatus string

// String returns the VersionRunningStatus as string value.
func (t VersionRunningStatus) String() string {
	return string(t)
}

// VersionRunningStatus values.
const (
	VersionRunningStatusStopped  VersionRunningStatus = "stopped"
	VersionRunningStatusStarting VersionRunningStatus = "starting"
	VersionRunningStatusRunning  VersionRunningStatus = "running"
	VersionRunningStatusStopping VersionRunningStatus = "stopping"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t VersionRunningStatus) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t VersionRunningStatus) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *VersionRunningStatus) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch VersionRunningStatus(in.String()) {
	case VersionRunningStatusStopped:
		*t = VersionRunningStatusStopped
	case VersionRunningStatusStarting:
		*t = VersionRunningStatusStarting
	case VersionRunningStatusRunning:
		*t = VersionRunningStatusRunning
	case VersionRunningStatusStopping:
		*t = VersionRunningStatusStopping

	default:
		in.AddError(errors.New("unknown VersionRunningStatus value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *VersionRunningStatus) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// VersionStatus [no description].
type VersionStatus string

// String returns the VersionStatus as string value.
func (t VersionStatus) String() string {
	return string(t)
}

// VersionStatus values.
const (
	VersionStatusNew        VersionStatus = "new"
	VersionStatusInstalling VersionStatus = "installing"
	VersionStatusInstalled  VersionStatus = "installed"
	VersionStatusActivating VersionStatus = "activating"
	VersionStatusActivated  VersionStatus = "activated"
	VersionStatusRedundant  VersionStatus = "redundant"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t VersionStatus) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t VersionStatus) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *VersionStatus) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch VersionStatus(in.String()) {
	case VersionStatusNew:
		*t = VersionStatusNew
	case VersionStatusInstalling:
		*t = VersionStatusInstalling
	case VersionStatusInstalled:
		*t = VersionStatusInstalled
	case VersionStatusActivating:
		*t = VersionStatusActivating
	case VersionStatusActivated:
		*t = VersionStatusActivated
	case VersionStatusRedundant:
		*t = VersionStatusRedundant

	default:
		in.AddError(errors.New("unknown VersionStatus value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *VersionStatus) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// Version serviceWorker version.
type Version struct {
	VersionID          string               `json:"versionId"`
	RegistrationID     string               `json:"registrationId"`
	ScriptURL          string               `json:"scriptURL"`
	RunningStatus      VersionRunningStatus `json:"runningStatus"`
	Status             VersionStatus        `json:"status"`
	ScriptLastModified float64              `json:"scriptLastModified,omitempty"` // The Last-Modified header value of the main script.
	ScriptResponseTime float64              `json:"scriptResponseTime,omitempty"` // The time at which the response headers of the main script were received from the server. For cached script it is the last time the cache entry was validated.
	ControlledClients  []target.ID          `json:"controlledClients,omitempty"`
	TargetID           target.ID            `json:"targetId,omitempty"`
}

// ErrorMessage serviceWorker error message.
type ErrorMessage struct {
	ErrorMessage   string `json:"errorMessage"`
	RegistrationID string `json:"registrationId"`
	VersionID      string `json:"versionId"`
	SourceURL      string `json:"sourceURL"`
	LineNumber     int64  `json:"lineNumber"`
	ColumnNumber   int64  `json:"columnNumber"`
}
