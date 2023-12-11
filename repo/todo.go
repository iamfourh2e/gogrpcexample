package repo

import (
	"context"
	"kevin/gogrpc/pb"
)

type TodoImpl struct {
	pb.UnimplementedTodoServiceServer
}

func NewTodoImpl() *TodoImpl {
	return &TodoImpl{}
}

func (t *TodoImpl) Create(c context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{
		Id:  "1",
		Msg: "Mother fker",
	}, nil
}
