package mysql_transaction_prac

type user struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

type address struct {
	ID      int    `db:"id"`
	UserID  string `db:"user_id"`
	Address string `db:"address"`
}

type MySQLConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	DB       string
}
