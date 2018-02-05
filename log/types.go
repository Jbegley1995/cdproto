package log

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/jbegley1995/cdproto/network"
	"github.com/jbegley1995/cdproto/runtime"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Entry log entry.
type Entry struct {
	Source           Source                  `json:"source"`                     // Log entry source.
	Level            Level                   `json:"level"`                      // Log entry severity.
	Text             string                  `json:"text"`                       // Logged text.
	Timestamp        *runtime.Timestamp      `json:"timestamp"`                  // Timestamp when this entry was added.
	URL              string                  `json:"url,omitempty"`              // URL of the resource if known.
	LineNumber       int64                   `json:"lineNumber,omitempty"`       // Line number in the resource.
	StackTrace       *runtime.StackTrace     `json:"stackTrace,omitempty"`       // JavaScript stack trace.
	NetworkRequestID network.RequestID       `json:"networkRequestId,omitempty"` // Identifier of the network request associated with this entry.
	WorkerID         string                  `json:"workerId,omitempty"`         // Identifier of the worker associated with this entry.
	Args             []*runtime.RemoteObject `json:"args,omitempty"`             // Call arguments.
}

// ViolationSetting violation configuration setting.
type ViolationSetting struct {
	Name      Violation `json:"name"`      // Violation type.
	Threshold float64   `json:"threshold"` // Time threshold to trigger upon.
}

// Source log entry source.
type Source string

// String returns the Source as string value.
func (t Source) String() string {
	return string(t)
}

// Source values.
const (
	SourceXML            Source = "xml"
	SourceJavascript     Source = "javascript"
	SourceNetwork        Source = "network"
	SourceStorage        Source = "storage"
	SourceAppcache       Source = "appcache"
	SourceRendering      Source = "rendering"
	SourceSecurity       Source = "security"
	SourceDeprecation    Source = "deprecation"
	SourceWorker         Source = "worker"
	SourceViolation      Source = "violation"
	SourceIntervention   Source = "intervention"
	SourceRecommendation Source = "recommendation"
	SourceOther          Source = "other"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t Source) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t Source) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *Source) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch Source(in.String()) {
	case SourceXML:
		*t = SourceXML
	case SourceJavascript:
		*t = SourceJavascript
	case SourceNetwork:
		*t = SourceNetwork
	case SourceStorage:
		*t = SourceStorage
	case SourceAppcache:
		*t = SourceAppcache
	case SourceRendering:
		*t = SourceRendering
	case SourceSecurity:
		*t = SourceSecurity
	case SourceDeprecation:
		*t = SourceDeprecation
	case SourceWorker:
		*t = SourceWorker
	case SourceViolation:
		*t = SourceViolation
	case SourceIntervention:
		*t = SourceIntervention
	case SourceRecommendation:
		*t = SourceRecommendation
	case SourceOther:
		*t = SourceOther

	default:
		in.AddError(errors.New("unknown Source value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *Source) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// Level log entry severity.
type Level string

// String returns the Level as string value.
func (t Level) String() string {
	return string(t)
}

// Level values.
const (
	LevelVerbose Level = "verbose"
	LevelInfo    Level = "info"
	LevelWarning Level = "warning"
	LevelError   Level = "error"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t Level) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t Level) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *Level) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch Level(in.String()) {
	case LevelVerbose:
		*t = LevelVerbose
	case LevelInfo:
		*t = LevelInfo
	case LevelWarning:
		*t = LevelWarning
	case LevelError:
		*t = LevelError

	default:
		in.AddError(errors.New("unknown Level value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *Level) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// Violation violation type.
type Violation string

// String returns the Violation as string value.
func (t Violation) String() string {
	return string(t)
}

// Violation values.
const (
	ViolationLongTask          Violation = "longTask"
	ViolationLongLayout        Violation = "longLayout"
	ViolationBlockedEvent      Violation = "blockedEvent"
	ViolationBlockedParser     Violation = "blockedParser"
	ViolationDiscouragedAPIUse Violation = "discouragedAPIUse"
	ViolationHandler           Violation = "handler"
	ViolationRecurringHandler  Violation = "recurringHandler"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t Violation) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t Violation) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *Violation) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch Violation(in.String()) {
	case ViolationLongTask:
		*t = ViolationLongTask
	case ViolationLongLayout:
		*t = ViolationLongLayout
	case ViolationBlockedEvent:
		*t = ViolationBlockedEvent
	case ViolationBlockedParser:
		*t = ViolationBlockedParser
	case ViolationDiscouragedAPIUse:
		*t = ViolationDiscouragedAPIUse
	case ViolationHandler:
		*t = ViolationHandler
	case ViolationRecurringHandler:
		*t = ViolationRecurringHandler

	default:
		in.AddError(errors.New("unknown Violation value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *Violation) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
