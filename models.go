package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/tema-front/go-aggregator/internal/database"
)

type UsersRow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
}

type User struct {
	UsersRow
	ApiKey 		string `json:"api_key"`
}

func (u User) GetUsersRow() UsersRow {
	return u.UsersRow
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		UsersRow: UsersRow{
			ID: 			 dbUser.ID,
			CreatedAt: dbUser.CreatedAt,
			UpdatedAt: dbUser.UpdatedAt,
			Name: 		 dbUser.Name,
		},
		ApiKey: dbUser.ApiKey,
	}
}

func databaseUsersRowToUsersRow(dbUser database.GetUsersRow) UsersRow {
	return UsersRow{
		ID: 			 dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: 		 dbUser.Name,
	}
}

func databaseUsersToUsers(dbUsers []database.GetUsersRow) []UsersRow {
	users := make([]UsersRow, len(dbUsers))

	for i, dbUser := range dbUsers {
		users[i] = databaseUsersRowToUsersRow(dbUser)
	}

	return users
}