package ekkorm

import (
	"database/sql"
	"ekkorm/log"
	"ekkorm/session"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	e = &Engine{
		db: db,
	}
	log.Info("Connect Database Success")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Close Database Failed")
		return
	}
	log.Info("Close Database Success")
	return
}

func (engin *Engine) NewSession() *session.Session {
	return session.New(engin.db)
}
