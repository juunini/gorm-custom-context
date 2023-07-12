package custom_context

import (
	"context"
	"net/http"

	"gorm.io/gorm"
)

type CustomContext struct {
	Database *gorm.DB
}

const CUSTOM_CONTEXT_KEY = "GORM_CUSTOM_CONTEXT"

func CreateContext(args *CustomContext, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customContext := &CustomContext{
			Database: args.Database,
		}
		requestWithCtx := r.WithContext(
			context.WithValue(
				r.Context(),
				CUSTOM_CONTEXT_KEY,
				customContext,
			),
		)
		next.ServeHTTP(w, requestWithCtx)
	})
}

func GetContext(ctx context.Context) *CustomContext {
	customContext, ok := ctx.Value(CUSTOM_CONTEXT_KEY).(*CustomContext)
	if !ok {
		return nil
	}
	return customContext
}
