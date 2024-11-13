module example/data-access

go 1.23.2

require github.com/go-sql-driver/mysql v1.8.1

require filippo.io/edwards25519 v1.1.0 // indirect

require (
    gorm.io/driver/sqlite v1.1.5 // version might vary
    gorm.io/gorm v1.21.12 // version might vary
)