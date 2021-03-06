package service_test

import (
	"github.com/Tweeter/src/domain"
	"github.com/Tweeter/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.TextTweet

	userName := "grupoesfera"
	text := "This is my first tweet"

	user := tweetManager.UserManager.NewUser(userName, userName,userName,userName)

	tweet = domain.NewTextTweet(user, text)

	// Operation
	id, _ := tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweet()

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.TextTweet

	var user *domain.User
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.TextTweet

	userName := "grupoesfera"
	var text string
	user := tweetManager.UserManager.NewUser(userName, userName,userName,userName)

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.TextTweet

	userName := "grupoesfera"
	text := `The Go project has grown considerably with over half a million users and community members 
	all over the world. To date all community oriented activities have been organized by the community
	with minimal involvement from the Go project. We greatly appreciate these efforts`
	user := tweetManager.UserManager.NewUser(userName, userName,userName,userName)

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text exceeds 140 characters" {
		t.Error("Expected error is text exceeds 140 characters")
	}
}
func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet *domain.TextTweet

	userName := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	user := tweetManager.UserManager.NewUser(userName, userName,userName,userName)

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)

	// Operation
	firstId, _ := tweetManager.PublishTweet(tweet)
	secondId, _ := tweetManager.PublishTweet(secondTweet)

	// Validation
	publishedTweets := tweetManager.GetTweets()

	if len(publishedTweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, user, secondText) {
		return
	}

}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.TextTweet
	var id int

	userName := "grupoesfera"
	text := "This is my first tweet"
	user := tweetManager.UserManager.NewUser(userName, userName,userName,userName)

	tweet = domain.NewTextTweet(user, text)

	// Operation
	id, _ = tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet *domain.TextTweet

	userName := "grupoesfera"
	anotherUserName := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	user := tweetManager.UserManager.NewUser(userName, userName,userName,userName)
	anotherUser := tweetManager.UserManager.NewUser(anotherUserName, anotherUserName,anotherUserName,anotherUserName)

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	count := tweetManager.CountTweetsByUser(user)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}

}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet *domain.TextTweet

	userName := "grupoesfera"
	anotherUserName := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	user := tweetManager.UserManager.NewUser(userName, userName,userName,userName)
	anotherUser := tweetManager.UserManager.NewUser(anotherUserName, anotherUserName,anotherUserName,anotherUserName)

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	firstId, _ := tweetManager.PublishTweet(tweet)
	secondId, _ := tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	tweets := tweetManager.GetTweetsByUser(user)

	// Validation
	if len(tweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(tweets))
		return
	}

	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, user, secondText) {
		return
	}

}

func isValidTweet(t *testing.T, tweet domain.Tweet, id int, user *domain.User, text string) bool {

	if tweet.GetId() != id {
		t.Errorf("Expected id is %v but was %v", id, tweet.GetId())
	}

	if tweet.GetUser() != user && tweet.GetText() != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, tweet.GetUser(), tweet.GetText())
		return false
	}

	if tweet.GetDate() == nil {
		t.Error("Expected date can't be nil")
		return false
	}

	return true

}

func TestOnlyRegisteredUsersCanTweet(t *testing.T) {
	tweetManager := service.NewTweetManager()

	var tweet *domain.TextTweet
	var tweet2 *domain.TextTweet

	userName := "user1"
	text := "This is my first tweet"
	user := tweetManager.UserManager.NewUser(userName, userName,userName,userName)

	userName2 := "user2"
	text2 := "This is my first tweet"
	user2 := domain.User{Name:userName2,Mail:userName2,Nick:userName2,Pass:userName2}

	tweet = domain.NewTextTweet(user, text)
	tweet2 = domain.NewTextTweet(&user2, text2)

	_, err := tweetManager.PublishTweet(tweet)
	if err != nil {
		t.Error("Unexpected Error")
	}

	_, err = tweetManager.PublishTweet(tweet2)
	if err == nil || err.Error() != "user is no registered" {
		t.Error("Expected error text: 'user is no registered'")
	}
}