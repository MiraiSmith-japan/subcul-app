.PHONY: build-dev
build-dev:
	docker-compose build

.PHONY: run-dev
run-dev:
	docker-compose up

.PHONY: bye-dev
bye-dev:
	docker-compose down