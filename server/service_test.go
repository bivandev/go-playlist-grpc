package main

import (
	"context"
	"testing"

	pb "github.com/bivandev/go-playlist-grpc/proto"
)

func TestNewPlaylist(t *testing.T) {
	s := &Server{}
	ctx := context.Background()
	req := &pb.Empty{}
	_, err := s.NewPlaylist(ctx, req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestAddSong(t *testing.T) {
	s := &Server{}
	ctx := context.Background()
	req := &pb.Song{Name: "song1", Duration: 120}
	_, err := s.AddSong(ctx, req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGetSongs(t *testing.T) {
	s := &Server{}
	ctx := context.Background()
	req := &pb.Empty{}
	_, err := s.GetSongs(ctx, req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGetSong(t *testing.T) {
	s := &Server{}
	ctx := context.Background()
	req := &pb.SongRequest{Name: "song1"}
	_, err := s.GetSong(ctx, req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestUpdateSong(t *testing.T) {
	s := &Server{}
	ctx := context.Background()
	req := &pb.Song{Name: "song1", Duration: 150}
	_, err := s.UpdateSong(ctx, req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestDeleteSong(t *testing.T) {
	s := &Server{}
	ctx := context.Background()
	req := &pb.SongRequest{Name: "song1"}
	_, err := s.DeleteSong(ctx, req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestPlay(t *testing.T) {
	s := &Server{}
	ctx := context.Background()
	req := &pb.Empty{}
	_, err := s.Play(ctx, req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestPause(t *testing.T) {
	s := &Server{}
	ctx := context.Background()
	req := &pb.Empty{}
	_, err := s.Pause(ctx, req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestNext(t *testing.T) {
	s := &Server{}
	ctx := context.Background()
	req := &pb.Empty{}
	_, err := s.Next(ctx, req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestPrev(t *testing.T) {
	s := &Server{}
	ctx := context.Background()
	req := &pb.Empty{}
	_, err := s.Prev(ctx, req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
