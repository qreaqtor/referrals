package userst

import "database/sql"

type UserStorage struct {
	table string
	db    *sql.DB
}

func NewUserStorage(db *sql.DB, tableName string) *UserStorage {
	return &UserStorage{
		table: tableName,
		db:    db,
	}
}
