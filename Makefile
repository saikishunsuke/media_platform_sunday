build:
	docker-compose build
dev:
	- docker image prune -f
	- @make build
	- docker-compose up