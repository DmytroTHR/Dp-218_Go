package services_test

import (
	"errors"
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"Dp218Go/models"
	"Dp218Go/repositories/mock"
	"Dp218Go/services"
)

type problemUseCasesMock struct {
	repoProblem  *mock.MockProblemRepo
	repoSolution *mock.MockSolutionRepo
	problemUC    *services.ProblemService
}

var _ = Describe("Problem/Solution", Label("problem", "solution"), func() {
	var problemMock *problemUseCasesMock
	var mockProblemExpect *mock.MockProblemRepoMockRecorder
	var mockSolutionExpect *mock.MockSolutionRepoMockRecorder
	var strToJustBefore string

	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		repoProblem := mock.NewMockProblemRepo(ctrl)
		repoSolution := mock.NewMockSolutionRepo(ctrl)
		problemMock = &problemUseCasesMock{
			repoProblem:  repoProblem,
			repoSolution: repoSolution,
			problemUC:    services.NewProblemService(repoProblem, repoSolution),
		}

		Expect(problemMock.problemUC).NotTo(BeNil())
	})

	JustBeforeEach(func() {
		mockProblemExpect = problemMock.repoProblem.EXPECT()
		mockSolutionExpect = problemMock.repoSolution.EXPECT()
		if len(strToJustBefore) > 0 {
			AddReportEntry(strToJustBefore, ReportEntryVisibilityAlways)
			strToJustBefore = ""
		}
	})

	FDescribe("get problem by id", Label("GetProblemByID"), func() {
		var id int
		var errorNotFound error

		BeforeEach(func() {
			id = 1
			errorNotFound = errors.New("not found in DB")
			strToJustBefore = fmt.Sprintf("pass to Just Before from GetProblemByID. Process # %d", GinkgoParallelProcess())
		})

		It("with ID in DB", func() {
			mockProblemExpect.GetProblemByID(id).Return(models.Problem{ID: id}, nil).Times(1)

			result, err := problemMock.problemUC.GetProblemByID(id)
			Expect(err).NotTo(HaveOccurred())
			Expect(result.ID).To(Equal(id))
		})

		It("no ID in DB", func() {
			mockProblemExpect.GetProblemByID(id).Return(models.Problem{}, errorNotFound).Times(1)

			result, err := problemMock.problemUC.GetProblemByID(id)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(errorNotFound))
			Expect(result.ID).To(BeZero())
		})
	})

	DescribeTable("add problem solution", Label("AddProblemSolution"),
		func(problem *models.Problem, solution *models.Solution, expectedError error) {
			alreadySolved := problem.IsSolved
			call1 := mockSolutionExpect.AddProblemSolution(problem.ID, solution).
				Return(expectedError).Times(1)
			call2 := mockProblemExpect.GetProblemByID(problem.ID).After(call1).
				Return(*problem, nil).MaxTimes(1)
			mockProblemExpect.MarkProblemAsSolved(problem).After(call2).
				Return(models.Problem{ID: problem.ID, IsSolved: true}, nil).MaxTimes(2)
			mockSolutionExpect.GetSolutionByProblem(*problem).
				Return(models.Solution{Description: solution.Description}, nil).MaxTimes(1)

			By("adding solution")
			err := problemMock.problemUC.AddProblemSolution(problem.ID, solution)
			if expectedError != nil {
				Expect(err).To(Equal(expectedError))
				return
			}
			Expect(err).To(BeNil())

			By("mark problem as solved")
			*problem, _ = problemMock.problemUC.MarkProblemAsSolved(problem)
			Expect(problem.IsSolved).To(BeTrue())

			By("check solution")
			sol, err := problemMock.problemUC.GetSolutionByProblem(models.Problem{ID: problem.ID, IsSolved: alreadySolved})
			Expect(sol.Description).To(Equal(solution.Description))
			Expect(err).NotTo(HaveOccurred())
		},

		Entry(nil, &models.Problem{ID: 1, IsSolved: false}, &models.Solution{Description: "solved"}, nil, Label("entry1"), Focus),
		Entry(nil, &models.Problem{ID: 1, IsSolved: true}, &models.Solution{Description: "new solution provided"}, nil, Label("entry2"), Focus),
		Entry(nil, &models.Problem{ID: 1, IsSolved: true}, &models.Solution{}, errors.New("no new solution to already solved problem"), Label("entry3")),
		Entry(nil, &models.Problem{ID: 1, IsSolved: false}, &models.Solution{}, errors.New("no description provided"), Label("entry4"), Focus),
	)

	When("some pending test", Pending, FlakeAttempts(2), Label("pending"), func() {
		Specify("some test case", func() {
			Expect(nil).Should(BeNil())
		})
	})
})
