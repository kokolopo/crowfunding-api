package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input InputRegistration) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input InputRegistration) (User, error) {
	var user User

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email
	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, errNew := s.repository.Save(user)
	if errNew != nil {
		return newUser, errNew
	}

	return newUser, nil
}
