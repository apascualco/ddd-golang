package user

import "context"

//go:generate mockgen -source=internal/user_repository.go -destination internal/platform/storage/mockmysql/mock_user_mysql_repository.go
type UserRepository interface {
	Save(ctx context.Context, user User) error
	SearchUserByID(ctx context.Context, uuid UUID) (User, error)
}
