package domain

import (
	"fmt"
	"time"
)

type Tweet struct {
	User *User
	Text string
	Date *time.Time
	Id int
}

var idCount int = 0

func NewTweet(user *User, text string) *Tweet {
	var now = time.Now();
	idCount++
	return &Tweet{user,text, &now, idCount}
}


func (t *Tweet) PrintableTweet() string {
	return t.String()
}

func (t *Tweet) PrintableTweetWithTime() string {
	return fmt.Sprintf("(%02d:%02d:%02d) @%s: %s", t.Date.Hour(), t.Date.Minute(),t.Date.Second(), t.User.Nick, t.Text)
}

func (t *Tweet) String() string {
	return fmt.Sprintf("@%s: %s", t.User.Nick, t.Text)
}