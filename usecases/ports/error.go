package ports

import (
	"context"
	"net/http"
)

type ErrOutputPort interface {
	WriteErrResp(ctx context.Context, err error, w http.ResponseWriter, statusCode int)
}
