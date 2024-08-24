package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/ladoblanco/learning_gui_apps_fyne/trevor_sawler_fyne/gold/repository"
	
	_ "github.com/glebarez/go-sqlite" 
)

type Config struct {
	App                 fyne.App
	InfoLog             *log.Logger
	ErrorLog            *log.Logger
	DB                  repository.Repository
	MainWindow          fyne.Window
	PriceContainer      *fyne.Container
	ToolBar             *widget.Toolbar
	PriceChartContainer *fyne.Container
	HTTPClient          *http.Client
}

func main() {
	var myApp Config // to store our configuration
	// create a fyne application
	fyneApp := app.NewWithID("test.ladoblanco.ID.unique.preferences")
	myApp.App = fyneApp
	myApp.HTTPClient = &http.Client{}

	// create our loggers
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// open a connection to the database - sqlite to embed the dbase in the application
	sqlDB, err := myApp.connecSQL()
	if err != nil {
		log.Panic(err)
	}

	// create a database respository (simply pattern for interacting with things)
	myApp.setupDB(sqlDB)

	// create and size a fyne window
	myApp.MainWindow = fyneApp.NewWindow("GoldWatcher")
	myApp.MainWindow.Resize(fyne.NewSize(770, 410))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()

	myApp.makeUI()

	// show and run the application
	myApp.MainWindow.ShowAndRun()
}

func (app *Config) connecSQL() (*sql.DB, error) {
	var path string

	if os.Getenv("DB_PATH") != "" {
		path = os.Getenv("DB_PATH")
	} else {
		path = app.App.Storage().RootURI().Path() + "/sql.db"
		app.InfoLog.Println("db in:", path)
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return db, nil

}

func (app *Config) setupDB(sqlDB *sql.DB) {
	app.DB = repository.NewSQLiteRepository(sqlDB)

	err := app.DB.Migrate()
	if err != nil {
		app.ErrorLog.Println(err)
		log.Panic()
	}
}
