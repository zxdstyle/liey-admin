package exception

import (
	"errors"
)

// kernel
var (
	ErrNotRegisterKernel      = errors.New("please register the kernel")
	ErrRepeatedRegisterKernel = errors.New("please do not repeated register the kernel")
)

// migration
var (
	ErrRepeatedMigration = errors.New("migration already exists")
)

// plugin
var (
	ErrPluginExists = errors.New("duplicate plugin name or duplicate registered plugin")
)
