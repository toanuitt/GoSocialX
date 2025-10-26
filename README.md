# GoSocialX
- Go 1.22
- Docker
- Postgres in docker
- swagger for docs
- golang migrate for migrations

go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
go get -u go.uber.org/zap
go get github.com/go-playground/validator/v10
sudo systemctl stop postgresql
