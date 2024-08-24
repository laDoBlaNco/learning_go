package repository

import (
	"database/sql"
	"errors"
	"time"
)

// a pool of connections to a sqlite dbase
type SQLiteRepository struct {
	Conn *sql.DB
}

// NewSQLiteRepository receives a connection to a sqlite dBase and returns a var instance of our
// SQLiteRepository struct with the necessary fields filled.
func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{ // simply returns a var instance of our db connection
		Conn: db,
	}
}

// This is the first 'method' function associated to our DBase type that fills our interface reqs
func (repo *SQLiteRepository) Migrate() error {
	query := `
	create table if not exists holdings(
		id integer primary key autoincrement,
		amount real not null,
		purchase_date integer not null,
		purchase_price integer not null);
	`
	// this kicks off no tables are found in our dbase and uses the sql instruction above to
	// create the table. we'll run this at the beginning of every start
	_, err := repo.Conn.Exec(query) // using our repo we exec the above query and return err if there is one
	return err                      // as estabilished returning an error if there is one
}

// this is the 2nd method to fulfill our  interface reqs. We take a holdings object and return
// a pointer to Holdings or an error
func (repo *SQLiteRepository) InsertHolding(holdings Holdings) (*Holdings, error) {
	stmt := "insert into holdings (amount,purchase_date,purchase_price) values (?,?,?)"

	// again we have our sql instruction above and we exec it below. NOTE how we use Unix() to convert
	// the date to a unix integer style time
	res, err := repo.Conn.Exec(stmt, holdings.Amount, holdings.PurchaseDate.Unix(), holdings.PurchasePrice)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	holdings.ID = id
	return &holdings, nil // as needed we are returning a pointer (ref) to holdings

}

// Here's the 3rd method of our interface which doesn't take anything and returns a slice of
// Holding types and possibly an error
func (repo *SQLiteRepository) AllHoldings() ([]Holdings, error) {
	query := "select id, amount, purchase_date, purchase_price from holdings order by purchase_date"
	rows, err := repo.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// again here we create our query, but not sure why we use .Query instead of .Exec here.
	// but we check for the error

	// we start by creating the var for our slice of holdings and then go through all of the rows
	// that were returned in our query, setting to vars (h, unixTim) pulling the info needed with
	// our holding struct fields as our model, also calculating the purchase_date
	var all []Holdings
	for rows.Next() {
		var h Holdings
		var unixTime int64
		err := rows.Scan(
			&h.ID,
			&h.Amount,
			&unixTime,
			&h.PurchasePrice,
		)
		if err != nil {
			return nil, err
		}
		h.PurchaseDate = time.Unix(unixTime, 0)
		all = append(all, h) // we append everything to our slice of holdings on each iteration
	}

	return all, nil // and as needed we return the slice of holding info filling our struct.

}

func (repo *SQLiteRepository) GetHoldingByID(id int) (*Holdings, error) {
	row := repo.Conn.QueryRow("select id,amount,purchase_date,purchase_price from holdings where id=?", id)

	var h Holdings
	var unixTime int64
	err := row.Scan(
		&h.ID,
		&h.Amount,
		&unixTime,
		&h.PurchasePrice,
	)
	if err != nil {
		return nil, err
	}

	h.PurchaseDate = time.Unix(unixTime, 0)

	return &h, nil
}

func (repo *SQLiteRepository) UpdateHolding(id int64, updated Holdings) error {
	if id == 0 {
		return errors.New("invalid updated id")
	}

	stmt := "update holdings set amount = ?, purchase_date = ?, purchase_price = ? where id = ?"
	res, err := repo.Conn.Exec(stmt, updated.Amount, updated.PurchaseDate.Unix(), updated.PurchasePrice, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errUpdateFailed
	}

	return nil

}

func (repo *SQLiteRepository) DeleteHolding(id int64) error {
	res, err := repo.Conn.Exec("delete from holdings where id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errDeleteFailed
	}

	return nil
}
