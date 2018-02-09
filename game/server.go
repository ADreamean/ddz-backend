package game

import (
	"context"
	pb "github.com/ADreamean/ddz-backend/proto"
)

//推送服务器
type PushServer struct{}

func (ps *PushServer) Fetch(pr *pb.PushRequest, pf pb.PushServer_FetchServer) error {
	return nil
}

//游戏逻辑服务器
type Server struct{}

func (s *Server) Request(context.Context, *pb.GameRequest) (*pb.GameResponse, error) {
	//type区分游戏的行为，准备，叫地主，抢地主。出牌




	return &pb.GameResponse{}, nil
}
