package main

import (
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
	"github.com/nadeem-baig/go-auth/cmd/api"
	"github.com/nadeem-baig/go-auth/db"
	"github.com/nadeem-baig/go-auth/utils/logger"
)



func main() {

    db,err := db.ConnectDB()
    if err!= nil {
        logger.Fatal("Failed to connect to database", err)
        return
    }
    api.StartServer(db)
}

