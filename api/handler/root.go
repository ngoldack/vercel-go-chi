package handler

import "github.com/jmoiron/sqlx"

type RootHandler struct {
	db *sqlx.DB
}

func NewRootHandler(db *sqlx.DB) *RootHandler {
	return &RootHandler{db: db}
}
