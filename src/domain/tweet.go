package domain

import (
	"fmt"
	"time"
)

type Tweet interface {
	PrintableTweet() string
	GetId() int
	GetUser() *User
	GetText() string
	GetDate() *time.Time
}

type TextTweet struct {
	User *User
	Text string
	Date *time.Time
	Id int
}

type ImageTweet struct {
	TextTweet
	Image string
}

type QuoteTweet struct {
	TextTweet
	twt Tweet
}

var idCount int = 0

func (t *TextTweet) PrintableTweet() string {
	return t.String()
}

func (t *ImageTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s %s", t.User.Nick, t.Text, t.Image)
}
func (t *QuoteTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s '%s'", t.User.Nick, t.Text, t.twt.PrintableTweet())
}

func (t *TextTweet) GetId() int {
	return t.Id
}
func (t *TextTweet) GetUser() *User {
	return t.User
}
func (t *TextTweet) GetText() string {
	return t.Text
}
func (t *TextTweet) GetDate() *time.Time {
	return t.Date
}

func (t *TextTweet) PrintableTweetWithTime() string {
	return fmt.Sprintf("(%02d:%02d:%02d) @%s: %s", t.Date.Hour(), t.Date.Minute(),t.Date.Second(), t.User.Nick, t.Text)
}

func (t *TextTweet) String() string {
	return fmt.Sprintf("@%s: %s", t.User.Nick, t.Text)
}

func NewTextTweet(user *User, text string) *TextTweet {
	var now = time.Now();
	idCount++
	return &TextTweet{user,text, &now, idCount}
}

func NewImageTweet(user *User, text string, image string) *ImageTweet {
	var now = time.Now();
	idCount++
	return &ImageTweet{TextTweet{user,text, &now, idCount},image}
}

func NewQuoteTweet(user *User, text string, t Tweet) Tweet {
	var now = time.Now();
	idCount++
	return &QuoteTweet{TextTweet{user,text, &now, idCount},t}
}