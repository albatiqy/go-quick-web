package main

import (
	"database/sql"
	"embed"
	"flag"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"

	// "time"

	"github.com/buger/jsonparser"
	_ "github.com/go-ozzo/ozzo-validation/v4"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/shopspring/decimal"
	_ "github.com/sony/sonyflake"
	_ "github.com/xuri/excelize/v2"
	_ "golang.org/x/sync/errgroup"

	"github.com/albatiqy/go-quick-web/handler"
)

type config struct {
	servicePort string
	fsRoot      string
}

//go:embed _embed
var embedFs embed.FS

func main() {
	var configFile string
	flag.StringVar(&configFile, "cfg", "", "configuration file")
	flag.Parse()
	cfg := getConfig(configFile)

	db := getDB(filepath.Join(cfg.fsRoot, "db/data.sqlite"), true)

	// contextTimeout := time.Duration(2) * time.Second
	// sonyflake := sonyflake.NewSonyflake(sonyflake.Settings{})

	fsUi, err := fs.Sub(embedFs, "_embed/ui")
	if err != nil {
		log.Fatal(err)
	}

	handlerParams := &handler.Params{
		DB: db,
	}
	router := handler.Router(handlerParams)
	// router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	// })
	router.NotFound = http.FileServer(http.FS(fsUi))

	log.Fatal(http.ListenAndServe(cfg.servicePort, router))
}

func getConfig(configFile string) config {
	workingDir, err := os.Getwd()
	if configFile == "" {
		if err == nil {
			configFile = filepath.Join(workingDir, "config.json")
		}
	}

	// default config
	cfg := config{
		servicePort: ":8085",
	}

	if _, err := os.Stat(configFile); err == nil {
		if json, err := os.ReadFile(configFile); err == nil {
			if str, err := jsonparser.GetString(json, "fsRoot"); err == nil {
				cfg.fsRoot = str
			}
			if str, err := jsonparser.GetString(json, "servicePort"); err == nil {
				cfg.servicePort = str
			}

			if cfg.fsRoot == "" {
				cfg.fsRoot = filepath.Join(workingDir, "appfs")
			}
			fInfo, err := os.Stat(cfg.fsRoot)
			if err != nil {
				log.Fatal(err)
			} else {
				if !fInfo.IsDir() {
					log.Fatal("main: resDir is not valid directory.")
				}
			}
		}
	}

	return cfg
}

func getDB(dbFile string, create bool) *sql.DB {
	createDB := false
	if _, err := os.Stat(dbFile); err != nil {
		if !os.IsNotExist(err) {
			log.Fatal(err)
		}
		if !create {
			log.Fatal(err)
		}
		createDB = true
	}

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}

	if createDB {
		sqlStmt := `
		create table foo (id integer not null primary key, name text);
		delete from foo;
		`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			log.Fatal(err)
		}
	}

	return db
}
