package main

import (
	"fmt"
	"github.com/fwidjaya20/ecommerce-point-system/cmd/container"
	"github.com/fwidjaya20/ecommerce-point-system/cmd/http"
	nats2 "github.com/fwidjaya20/ecommerce-point-system/cmd/nats"
	"github.com/fwidjaya20/ecommerce-point-system/config"
	"github.com/fwidjaya20/ecommerce-point-system/internal/globals"
	"github.com/fwidjaya20/ecommerce-point-system/lib/database"
	"github.com/fwidjaya20/ecommerce-point-system/lib/nats"
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/stan.go"
	"github.com/oklog/oklog/pkg/group"
	"github.com/rs/cors"
	netHttp "net/http"
	"os"
)

func main() {
	var logger log.Logger
	var g group.Group

	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestamp)
	logger = log.With(logger, "caller", log.DefaultCaller)

	con := globals.DB()
	defer con.Close()

	di := container.New(logger)

	initMigration(con)
	initNATS(logger, &g, di)
	initHTTP(logger, &g, di)

	_ = logger.Log("exit", g.Run())
}

func initMigration(dbConn *sqlx.DB) {
	root, err := os.Getwd()
	if nil != err {
		panic(fmt.Sprintf("failed retrieve root path : %v", err.Error()))
	}

	migrationPath := fmt.Sprintf("%s/%s", root, config.GetEnv(config.MIGRATION_PATH))
	database.Migrate(dbConn.DB, config.GetEnv(config.DB_NAME), migrationPath)
}

func initHTTP(
	logger log.Logger,
	g *group.Group,
	container container.Container,
) {
	_ = logger.Log(logger, "Component", "HTTP")

	HTTP_ADDR := config.GetEnv(config.HTTP_ADDR)

	if len(HTTP_ADDR) < 1 {
		panic(fmt.Sprintf("Environment Missing!\n*%s* is required", HTTP_ADDR))
	}

	var router *chi.Mux
	router = chi.NewRouter()
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(corsHandler.Handler)
	router.Mount("/v1", http.MakeHandler(router, container))

	server := &netHttp.Server{
		Addr:    HTTP_ADDR,
		Handler: router,
	}

	g.Add(
		func() error {
			_ = logger.Log("transport", "debug/HTTP", "addr", HTTP_ADDR)
			return server.ListenAndServe()
		},
		func(err error) {
			if nil != err {
				_ = logger.Log("transport", "debug/HTTP", "addr", HTTP_ADDR, "values", err)
				panic(err)
			}
		},
	)
}

func initNATS(
	logger log.Logger,
	g *group.Group,
	container container.Container,
) stan.Conn {
	var natsAddr = config.GetEnv(config.NATS_ADDR)
	var natsClient = config.GetEnv(config.NATS_CLIENT)
	var natsCluster = config.GetEnv(config.NATS_CLUSTER)

	natsConn, err := stan.Connect(natsCluster, natsClient, stan.NatsURL(natsAddr))

	if nil != err {
		_ = logger.Log("transport", "nats", err)
		os.Exit(1)
	}

	pub := nats.NewPublisher(natsConn, &logger)
	nats.SetGlobalPublisher(pub)

	var natsSubs []stan.Subscription

	g.Add(func() error {
		_ = logger.Log("transport", "nats", "addr", natsAddr)
		natsSubs, err = nats2.NATSSubscribers(natsConn, container, logger)
		if err != nil {
			_ = logger.Log("transport", "nats", "err", err)
			panic(err)
		}
		return nil
	}, func(err error) {
		if err != nil {
			for _, sub := range natsSubs {
				_ = sub.Close()
			}
			_ = natsConn.Close()
			panic(err)
		}
	})

	return natsConn
}