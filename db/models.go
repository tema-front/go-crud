package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/tema-front/go-crud/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	apiKey 		string
}

func DatabaseUserToUser(dbUser database.User) User {
	return User{
		ID: 			 dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: 		 dbUser.Name,
		apiKey: 	 dbUser.ApiKey,
	}
}

func DatabaseUsersToUsers(dbUsers []database.User) []User {
	users := make([]User, len(dbUsers))

	for i, dbUser := range dbUsers {
		users[i] = DatabaseUserToUser(dbUser)
	}

	return users
}