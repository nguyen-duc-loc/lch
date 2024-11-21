package leetcode

import (
	"time"

	"github.com/machinebox/graphql"
)

const (
	queryContextTime time.Duration = 3 * time.Second
	maxRetries                     = 3
)

type Leetcode struct {
	Problems interface {
		GetByID(int32) (*Problem, error)
		GetToday() (*Problem, error)
	}
	Solutions interface {
		GetByID(int32, string) (*Solution, error)
	}
}

func New() *Leetcode {
	client := graphql.NewClient("https://leetcode.com/graphql")

	return &Leetcode{
		Problems:  &ProblemService{client},
		Solutions: &SolutionService{client},
	}
}
