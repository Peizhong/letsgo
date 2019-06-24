package main

import (
	"context"
	"fmt"
	"github.com/peizhong/letsgo/internal"
	"github.com/peizhong/letsgo/pkg/config"
	"github.com/peizhong/letsgo/pkg/log"
	"github.com/peizhong/letsgo/playground/webapi/gateway/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	srv *http.Server
)

const apiPreFix = "/api/"

/*
consul for service discovery
*/

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("I do nothing")
}

func Start() {
	r := mux.NewRouter()
	r.PathPrefix(apiPreFix).HandlerFunc(homeHandler)
	r.Use(middlewares.ErrorMiddleware, middlewares.LoggingMiddleware, middlewares.TracingMiddleware, middlewares.ReRoutingMiddleware, middlewares.RequestMiddleware, middlewares.ResponseMiddleware)
	srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", config.GatewayPort),
		Handler: r,
	}
	log.Info("api gateway service is on")
	if err := srv.ListenAndServe(); err != nil {
		log.Errorf(err.Error())
	}
}

func Stop() {
	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}

func main() {
	internal.Host(Start, Stop)
}