// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/core/v3/backoff.proto

package envoy_config_core_v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on BackoffStrategy with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *BackoffStrategy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BackoffStrategy with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BackoffStrategyMultiError, or nil if none found.
func (m *BackoffStrategy) ValidateAll() error {
	return m.validate(true)
}

func (m *BackoffStrategy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetBaseInterval() == nil {
		err := BackoffStrategyValidationError{
			field:  "BaseInterval",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if d := m.GetBaseInterval(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = BackoffStrategyValidationError{
				field:  "BaseInterval",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gte := time.Duration(0*time.Second + 1000000*time.Nanosecond)

			if dur < gte {
				err := BackoffStrategyValidationError{
					field:  "BaseInterval",
					reason: "value must be greater than or equal to 1ms",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if d := m.GetMaxInterval(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = BackoffStrategyValidationError{
				field:  "MaxInterval",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gt := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur <= gt {
				err := BackoffStrategyValidationError{
					field:  "MaxInterval",
					reason: "value must be greater than 0s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if len(errors) > 0 {
		return BackoffStrategyMultiError(errors)
	}
	return nil
}

// BackoffStrategyMultiError is an error wrapping multiple validation errors
// returned by BackoffStrategy.ValidateAll() if the designated constraints
// aren't met.
type BackoffStrategyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BackoffStrategyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BackoffStrategyMultiError) AllErrors() []error { return m }

// BackoffStrategyValidationError is the validation error returned by
// BackoffStrategy.Validate if the designated constraints aren't met.
type BackoffStrategyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BackoffStrategyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BackoffStrategyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BackoffStrategyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BackoffStrategyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BackoffStrategyValidationError) ErrorName() string { return "BackoffStrategyValidationError" }

// Error satisfies the builtin error interface
func (e BackoffStrategyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBackoffStrategy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BackoffStrategyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BackoffStrategyValidationError{}
