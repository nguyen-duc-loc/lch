package leetcode

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetProblemByID(t *testing.T) {
	lc := New()

	testCases := []struct {
		name          string
		problemID     int32
		checkResponse func(t *testing.T, problem *Problem, err error)
	}{
		{
			name:      "ProblemNotExists",
			problemID: math.MaxInt32,
			checkResponse: func(t *testing.T, problem *Problem, err error) {
				require.Error(t, err)
				require.EqualError(t, err, ErrProblemNotFound.Error())
				require.Empty(t, problem)
			},
		},
		{
			name:      "OK",
			problemID: 1,
			checkResponse: func(t *testing.T, problem *Problem, err error) {
				expectProblem := &Problem{
					Title:     "Two Sum",
					TitleSlug: "two-sum",
					ID:        1,
				}

				require.NoError(t, err)
				require.NotEmpty(t, problem)
				require.Equal(t, problem.Title, expectProblem.Title)
				require.Equal(t, problem.TitleSlug, expectProblem.TitleSlug)
				require.Equal(t, problem.ID, expectProblem.ID)
				require.NotEmpty(t, problem.Difficulty)
				require.Positive(t, problem.AcRate)
				require.NotEmpty(t, problem.Topics)
			},
		},
		{
			name:      "EasyProblem",
			problemID: 9,
			checkResponse: func(t *testing.T, problem *Problem, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, problem)
				require.Equal(t, problem.Difficulty, "Easy")
			},
		},
		{
			name:      "MediumProblem",
			problemID: 2,
			checkResponse: func(t *testing.T, problem *Problem, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, problem)
				require.Equal(t, problem.Difficulty, "Medium")
			},
		},
		{
			name:      "HardProblem",
			problemID: 4,
			checkResponse: func(t *testing.T, problem *Problem, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, problem)
				require.Equal(t, problem.Difficulty, "Hard")
			},
		},
		{
			name:      "PremiumProblem",
			problemID: 158,
			checkResponse: func(t *testing.T, problem *Problem, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, problem)
				require.True(t, problem.PaidOnly)
			},
		},
		{
			name:      "NonPremiumProblem",
			problemID: 10,
			checkResponse: func(t *testing.T, problem *Problem, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, problem)
				require.False(t, problem.PaidOnly)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			problem, err := lc.Problems.GetByID(tc.problemID)
			tc.checkResponse(t, problem, err)
		})
	}
}
