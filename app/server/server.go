package main

import pb "github.com/go-playlist-grpc/proto"

type Server struct {
	pb.PlaylistServiceServer
}
