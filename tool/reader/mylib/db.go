// @see http://golang.org/pkg/database/sql/
// @see https://github.com/mattn/go-sqlite3
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
