package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var Pool *pgxpool.Pool

func Db() (*pgxpool.Pool, error) {
	fmt.Println("Connecting to DB...")

	cs := "host='%v' user='%v' password='%v' dbname='%v'"

	connectionString := fmt.Sprintf(cs,
		"localhost",
		"samuelmartel",
		"",
		"samuelmartel",
	)
	//pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	fmt.Printf("Using DB Connection String [%s] \n", connectionString)
	var err error
	Pool, err = pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}
	return Pool, nil
}
