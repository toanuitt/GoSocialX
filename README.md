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
go install github.com/swaggo/swag/cmd/swag@latest
http://localhost:8080/v1/swagger/index.html#/
go get github.com/google/uuid
go get gopkg.in/mail.v2
go get github.com/sendgrid/sendgrid-go
go get -u github.com/golang-jwt/jwt/v5