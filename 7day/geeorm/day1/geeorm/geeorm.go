package geeorm

import (
	"database/sql"
	"day1/log"
	"day1/session"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (e *Engine, err error) {
	// 初始化数据库连接
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Errorln(err.Error())
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Errorln(err.Error())
		return nil, err
	}

	log.Infoln("Connect database success")
	return &Engine{db: db}, nil
}

func (engine *Engine) Close() {

	if err := engine.db.Close(); err != nil {
		log.Errorln(err.Error())
	}
	log.Infoln("Close database success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.NewSession(engine.db)
}
