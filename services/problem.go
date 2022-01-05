package services

import (
	"Dp218Go/models"
	"Dp218Go/repositories"
	"context"
	"google.golang.org/grpc"
	"time"
	"problem.micro/proto"
)

// ProblemService - structure for implementing user problem service
type ProblemService struct {
	repoProblem  repositories.ProblemRepo
	repoSolution repositories.SolutionRepo
	microservice proto.ProblemServiceClient
}
func (problserv *ProblemService) unmarshallProblem(problemGRPC *proto.Problem) models.Problem  {
	problem := models.Problem{
		ID:           int(problemGRPC.Id),
		DateReported: time.Unix(problemGRPC.ReportedAt.Seconds, 0),
		Description:  problemGRPC.Description,
		IsSolved:     problemGRPC.IsSolved,
	}
	problserv.AddProblemComplexFields(&problem, int(problemGRPC.Type.Id), 0, int(problemGRPC.UserId))
	return problem
}

func (problserv *ProblemService) marshallProblem(problem *models.Problem) *proto.Problem  {
	return &proto.Problem{
		Id:          int64(problem.ID),
		UserId:      int64(problem.User.ID),
		Description: problem.Description,
		Type:        &proto.ProblemType{Id: int32(problem.Type.ID), Name: problem.Type.Name},
		IsSolved:    problem.IsSolved,
		ReportedAt:  &proto.DateTime{Seconds: problem.DateReported.Unix()},
	}
}

// NewProblemService - initialization of ProblemService
func NewProblemService(repoProblem repositories.ProblemRepo, repoSolution repositories.SolutionRepo, grpcConn grpc.ClientConnInterface) *ProblemService {
	return &ProblemService{
		repoProblem: repoProblem,
		repoSolution: repoSolution,
		microservice: proto.NewProblemServiceClient(grpcConn),
	}
}

// AddNewProblem - add new user problem record
func (problserv *ProblemService) AddNewProblem(problem *models.Problem) error {
	problemType := &proto.ProblemType{
		Id: int32(problem.Type.ID),
	}
	problemToAdd := &proto.Problem{
		UserId:      int64(problem.User.ID),
		Description: problem.Description,
		Type:        problemType,
		IsSolved:    problem.IsSolved,
	}
	_, err := problserv.microservice.AddNewProblem(context.Background(), problemToAdd)
	return err
}

// GetProblemByID - get problem information by its ID
func (problserv *ProblemService) GetProblemByID(problemID int) (models.Problem, error) {
	request := &proto.ProblemRequest{Id: int64(problemID)}
	response, err := problserv.microservice.GetProblemByID(context.Background(), request)

	return problserv.unmarshallProblem(response.Problem), err
}

// MarkProblemAsSolved - update problem record to make problem solved
func (problserv *ProblemService) MarkProblemAsSolved(problem *models.Problem) (models.Problem, error) {
	problem.IsSolved = true
	response, err := problserv.microservice.UpdateProblem(context.Background(), problserv.marshallProblem(problem))

	return problserv.unmarshallProblem(response.Problem), err
}

// GetProblemTypeByID - get problem type record by its ID
func (problserv *ProblemService) GetProblemTypeByID(typeID int) (models.ProblemType, error) {
	return problserv.repoProblem.GetProblemTypeByID(typeID)
}

func (problserv *ProblemService) GetAllProblemTypes() ([]models.ProblemType, error) {
	return problserv.repoProblem.GetAllProblemTypes()
}

// GetProblemsByUserID - get problem list for given user (by user ID)
func (problserv *ProblemService) GetProblemsByUserID(userID int) (*models.ProblemList, error) {
	request := &proto.ProblemRequest{UserId: int64(userID)}
	response, err := problserv.microservice.GetProblemsByUserID(context.Background(), request)
	var resultingList *models.ProblemList
	if err != nil {
		return resultingList, err
	}
	for _, val := range response.Problems {
		resultingList.Problems = append(resultingList.Problems, problserv.unmarshallProblem(val))
	}
	return resultingList, err
}

// GetProblemsByTypeID - get problem list by given problem type ID
func (problserv *ProblemService) GetProblemsByTypeID(typeID int) (*models.ProblemList, error) {
	request := &proto.ProblemRequest{TypeId: int32(typeID)}
	response, err := problserv.microservice.GetProblemsByTypeID(context.Background(), request)
	var resultingList *models.ProblemList
	if err != nil {
		return resultingList, err
	}
	for _, val := range response.Problems {
		resultingList.Problems = append(resultingList.Problems, problserv.unmarshallProblem(val))
	}
	return resultingList, err
}

// GetProblemsByBeingSolved - get problem list by is_solved field value
func (problserv *ProblemService) GetProblemsByBeingSolved(solved bool) (*models.ProblemList, error) {
	request := &proto.ProblemRequest{IsSolved: solved}
	response, err := problserv.microservice.GetProblemsBySolved(context.Background(), request)
	var resultingList *models.ProblemList
	if err != nil {
		return resultingList, err
	}
	for _, val := range response.Problems {
		resultingList.Problems = append(resultingList.Problems, problserv.unmarshallProblem(val))
	}
	return resultingList, err
}

// GetProblemsByTimePeriod - get problem list from time start to time end
func (problserv *ProblemService) GetProblemsByTimePeriod(start, end time.Time) (*models.ProblemList, error) {
	return problserv.repoProblem.GetProblemsByTimePeriod(start, end)
}

// AddProblemComplexFields - fulfill problem model with problem type, scooter, user (by their IDs)
func (problserv *ProblemService) AddProblemComplexFields(problem *models.Problem, typeID, scooterID, userID int) error {
	return problserv.repoProblem.AddProblemComplexFields(problem, typeID, scooterID, userID)
}

// AddProblemSolution - make solution record for given problem (by ID)
func (problserv *ProblemService) AddProblemSolution(problemID int, solution *models.Solution) error {
	err := problserv.repoSolution.AddProblemSolution(problemID, solution)
	if err != nil {
		return err
	}
	problem, err := problserv.repoProblem.GetProblemByID(problemID)
	if err != nil {
		return err
	}
	_, err = problserv.repoProblem.MarkProblemAsSolved(&problem)
	return err
}

// GetSolutionByProblem - get solution for given problem
func (problserv *ProblemService) GetSolutionByProblem(problem models.Problem) (models.Solution, error) {
	return problserv.repoSolution.GetSolutionByProblem(problem)
}

// GetSolutionsByProblems - get solutions for given problems list
func (problserv *ProblemService) GetSolutionsByProblems(problems models.ProblemList) (map[models.Problem]models.Solution, error) {
	return problserv.repoSolution.GetSolutionsByProblems(problems)
}
