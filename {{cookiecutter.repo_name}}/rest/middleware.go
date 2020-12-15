package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/suryakencana007/mimir"
	"github.com/suryakencana007/mimir/ruuto"
	"{{cookiecutter.repo_path}}/{{cookiecutter.repo_name}}/config"
	"{{cookiecutter.repo_path}}/{{cookiecutter.repo_name}}/service"
)

// Options struct Middleware func
type Options struct {
	*config.Config
	Tracer  opentracing.Tracer
	Service service.Atomic
}

// Middleware func router rest
func Middleware(opts Options) ruuto.Router {
	router := ruuto.NewChiRouter()
	router.Use(GenerateCallID())
	router.Use(mimir.Recovery())
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mimir.SemanticVersion(r, fmt.Sprintf("%s.%s", opts.App.Name, opts.Rest.Version), opts.App.Version)
			next.ServeHTTP(w, r)
		})
	})
	router.Use(mimir.Logger())

	return Router(Handlers{
		Config:           opts.Config,
		TracerMiddleware: mimir.TracerServer(opts.Tracer, opts.App.Name),
		Router:           router,
		Health:           Health(),
	})
}

// GenerateCallID func for generator call id request
func GenerateCallID() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			if len(r.Header.Get("X-Call-Id")) == 0 {
				callID := mimir.UUID()
				r.Header.Add("X-Call-Id", callID)
				w.Header().Add("X-Call-Id", callID)
			}

			r.Header.Add("X-Start-Time", fmt.Sprintf("%d", start.UTC().UnixNano()))
			w.Header().Add("X-Start-Time", fmt.Sprintf("%d", start.UTC().UnixNano()))

			next.ServeHTTP(w, r)
		})
	}
}
