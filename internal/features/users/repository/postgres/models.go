package users_postgres_repository

import "github.com/Akimpupupuu/ToDoApp/internal/core/domain"

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func userDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))

	for i, val := range users {
		userDomains[i] = domain.NewUser(val.ID, val.Version, val.FullName, val.PhoneNumber)
	}

	return userDomains
}
