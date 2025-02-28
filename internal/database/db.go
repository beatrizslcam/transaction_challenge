package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}


func DBConnect(dsn string) (*DB, error){
	dsn = os.Getenv("DATABASE_URL")
	
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	ctx, cancel:= context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err:= pgxpool.New(ctx, dsn)
	if err!= nil{
		log.Printf("Unable to connect to database due to: %v\n", err)
		return nil, err
	}

	if err:= pool.Ping(ctx); err!= nil{
		log.Printf("Unable to ping database due to: %v\n", err)
		return nil, err
	}
	log.Println("Connected to database!")
	return &DB{Pool: pool}, nil

}

func (db *DB) Close(){
	db.Pool.Close()
	log.Println("Closed database connection")
}