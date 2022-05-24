package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

func QueryNameById(db *sql.DB, id string) (string, error) {
	var name string
	err := db.QueryRow("select name from users where id = ?", id).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			errors.Wrap(err, "empty data")
		} else {
			errors.Wrap(err, "query failed")
		}
	}
	return name, nil
}
