package leetcode

import (
	"fmt"
	"time"

	"github.com/machinebox/graphql"
	"github.com/nguyen-duc-loc/leetcode-helper/utils"
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
	Profiles interface {
		GetByUsername(string) (*Profile, error)
	}
	UserContestInfo interface {
		GetByUsername(string) (*UserContestInfo, error)
	}
	Stats interface {
		GetByUsername(string) (*Stats, error)
	}
	Submissions interface {
		GetByUsername(string) ([]*Submission, error)
	}
}

func New() *Leetcode {
	client := graphql.NewClient("https://leetcode.com/graphql")

	return &Leetcode{
		Problems:        &ProblemService{client},
		Solutions:       &SolutionService{client},
		Profiles:        &ProfileService{client},
		UserContestInfo: &UserContestInfoService{client},
		Stats:           &StatsService{client},
		Submissions:     &SubmissionsService{client},
	}
}

func FormatGlobalRanking(rank uint64) string {
	rankInText := fmt.Sprintf("%d", rank)

	switch {
	case rank <= 3000:
		return utils.RedText(rankInText)

	case rank <= 10000:
		return utils.OrangeText(rankInText)

	case rank <= 25000:
		return utils.BlueText(rankInText)

	case rank <= 50000:
		return utils.MagentaText(rankInText)

	default:
		return utils.WhiteText(rankInText)
	}
}

func FormatContestRating(rating uint16) string {
	ratingInText := fmt.Sprintf("%d", rating)

	switch {
	case rating < 1700:
		return utils.WhiteText(ratingInText)

	case rating < 1900:
		return utils.BlueText(ratingInText)

	case rating < 2100:
		return utils.MagentaText(ratingInText)

	case rating < 2400:
		return utils.OrangeText(ratingInText)

	default:
		return utils.RedText(ratingInText)
	}
}

func FormatContestLevel(level string) string {
	switch {
	case level == "":
		return ""

	case level == "Guardian":
		return utils.BlueText(level)

	case level == "Knight":
		return utils.GreenText(level)

	default:
		return utils.WhiteText(level)
	}
}
