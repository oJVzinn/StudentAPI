module studentapi/database

require (
	gorm.io/driver/sqlite v1.5.7
	gorm.io/gorm v1.25.12
	studentapi/person v0.0.0-00010101000000-000000000000
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	golang.org/x/text v0.14.0 // indirect
)

go 1.23.5

replace studentapi/person => ../person
