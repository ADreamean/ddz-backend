package login

import (
	pb "github.com/ADreamean/ddz-backend/proto"
	"context"
	"google.golang.org/grpc"
	"github.com/ADreamean/ddz-backend/user"
)

type User struct {
	id   int
	name string
}

func RegisterGRPC(grsv *grpc.Server) {
	pb.RegisterLoginServer(grsv, &Server{})
}

type Server struct{}

func (s *Server) Login(ctx context.Context, info *pb.LoginInfo) (*pb.LoginResult, error) {
	u := user.Create(info.Name)
	return &pb.LoginResult{Code: u.Id}, nil
}
