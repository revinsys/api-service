package store

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//Store ...
type Store struct {
	config               *StoreConfig
	db                   *sqlx.DB
}

//NewStore ...
func NewStore() *Store {
	return &Store{
		config: NewStoreConfig(),
	}
}

//Open ...
func (s *Store) Open() error {
	db, err := sqlx.Connect("postgres", s.config.DbUrl)
	if err != nil {
		return err
	}

	s.db = db
	return nil
}

