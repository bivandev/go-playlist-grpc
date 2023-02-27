package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/bivandev/go-playlist-grpc/proto"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.PlaylistServiceServer
	playlist *Playlist
	db       *sql.DB
}

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"db", 5432, "myuser", "secret", "mydb")

	for i := 0; i < 10; i++ { // try connecting 10 times
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Println("Failed to connect to database:", err)
			time.Sleep(5 * time.Second) // wait for 5 seconds before trying again
			continue
		}
		if err := db.Ping(); err != nil {
			log.Println("Failed to ping database:", err)
			time.Sleep(5 * time.Second) // wait for 5 seconds before trying again
			continue
		}
		log.Println("Successfully connected to database")
		break
	}
	defer db.Close()

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database")

	db.Exec(`
		CREATE TABLE IF NOT EXISTS playlist (
			id SERIAL PRIMARY KEY,
			current_song VARCHAR(255) DEFAULT NULL,
			playing BOOLEAN DEFAULT false,
			pause BOOLEAN DEFAULT false
		)CREATE TABLE IF NOT EXISTS songs (
			id SERIAL PRIMARY KEY,
			playlist_id INTEGER NOT NULL REFERENCES playlist(id) ON DELETE CASCADE,
			name VARCHAR(255) NOT NULL,
			duration INTEGER NOT NULL
		)`)

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
