package leetcode

import (
	"context"
	"errors"
	"time"

	"github.com/machinebox/graphql"
)

var (
	ErrUserNotExists = errors.New("user not found")
)

type Profile struct {
	Username      string
	RealName      string
	Bio           string
	GlobalRanking uint64
	ContestLevel  string
	Country       string
	School        string
	Skills        []string
	SocialURL     SocialURl
}

type SocialURl struct {
	Github   string
	X        string
	Linkedin string
}

type ProfileService struct {
	client *graphql.Client
}

func (s *ProfileService) GetByUsername(username string) (*Profile, error) {
	type gqlResp struct {
		MatchedUser struct {
			ContestBadge struct {
				Name string
			}
			Username    string
			GithubUrl   string
			TwitterUrl  string
			LinkedinUrl string
			Profile     struct {
				Ranking     uint64
				RealName    string
				AboutMe     string
				CountryName string
				School      string
				SkillTags   []string
			}
		}
	}

	gqlReq := graphql.NewRequest(`
		query userPublicProfile($username: String!) {
			matchedUser(username: $username) {
				contestBadge {
					name
				}
				username
				githubUrl
				twitterUrl
				linkedinUrl
				profile {
					ranking
					realName
					aboutMe
					school
					countryName
					skillTags
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

	matchedUser := respData.MatchedUser

	if matchedUser.Username == "" {
		return nil, ErrUserNotExists
	}

	return &Profile{
		Username:      matchedUser.Username,
		RealName:      matchedUser.Profile.RealName,
		Bio:           matchedUser.Profile.AboutMe,
		GlobalRanking: matchedUser.Profile.Ranking,
		ContestLevel:  matchedUser.ContestBadge.Name,
		Country:       matchedUser.Profile.CountryName,
		School:        matchedUser.Profile.School,
		Skills:        matchedUser.Profile.SkillTags,
		SocialURL: SocialURl{
			Github:   matchedUser.GithubUrl,
			X:        matchedUser.TwitterUrl,
			Linkedin: matchedUser.LinkedinUrl,
		},
	}, nil
}
