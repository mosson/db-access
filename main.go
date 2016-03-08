package main

import (
	"db-access/db"

	"github.com/naoina/genmai"
)

func main() {
	_main()
}

func _main() {
	db.Dial(&genmai.SQLite3Dialect{}, ":memory")
	defer db.Close()
}
