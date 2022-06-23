run: up
	docker-compose exec app go run main.go

up:
	docker-compose up -d

down:
	docker-compose down

sh:
	docker-compose exec app bash
