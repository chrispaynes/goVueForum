phony: postgresRebuild

default: dockerRestart

build:
	docker-compose build

dockerRestart:
	@docker-compose down
	@docker-compose up -d

postgresRebuild:
	docker rm $$(docker ps -a | grep postgres | awk {'print $$1'}) -f && docker-compose up --build --force-recreate postgres
