package main

import (
	"net/http"
)

var ErrUserInvalid = apiError{Err: "User not found", Status: http.StatusForbidden}

func (e apiError) Error() string {
	return e.Err
}