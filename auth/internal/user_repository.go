package auth

import "context"

//go:generate mockgen -source=internal/user_repository.go -destination internal/platform/storage/mockmysql/mock_user_mysql_repository.go
type UserRepository interface {
	Save(ctx context.Context, user User) error
	SearchUserByEmail(ctx context.Context, email string) (User, error)
}
