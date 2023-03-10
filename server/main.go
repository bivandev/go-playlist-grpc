package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "github.com/bivandev/go-playlist-grpc/proto"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.PlaylistServiceServer
}

var (
	db   *sql.DB
	addr string = "0.0.0.0:50051"
	err  error
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"database", 5432, "myuser", "secret", "mydb")

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS playlist (
			id INTEGER PRIMARY KEY,
			current_song VARCHAR(255) DEFAULT NULL,
			playing BOOLEAN DEFAULT false,
			pause BOOLEAN DEFAULT false
		);
		CREATE TABLE IF NOT EXISTS songs (
			id SERIAL PRIMARY KEY,
			playlist_id INTEGER NOT NULL REFERENCES playlist(id) ON DELETE CASCADE,
			name VARCHAR(255) NOT NULL,
			duration BIGINT NOT NULL
		);`)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	log.Printf("Listening at %s\n", addr)

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	pb.RegisterPlaylistServiceServer(s, &Server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
