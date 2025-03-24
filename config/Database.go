package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Database() *gorm.DB {
	createDirStorageLogsDatabase()
	filePath := "./storage/logs/database/" + time.Now().Format("01-02-2006") + ".log"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0744)
	if err != nil {
		panic(fmt.Sprintf("error opening file database: %v", err))
	}

	logWriter := &reopenableWriter{
		filePath: filePath,
		file:     file,
	}
	// Connect to the database using GORM
	// migrate -path db/migrations -database "mysql://arifin:Arifin123\!@tcp(10.217.18.4:3306)/lennadb" down
	// migrate -path db/migrations -database "mysql://arifin:Arifin123\!@tcp(10.217.18.4:3306)/lennadb" down
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	newLogger := logger.New(
		log.New(logWriter, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, // Set the custom GORM logger
	})

	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
		return db
	}

	return db
}

func createDirStorageLogsDatabase() {
	dir := "./storage/logs/database"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0744)
		if err != nil {
			fmt.Println(dir, "can't created directory")
		}
		fmt.Println("success created directory", dir)
	} else {
		fmt.Println("The provided directory named", dir, "exists")
	}
}

// reopenableWriter is a custom writer that reopens the file if it's deleted
type reopenableWriter struct {
	filePath string
	file     *os.File
}

func (w *reopenableWriter) Write(p []byte) (n int, err error) {
	if _, err := os.Stat(w.filePath); os.IsNotExist(err) {
		// Reopen the file if it was deleted
		w.file, err = os.OpenFile(w.filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0744)
		if err != nil {
			return 0, fmt.Errorf("error reopening file: %v", err)
		}
	}
	return w.file.Write(p)
}
