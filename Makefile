run-db:
	docker run --name my-postgres-container -p 8624:5432 -e POSTGRES_PASSWORD=postgres -d postgres:16.2
drop-db:
	docker rm -f my-postgres-container