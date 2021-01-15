package main

import (
	"api-trainning-center/database"
	"api-trainning-center/handlers"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var client *redis.Client

func init() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr:       dsn,
		Password:   "admin123",
		DB:         0,
		MaxRetries: 3,
	})
	pong, err := client.Ping().Result()

	if err != nil {
		fmt.Printf("Cannot Ping: %v\n", err.Error())
	} else {
		fmt.Printf("Pong: %v\n", pong)
	}
}

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
		ForceColors:   true,
	})
	logrus.SetReportCaller(true)
}

func main() {
	addr := ":8080"
	listener, err := net.Listen("tcp", addr)
	db, err := database.Initialize()

	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer db.Close()

	maxConn, err := strconv.Atoi(os.Getenv("MAX_OPEN_CONNS"))
	if err != nil {
		log.Fatalf("Can not parse maxConn: %v", err)
	}
	maxIdleConn, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONNS"))
	if err != nil {
		log.Fatalf("Can not parse maxConn: %v", err)
	}
	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(5 * time.Minute)

	httpHandler := handlers.NewHandler(db, client)

	server := &http.Server{
		Handler: httpHandler,
	}
	go func() {
		server.Serve(listener)
	}()
	defer Stop(server)
	log.Printf("Started server on %s", addr)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping API server.")

}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}
