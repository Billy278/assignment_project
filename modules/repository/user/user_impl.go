package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	models "github.com/Billy278/assignment_project/modules/models/user"
)

type UserRepoImpl struct {
	DB *sql.DB
}

func NewUserRepoImpl(db *sql.DB) UserRepo {
	return &UserRepoImpl{
		DB: db,
	}
}
func (repo *UserRepoImpl) CreatedUser(ctx context.Context, userIn models.User) (err error) {
	fmt.Println("Repo CreatedUSer")
	sqlCreate := "INSERT INTO users(name,dob,gmail,username,password,created_at,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7)"
	_, err = repo.DB.ExecContext(ctx, sqlCreate, userIn.Name, userIn.DoB, userIn.Gmail, userIn.Username, userIn.Password, userIn.Created_at, userIn.Updated_at)
	if err != nil {
		return
	}

	return
}
func (repo *UserRepoImpl) RepoFindUser(ctx context.Context, username string) (err error) {
	fmt.Println("RepoFindUser")
	sqlFind := "SELECT username,password FROM users WHERE username=$1"
	row, err := repo.DB.QueryContext(ctx, sqlFind, username)
	if err != nil {
		return
	}
	defer row.Close()
	if row.Next() {
		err = errors.New("NOT FOUND")
	}
	return
}
func (repo *UserRepoImpl) GetAllUserIsBirthday(ctx context.Context) (resUser []models.User, err error) {
	fmt.Println("repo GetAllUserBitrhday")
	sql := `SELECT id,name,gmail FROM users WHERE
		EXTRACT(
			MONTH
			FROM dob
		) = EXTRACT(
			MONTH
			FROM CURRENT_DATE
		)
		AND EXTRACT(
			DAY
			FROM dob
		) = EXTRACT(
			DAY
			FROM CURRENT_DATE
		)`
	rows, err := repo.DB.QueryContext(ctx, sql)
	if err != nil {
		return
	}
	defer rows.Close()

	user := models.User{}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Gmail)
		if err != nil {
			return
		}
		resUser = append(resUser, user)
	}
	fmt.Println("repo")
	fmt.Println(resUser)

	return
}
