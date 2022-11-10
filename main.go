package main

import (
	"context"
	"fetcher/internal/argument"
	"fetcher/internal/config"
	"fetcher/internal/http"
	"fetcher/internal/repository"
	"fetcher/internal/service"
	"fetcher/internal/sqlite"
	"go.uber.org/fx"
	"log"
	"time"
)

func main() {
	var startTimeout config.StartTimeoutType
	var stopTimeout config.StopTimeoutType

	fetcher := fx.New(
		fx.NopLogger,
		fx.Provide(config.NewConfig),
		fx.Provide(config.StartTimeout),
		fx.Provide(config.StopTimeout),
		fx.Populate(&startTimeout),
		fx.Populate(&stopTimeout),
		fx.StartTimeout(time.Duration(startTimeout)),
		fx.StopTimeout(time.Duration(stopTimeout)),
		fx.Provide(argument.NewArgument),
		fx.Provide(service.NewFetchService),
		fx.Provide(http.NewHttpClient),
		fx.Provide(repository.NewNetworkRepository),
		fx.Provide(repository.NewBrowserRepository),
		fx.Provide(repository.NewStorageRepository),
		fx.Provide(sqlite.NewSQLite),
		fx.Provide(repository.NewMetadataRepository),
		fx.Invoke(func(shutdowner fx.Shutdowner, lc fx.Lifecycle, fetchService *service.FetchService) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go func() {
						if err := fetchService.Run(ctx); err != nil {
							log.Fatal(err)
						}

						if err := shutdowner.Shutdown(); err != nil {
							log.Fatal(err)
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return nil
				},
			})
		}),
	)

	startCtx, startCancel := context.WithTimeout(context.Background(), time.Duration(startTimeout))
	defer startCancel()

	if err := fetcher.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	<-fetcher.Done()

	stopCtx, stopCancel := context.WithTimeout(context.Background(), time.Duration(stopTimeout))
	defer stopCancel()

	if err := fetcher.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}
