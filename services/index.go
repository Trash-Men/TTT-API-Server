package services

import "github.com/Trash-Men/api-server/repositories"

type ServicesStruct struct {
	UserService     UserService
	JwtService      JwtService
	PasswordService PasswordService
	PhotoService    PhotoService
	TrashService    TrashService
	TrashCanService TrashCanService
}

func Services() ServicesStruct {
	return ServicesStruct{
		UserService:     UserService{},
		JwtService:      JwtService{},
		PasswordService: PasswordService{},
		PhotoService:    PhotoService{},
		TrashService:    TrashService{},
		TrashCanService: TrashCanService{},
	}
}

var _repositories repositories.RepositoriesStruct = repositories.Repositories()
