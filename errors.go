package assembly

import "fmt"

type ErrBadRequest struct {
	Message  string
	HTTPCode int
}

func (e *ErrBadRequest) Error() string {
	return fmt.Sprintf("HTTP response code: %d\nError message: %s", e.HTTPCode, e.Message)
}

type ErrUnauthorized struct {
	Message  string
	HTTPCode int
}

func (e *ErrUnauthorized) Error() string {
	return fmt.Sprintf("HTTP response code: %d\nError message: %s", e.HTTPCode, e.Message)
}

type ErrServerError struct {
	Message  string
	HTTPCode int
}

func (e *ErrServerError) Error() string {
	return fmt.Sprintf("HTTP response code: %d\nError message: %s", e.HTTPCode, e.Message)
}

// We throw the following errors based on our validation

type ErrInvalidURL struct {
	Message string
}

func (e *ErrInvalidURL) Error() string {
	return fmt.Sprintf("Error message: %s", e.Message)
}

type ErrUnsupportedFileType struct {
	Message string
}

func (e *ErrUnsupportedFileType) Error() string {
	return fmt.Sprintf("Error message: %s", e.Message)
}

type ErrProcessingNotComplete struct {
	Message string
}

func (e *ErrProcessingNotComplete) Error() string {
	return fmt.Sprintf("Error message: %s", e.Message)
}
