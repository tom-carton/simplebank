package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"

var (
	testQueries *Queries
	testDB      *pgxpool.Pool
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = pgxpool.New(
		context.Background(),
		dbSource,
	)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
