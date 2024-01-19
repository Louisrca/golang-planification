DOCKER_COMPOSE = docker compose

start:
	${DOCKER_COMPOSE} up -d --remove-orphans

stop:
	${DOCKER_COMPOSE} down