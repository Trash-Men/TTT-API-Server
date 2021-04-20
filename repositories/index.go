package repositories

import (
	"github.com/go-pg/pg/v10"
)

type RepositoriesStruct struct {
	UserRepository     UserRepository
	TrashRepository    TrashRepository
	TrashCanRepository TrashCanRepository
}

func Repositories() RepositoriesStruct {
	return RepositoriesStruct{
		UserRepository:     UserRepository{},
		TrashRepository:    TrashRepository{},
		TrashCanRepository: TrashCanRepository{},
	}
}

var dbClient *pg.DB

func SetDBClient(client *pg.DB) {
	dbClient = client
}
