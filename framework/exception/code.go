package exception

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"net/http"
)

var (
	CodeOK               = gcode.New(http.StatusOK, "OK", nil)                                 // It is OK.
	CodeCreated          = gcode.New(http.StatusCreated, "Created", nil)                       // It is OK.
	CodeInternalError    = gcode.New(http.StatusInternalServerError, "Internal Error", nil)    // An error occurred internally.
	CodeValidationFailed = gcode.New(http.StatusUnprocessableEntity, "Validation Failed", nil) // Data validation failed.
	CodeBadRequest       = gcode.New(http.StatusBadRequest, "Invalid Parameter", nil)          // The given parameter for current operation is invalid.
	CodeUnauthorized     = gcode.New(http.StatusUnauthorized, "Not Authorized", nil)           // Not Authorized.
	CodeNotFound         = gcode.New(http.StatusNotFound, "Not Found", nil)                    // Resource does not exist.
)
