package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
	pb "github.com/nabin3/userInfo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":8080"
)

type server struct {
	pb.UnimplementedUserServiceServer
	usersMap map[string]*pb.User
}

func (s *server) AddUser(ctx context.Context, in *pb.User) (*pb.UserID, error) {
	out, err := uuid.NewV6()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating User Id: %v", err)
	}

	in.Id = out.String()
	if s.usersMap == nil {
		s.usersMap = make(map[string]*pb.User)
	}
	s.usersMap[in.Id] = in

	return &pb.UserID{Id: in.Id}, nil
}

func (s *server) RetrieveOneUser(ctx context.Context, in *pb.UserID) (*pb.User, error) {
	if out, exist := s.usersMap[in.Id]; !exist {
		return nil, status.Errorf(codes.NotFound, "Given user_id doesn't exist")
	} else {
		return out, nil
	}
}

func (s *server) RetrieveMultipleUsers(ctx context.Context, in *pb.UserIDList) (*pb.UserList, error) {
	out := make([]*pb.User, 0)

	for _, id := range in.Ids {
		if user, exist := s.usersMap[id]; exist {
			out = append(out, user)
		}
	}

	if len(out) == 0 {
		return nil, status.Error(codes.NotFound, "not a single user has been founded with given user_id_list")
	}

	return &pb.UserList{Users: out}, nil
}

func (s *server) SearchUsers(ctx context.Context, in *pb.UserSearchCriteria) (*pb.UserList, error) {
	out := make([]*pb.User, 0)

	for _, user := range s.usersMap {
		if matchCriteria(user, in) {
			out = append(out, user)
		}
	}

	if len(out) == 0 {
		return nil, status.Error(codes.NotFound, "not a single user has been founded with given serch_phrase")
	}

	return &pb.UserList{Users: out}, nil
}

func matchCriteria(user *pb.User, req *pb.UserSearchCriteria) bool {
	if req.GetFname() != "" && user.GetFname() == req.GetFname() {
		return true
	}

	if req.GetCity() != "" && user.GetCity() == req.GetCity() {
		return true
	}

	if req.GetPhone() != "" && user.GetPhone() == req.GetPhone() {
		return true
	}

	if req.GetHeight() != 0.0 && user.GetHeight() == req.GetHeight() {
		return true
	}

	if req.GetIsmarried() && user.GetIsmarried() == req.GetIsmarried() {
		return true
	}

	return false
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Printf("service serving at port: %s", port)
}
