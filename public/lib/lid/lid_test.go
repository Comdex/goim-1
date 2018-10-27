package lid

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestLid(t *testing.T) {
	db, err := sql.Open("mysql", "root:Liu123456@tcp(localhost:3306)/im?charset=utf8")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	lid, err := NewLid(db, "test")
	if err != nil {
		fmt.Println(err)
		return
	}
	i := 0
	for i < 100 {
		id, err := lid.Get()
		if err != nil {
			fmt.Println("fail")
			continue
		}
		fmt.Println(id)
		i++
	}
}

func TestLid_Get(t *testing.T) {
	go getLid("one")
	go getLid("two")
	go getLid("three")
	select {}
}

func getLid(index string) {
	db, err := sql.Open("mysql", "root:Liu123456@tcp(localhost:3306)/im?charset=utf8")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	lid, err := NewLid(db, "test")
	if err != nil {
		fmt.Println(err)
		return
	}
	i := 0
	for i < 100 {
		id, err := lid.Get()
		if err != nil {
			fmt.Println("fail")
			continue
		}
		fmt.Println(index, id)
		i++
	}
}

func BenchmarkLeafKey(b *testing.B) {
	db, err := sql.Open("mysql", "root:Liu123456@tcp(localhost:3306)/im?charset=utf8")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	leafKey, err := NewLid(db, "test")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < b.N; i++ {
		_, err := leafKey.Get()
		if err != nil {
			fmt.Println(err)
		}
	}
}
