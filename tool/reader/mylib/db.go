/*
RSS:
http://stackoverflow.com/questions/15245896/rss-update-single-item
http://stackoverflow.com/questions/164124/rss-item-updates
SQLite:
http://golang.org/pkg/database/sql/
https://github.com/mattn/go-sqlite3
http://stackoverflow.com/questions/3634984/insert-if-not-exists-else-update
http://stackoverflow.com/questions/2251699/sqlite-insert-or-replace-into-vs-update-where
http://stackoverflow.com/questions/1601151/how-do-i-check-in-sqlite-whether-a-table-exists
http://stackoverflow.com/questions/19337029/insert-if-not-exists-statement-in-sqlite
http://stackoverflow.com/questions/6740733/insert-or-replace-is-creating-duplicates
http://stackoverflow.com/questions/12105198/sqlite-how-to-get-insert-or-ignore-to-work
http://stackoverflow.com/questions/19134274/sqlitedatabase-insert-or-replace-if-changed
http://stackoverflow.com/questions/14091183/sqlite-order-by-date
http://stackoverflow.com/questions/15473325/inserting-current-date-and-time-in-sqlite-database
http://stackoverflow.com/questions/7316747/sqlite-inserting-bool-value
*/
package mylib

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/url"
	"fmt"
)


func XmlUrl2DBPath(xmlUrl string) string {
	return fmt.Sprint("sqlite3/", url.QueryEscape(xmlUrl), ".db")
}

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil { panic(err) }
	if db == nil { panic("db nil") }
	return db
}

func createSitesTable(db *sql.DB) {
	sql_table := `
	CREATE TABLE sites(
		XmlUrl TEXT NOT NULL PRIMARY KEY,
		Title TEXT,
		Type TEXT,
		Text TEXT,
		HtmlUrl TEXT,
		Favicon TEXT
	);
	`

	_, err := db.Exec(sql_table)
	if err != nil { panic(err) }
}

func storeSites(db *sql.DB, sites []OpmlOutline) {
	sql_additem := `
	INSERT INTO sites(
		XmlUrl,
		Title,
		Type,
		Text,
		HtmlUrl,
		Favicon
	) values(?, ?, ?, ?, ?, ?)
	`

	stmt, err := db.Prepare(sql_additem)
	if err != nil { panic(err) }
	defer stmt.Close()

	for _, site := range sites {
		_, err2 := stmt.Exec(site.XmlUrl, site.Title, site.Type,
				site.Text, site.HtmlUrl, site.Favicon)
		if err2 != nil { panic(err2) }
	}
}

func readSites(db *sql.DB) []OpmlOutline {
	sql_readall := `
	SELECT XmlUrl, Title, Type, Text, HtmlUrl, Favicon FROM sites
	`
	rows, err := db.Query(sql_readall)
	if err != nil { panic(err) }
	defer rows.Close()

	var result []OpmlOutline
	for rows.Next() {
		site := OpmlOutline{}
		err2 := rows.Scan(&site.XmlUrl, &site.Title, &site.Type,
			&site.Text, &site.HtmlUrl, &site.Favicon)
		if err2 != nil { panic(err2) }
		result = append(result, site)
	}
	if err3 := rows.Err(); err3 != nil { panic(err3) }
	return result
}

func GetSites(dbpath string) []OpmlOutline {
	db := InitDB(dbpath)
	defer db.Close()
	return readSites(db)
}

func storeItems(db *sql.DB, items []Item) {
	// FIXME: add IsRead & InsertDatetime
	sql_table := `
	CREATE TABLE IF NOT EXISTS items(
		Link TEXT NOT NULL PRIMARY KEY,
		Title TEXT,
		Description TEXT,
		PubDate TEXT,
		Comments TEXT,
		IsRead INTEGER,
		InsertedDatetime DATETIME
	);
	`

	_, err := db.Exec(sql_table)
	if err != nil { panic(err) }

	// insert items into db
	sql_qryItem := `SELECT PubDate FROM items WHERE Link=?`
	sql_additem := `
	INSERT OR REPLACE INTO items(
		Link,
		Title,
		Description,
		PubDate,
		Comments,
		IsRead,
		InsertedDatetime
	) values(?, ?, ?, ?, ?, 0, CURRENT_TIMESTAMP)
	`

	stmt, err2 := db.Prepare(sql_additem)
	if err2 != nil { panic(err2) }
	defer stmt.Close()

	for _, item := range items {
		// query if the item exists
		var pubDate string
		err3 := db.QueryRow(sql_qryItem, item.Link).Scan(&pubDate)
		switch {
		case err3 == sql.ErrNoRows:
			// add item to table in db
			_, err4 := stmt.Exec(item.Link, item.Title,
				string(item.Description), item.PubDate,
				item.Comments)
			if err4 != nil { panic(err4) }
		case err3 != nil:
			panic(err3)
		default:
			// TODO: check if the same pubDate
			if pubDate != item.PubDate {
				_, err4 := stmt.Exec(
					item.Link, item.Title,
					string(item.Description),
					item.PubDate, item.Comments)
				if err4 != nil { panic(err4) }
			}
		}
	}
}

func ReadItems(db *sql.DB) []Item {
	sql_readall := `
	SELECT Link, Title, Description, PubDate, Comments FROM items
	ORDER BY datetime(InsertedDatetime) DESC
	`
	rows, err := db.Query(sql_readall)
	if err != nil { panic(err) }
	defer rows.Close()

	var result []Item
	for rows.Next() {
		item := Item{}
		var rawHtml string
		err2 := rows.Scan(&item.Link, &item.Title,
			&rawHtml, &item.PubDate, &item.Comments)
		item.Description = template.HTML(rawHtml)
		if err2 != nil { panic(err2) }
		result = append(result, item)
	}
	if err3 := rows.Err(); err3 != nil { panic(err3) }
	return result
}
