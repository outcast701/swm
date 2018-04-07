package controller

import (
  "database/sql"

  _ "github.com/lib/pq"
  "github.com/outcast701/swm/pkg/model"
)

type PController struct {
  db *sql.DB
}

func (pController *PController) CreateTable() {
  sqlCreateTable := model.GetQuery("cP")
  pController.db.Exec(sqlCreateTable)
}
