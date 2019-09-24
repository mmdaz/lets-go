package repositories

type Database interface {
	GetUserRepo() UserRepo
}
