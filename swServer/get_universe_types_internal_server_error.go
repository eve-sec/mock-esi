package swServer

import "time"
var _ time.Time

/* 
Internal server error */
type GetUniverseTypesInternalServerError struct {
/*
	 Internal server error message */
	_error string `json:"error,omitempty"`
}
