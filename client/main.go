package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/bivandev/go-playlist-grpc/proto"
)

func main() {
	conn, err := grpc.Dial("grpcserver:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPlaylistServiceClient(conn)

	// Create new playlist
	_, err = client.NewPlaylist(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("could not create new playlist: %v", err)
	}

	// Add songs to the playlist
	songs := []struct {
		name     string
		duration time.Duration
	}{
		{"Song 1", 2 * time.Minute},
		{"Song 2", 3 * time.Minute},
		{"Song 3", 4 * time.Minute},
		{"Song 4", 3 * time.Minute},
		{"Song 3", 4 * time.Minute},
	}

	for _, song := range songs {
		_, err = client.AddSong(context.Background(), &pb.Song{
			Name:     song.name,
			Duration: int64(song.duration.Milliseconds()),
		})
		if err != nil {
			log.Fatalf("could not add song %s to playlist: %v", song.name, err)
		}
	}

	// Get all songs in the playlist
	resp, err := client.GetSongs(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("could not get songs from playlist: %v", err)
	}
	for _, song := range resp.Songs {
		fmt.Printf("Song name: %s, duration: %v\n", song.Name, time.Duration(song.Duration))
	}

	// Play the playlist
	_, err = client.Play(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("could not start playing playlist: %v", err)
	}

	// Pause the playlist
	_, err = client.Pause(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("could not pause playlist: %v", err)
	}

	// Skip to the next song in the playlist
	_, err = client.Next(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("could not skip to the next song: %v", err)
	}

	// Skip back to the previous song in the playlist
	_, err = client.Prev(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("could not skip back to the previous song: %v", err)
	}

	// Get information about a specific song in the playlist
	resp2, err := client.GetSong(context.Background(), &pb.SongRequest{Name: "Song 2"})
	if err != nil {
		log.Fatalf("could not get song information: %v", err)
	}
	fmt.Printf("Song name: %s, duration: %v\n", resp2.Name, time.Duration(resp2.Duration))
}
