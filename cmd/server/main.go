package main

import (
    "log"
    "net/http"

    "github.com/pp00x/hydrate/config"
    "github.com/pp00x/hydrate/internal/router"
    "github.com/pp00x/hydrate/pkg/util"
)

func main() {
    config.InitConfig()
    util.SetupLogger()

    r := router.SetupRouter()

    srv := &http.Server{
        Addr:         ":" + config.AppConfig.Server.Port,
        Handler:      r,
        ReadTimeout:  config.AppConfig.Server.ReadTimeout,
        WriteTimeout: config.AppConfig.Server.WriteTimeout,
    }

    log.Printf("Starting server on port %s", config.AppConfig.Server.Port)
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatalf("Server failed to start: %v", err)
    }
}