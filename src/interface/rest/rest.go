package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	usecases "dmoniac/src/app/usecase"
	"dmoniac/src/infra/config"

	riwayatHandler "dmoniac/src/interface/rest/handlers/riwayat"
	"dmoniac/src/interface/rest/response"
	"dmoniac/src/interface/rest/route"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

// HttpServer holds the dependencies for a HTTP server.
type HttpServer struct {
	*http.Server
}

// New creates and configures a server serving all application routes.
// The server implements a graceful shutdown.
// chi.Mux is used for registering some convenient middlewares and easy configuration of
// routes using different http verbs.
func New(
	conf config.HttpConf,
	isProd bool,
	useCases usecases.AllUseCases,
) (*HttpServer, error) {
	// wrap all the routes
	routeHandler := makeRoute(conf.XRequestID, conf.Timeout, isProd, useCases)

	// http service
	srv := http.Server{
		Addr:    ":" + conf.Port,
		Handler: routeHandler,
	}

	return &HttpServer{&srv}, nil
}

// makeRoute register routes
func makeRoute(
	xRequestID string,
	timeout int,
	isProd bool,
	useCases usecases.AllUseCases,
) *chi.Mux {

	r := chi.NewRouter()

	// apply common middleware here ...

	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	// logging middleware
	// if !isProd {
	// 	r.Use(middleware.Logger)
	// }

	// if isProd {
	// 	// relic middleware
	// 	r.Use(ms_middleware.InitNewRelicHandler(newRelic.AppName, newRelic.License))
	// }

	// timeout middleware
	if timeout <= 0 {
		log.Fatalf("invalid http timeout")
	}

	// instantiate the handlers here ...

	respClient := response.NewResponseClient()

	rh := riwayatHandler.NewRiwayatHandler(respClient, useCases.RiwayatUC)

	r.Route("/api", func(r chi.Router) {

		r.Mount("/dmoniacs", route.DmoniacRouter(rh))
	})
	return r
}

// Start runs ListenAndServe on the http.Server with graceful shutdown
func (srv *HttpServer) Start(ctx context.Context) {
	// run HTTP service
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// ready to serve
	log.Printf("listen on %s", srv.Addr)

	srv.gracefulShutdown(ctx)
}

func (srv *HttpServer) gracefulShutdown(ctx context.Context) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// wait forever until 1 of signals above are received
	<-quit
	log.Printf("got signal: %v, shutting down server ...", quit)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}

	log.Println("server exiting")
}
