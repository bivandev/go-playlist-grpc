package main

import pb "github.com/bivandev/go-playlist-grpc/proto"

type Server struct {
	pb.PlaylistServiceServer
}
