package dboperations

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/rushteam/gosql"
  "os"
)

type UserModel struct {
  ID int `db:"id"`
  Station string `db:"station"`
}
