package user

import (
	"demo/app/pkg/db"
)

type UserRepositoryDeps struct {
	DB *db.Db
}

type UserRepository struct {
	User     *User
	Database *db.Db
}

func (repo *UserRepository) CreateUser(user *User) (*User, error) {
	result := repo.Database.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) NewUserRepository(connection *db.Db) *UserRepository {
	return &UserRepository{
		Database: connection,
	}
}
