package repository

import (
	"context"
	"database/sql"
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
	fmt.Println(resUser)
	return
}
