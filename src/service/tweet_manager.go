package service

import (
	"errors"
	"github.com/Tweeter/src/domain"
)

type TweetManager struct {
	tweets []domain.Tweet
	UserManager *UserManager
}

func (tm *TweetManager) PublishTweet(s domain.Tweet) (int, error) {
	if s.GetUser() == nil {
		return 0, errors.New("user is required");
	}

	if s.GetText() == "" {
		return 0, errors.New("text is required");
	}

	if len(s.GetText()) > 140 {
		return 0, errors.New("text exceeds 140 characters");
	}
	if !tm.UserManager.IsRegisteredUser(s.GetUser()) {
		return 0, errors.New("user is no registered");
	}

	tm.tweets = append(tm.tweets, s)

	return s.GetId(), nil
}

func (tm *TweetManager) GetTweet() domain.Tweet{
	return tm.tweets[0]
}

func NewTweetManager() *TweetManager {
	tweetManager := TweetManager{ tweets: make([]domain.Tweet, 0), UserManager: NewUserManager()}
	return &tweetManager
}

func (tm *TweetManager) GetTweets() []domain.Tweet {
	return tm.tweets
}

func (tm *TweetManager) GetTweetById(id int) domain.Tweet {
	var result domain.Tweet
	for _, value := range tm.tweets {
		if value.GetId() == id {
			result = value
			break
		}
	}

	return result
}

func (tm *TweetManager) CountTweetsByUser(user *domain.User) int {
	var result int = 0
	for _, value := range tm.tweets {
		if value.GetUser() == user {
			result++
		}
	}

	return result
}

func (tm *TweetManager) GetTweetsByUser(user *domain.User) []domain.Tweet {
	var result map[*domain.User][]domain.Tweet
	result = make(map[*domain.User][]domain.Tweet)
	for _, value := range tm.tweets {
		result[value.GetUser()] = append(result[value.GetUser()], value)
	}

	return result[user]
}