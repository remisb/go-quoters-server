module github.com/remisb/go-quoters-server

go 1.15

replace github.com/remisb/go-quoters-server/server => ./server

require (
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/pkg/errors v0.9.1
	go.uber.org/zap v1.15.0
)
