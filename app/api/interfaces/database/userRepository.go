package database

import (
	"fmt"

	"github.com/Code0716/clean_architecture/app/api/domain"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id string, err error) {
	_, err = repo.Execute(
		"INSERT INTO users (ID,Name,Email,Password,CreatedDate,DeletedDate) VALUES (?,?,?,?,?,?)", u.ID, u.Name, u.Email, u.Password, u.CreatedDate, nil)

	if err != nil {
		return
	}

	id = u.ID

	return
}

func (repo *UserRepository) FindById(identifier string) (user domain.User, err error) {
	row, err := repo.Query("SELECT * FROM users WHERE ID = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	row.Next()
	if err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedDate); err != nil {
		return
	}

	return
}

func (repo *UserRepository) FindAll() (users domain.UserInfo, err error) {
	rows, err := repo.Query("SELECT * FROM users")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		user := new(domain.User)

		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedDate, user.DeletedDate); err != nil {
			fmt.Println(err)
		}
		users = append(users, *user)
	}
	return
}
