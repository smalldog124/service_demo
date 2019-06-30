
all: build-app up

build-app:
	docker build -t sales-api .

up:
	docker-compose up -d

down:
	docker-compose down