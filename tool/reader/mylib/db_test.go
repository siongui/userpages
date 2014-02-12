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

	createSitesTable(db)
	t.Log("create sites table in test db: OK!")

	// insert sites data into test db
	const filepath = "Feeder_test.opml"
	siteList := GetOutlineList(filepath)
	storeSites(db, siteList)
	t.Log("store sites table in test db: OK!")

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
	// test storeItems
	// FIXME: make url.QueryEscape(hnUrl) a function
	os.Remove(url.QueryEscape(hnUrl))
	dbHN := InitDB(url.QueryEscape(hnUrl))
	defer dbHN.Close()
	storeItems(dbHN, tmp.Channel.ItemList)
	storeItems(dbHN, tmp2.Channel.ItemList)

	// test ReadItems
	items := ReadItems(dbHN)
	for index, item := range items {
		t.Log(index)
		t.Log(item)
	}
}
