DOCKER_COMPOSE = docker-compose

.PHONY: start
start:
	${DOCKER_COMPOSE} up -d --remove-orphans

.PHONY: stop
stop:
	${DOCKER_COMPOSE} down