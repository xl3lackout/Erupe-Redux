// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53resolver

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// The current account doesn't have the IAM permissions required to perform
	// the specified Resolver operation.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServiceErrorException for service response error code
	// "InternalServiceErrorException".
	//
	// We encountered an unknown error. Try again in a few minutes.
	ErrCodeInternalServiceErrorException = "InternalServiceErrorException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The value that you specified for NextToken in a List request isn't valid.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// One or more parameters in this request are not valid.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeInvalidPolicyDocument for service response error code
	// "InvalidPolicyDocument".
	//
	// The specified Resolver rule policy is invalid.
	ErrCodeInvalidPolicyDocument = "InvalidPolicyDocument"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// The request is invalid.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeInvalidTagException for service response error code
	// "InvalidTagException".
	//
	// The specified tag is invalid.
	ErrCodeInvalidTagException = "InvalidTagException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The request caused one or more limits to be exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeResourceExistsException for service response error code
	// "ResourceExistsException".
	//
	// The resource that you tried to create already exists.
	ErrCodeResourceExistsException = "ResourceExistsException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// The resource that you tried to update or delete is currently in use.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified resource doesn't exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeResourceUnavailableException for service response error code
	// "ResourceUnavailableException".
	//
	// The specified resource isn't available.
	ErrCodeResourceUnavailableException = "ResourceUnavailableException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request was throttled. Try again in a few minutes.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeUnknownResourceException for service response error code
	// "UnknownResourceException".
	//
	// The specified resource doesn't exist.
	ErrCodeUnknownResourceException = "UnknownResourceException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":         newErrorAccessDeniedException,
	"ConflictException":             newErrorConflictException,
	"InternalServiceErrorException": newErrorInternalServiceErrorException,
	"InvalidNextTokenException":     newErrorInvalidNextTokenException,
	"InvalidParameterException":     newErrorInvalidParameterException,
	"InvalidPolicyDocument":         newErrorInvalidPolicyDocument,
	"InvalidRequestException":       newErrorInvalidRequestException,
	"InvalidTagException":           newErrorInvalidTagException,
	"LimitExceededException":        newErrorLimitExceededException,
	"ResourceExistsException":       newErrorResourceExistsException,
	"ResourceInUseException":        newErrorResourceInUseException,
	"ResourceNotFoundException":     newErrorResourceNotFoundException,
	"ResourceUnavailableException":  newErrorResourceUnavailableException,
	"ThrottlingException":           newErrorThrottlingException,
	"UnknownResourceException":      newErrorUnknownResourceException,
	"ValidationException":           newErrorValidationException,
}
