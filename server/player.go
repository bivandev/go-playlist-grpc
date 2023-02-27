package main

import (
	"context"
	"fmt"

	pb "github.com/bivandev/go-playlist-grpc/proto"
)

func (s *Server) NewPlaylist(ctx context.Context, req *pb.Empty) (*pb.NewPlaylistResponse, error) {
	var err error
	s.playlist, err = NewPlaylist(s.db)
	if err != nil {
		return nil, err
	}

	return &pb.NewPlaylistResponse{Success: true}, nil
}

func (s *Server) AddSong(ctx context.Context, req *pb.Song) (*pb.AddSongResponse, error) {
	s.playlist.AddSong(&Song{Name: req.Name, Duration: req.Duration})
	return &pb.AddSongResponse{Success: true}, nil
}

func (s *Server) GetSongs(ctx context.Context, req *pb.Empty) (*pb.GetSongsResponse, error) {
	songs := make([]*pb.Song, 0)
	node := s.playlist.head
	for node != nil {
		songs = append(songs, &pb.Song{Name: node.Song.Name, Duration: int64(node.Song.Duration)})
		node = node.Next
	}
	return &pb.GetSongsResponse{Songs: songs}, nil
}

func (s *Server) GetSong(ctx context.Context, req *pb.SongRequest) (*pb.SongResponse, error) {
	node := s.playlist.head
	for node != nil {
		if node.Song.Name == req.Name {
			return &pb.SongResponse{Name: node.Song.Name, Duration: int64(node.Song.Duration)}, nil
		}
		node = node.Next
	}
	return nil, fmt.Errorf("song not found")
}

func (s *Server) UpdateSong(ctx context.Context, req *pb.Song) (*pb.UpdateSongResponse, error) {
	node := s.playlist.head
	for node != nil {
		if node.Song.Name == req.Name {
			node.Song.Duration = req.Duration
			return &pb.UpdateSongResponse{Success: true}, nil
		}
		node = node.Next
	}
	return nil, fmt.Errorf("song not found")
}

func (s *Server) DeleteSong(ctx context.Context, req *pb.SongRequest) (*pb.DeleteSongResponse, error) {
	err := s.playlist.DelSong(req.Name)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteSongResponse{Success: true}, nil
}

func (s *Server) Play(ctx context.Context, req *pb.Empty) (*pb.PlayResponse, error) {
	err := s.playlist.Play()
	if err != nil {
		return nil, err
	}
	return &pb.PlayResponse{Success: true}, nil
}

func (s *Server) Pause(ctx context.Context, req *pb.Empty) (*pb.PauseResponse, error) {
	s.playlist.Pause()
	return &pb.PauseResponse{Success: true}, nil
}

func (s *Server) Next(ctx context.Context, req *pb.Empty) (*pb.NextResponse, error) {
	err := s.playlist.Next()
	if err != nil {
		return nil, err
	}
	return &pb.NextResponse{Success: true}, nil
}

func (s *Server) Prev(ctx context.Context, req *pb.Empty) (*pb.PrevResponse, error) {
	err := s.playlist.Prev()
	if err != nil {
		return nil, err
	}
	return &pb.PrevResponse{Success: true}, nil
}
