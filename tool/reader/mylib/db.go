// @see http://golang.org/pkg/database/sql/
// @see https://github.com/mattn/go-sqlite3
// @see http://stackoverflow.com/questions/3634984/insert-if-not-exists-else-update
// @see http://stackoverflow.com/questions/2251699/sqlite-insert-or-replace-into-vs-update-where
// @see http://stackoverflow.com/questions/1601151/how-do-i-check-in-sqlite-whether-a-table-exists
package mylib

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
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
	return ReadSites(db)
}

func updateOrInsertIfNotExist(db *sql.DB, items []Item) {
	sql_table := `
	CREATE TABLE IF NOT EXISTS items(
		Link NOT NULL PRIMARY KEY,
		Title TEXT,
		Comments TEXT,
		Description TEXT,
		PubDate TEXT,
		IsRead INTEGER
	);
	`

	_, err := db.Exec(sql_table)
	if err != nil { panic(err) }

	// insert items into db
	sql_additem := `
	INSERT OR REPLACE INTO items(
		Link,
		Title,
		Comments,
		Description,
		PubDate,
		IsRead
	) values(?, ?, ?, ?, ?, ?)
	`

	stmt, err2 := db.Prepare(sql_additem)
	if err2 != nil { panic(err2) }
	defer stmt.Close()

	for _, item := range items {
		_, err3 := stmt.Exec(item.Link, item.Title, item.Comments,
			string(item.Description), item.PubDate, 0)
		if err3 != nil { panic(err3) }
	}
}

func ReadItems(db *sql.DB) []Item {
	sql_readall := `
	SELECT Link, Title, Comments, Description, PubDate, IsRead FROM items
	`
	rows, err := db.Query(sql_readall)
	if err != nil { panic(err) }
	defer rows.Close()

	var result []Item
	for rows.Next() {
		item := Item{}
		var rawHtml string
		var isRead int
		err2 := rows.Scan(&item.Link, &item.Title, &item.Comments,
			&rawHtml, &item.PubDate, &isRead)
		item.Description = template.HTML(rawHtml)
		if err2 != nil { panic(err2) }
		result = append(result, item)
	}
	if err3 := rows.Err(); err3 != nil { panic(err3) }
	return result
}
