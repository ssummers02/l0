package app

import (
	"context"
	"encoding/json"
	"l0/internal/bootstrap"
	"l0/internal/delivery/api/handler"
	"l0/internal/delivery/api/restmodel"
	"l0/internal/domain/dto"
	"l0/internal/domain/service"
	"l0/internal/infrastructure/repository"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gocraft/dbr"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

const pathToMigrations = "migrations"

func Run() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	c, err := bootstrap.NewConfig()
	if err != nil {
		log.Fatalln("Error loading config:", err)
	}

	dbPool, err := bootstrap.NewDBConn(c.DB.Username, c.DB.Password, c.DB.Name, c.DB.Host, c.DB.Port)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	defer func(dbPool *dbr.Connection) {
		err := dbPool.Close()
		if err != nil {
			log.Println("Error closing DB pool:", err)
		}
	}(dbPool)

	err = bootstrap.UpMigrations(dbPool.DB, c.DB.Name, pathToMigrations)
	if err != nil {
		log.Fatal("Error up migrations:", err)
	}

	repo := repository.NewRepository(dbPool)
	services := service.NewService(repo)
	srv := handler.NewServer(c.HTTPPort, services)

	go listenNuts(services, c)

	go func() {
		if err := srv.Run(); err != nil {
			logrus.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()

	logrus.Print("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutting down: %s", err.Error())
	}
}

//nolint:errcheck
func listenNuts(s *service.Service, c *bootstrap.Config) {
	// Connect to a server
	nc, err := nats.Connect(c.NatsURL)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	// Simple Async Subscriber
	ec.Subscribe("orders", func(msg *nats.Msg) {
		var order restmodel.Order
		// Unmarshal JSON that represents the Order data
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Printf("Error: %s", err)

			return
		}
		log.Println("Received order:", order.OrderUID)
		err = s.Orders.InsertOrder(context.Background(), dto.OrderFromRest(order))
		if err != nil {
			log.Printf("Error: %s", err)

			return
		}
	})
}
