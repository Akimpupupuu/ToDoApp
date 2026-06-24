package users_transport_http

import "github.com/Akimpupupuu/ToDoApp/internal/core/domain"

type UserDTOResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}

func usersDTOFromDomain(users []domain.User) []UserDTOResponse {
	usersDto := make([]UserDTOResponse, len(users))

	for i, user := range users {
		usersDto[i] = userDTOFromDomain(user)
	}

	return usersDto
}
