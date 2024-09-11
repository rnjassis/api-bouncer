package main

import (
	"github.com/rnjassis/api-bouncer/sqllite"
)

func main() {
	db := sqllite.Init()
	defer db.Close()
}
