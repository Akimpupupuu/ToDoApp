package users_postgres_repository

import core_postdgres_pool "github.com/Akimpupupuu/ToDoApp/internal/core/repository/posgres/pool"

type UsersRepository struct {
	//для того, чтобы закинуть заглушку и тестировать без бд
	pool core_postdgres_pool.Pool
}

func NewUsersRepository(pool core_postdgres_pool.Pool) *UsersRepository {
	return &UsersRepository{
		pool: pool,
	}
}
