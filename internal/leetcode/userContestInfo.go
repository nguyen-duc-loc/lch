package leetcode

import (
	"context"
	"math"
	"sort"
	"time"

	"github.com/machinebox/graphql"
)

type UserContestInfoService struct {
	client *graphql.Client
}

type UserContestInfo struct {
	AttendedContestsCount uint64
	Rating                uint16
	GlobalRanking         uint64
	TopPercentage         float64
	ContestLevel          string
	AttendedContests      []Contest
}

type Contest struct {
	ProblemsSolved      uint8
	FinishTimeInSeconds uint16
	Rating              uint16
	Ranking             uint64
	Metadata            ContestMetadata
}

type ContestMetadata struct {
	StartTime     uint64
	Title         string
	TotalProblems uint8
}

func (s *UserContestInfoService) GetByUsername(username string) (*UserContestInfo, error) {
	type gqlResp struct {
		UserContestRanking struct {
			AttendedContestsCount uint64
			Rating                float64
			GlobalRanking         uint64
			TopPercentage         float64
			Badge                 struct {
				Name string
			}
		}
		UserContestRankingHistory []struct {
			Attended            bool
			ProblemsSolved      uint8
			TotalProblems       uint8
			FinishTimeInSeconds uint16
			Rating              float64
			Ranking             uint64
			Contest             struct {
				StartTime uint64
				Title     string
			}
		}
	}

	gqlReq := graphql.NewRequest(`
		query userContestRankingInfo($username: String!) {
			userContestRanking(username: $username) {
				attendedContestsCount
				rating
				globalRanking
				topPercentage
				badge {
					name
				}
			}
			userContestRankingHistory(username: $username) {
				attended
				problemsSolved
				totalProblems
				finishTimeInSeconds
				rating
				ranking
				contest {
					title
					startTime
				}
			}
		}
	`)

	gqlReq.Var("username", username)

	var respData gqlResp

	ctx, cancel := context.WithTimeout(context.Background(), queryContextTime)
	defer cancel()

	for i := range maxRetries {
		err := s.client.Run(ctx, gqlReq, &respData)
		if nil == err {
			break
		}
		if i == maxRetries-1 {
			return nil, ErrUserNotExists
		} else {
			time.Sleep(time.Second * time.Duration(i+1))
		}
	}

	userContestRanking := respData.UserContestRanking

	allContests := respData.UserContestRankingHistory
	attendedContests := []Contest{}
	for _, contest := range allContests {
		if contest.Attended {
			attendedContests = append(attendedContests, Contest{
				ProblemsSolved:      contest.ProblemsSolved,
				FinishTimeInSeconds: contest.FinishTimeInSeconds,
				Rating:              uint16(math.Round(contest.Rating)),
				Ranking:             contest.Ranking,
				Metadata: ContestMetadata{
					StartTime:     contest.Contest.StartTime,
					Title:         contest.Contest.Title,
					TotalProblems: contest.TotalProblems,
				},
			})
		}
	}
	sort.Slice(attendedContests, func(i, j int) bool {
		return attendedContests[i].Metadata.StartTime > attendedContests[j].Metadata.StartTime
	})

	return &UserContestInfo{
		AttendedContestsCount: userContestRanking.AttendedContestsCount,
		Rating:                uint16(math.Round(userContestRanking.Rating)),
		GlobalRanking:         userContestRanking.GlobalRanking,
		TopPercentage:         userContestRanking.TopPercentage,
		ContestLevel:          userContestRanking.Badge.Name,
		AttendedContests:      attendedContests,
	}, nil
}
