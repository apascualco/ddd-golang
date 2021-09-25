.DEFAULT_GOAL := up
mock:
	cd auth; mockgen -source=internal/user_repository.go -destination internal/platform/storage/mockmysql/mock_user_mysql_repository.go
up:
	export BUILDKIT_PROGRESS=tty
	cd infra; docker-compose up --build -d
	cd auth; docker-compose up --build -d
	cd user; docker-compose up --build -d
