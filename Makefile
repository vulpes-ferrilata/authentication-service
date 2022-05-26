dns := mysql://auth:123456@tcp(localhost:3306)/auth?charset=utf8mb4
path := ./migration

migrate-up:
	migrate -database ${dns} -path ${path} up

migrate-down:
	migrate -database ${dns} -path ${path} down

.PHONY: migrate-up migrate-down