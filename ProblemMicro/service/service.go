package service

import (
	"ProblemMicro/proto"
	"ProblemMicro/repository"
	"context"
	"database/sql"
)

type ProblemService struct {
	Repo *repository.ProblemRepo
	*proto.UnimplementedProblemServiceServer
}

func NewProblemService(repo *sql.DB) *ProblemService {
	return &ProblemService{
		Repo: repository.NewProblemRepo(repo),
	}
}

func (serv *ProblemService) AddNewProblem(ctx context.Context, problem *proto.Problem) (*proto.Response, error) {
	problemCreated, err := serv.Repo.Create(ctx, problem)
	return &proto.Response{
		Success: err == nil,
		Problem: problemCreated,
	}, err
}

func (serv *ProblemService) UpdateProblem(ctx context.Context, problem *proto.Problem) (*proto.Response, error) {
	problemUpdated, err := serv.Repo.Update(ctx, problem)
	return &proto.Response{
		Success: err == nil,
		Problem: problemUpdated,
	}, err
}

func (serv *ProblemService) DeleteProblem(ctx context.Context, problem *proto.Problem) (*proto.Response, error) {
	err := serv.Repo.DeleteByID(ctx, problem.Id)
	return &proto.Response{Success: err == nil}, err
}

func (serv *ProblemService) GetProblemByID(ctx context.Context, request *proto.ProblemRequest) (*proto.Response, error) {
	problem, err := serv.Repo.ReadByID(ctx, request.Id)
	return &proto.Response{
		Success: err == nil,
		Problem: problem,
	}, err
}

func (serv *ProblemService) GetAllProblems(ctx context.Context, request *proto.ProblemRequest) (*proto.Response, error) {
	_ = request
	problems, err := serv.Repo.ReadAll(ctx)
	return &proto.Response{
		Success:  err == nil,
		Problems: problems,
	}, err
}

func (serv *ProblemService) GetProblemsByTypeID(ctx context.Context, request *proto.ProblemRequest) (*proto.Response, error) {
	problems, err := serv.Repo.ReadByTypeID(ctx, request.TypeId)
	return &proto.Response{
		Success:  err == nil,
		Problems: problems,
	}, err
}

func (serv *ProblemService) GetProblemsByUserID(ctx context.Context, request *proto.ProblemRequest) (*proto.Response, error) {
	problems, err := serv.Repo.ReadByUserID(ctx, request.UserId)
	return &proto.Response{
		Success:  err == nil,
		Problems: problems,
	}, err
}

func (serv *ProblemService) GetProblemsBySolved(ctx context.Context, request *proto.ProblemRequest) (*proto.Response, error) {
	problems, err := serv.Repo.ReadBySolved(ctx, request.IsSolved)
	return &proto.Response{
		Success:  err == nil,
		Problems: problems,
	}, err
}
