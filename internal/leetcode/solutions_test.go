package leetcode

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSolutionByID(t *testing.T) {
	lc := New()

	testCases := []struct {
		name          string
		problemID     int32
		lang          string
		checkResponse func(t *testing.T, solution *Solution, err error)
	}{
		{
			name:      "ProblemNotExists",
			problemID: math.MaxInt32,
			lang:      "cpp",
			checkResponse: func(t *testing.T, solution *Solution, err error) {
				require.Error(t, err)
				require.EqualError(t, err, ErrProblemNotFound.Error())
				require.Empty(t, solution)
			},
		},
		{
			name:      "UnsupportedLanguage",
			problemID: 1,
			lang:      "unsupportedLanguage",
			checkResponse: func(t *testing.T, solution *Solution, err error) {
				require.Error(t, err)
				require.EqualError(t, err, ErrLanguageNotSupported.Error())
				require.Empty(t, solution)
			},
		},
		{
			name:      "SolutionNotWrittenInThisLanguage",
			problemID: 1,
			lang:      "sql",
			checkResponse: func(t *testing.T, solution *Solution, err error) {
				require.Error(t, err)
				require.EqualError(t, err, ErrUnableToGetSolution(solution.ProblemID, solution.Language).Error())
				require.Empty(t, solution.Content)
			},
		},
		{
			name:      "CppSolutionOK",
			problemID: 1,
			lang:      "cpp",
			checkResponse: func(t *testing.T, solution *Solution, err error) {
				expectSolution := &Solution{
					ProblemID: 1,
					Language:  "cpp",
				}

				require.NoError(t, err)
				require.NotEmpty(t, solution)
				require.Equal(t, solution.ProblemID, expectSolution.ProblemID)
				require.Equal(t, solution.Language, expectSolution.Language)
				require.NotEmpty(t, solution.Content)
			},
		},
		{
			name:      "CppSolutionOK",
			problemID: 1,
			lang:      "cpp",
			checkResponse: func(t *testing.T, solution *Solution, err error) {
				expectSolution := &Solution{
					ProblemID: 1,
					Language:  "cpp",
				}

				require.NoError(t, err)
				require.NotEmpty(t, solution)
				require.Equal(t, solution.ProblemID, expectSolution.ProblemID)
				require.Equal(t, solution.Language, expectSolution.Language)
				require.NotEmpty(t, solution.Content)
			},
		},
		{
			name:      "TsSolutionOK",
			problemID: 2630,
			lang:      "ts",
			checkResponse: func(t *testing.T, solution *Solution, err error) {
				expectSolution := &Solution{
					ProblemID: 2630,
					Language:  "ts",
				}

				require.NoError(t, err)
				require.NotEmpty(t, solution)
				require.Equal(t, solution.ProblemID, expectSolution.ProblemID)
				require.Equal(t, solution.Language, expectSolution.Language)
				require.NotEmpty(t, solution.Content)
			},
		},
		{
			name:      "JavaSolutionOK",
			problemID: 1,
			lang:      "java",
			checkResponse: func(t *testing.T, solution *Solution, err error) {
				expectSolution := &Solution{
					ProblemID: 1,
					Language:  "java",
				}

				require.NoError(t, err)
				require.NotEmpty(t, solution)
				require.Equal(t, solution.ProblemID, expectSolution.ProblemID)
				require.Equal(t, solution.Language, expectSolution.Language)
				require.NotEmpty(t, solution.Content)
			},
		},
		{
			name:      "PythonSolutionOK",
			problemID: 1,
			lang:      "py",
			checkResponse: func(t *testing.T, solution *Solution, err error) {
				expectSolution := &Solution{
					ProblemID: 1,
					Language:  "py",
				}

				require.NoError(t, err)
				require.NotEmpty(t, solution)
				require.Equal(t, solution.ProblemID, expectSolution.ProblemID)
				require.Equal(t, solution.Language, expectSolution.Language)
				require.NotEmpty(t, solution.Content)
			},
		},
		{
			name:      "MySQLSolutionOK",
			problemID: 175,
			lang:      "sql",
			checkResponse: func(t *testing.T, solution *Solution, err error) {
				expectSolution := &Solution{
					ProblemID: 175,
					Language:  "sql",
				}

				require.NoError(t, err)
				require.NotEmpty(t, solution)
				require.Equal(t, solution.ProblemID, expectSolution.ProblemID)
				require.Equal(t, solution.Language, expectSolution.Language)
				require.NotEmpty(t, solution.Content)
			},
		},
		{
			name:      "ShellSolutionOK",
			problemID: 192,
			lang:      "sh",
			checkResponse: func(t *testing.T, solution *Solution, err error) {
				expectSolution := &Solution{
					ProblemID: 192,
					Language:  "sh",
				}

				require.NoError(t, err)
				require.NotEmpty(t, solution)
				require.Equal(t, solution.ProblemID, expectSolution.ProblemID)
				require.Equal(t, solution.Language, expectSolution.Language)
				require.NotEmpty(t, solution.Content)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			solution, err := lc.Solutions.GetByID(tc.problemID, tc.lang)
			tc.checkResponse(t, solution, err)
		})
	}
}
