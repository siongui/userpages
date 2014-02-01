// @see http://golang.org/pkg/database/sql/
// @see https://github.com/mattn/go-sqlite3
// @see http://stackoverflow.com/questions/3634984/insert-if-not-exists-else-update
// @see http://stackoverflow.com/questions/2251699/sqlite-insert-or-replace-into-vs-update-where
// @see http://stackoverflow.com/questions/1601151/how-do-i-check-in-sqlite-whether-a-table-exists
package mylib

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)


func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil { panic(err) }
	if db == nil { panic("db nil") }
	return db
}

func ReadSites(db *sql.DB) []OpmlOutline {
	sql_readall := `
	SELECT XmlUrl, Title, Type, Text, HtmlUrl, Favicon FROM sites
	`
	rows, err := db.Query(sql_readall)
	if err != nil { log.Fatal(err) }
	defer rows.Close()

	var result []OpmlOutline
	for rows.Next() {
		site := OpmlOutline{}
		err := rows.Scan(&site.XmlUrl, &site.Title, &site.Type,
			&site.Text, &site.HtmlUrl, &site.Favicon)
		if err != nil { log.Fatal(err) }
		result = append(result, site)
	}
	if err := rows.Err(); err != nil { log.Fatal(err) }
	return result
}

func GetSites(dbpath string) []OpmlOutline {
	db := InitDB(dbpath)
	defer db.Close()
	return ReadSites(db)
}

func updateOrInsertIfNotExist(db *sql.DB, items []Item) {
	sql_table := `
	CREATE TABLE IF NOT EXISTS items(
		Link NOT NULL PRIMARY KEY,
		Title TEXT,
		Comments TEXT,
		Description TEXT,
		PubDate TEXT
	);
	`

	_, err := db.Exec(sql_table)
	if err != nil {
		log.Fatal(err)
		return
	}
}
