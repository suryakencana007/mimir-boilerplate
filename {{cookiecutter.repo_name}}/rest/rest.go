package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/suryakencana007/mimir"
	"github.com/suryakencana007/mimir/ruuto"
	"{{cookiecutter.repo_path}}/{{cookiecutter.repo_name}}/config"
	"{{cookiecutter.repo_path}}/{{cookiecutter.repo_name}}/service"
)

// Handlers struct rest
type Handlers struct {
	*config.Config
	TracerMiddleware ruuto.Constructor
	Router ruuto.Router
	Health http.HandlerFunc
}

// Router handlers for rest
func Router(handlers Handlers) ruuto.Router {
	router := handlers.Router

	version := router.Group(fmt.Sprintf("/%s", handlers.Rest.Prefix))
	version.GET("/healthz", handlers.Health)

	return router
}

// Health handler func
func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := mimir.Response(r)

		body := map[string]interface{}{
			"Status": "Health is Okay",
			"Method": r.Method,
		}
		// ping DB

		response.Body(body)
		response.APIStatusSuccess(w, r).WriteJSON()
	}
}

// Rest Application func
func Application(cfg *config.Config, logger mimir.Logging) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tracer, cleanupTrace, erTracer := mimir.Tracer(
		fmt.Sprintf("%s-rest", cfg.App.Name),
		cfg.App.Version,
		logger)
	if erTracer != nil {
		return erTracer
	}

	opentracing.SetGlobalTracer(tracer)

	span, ctxTracer := opentracing.StartSpanFromContextWithTracer(ctx, tracer, "RestApplication")
	defer span.Finish()

	router := Middleware(Options{
		Config: cfg,
		Tracer: tracer,
		Service: service.Atomic{},
	})

	// http serve
	serve, cleanup := mimir.ListenAndServe(mimir.ServeOpts{
		Logger: logger,
		Port:   mimir.WebPort(cfg.App.Port),
		Router: router,
	})

	errChan := make(chan error)
	go func(c context.Context) {
		errChan <- serve(c)
	}(ctx)

	interrupt := mimir.InterruptChannelFunc()
	select {
	case <-interrupt:
		cleanupTrace()
		cleanup(ctxTracer)
		cancel()
		return fmt.Errorf("interrupt received, shutting down")
	case err := <-errChan:
		cleanupTrace()
		cleanup(ctxTracer)
		cancel()
		return err
	}
}
