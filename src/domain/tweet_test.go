package domain_test

import (
	"github.com/Tweeter/src/domain"
	"testing"
)

func TestCanGetAPrintableTweet(t *testing.T) {

	userName := "grupoesfera"
	user := domain.User{Name:userName,Mail:userName,Nick:userName,Pass:userName}
	tweet := domain.NewTweet(&user, "This is my tweet")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestCanGetAStringFromATweet(t *testing.T) {

	userName := "grupoesfera"
	user := domain.User{Name:userName,Mail:userName,Nick:userName,Pass:userName}
	tweet := domain.NewTweet(&user, "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}