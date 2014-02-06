package mylib

import (
	"testing"
	"os"
	"net/url"
	"io/ioutil"
)


func TestDbAll(t *testing.T) {
	const dbpath = "sites.sqlite3"
	os.Remove(dbpath)

	// test InitDB
	db := InitDB(dbpath)
	defer db.Close()

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
	t.Log("create table in test db: OK!")

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
	if err2 != nil { panic(err2) }
	defer stmt.Close()

	const filepath = "Feeder_test.opml"
	siteList := GetOutlineList(filepath)
	for _, site := range siteList {
		_, err3 := stmt.Exec(site.XmlUrl, site.Title, site.Type,
				site.Text, site.HtmlUrl, site.Favicon)
		if err3 != nil { panic(err3) }
	}

	// test readSites
	sites := readSites(db)
	for _, site := range sites {
		t.Log(site)
	}

	// query HN in test db
	sql_queryHN := `SELECT XmlUrl FROM sites WHERE Title = ?`
	var hnUrl string
	err4 := db.QueryRow(sql_queryHN, "Hacker News").Scan(&hnUrl)
	if err4 != nil { panic(err4) }
	t.Log("query Hacker News site in test db: OK!", hnUrl)

	// read HN seed
	content, err5 := ioutil.ReadFile("hn_test.rss")
	if err5 != nil { panic(err5) }
	content2, err6 := ioutil.ReadFile("hn_test2.rss")
	if err6 != nil { panic(err6) }
	tmp := parseSeedContent(content)
	tmp2 := parseSeedContent(content2)
	// test updateOrInsertIfNotExist
	// FIXME: make url.QueryEscape(hnUrl) a function
	os.Remove(url.QueryEscape(hnUrl))
	dbHN := InitDB(url.QueryEscape(hnUrl))
	defer dbHN.Close()
	updateOrInsertIfNotExist(dbHN, tmp.Channel.ItemList)
	updateOrInsertIfNotExist(dbHN, tmp2.Channel.ItemList)

	// test ReadItems
	items := ReadItems(dbHN)
	for _, item := range items {
		t.Log(item)
	}
}
