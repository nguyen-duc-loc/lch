package leetcode

import (
	"context"
	"sync"
	"time"

	"github.com/machinebox/graphql"
)

var (
	problemDifficultyIdx = map[string]int{
		"Easy":   0,
		"Medium": 1,
		"Hard":   2,
	}
)

type StatsService struct {
	client *graphql.Client
}

type LanguageStats struct {
	Name           string
	ProblemsSolved uint64
}

type ProblemStats struct {
	Difficulty     string
	ProblemsSolved uint64
	Total          uint64
	Beat           float64
}

type Stats struct {
	Languages []LanguageStats
	Problems  [3]ProblemStats
}

func (s *StatsService) getProblemsStats(username string) ([3]ProblemStats, error) {
	type problemsStatsResp struct {
		AllQuestionsCount []struct {
			Difficulty string
			Count      uint64
		}
		MatchedUser struct {
			ProblemsSolvedBeatsStats []struct {
				Difficulty string
				Percentage float64
			}
			SubmitStatsGlobal struct {
				AcSubmissionNum []struct {
					Difficulty string
					Count      uint64
				}
			}
		}
	}

	problemsStatsReq := graphql.NewRequest(`
		query userProblemsSolved($username: String!) {
			allQuestionsCount {
				difficulty
				count
			}
			matchedUser(username: $username) {
				problemsSolvedBeatsStats {
					difficulty
					percentage
				}
				submitStatsGlobal {
					acSubmissionNum {
						difficulty
						count
					}
				}
			}
		}
	`)

	problemsStatsReq.Var("username", username)

	var respData problemsStatsResp

	ctx, cancel := context.WithTimeout(context.Background(), queryContextTime)
	defer cancel()

	problemsStats := [3]ProblemStats{}

	for i := range maxRetries {
		err := s.client.Run(ctx, problemsStatsReq, &respData)
		if nil == err {
			break
		}
		if i == maxRetries-1 {
			return problemsStats, ErrUserNotExists
		} else {
			time.Sleep(time.Second * time.Duration(i+1))
		}
	}

	for difficulty, idx := range problemDifficultyIdx {
		problemsStats[idx].Difficulty = difficulty
	}

	for _, problem := range respData.AllQuestionsCount {
		problemsStats[problemDifficultyIdx[problem.Difficulty]].Total = problem.Count
	}

	for _, problemBeat := range respData.MatchedUser.ProblemsSolvedBeatsStats {
		problemsStats[problemDifficultyIdx[problemBeat.Difficulty]].Beat = problemBeat.Percentage
	}

	for _, problemSolved := range respData.MatchedUser.SubmitStatsGlobal.AcSubmissionNum {
		problemsStats[problemDifficultyIdx[problemSolved.Difficulty]].ProblemsSolved = problemSolved.Count
	}

	return problemsStats, nil
}

func (s *StatsService) getLanguagesStats(username string) ([]LanguageStats, error) {
	type languagesStatsResp struct {
		MatchedUser struct {
			LanguageProblemCount []struct {
				LanguageName   string
				ProblemsSolved uint64
			}
		}
	}

	languagesStatsReq := graphql.NewRequest(`
		query languageStats($username: String!) {
			matchedUser(username: $username) {
				languageProblemCount {
					languageName
					problemsSolved
				}
			}
		}
	`)

	languagesStatsReq.Var("username", username)

	var respData languagesStatsResp

	ctx, cancel := context.WithTimeout(context.Background(), queryContextTime)
	defer cancel()

	for i := range maxRetries {
		err := s.client.Run(ctx, languagesStatsReq, &respData)
		if nil == err {
			break
		}
		if i == maxRetries-1 {
			return nil, ErrUserNotExists
		} else {
			time.Sleep(time.Second * time.Duration(i+1))
		}
	}

	languagesStats := []LanguageStats{}

	for _, lang := range respData.MatchedUser.LanguageProblemCount {
		languagesStats = append(languagesStats, LanguageStats{
			Name:           lang.LanguageName,
			ProblemsSolved: lang.ProblemsSolved,
		})
	}

	return languagesStats, nil
}

func (s *StatsService) GetByUsername(username string) (*Stats, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(2)

	var stats Stats
	errs := make(chan error, 2)

	go func() {
		defer wg.Done()
		languagesStats, err := s.getLanguagesStats(username)
		mu.Lock()
		stats.Languages = languagesStats
		errs <- err
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		problemsStats, err := s.getProblemsStats(username)
		mu.Lock()
		stats.Problems = problemsStats
		errs <- err
		mu.Unlock()
	}()

	wg.Wait()

	for i := 0; i < 2; i++ {
		if err := <-errs; err != nil {
			return nil, err
		}
	}

	return &stats, nil
}
