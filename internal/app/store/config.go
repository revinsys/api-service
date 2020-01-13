package store

import "{GIT_PATH}/internal/app/helpers"

// StoreConfig ...
type StoreConfig struct {
	DbUrl string
}

// NewStoreConfig ...
func NewStoreConfig() *StoreConfig {
	sc := &StoreConfig{}
	sc.DbUrl = helpers.GetEnv("DB_URL", "postgres://user:password@localhost:5432/database?sslmode=disable")

	return sc
}
