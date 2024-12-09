package leetcode

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/machinebox/graphql"
)

var (
	ErrProblemNotFound = errors.New("problem cannot be found")
)

type ProblemService struct {
	client *graphql.Client
}

type Problem struct {
	Title      string
	TitleSlug  string
	ID         int32
	PaidOnly   bool
	Difficulty string
	AcRate     float64
	Topics     []Topic
}

type Topic struct {
	Name string
	Slug string
	ID   string
}

func (s *ProblemService) GetByID(problemID int32) (*Problem, error) {
	type gqlResp struct {
		ProblemsetQuestionList struct {
			Questions []struct {
				Title              string
				TitleSlug          string
				FrontendQuestionId string
				PaidOnly           bool
				Difficulty         string
				AcRate             float64
				TopicTags          []Topic
			}
		}
	}

	type QuestionListFilterInput struct{}

	gqlReq := graphql.NewRequest(`
		query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {
			problemsetQuestionList: questionList(
				categorySlug: $categorySlug
				limit: $limit
				skip: $skip
				filters: $filters
			) {
				questions: data {
					acRate
					difficulty
					frontendQuestionId: questionFrontendId
					paidOnly: isPaidOnly
					title
					titleSlug
					topicTags {
						name
						slug
					}
				}
			}
		}
	`)

	skip := problemID - 1
	if skip >= 358 {
		skip--
	}

	gqlReq.Var("categorySlug", "")
	gqlReq.Var("skip", skip)
	gqlReq.Var("limit", 1)
	gqlReq.Var("filters", QuestionListFilterInput{})

	var respData gqlResp

	ctx, cancel := context.WithTimeout(context.Background(), queryContextTime)
	defer cancel()

	for i := range maxRetries {
		err := s.client.Run(ctx, gqlReq, &respData)
		if nil == err {
			break
		}
		if i == maxRetries-1 {
			return nil, err
		} else {
			time.Sleep(time.Second * time.Duration(i+1))
		}
	}

	problems := respData.ProblemsetQuestionList.Questions

	if len(problems) == 0 {
		return nil, ErrProblemNotFound
	}

	problem := problems[0]
	id, _ := strconv.ParseInt(problem.FrontendQuestionId, 10, 32)

	return &Problem{
		Title:      problem.Title,
		TitleSlug:  problem.TitleSlug,
		ID:         int32(id),
		PaidOnly:   problem.PaidOnly,
		Difficulty: problem.Difficulty,
		AcRate:     problem.AcRate,
		Topics:     problem.TopicTags,
	}, nil
}

func (s *ProblemService) GetToday() (*Problem, error) {
	type gqlResp struct {
		ActiveDailyCodingChallengeQuestion struct {
			Link     string
			Question struct {
				Title              string
				TitleSlug          string
				FrontendQuestionId string
				PaidOnly           bool
				Difficulty         string
				AcRate             float64
				TopicTags          []Topic
			}
		}
	}

	gqlReq := graphql.NewRequest(`
		query questionOfToday {
			activeDailyCodingChallengeQuestion {
				link
				question {
					acRate
					difficulty
					frontendQuestionId: questionFrontendId
					paidOnly: isPaidOnly
					title
					titleSlug
					topicTags {
						name
						slug
					}
				}
			}
		}
	`)

	var respData gqlResp

	ctx, cancel := context.WithTimeout(context.Background(), queryContextTime)
	defer cancel()

	for i := range maxRetries {
		err := s.client.Run(ctx, gqlReq, &respData)
		if nil == err {
			break
		}
		if i == maxRetries-1 {
			return nil, err
		} else {
			time.Sleep(time.Second * time.Duration(i+1))
		}
	}

	if len(respData.ActiveDailyCodingChallengeQuestion.Link) == 0 {
		return nil, ErrProblemNotFound
	}

	problem := respData.ActiveDailyCodingChallengeQuestion.Question
	id, _ := strconv.ParseInt(problem.FrontendQuestionId, 10, 32)

	return &Problem{
		Title:      problem.Title,
		TitleSlug:  problem.TitleSlug,
		ID:         int32(id),
		PaidOnly:   problem.PaidOnly,
		Difficulty: problem.Difficulty,
		AcRate:     problem.AcRate,
		Topics:     problem.TopicTags,
	}, nil
}
