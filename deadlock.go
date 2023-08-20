package mysql_transaction_prac

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// createDeadlock
// Suppose tx1 locks row1 of users
// Then tx2 locks row2 of addresses
// Then tx1 updates row2 of addresses
// Then tx2 updates row1 of users
// The tx1 needs row2 of addresses lock to be released which is held by tx2.
// The tx2 needs row1 of users lock to be released which is held by tx1.
// Boom, the tx1 and tx2 can not proceed. and deadlock happens for tx1.
func createDeadlock(store *MySQLStore) {
	var wg sync.WaitGroup

	wg.Add(1)
	go doTx1(store, &wg)
	wg.Add(1)
	go doTx2(store, &wg)

	wg.Wait()
}

func doTx1(store *MySQLStore, wg *sync.WaitGroup) {
	defer wg.Done()

	tx1, err := store.db.Begin()
	r := tx1.QueryRow("select * from users where id=1 for update")

	var user user
	err = r.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("row 1 of users locked by tx1")
	time.Sleep(2 * time.Second)

	if err != nil {
		log.Fatal(err)
	}
	_, err = tx1.Exec("update addresses set address='hello' where id=2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("row2 of addresses updated by tx1")

	err = tx1.Commit()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("row1 done")
}

func doTx2(store *MySQLStore, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	tx1, err := store.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	var user user
	r := tx1.QueryRow("select * from addresses where id=2 for update")
	err = r.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("row 2 of addresses locked by tx2")

	_, err = tx1.Exec("update users set name='hossein' where id=1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("row 1 of users updated by tx2")
	err = tx1.Commit()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("row2 done")
}
