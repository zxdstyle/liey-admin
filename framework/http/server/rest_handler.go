package server

import (
	"context"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
)

type (
	RestHandlerFunc func(ctx context.Context, req requests.Request) (*responses.Response, error)

	RestHandler interface {
		Index(ctx context.Context, req requests.Request) (*responses.Response, error)
		Show(ctx context.Context, req requests.Request) (*responses.Response, error)
		Update(ctx context.Context, req requests.Request) (*responses.Response, error)
		Create(ctx context.Context, req requests.Request) (*responses.Response, error)
		Destroy(ctx context.Context, req requests.Request) (*responses.Response, error)
	}

	BatchRestHandler interface {
		BatchUpdate(ctx context.Context, req requests.Request) (*responses.Response, error)
		BatchCreate(ctx context.Context, req requests.Request) (*responses.Response, error)
		BatchDestroy(ctx context.Context, req requests.Request) (*responses.Response, error)
	}
)
