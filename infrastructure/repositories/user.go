package repositories

import "github.com/rennanbadaro/discount-calculator/infrastructure/storage"

type IUserRepository interface {
	FetchByID(id string) (*User, error)
}

type UserRepository struct {
	dbClient *storage.PostgresClient
}

type User struct {
	Id          string
	FirstName   string
	LastName    string
	DateOfBirth string
}

func (ur *UserRepository) FetchByID(id string) (*User, error) {
	query := "SELECT id, first_name, last_name, date_of_birth from users WHERE id=$1"
	row := ur.dbClient.Conn.QueryRow(query, id)

	u := &User{}
	if err := row.Scan(&u.Id, &u.FirstName, &u.LastName, &u.DateOfBirth); err != nil {
		return nil, err
	}

	return u, nil
}

func NewUserRepository(dbClient *storage.PostgresClient) *UserRepository {
	return &UserRepository{dbClient: dbClient}
}
