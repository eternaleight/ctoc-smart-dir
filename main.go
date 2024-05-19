package main

func main() {
	dsn := loadConfig()
	db := initializeDatabase(dsn)
	r := setupRouter(db)

	r.Run(":8001")
}
