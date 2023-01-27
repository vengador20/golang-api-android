package infraestructure

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

const create string = `
		CREATE TABLE IF NOT EXISTS cache (
		id INTEGER NOT NULL PRIMARY KEY,
		nombre TEXT,
		datos TEXT NOT NULL
		);
`

// const file string = "activities.db"
var (
	one           sync.Once
	wg            sync.Mutex
	connectSqlite *Sqlite
)

type Sqlite struct {
	//Client
	Ctx context.Context
	Db  *sql.DB
}

func Open() *Sqlite {
	one.Do(func() {
		conf := "file:cache.sqlite?cache=shared&mode=memory&_cache_size=6000&_journal=PERSIST&_locking=EXCLUSIVE"
		db, err := sql.Open("sqlite3", conf)

		if err != nil {
			panic("")
		}

		if _, err := db.Exec(create); err != nil {
			panic(err)
		}

		error := db.Ping()

		if error != nil {
			panic("")
		}
		connectSqlite = &Sqlite{
			Db: db,
		}
	})

	return connectSqlite //&Sqlite{Db: db}
}

func (s *Sqlite) Insert(nombre string, datos interface{}) {
	wg.Lock()
	prepare, err := s.Db.Prepare("INSERT INTO cache(nombre, datos) values(?,?)")
	if err != nil {
		panic(err)
	}

	val, err := json.Marshal(datos)

	if err != nil {
		panic(err)
	}

	res, err := prepare.Exec(nombre, val)

	if err != nil {
		panic(err)
	}

	_, error := res.LastInsertId()

	if error != nil {
		panic(error)
	}
	defer wg.Unlock()
}

func (s *Sqlite) Query(nombre string) (map[string]interface{}, error) {
	wg.Lock()
	defer wg.Unlock()
	row := s.Db.QueryRowContext(s.Ctx, "SELECT datos FROM cache WHERE nombre = ? LIMIT 1", nombre)

	var datos string

	err := row.Scan(&datos)
	if err != nil {
		return make(map[string]interface{}), fmt.Errorf(err.Error())
	}

	Mp := make(map[string]interface{})

	Mp["datos"] = datos

	return Mp, nil

}

func (s *Sqlite) Clear() {
	_, err := s.Db.PrepareContext(s.Ctx, "delete from cache")
	if err != nil {
		panic(err)
	}
}
