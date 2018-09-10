phony: postgresRebuild

default: dockerRestart

build:
	docker-compose build

dockerRestart:
	@docker-compose down
	@docker-compose up -d

postgresRebuild:
	(docker rm $$(docker ps -a | grep postgres | awk {'print $$1'}) -f || true) \
	&& sudo rm -rf ./api/pgdata \
	&& mkdir -p ./api/pgdata \
	&& sudo chown -R root:$$(id -g $$(whoami)) ./api/pgdata \
	&& sudo chmod -R 770 ./api/pgdata \
	&& docker-compose up --build --force-recreate postgres
