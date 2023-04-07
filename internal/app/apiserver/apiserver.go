package apiserver

import (
	"database/sql"

	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/cache"
	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/store/sqlstore"
)

// Start ...
func Start(config *ConfigT) error {
	db, err := newDB(config.DataBase.URL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)

	if config.UseDBCache {
		// Create a cache and connect it to the store
		c := cache.New()
		store.SetCache(c)
		store.EnableCache()
	}

	srv, err := newServer(store, config)
	if err != nil {
		return err
	}

	return srv.frontend.ListenAndServeTLS("", "")
}

// newDB ...
func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
