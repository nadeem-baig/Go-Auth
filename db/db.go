package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/nadeem-baig/go-auth/utils/logger"
)


func DBConnect(cfg mysql.Config)  (*sql.DB,error){
	db,err := sql.Open("mysql",cfg.FormatDSN())
	if err!=nil {
		logger.Fatal(err)
        return nil,err
    }
	return db,nil
}