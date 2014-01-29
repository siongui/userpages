package mylib

import (
	"testing"
	"os"
	"log"
)


func TestInitDB(t *testing.T) {
	const dbpath = "sites.sqlite3"
	os.Remove(dbpath)

	db := InitDB(dbpath)
	defer db.Close()

	sql_table := `
	CREATE TABLE sites(
		XmlUrl NOT NULL PRIMARY KEY,
		Title TEXT,
		Type TEXT,
		Text TEXT,
		HtmlUrl TEXT,
		Favicon TEXT
	);
	`

	_, err := db.Exec(sql_table)
	if err != nil {
		log.Fatal(err)
		return
	}

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
	if err != nil { log.Fatal(err) }
	defer stmt.Close()

	const filepath = "Feeder_test.opml"
	siteList := GetOutlineList(filepath)
	for _, site := range siteList {
		_, err := stmt.Exec(site.XmlUrl, site.Title, site.Type,
				site.Text, site.HtmlUrl, site.Favicon)
		if err != nil { log.Fatal(err) }
	}

	sites := ReadSites(db)
	for _, site := range sites {
		log.Println(site)
	}
}
