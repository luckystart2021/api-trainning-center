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
)

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

	httpHandler := handlers.NewHandler(db)

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
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}
