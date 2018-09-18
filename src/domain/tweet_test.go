package domain_test

import (
	"github.com/Tweeter/src/domain"
	"testing"
)

func TestTextTweetPrintsUserAndText(t *testing.T) {

	// Initialization
	userName := "grupoesfera"
	user := domain.User{Name:userName, Mail:userName,Nick:userName,Pass:userName}
	tweet := domain.NewTextTweet(&user, "This is my tweet")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

	// Initialization
	userName := "grupoesfera"
	user := domain.User{Name:userName, Mail:userName,Nick:userName,Pass:userName}
	tweet := domain.NewImageTweet(&user, "This is my image", "http://www.grupoesfera.com.ar/common/img/grupoesfera.png")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my image http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {

	// Initialization
	userName := "grupoesfera"
	user := domain.User{Name:userName, Mail:userName,Nick:userName,Pass:userName}
	quotedTweet := domain.NewTextTweet(&user, "This is my tweet")
	userName2 := "nick"
	user2 := domain.User{Name:userName2, Mail:userName2,Nick:userName2,Pass:userName2}
	tweet := domain.NewQuoteTweet(&user2, "Awesome", quotedTweet)

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := `@nick: Awesome '@grupoesfera: This is my tweet'`
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestCanGetAStringFromATweet(t *testing.T) {

	// Initialization
	userName := "grupoesfera"
	user := domain.User{Name:userName, Mail:userName,Nick:userName,Pass:userName}
	tweet := domain.NewTextTweet(&user, "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}