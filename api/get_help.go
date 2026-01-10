package api

func GetHelp() {
	println(`
	go run main.go pertdb:create-table <migration-name> <table-name>
	go run main.go pertdb:alter-table <migration-name> <table-name>
	go run main.go pertdb:run
	go run main.go pertdb:rollback`)
}
