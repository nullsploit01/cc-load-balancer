COMPOSE_FILE = docker-compose.yml
DOCKER_COMPOSE = docker-compose -f $(COMPOSE_FILE)
IMAGE_NAME = go-http-server
SERVER_DOCKERFILE = server.dockerfile

build-server-image:
	docker build -f $(SERVER_DOCKERFILE) -t $(IMAGE_NAME) .

up: clean build-server-image
	$(DOCKER_COMPOSE) up -d

down:
	$(DOCKER_COMPOSE) down

clean:
	$(DOCKER_COMPOSE) down --volumes --remove-orphans

build: build-image
	$(DOCKER_COMPOSE) build

rebuild: clean build-image build up

logs:
	$(DOCKER_COMPOSE) logs -f

restart:
	$(DOCKER_COMPOSE) restart

status:
	$(DOCKER_COMPOSE) ps

logs-service:
	$(DOCKER_COMPOSE) logs -f $(SERVICE)