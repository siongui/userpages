package mylib

import "testing"

func TestAll(t *testing.T) {
	const dbpath = "foo.db"

	db := InitDB(dbpath)
	defer db.Close()
	CreateTable(db)

	items := []TestItem{
		TestItem{"1", "A", "213"},
		TestItem{"2", "B", "214"},
	}
	StoreItem(db, items)

	readItems := ReadItem(db)
	t.Log(readItems)

	items2 := []TestItem{
		TestItem{"1", "C", "215"},
		TestItem{"3", "D", "216"},
	}
	StoreItem(db, items2)

	readItems2 := ReadItem(db)
	t.Log(readItems2)
}
