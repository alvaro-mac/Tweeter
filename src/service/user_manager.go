package service

import "github.com/Tweeter/src/domain"

type UserManager struct {
	Users []*domain.User
}

func NewUserManager() *UserManager {
	userManager := UserManager{Users:make([]*domain.User,0)}
	return &userManager
}

func (um *UserManager) NewUser(name string, mail string, nick string, pass string) *domain.User {
	newUser := domain.User{Name:name,Mail:mail,Nick:nick,Pass:pass}
	um.Users = append(um.Users, &newUser)
	return &newUser
}

func (um *UserManager) IsRegisteredUser(user *domain.User) bool {
	return contains(um.Users, user)
}

func contains(s []*domain.User, e *domain.User) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}