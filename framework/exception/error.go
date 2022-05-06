package exception

import (
	"errors"
)

// kernel
var (
	ErrNotRegisterKernel      = errors.New("please register the kernel")
	ErrRepeatedRegisterKernel = errors.New("please do not repeated register the kernel")

	ErrInvalidOrderArgument = errors.New("invalid sort parameter")
)

// migration
var (
	ErrRepeatedMigration = errors.New("migration already exists")
)

// plugin
var (
	ErrPluginExists = errors.New("duplicate plugin name or duplicate registered plugin")
)

// auth
var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrMissingToken = errors.New("token not provider")
	ErrInvalidToken = errors.New("invalid token")
)
