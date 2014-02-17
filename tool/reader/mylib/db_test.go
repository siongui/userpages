package mylib

import (
	"testing"
	"os"
	"io/ioutil"
)


func TestDbAll(t *testing.T) {
	const dbpath = "sqlite3/sites.db"
	os.Remove(dbpath)

	// test InitDB
	db := InitDB(dbpath)
	defer db.Close()

	createSitesTable(db)
	t.Log("create sites table in test db: OK!")

	// insert sites data into test db
	const filepath = "sqlite3/Feeder_test.opml"
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
	content, err5 := ioutil.ReadFile("sqlite3/hn_test.rss")
	if err5 != nil { panic(err5) }
	content2, err6 := ioutil.ReadFile("sqlite3/hn_test2.rss")
	if err6 != nil { panic(err6) }
	tmp, _ := parseSeedContent(content)
	tmp2, _ := parseSeedContent(content2)
	// test storeItems
	os.Remove(XmlUrl2DBPath(hnUrl))
	dbHN := InitDB(XmlUrl2DBPath(hnUrl))
	defer dbHN.Close()
	storeItems(dbHN, tmp.ItemList)
	storeItems(dbHN, tmp2.ItemList)

	// test ReadItems
	items := ReadItems(dbHN)
	for index, item := range items {
		t.Log(index)
		t.Log(item)
	}
}
