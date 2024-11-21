package leetcode

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/machinebox/graphql"
)

type SolutionService struct {
	client *graphql.Client
}

type Solution struct {
	ProblemID int32
	Language  string
	Content   string
}

var (
	ErrLanguageNotSupported = errors.New("language is not supported")
)

func ErrUnableToGetSolution(problemID int32, language string) error {
	return fmt.Errorf("unable to get solution to problem with ID %d written in %s", problemID, AvailableLanguage[language])
}

func (s *SolutionService) GetByID(problemID int32, language string) (*Solution, error) {
	if !ValidLanguage(language) {
		return nil, ErrLanguageNotSupported
	}

	problemService := &ProblemService{s.client}

	problem, err := problemService.GetByID(problemID)
	if err != nil {
		return nil, err
	}

	solution := &Solution{}
	solution.ProblemID = problemID
	solution.Language = language

	url := fmt.Sprintf("https://raw.githubusercontent.com/walkccc/LeetCode/refs/heads/main/solutions/%d. %s/%d.%s",
		problemID,
		problem.Title,
		problemID,
		language,
	)

	httpResp, err := http.Get(url)
	if err != nil || httpResp.StatusCode != http.StatusOK {
		return solution, ErrUnableToGetSolution(problemID, language)
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return solution, ErrUnableToGetSolution(problemID, language)
	}

	solution.Content = string(body)

	return solution, nil
}
