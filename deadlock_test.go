package mysql_transaction_prac

import (
	"log"
	"testing"
)

func Test_tx1(t *testing.T) {
	store, err := NewMySQLStore(&MySQLConfig{
		Username: "txprac_user",
		Password: "txprac",
		Host:     "localhost",
		Port:     6036,
		DB:       "txprac",
	})
	if err != nil {
		log.Fatal(err)
	}

	err = store.initializeDB()

	createDeadlock(store)
}
