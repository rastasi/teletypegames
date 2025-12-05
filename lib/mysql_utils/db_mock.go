package mysql_utils

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBMockSuite struct {
	DB   *gorm.DB
	Mock sqlmock.Sqlmock
}

func NewDBMockSuite() *DBMockSuite {
	var (
		db     *sql.DB
		gormdb *gorm.DB
		mock   sqlmock.Sqlmock
		err    error
	)

	db, mock, err = sqlmock.New()

	if err != nil {
		panic("SQL Mock error")
	}

	mock.ExpectQuery("SELECT VERSION()").WithArgs().WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	gormdb, err = gorm.Open(mysql.New(mysql.Config{
		DSN:        "sqlmock_db_0",
		DriverName: "mysql",
		Conn:       db,
	}), &gorm.Config{})

	if err != nil {
		panic("GORM open error")
	}

	// gormdb.AutoMigrate()

	return &DBMockSuite{
		DB:   gormdb,
		Mock: mock,
	}
}
