package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/ssr0016/gobank/types"
)

// Step 2 Create Database

type Storage interface {
	CreateAccount(*types.Account) error
	DeleteAccount(int) error
	UpdateAccount(*types.Account) error
	GetAccounts() ([]*types.Account, error)
	GetAccountById(int) (*types.Account, error)
	GetAccountByNumber(int) (*types.Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connstr := "user=postgres dbname=postgres password=secret  sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil

}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

// Create DB migrations manually
func (s *PostgresStore) createAccountTable() error {
	query := `create table if not exists accounts(
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		encrypted_password varchar(100),
		balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(acc *types.Account) error {

	query := `INSERT INTO accounts
	(
		first_name,
		last_name,
		number,
		encrypted_password,
		balance,
		created_at
	)VALUES
	(
		$1, $2, $3, $4, $5, $6)`

	_, err := s.db.Query(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.EncryptedPassword,
		acc.Balance,
		acc.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) UpdateAccount(*types.Account) error {
	return nil
}
func (s *PostgresStore) DeleteAccount(id int) error {

	_, err := s.db.Query(`
		DELETE
		FROM
		accounts
		WHERE
		id = $1`, id)

	return err
}
func (s *PostgresStore) GetAccountById(id int) (*types.Account, error) {

	rows, err := s.db.Query(`
			SELECT
			 	* 
			FROM accounts
			WHERE 
			id = $1`, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgresStore) GetAccounts() ([]*types.Account, error) {

	rows, err := s.db.Query(`
			SELECT
				* 
			FROM
			accounts`)
	if err != nil {
		return nil, err
	}

	accounts := []*types.Account{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

// Step 5
func (s *PostgresStore) GetAccountByNumber(number int) (*types.Account, error) {
	rows, err := s.db.Query(`
			SELECT
				* 
			FROM
			accounts
			WHERE
			number = $1`, number)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account with number [%d] not found", number)
}

// end

func scanIntoAccount(rows *sql.Rows) (*types.Account, error) {

	account := new(types.Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.EncryptedPassword,
		&account.Balance,
		&account.CreatedAt,
	)

	return account, err
}
