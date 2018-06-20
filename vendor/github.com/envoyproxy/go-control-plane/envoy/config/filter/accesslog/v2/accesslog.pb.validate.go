// Code generated by protoc-gen-validate
// source: envoy/config/filter/accesslog/v2/accesslog.proto
// DO NOT EDIT!!!

package v2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// Validate checks the field values on AccessLog with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *AccessLog) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if v, ok := interface{}(m.GetFilter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogValidationError{
				Field:  "Filter",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogValidationError{
				Field:  "Config",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// AccessLogValidationError is the validation error returned by
// AccessLog.Validate if the designated constraints aren't met.
type AccessLogValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e AccessLogValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccessLog.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = AccessLogValidationError{}

// Validate checks the field values on AccessLogFilter with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AccessLogFilter) Validate() error {
	if m == nil {
		return nil
	}

	switch m.FilterSpecifier.(type) {

	case *AccessLogFilter_StatusCodeFilter:

		if v, ok := interface{}(m.GetStatusCodeFilter()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogFilterValidationError{
					Field:  "StatusCodeFilter",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *AccessLogFilter_DurationFilter:

		if v, ok := interface{}(m.GetDurationFilter()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogFilterValidationError{
					Field:  "DurationFilter",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *AccessLogFilter_NotHealthCheckFilter:

		if v, ok := interface{}(m.GetNotHealthCheckFilter()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogFilterValidationError{
					Field:  "NotHealthCheckFilter",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *AccessLogFilter_TraceableFilter:

		if v, ok := interface{}(m.GetTraceableFilter()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogFilterValidationError{
					Field:  "TraceableFilter",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *AccessLogFilter_RuntimeFilter:

		if v, ok := interface{}(m.GetRuntimeFilter()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogFilterValidationError{
					Field:  "RuntimeFilter",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *AccessLogFilter_AndFilter:

		if v, ok := interface{}(m.GetAndFilter()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogFilterValidationError{
					Field:  "AndFilter",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *AccessLogFilter_OrFilter:

		if v, ok := interface{}(m.GetOrFilter()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogFilterValidationError{
					Field:  "OrFilter",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *AccessLogFilter_HeaderFilter:

		if v, ok := interface{}(m.GetHeaderFilter()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogFilterValidationError{
					Field:  "HeaderFilter",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	default:
		return AccessLogFilterValidationError{
			Field:  "FilterSpecifier",
			Reason: "value is required",
		}

	}

	return nil
}

// AccessLogFilterValidationError is the validation error returned by
// AccessLogFilter.Validate if the designated constraints aren't met.
type AccessLogFilterValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e AccessLogFilterValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccessLogFilter.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = AccessLogFilterValidationError{}

// Validate checks the field values on ComparisonFilter with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ComparisonFilter) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := ComparisonFilter_Op_name[int32(m.GetOp())]; !ok {
		return ComparisonFilterValidationError{
			Field:  "Op",
			Reason: "value must be one of the defined enum values",
		}
	}

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ComparisonFilterValidationError{
				Field:  "Value",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// ComparisonFilterValidationError is the validation error returned by
// ComparisonFilter.Validate if the designated constraints aren't met.
type ComparisonFilterValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ComparisonFilterValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sComparisonFilter.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ComparisonFilterValidationError{}

// Validate checks the field values on StatusCodeFilter with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *StatusCodeFilter) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetComparison() == nil {
		return StatusCodeFilterValidationError{
			Field:  "Comparison",
			Reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetComparison()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StatusCodeFilterValidationError{
				Field:  "Comparison",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// StatusCodeFilterValidationError is the validation error returned by
// StatusCodeFilter.Validate if the designated constraints aren't met.
type StatusCodeFilterValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e StatusCodeFilterValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatusCodeFilter.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = StatusCodeFilterValidationError{}

// Validate checks the field values on DurationFilter with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DurationFilter) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetComparison() == nil {
		return DurationFilterValidationError{
			Field:  "Comparison",
			Reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetComparison()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DurationFilterValidationError{
				Field:  "Comparison",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// DurationFilterValidationError is the validation error returned by
// DurationFilter.Validate if the designated constraints aren't met.
type DurationFilterValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e DurationFilterValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDurationFilter.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = DurationFilterValidationError{}

// Validate checks the field values on NotHealthCheckFilter with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *NotHealthCheckFilter) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// NotHealthCheckFilterValidationError is the validation error returned by
// NotHealthCheckFilter.Validate if the designated constraints aren't met.
type NotHealthCheckFilterValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e NotHealthCheckFilterValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNotHealthCheckFilter.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = NotHealthCheckFilterValidationError{}

// Validate checks the field values on TraceableFilter with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *TraceableFilter) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// TraceableFilterValidationError is the validation error returned by
// TraceableFilter.Validate if the designated constraints aren't met.
type TraceableFilterValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e TraceableFilterValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTraceableFilter.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = TraceableFilterValidationError{}

// Validate checks the field values on RuntimeFilter with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RuntimeFilter) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetRuntimeKey()) < 1 {
		return RuntimeFilterValidationError{
			Field:  "RuntimeKey",
			Reason: "value length must be at least 1 bytes",
		}
	}

	if v, ok := interface{}(m.GetPercentSampled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RuntimeFilterValidationError{
				Field:  "PercentSampled",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for UseIndependentRandomness

	return nil
}

// RuntimeFilterValidationError is the validation error returned by
// RuntimeFilter.Validate if the designated constraints aren't met.
type RuntimeFilterValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e RuntimeFilterValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRuntimeFilter.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = RuntimeFilterValidationError{}

// Validate checks the field values on AndFilter with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *AndFilter) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetFilters()) < 2 {
		return AndFilterValidationError{
			Field:  "Filters",
			Reason: "value must contain at least 2 item(s)",
		}
	}

	for idx, item := range m.GetFilters() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AndFilterValidationError{
					Field:  fmt.Sprintf("Filters[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// AndFilterValidationError is the validation error returned by
// AndFilter.Validate if the designated constraints aren't met.
type AndFilterValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e AndFilterValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAndFilter.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = AndFilterValidationError{}

// Validate checks the field values on OrFilter with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *OrFilter) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetFilters()) < 2 {
		return OrFilterValidationError{
			Field:  "Filters",
			Reason: "value must contain at least 2 item(s)",
		}
	}

	for idx, item := range m.GetFilters() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OrFilterValidationError{
					Field:  fmt.Sprintf("Filters[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// OrFilterValidationError is the validation error returned by
// OrFilter.Validate if the designated constraints aren't met.
type OrFilterValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e OrFilterValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrFilter.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = OrFilterValidationError{}

// Validate checks the field values on HeaderFilter with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HeaderFilter) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetHeader() == nil {
		return HeaderFilterValidationError{
			Field:  "Header",
			Reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetHeader()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HeaderFilterValidationError{
				Field:  "Header",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// HeaderFilterValidationError is the validation error returned by
// HeaderFilter.Validate if the designated constraints aren't met.
type HeaderFilterValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e HeaderFilterValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeaderFilter.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = HeaderFilterValidationError{}
