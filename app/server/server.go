package main

import pb "example.com/go-playlist-grpc/proto"

type Server struct {
	pb.PlaylistServiceServer
}
