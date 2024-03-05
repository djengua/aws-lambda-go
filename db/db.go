package db

import (
	"database/sql"
	"log"

	"github.com/djengua/aws-lambda-go/models"
)

type ContextDB struct {
	db *sql.DB
}

// Coneccion a la BD, se recibe configuracion, driver y connection string
func ConnectDB(config models.ConfigInfo) (*ContextDB, error) {
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot conect to db: ", err)
	}

	ctx := &ContextDB{
		db: conn,
	}

	return ctx, nil
}
