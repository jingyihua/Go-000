package dao

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func GetUserAge(id int) (int, error) {
	age, err := excSqlAge(id)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("find id: %d", id))
	}

	return age, nil
}

//执行数据库语句
func excSqlAge(id int) (int, error) {
	if id == 0 {
		return 18, nil
	}
	return 0, sql.ErrNoRows
}
