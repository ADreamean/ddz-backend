package room

import (
	"sync"
	"errors"
	"google.golang.org/grpc"
	pb "github.com/ADreamean/ddz-backend/proto"
	"context"
	"github.com/ADreamean/ddz-backend/user"
	"github.com/ADreamean/ddz-backend/until/redis"
	"log"
)

var userFull = errors.New("人数已满")
var roomNotExist = errors.New("房间不存在")

type Server struct {
	rooms map[int32]*Room
	lock  sync.Mutex
}

func (s *Server) Join(ctx context.Context, info *pb.RoomJoinRequest) (*pb.RoomInfo, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	room, ok := s.rooms[info.Id]
	if !ok {
		return &pb.RoomInfo{Id: 0, Users: []*pb.User{}}, nil
	}
	room.users = append(room.users, user.User{Id: 0, Name: ""})
	users := make([]*pb.User, 0, len(room.users))

	for _, value := range room.users {
		users = append(users, &pb.User{Id: value.Id, Name: value.Name})
	}

	return &pb.RoomInfo{Id: room.id, Users: users}, nil
}

func (s *Server) Create(ctx context.Context, u *pb.User) (*pb.RoomInfo, error) {
	room := newRoom()
	info := user.Find(u.Id)
	room.Join(*info)
	s.lock.Lock()
	defer s.lock.Unlock()
	s.rooms[room.id] = room

	users := make([]*pb.User, 0, len(room.users))
	for _, user := range room.users {
		users = append(users, &pb.User{Id: user.Id, Name: user.Name})
	}
	return &pb.RoomInfo{Id: room.id, Users: users}, nil
}

type Room struct {
	id    int32       `redis:"id"`
	users []user.User `redis:"users"`
	lock  sync.Mutex  `redis:"-"`
}

func (r *Room) Join(u user.User) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	if len(r.users) >= 3 {
		return userFull
	}
	r.users = append(r.users, u)
	return nil
}

func RegisterGRPC(grsv *grpc.Server) {
	pb.RegisterRoomServer(grsv, &Server{})
}

func newRoom() *Room {
	id, err := redis.Int("INCR", "room")
	if err != nil {
		log.Panic(err)
	}

	return &Room{id: int32(id), users: []user.User{}}
}
