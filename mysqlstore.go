package mysql_transaction_prac

import (
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLStore struct {
	db     *sqlx.DB
	config *MySQLConfig
}

func (m *MySQLStore) initializeDB() error {
	stmt := `
create table if not exists users
(
    id    int auto_increment
        primary key,
    name  varchar(128) null,
    email varchar(128) null
);
`
	_, err := m.db.Exec(stmt)
	if err != nil {
		return err
	}

	stmt = `
create table if not exists addresses
(
    id      int auto_increment
        primary key,
    user_id int        not null,
    address mediumtext null
);

`
	_, err = m.db.Exec(stmt)
	if err != nil {
		return err
	}

	stmt = `
INSERT INTO users (id, name, email) VALUES (1, 'hossein', 'hossein@gmail.com'), (2, 'hossein2', 'hossein2@gmail.com');
`
	_, err = m.db.Exec(stmt)
	if err != nil {
		return err
	}

	stmt = `
INSERT INTO addresses (id, user_id, address) VALUES (1, 1, 'hossein address'), (2, 2, 'hossein2 address');
`
	_, err = m.db.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}

func NewMySQLStore(config *MySQLConfig) (*MySQLStore, error) {
	dsn := config.Username + ":" + config.Password + "@" + "(" + config.Host + ":" + strconv.Itoa(config.Port) + ")/" + config.DB + "?parseTime=true"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {

		return nil, err
	}

	return &MySQLStore{
		db:     db,
		config: config,
	}, nil
}
