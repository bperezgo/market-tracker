.PHONY:

build-tracker:
	go build -o tracker/executable tracker/cmd/main.go

migrate-create:
	migrate create -ext sql -dir tracker/migrations $(migration_name)

migrate-tracker-up:
	migrate -path tracker/migrations -database postgresql://tracker:dummy-password@localhost:5432/tracker?sslmode=disable up

migrate-tracker-down:
	migrate -path tracker/migrations -database postgresql://tracker:dummy-password@localhost:5432/tracker?sslmode=disable down