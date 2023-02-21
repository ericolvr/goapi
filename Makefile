docker-dev:
	docker run -d --name api-todo -p 5432:5432 -e POSTGRES_PASSWORD=secret postgres:13.5

.PHONY: docker-dev