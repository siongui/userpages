package mylib

import (
	"testing"
	"os"
	"log"
	"net/url"
)


func TestDbAll(t *testing.T) {
	const dbpath = "sites.sqlite3"
	os.Remove(dbpath)

	// test InitDB
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
	log.Println("create table in test db: OK!")

	// insert data into test db
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

	stmt, err2 := db.Prepare(sql_additem)
	if err2 != nil { log.Fatal(err2) }
	defer stmt.Close()

	const filepath = "Feeder_test.opml"
	siteList := GetOutlineList(filepath)
	for _, site := range siteList {
		_, err3 := stmt.Exec(site.XmlUrl, site.Title, site.Type,
				site.Text, site.HtmlUrl, site.Favicon)
		if err3 != nil { log.Fatal(err3) }
	}

	// test ReadSites
	sites := ReadSites(db)
	for _, site := range sites {
		log.Println(site)
	}

	// query HN in test db
	sql_queryHN := `SELECT XmlUrl FROM sites WHERE Title = ?`
	var hnUrl string
	err4 := db.QueryRow(sql_queryHN, "Hacker News").Scan(&hnUrl)
	if err4 != nil {
		log.Fatal(err4)
		return
	}
	log.Println("query Hacker News site in test db: OK!", hnUrl)

	// read HN seed
	tmp := GetSeed(hnUrl)
	// test updateOrInsertIfNotExist
	dbHN := InitDB(url.QueryEscape(hnUrl))
	defer dbHN.Close()
	updateOrInsertIfNotExist(dbHN, tmp.Channel.ItemList)

	// test ReadItems
	items := ReadItems(dbHN)
	for _, item := range items {
		log.Println(item)
	}
}
