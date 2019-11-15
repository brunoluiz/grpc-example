package service

import (
	"context"

	"github.com/brunoluiz/grpc-example/simple-server/generated/api"
	"github.com/brunoluiz/grpc-example/simple-server/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func New() *GRPC {
	return &GRPC{
		current: 0,
		tasks:   map[int32]*model.Task{},
	}
}

type GRPC struct {
	api.TaskServer
	tasks   map[int32]*model.Task
	current int32
}

func (g *GRPC) GetTask(_ context.Context, req *api.GetTaskRequest) (*api.GetTaskResponse, error) {
	task, ok := g.tasks[req.GetTaskId()]
	if !ok {
		return nil, status.Error(codes.NotFound, "id was not found")
	}

	return &api.GetTaskResponse{
		Title: task.Title,
		Done:  task.Done,
	}, nil
}

func (g *GRPC) CreateTask(_ context.Context, req *api.CreateTaskRequest) (*api.CreateTaskResponse, error) {
	g.current++
	g.tasks[g.current] = &model.Task{
		ID:    g.current,
		Title: req.GetTitle(),
		Done:  false,
	}

	return &api.CreateTaskResponse{TaskId: g.current}, nil
}

func (g *GRPC) UpdateTask(_ context.Context, req *api.UpdateTaskRequest) (*api.UpdateTaskResponse, error) {
	task, ok := g.tasks[req.GetTaskId()]
	if !ok {
		return nil, status.Error(codes.NotFound, "id was not found")
	}

	g.tasks[g.current] = &model.Task{
		ID:    task.ID,
		Title: req.GetTitle(),
		Done:  req.GetDone(),
	}

	return &api.UpdateTaskResponse{}, nil
}
