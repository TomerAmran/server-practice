package provider

import (
	"github.com/TomerAmran/server-practice/cache"
	"github.com/TomerAmran/server-practice/database"
)

type contextKeyType string

const (
	ContextKey contextKeyType = "dep-provider"
)

type RepositoryProvider interface {
	Database()  (database.Executor)
	Cache() (cache.Executer)
}