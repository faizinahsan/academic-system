// Package app configures and runs application.
package app

import (
	"fmt"
	studentsRepo "github.com/faizinahsan/academic-system/internal/repo/students"
	userRepo "github.com/faizinahsan/academic-system/internal/repo/user"
	"github.com/faizinahsan/academic-system/internal/usecase/students"
	"github.com/faizinahsan/academic-system/internal/usecase/translation"
	"os"
	"os/signal"
	"syscall"

	"github.com/faizinahsan/academic-system/config"
	amqprpc "github.com/faizinahsan/academic-system/internal/controller/amqp_rpc"
	"github.com/faizinahsan/academic-system/internal/controller/grpc"
	"github.com/faizinahsan/academic-system/internal/controller/http"
	"github.com/faizinahsan/academic-system/internal/repo/persistent"
	"github.com/faizinahsan/academic-system/internal/repo/webapi"
	"github.com/faizinahsan/academic-system/internal/usecase/user"
	"github.com/faizinahsan/academic-system/pkg/grpcserver"
	"github.com/faizinahsan/academic-system/pkg/httpserver"
	"github.com/faizinahsan/academic-system/pkg/logger"
	"github.com/faizinahsan/academic-system/pkg/postgres"
	"github.com/faizinahsan/academic-system/pkg/rabbitmq/rmq_rpc/server"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use-Case
	translationUseCase := translation.New(
		persistent.New(pg),
		webapi.New(),
	)

	userUseCase := user.New(
		userRepo.New(pg))

	studentsUseCase := students.New(
		studentsRepo.New(pg))

	// RabbitMQ RPC Server
	rmqRouter := amqprpc.NewRouter(translationUseCase, l)

	rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	}

	// gRPC Server
	grpcServer := grpcserver.New(grpcserver.Port(cfg.GRPC.Port))
	grpc.NewRouter(grpcServer.App, translationUseCase, l)

	// HTTP Server
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))
	http.NewRouter(httpServer.App,
		cfg,
		translationUseCase,
		l,
		userUseCase,
		studentsUseCase)

	// Start servers
	rmqServer.Start()
	grpcServer.Start()
	httpServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: %s", s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	case err = <-grpcServer.Notify():
		l.Error(fmt.Errorf("app - Run - grpcServer.Notify: %w", err))
	case err = <-rmqServer.Notify():
		l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	err = grpcServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - grpcServer.Shutdown: %w", err))
	}

	err = rmqServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	}

}

//
//func handleArgs() {
//	flag.Parse()
//	args := flag.Args()
//
//	if len(args) >= 1 {
//		switch args[0] {
//		case "seed":
//			connString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&multiStatements=true", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
//			// connect DB
//			db, err := sql.Open("mysql", connString)
//			if err != nil {
//				log.Fatalf("Error opening DB: %v", err)
//			}
//			seeds.Execute(db, args[1:]...)
//			os.Exit(0)
//		}
//	}
//}
