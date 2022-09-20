dns := mongodb://root:123456@127.0.0.1:27017/authentication?authSource=admin&directConnection=true
path := ./migration

migrate-up:
	migrate -database "${dns}" -path ${path} up

migrate-down:
	migrate -database "${dns}" -path ${path} down

.PHONY: migrate-up migrate-down