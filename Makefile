build:
	docker-compose build
dev:
	- docker image prune -f
	- @make build
	- docker-compose up
exec-db:
	- docker-compose exec mysql bash
exec-web:
	- docker-compose exec web bash
