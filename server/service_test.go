package main

import (
	"context"
	"testing"

	pb "github.com/bivandev/go-playlist-grpc/proto"
	"github.com/stretchr/testify/assert"
)

func TestServer_NewPlaylist(t *testing.T) {
	// Initialize a new server and context
	s := &Server{}
	ctx := context.Background()

	// Call the NewPlaylist method and check if the response is successful
	res, err := s.NewPlaylist(ctx, &pb.Empty{})
	assert.NoError(t, err)
	assert.Equal(t, true, res.Success)
}

func TestServer_AddSong(t *testing.T) {
	// Initialize a new server and context
	s := &Server{}
	ctx := context.Background()

	// Call the AddSong method and check if the response is successful
	res, err := s.AddSong(ctx, &pb.Song{Name: "song1", Duration: 180})
	assert.NoError(t, err)
	assert.Equal(t, true, res.Success)
}

func TestServer_GetSongs(t *testing.T) {
	// Initialize a new server and context
	s := &Server{}
	ctx := context.Background()

	// Call the GetSongs method and check if the response is not empty
	res, err := s.GetSongs(ctx, &pb.Empty{})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Songs)
}

func TestServer_GetSong(t *testing.T) {
	// Initialize a new server and context
	s := &Server{}
	ctx := context.Background()

	// Call the GetSong method and check if the response is not empty
	res, err := s.GetSong(ctx, &pb.SongRequest{Name: "song1"})
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestServer_UpdateSong(t *testing.T) {
	// Initialize a new server and context
	s := &Server{}
	ctx := context.Background()

	// Call the UpdateSong method and check if the response is successful
	res, err := s.UpdateSong(ctx, &pb.Song{Name: "song1", Duration: 120})
	assert.NoError(t, err)
	assert.Equal(t, true, res.Success)
}

func TestServer_DeleteSong(t *testing.T) {
	// Initialize a new server and context
	s := &Server{}
	ctx := context.Background()

	// Call the DeleteSong method and check if the response is successful
	res, err := s.DeleteSong(ctx, &pb.SongRequest{Name: "song1"})
	assert.NoError(t, err)
	assert.Equal(t, true, res.Success)
}

func TestServer_Play(t *testing.T) {
	// Initialize a new server and context
	s := &Server{}
	ctx := context.Background()

	// Call the Play method and check if the response is successful
	res, err := s.Play(ctx, &pb.Empty{})
	assert.NoError(t, err)
	assert.Equal(t, true, res.Success)
}

func TestServer_Pause(t *testing.T) {
	// Initialize a new server and context
	s := &Server{}
	ctx := context.Background()

	// Call the Pause method and check if the response is successful
	res, err := s.Pause(ctx, &pb.Empty{})
	assert.NoError(t, err)
	assert.Equal(t, true, res.Success)
}

func TestServer_Next(t *testing.T) {
	// Initialize a new server and context
	s := &Server{}
	ctx := context.Background()

	// Call the Next method and check if the response is successful
	res, err := s.Next(ctx, &pb.Empty{})
	assert.NoError(t, err)
	assert.Equal(t, true, res.Success)
}

func TestServer_Prev(t *testing.T) {
	// Initialize a new server and context
	s := &Server{}
	ctx := context.Background()
	// Call the Prev method and check if the response is successful
	res, err := s.Prev(ctx, &pb.Empty{})
	assert.NoError(t, err)
	assert.Equal(t, true, res.Success)
}
