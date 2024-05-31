package main

import (
	"fmt"
	"log"
	"net/http"

	pb "github.com/nabin3/userInfo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	service_port = ":8080"
	api_port     = ":3000"
)

type apiConfig struct {
	client pb.UserServiceClient
}

func main() {
	conn, err := grpc.NewClient("localhost"+service_port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// Initializing apiConfig struct
	cfg := &apiConfig{
		client: client,
	}

	// Defining a router for our server
	mux := http.NewServeMux()

	// "POST /adduser" endpoint
	mux.HandleFunc("POST /adduser", cfg.handlerAddUser)

	// "GET /getuser" endpoint
	mux.HandleFunc("GET /getuser", cfg.handlerRetrieveSingleUser)

	// "GET /get_multiple_users" endpoint
	mux.HandleFunc("GET /get_multiple_users", cfg.handlerRetrieveMultipleUsers)

	// "GET /search_users" endpoint
	mux.HandleFunc("GET /search_users", cfg.handlerSearchUsers)

	// Setting access control headers
	corsMux := middlewareCors(mux)

	// Defining our server
	ourServer := &http.Server{
		Addr:    "localhost" + api_port,
		Handler: corsMux,
	}

	// Server starts to listen
	fmt.Printf("api_server listening on Port: %s\n", api_port)
	log.Fatal(ourServer.ListenAndServe())
}
