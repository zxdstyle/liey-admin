package instances

import (
	"github.com/zxdstyle/liey-admin/framework/exception"
	"github.com/zxdstyle/liey-admin/framework/plugins"
	"github.com/zxdstyle/liey-admin/framework/queue/job"
	"sync"
)

type (
	Kernel interface {
		Boot()
		Plugins() []plugins.Plugin
		Queues() []job.Job
	}
)

var (
	admKernel Kernel
	mu        = &sync.RWMutex{}
)

func RegisterKernel(kernel Kernel) error {
	mu.Lock()
	defer mu.Unlock()

	if admKernel != nil {
		return exception.ErrRepeatedRegisterKernel
	}
	admKernel = kernel
	return nil
}

func GetKernel() (Kernel, error) {
	mu.RLock()
	defer mu.RUnlock()

	if admKernel == nil {
		return nil, exception.ErrNotRegisterKernel
	}
	return admKernel, nil
}
