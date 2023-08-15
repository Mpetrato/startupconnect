migrate:
	migrate -source file://db/migrations \
					-database postgres://teste:teste@localhost:5432/postgres?sslmode=disable up