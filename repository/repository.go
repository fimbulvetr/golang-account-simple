package repository

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

func GetEngine(connStr string) (*xorm.Engine, error) {
	return xorm.NewEngine("mysql", connStr)
}
