# template_app
crud
## web tool
    go get -u github.com/gin-gonic/gin

## env   
    go get -u github.com/spf13/viper
    go get -u github.com/joho/godotenv

## migrations
    migrate create -ext sql -dir ./migrations -seq init

## sql
    go get -u github.com/jmoiron/sqlx
