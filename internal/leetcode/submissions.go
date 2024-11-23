package leetcode

import (
	"context"
	"strconv"
	"time"

	"github.com/machinebox/graphql"
)

type SubmissionsService struct {
	client *graphql.Client
}

type Submission struct {
	ProblemTitle string
	AcTime       uint64
}

func (s *SubmissionsService) GetByUsername(username string) ([]*Submission, error) {
	type gqlResp struct {
		RecentAcSubmissionList []struct {
			Title     string
			Timestamp string
		}
	}

	gqlReq := graphql.NewRequest(`
		query recentAcSubmissions($username: String!, $limit: Int!) {
			recentAcSubmissionList(username: $username, limit: $limit) {
				title
				timestamp
			}
		}
	`)

	gqlReq.Var("username", username)
	gqlReq.Var("limit", 10)

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

	submissions := []*Submission{}
	recentAcSubmissionList := respData.RecentAcSubmissionList

	for _, s := range recentAcSubmissionList {
		timestampInSeconds, _ := strconv.ParseInt(s.Timestamp, 10, 64)

		submissions = append(submissions, &Submission{
			ProblemTitle: s.Title,
			AcTime:       uint64(timestampInSeconds),
		})
	}

	return submissions, nil
}
