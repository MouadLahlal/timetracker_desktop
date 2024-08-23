package main

import (
	"context"
	"fmt"
	"os"
	"database/sql"
	
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/joho/godotenv/autoload"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

type Record struct {
	Program string
	Day string
	Time float64
}

func (a *App) GetTodayUsage() []Record {
	db, err := sql.Open("sqlite3", os.Getenv("DB_PATH"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query("select * from programs")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var records []Record

	for rows.Next() {
		var program string
		var day string
		var time float64

		err = rows.Scan(&program, &day, &time)
		if err != nil {
			panic(err)
		}

		records = append(records, Record{
			Program: program,
			Day: day,
			Time: time,
		})
	}

	return records
}
