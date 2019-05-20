package dao

import (
	"database/sql"
	"fmt"
	"github.com/Sirupsen/logrus"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sync"
)


var(
	db *sql.DB
	mu = sync.Mutex{}
)


func init() {
	db, _ = sql.Open("sqlite3", "./transData.db")

	sqlTable := `
    CREATE TABLE IF NOT EXISTS btc(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        hash VARCHAR(80),
        height INTEGER NOT NULL,
        created DATE,
        address VARCHAR(32)NOT NULL,
        amount DECIMAL(32, 20) DEFAULT 0,
        symbol amount DECIMAL(32, 20) DEFAULT 0
	);
	CREATE TABLE IF NOT EXISTS bitcny(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        hash VARCHAR(80),
        height INTEGER NOT NULL,
        created DATE,
        address VARCHAR(32)NOT NULL,
        amount DECIMAL(32, 20) DEFAULT 0,
        symbol amount DECIMAL(32, 20) DEFAULT 0
	);
	CREATE TABLE IF NOT EXISTS eth(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        hash VARCHAR(80),
        height INTEGER NOT NULL,
        created DATE,
        address VARCHAR(32)NOT NULL,
        amount DECIMAL(32, 20) DEFAULT 0,
        symbol amount DECIMAL(32, 20) DEFAULT 0
	);
	CREATE TABLE IF NOT EXISTS eos(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        hash VARCHAR(80),
        height INTEGER NOT NULL,
        created DATE,
        address VARCHAR(32)NOT NULL,
        amount DECIMAL(32, 20) DEFAULT 0,
        symbol amount DECIMAL(32, 20) DEFAULT 0
	);
	CREATE TABLE IF NOT EXISTS trx(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        hash VARCHAR(80),
        height INTEGER NOT NULL,
        created DATE,
        address VARCHAR(32)NOT NULL,
        amount DECIMAL(32, 20) DEFAULT 0,
        symbol amount DECIMAL(32, 20) DEFAULT 0
	);
	`

	db.Exec(sqlTable)
}
//=================insert=====================//
func InsertBtc(tx map[string]interface{}) bool {
	mu.Lock()
	defer mu.Unlock()
	if !rowExists("SELECT height from btc where height=?", tx["block_height"].(int64)) { //根据txid判断是否存在这个txid，如果存在，就不插入到数据库中

		sqlStr := "INSERT INTO btc(hash, height,created, address, amount,symbol) values(?,?,?,?,?,?)"

		return execute(sqlStr, tx["hash"].(string),tx["block_height"].(uint64), tx["block_time"].(string), tx["address"].(string), tx["amount"].(float64),
			tx["symbol"].(uint64))
	}
	return false
}


//====================tool==================//
func rowExists(query string, args ...interface{}) bool {
	// mu.Lock()
	// defer mu.Unlock()
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	row := db.QueryRow(query, args...)
	err := row.Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("error checking if row exists '%s' %v", args, err)
	}
	return exists
}
func execute(sqlStr string, args ...interface{}) bool {
	// mu.Lock()
	// defer mu.Unlock()
	logrus.Info("execute sql", sqlStr, args)
	stmt, err := db.Prepare(sqlStr)
	checkErr(err)

	res, err := stmt.Exec(args...)
	// fmt.Println(res, err)
	checkErr(err)

	id, err := res.RowsAffected()
	checkErr(err)

	return id >= 1
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
