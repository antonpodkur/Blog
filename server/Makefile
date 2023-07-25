BINARY_NAME=blog
POSTGRES_URL=postgresql://postgres:mypassword@localhost:5432/blog?sslmode=disable

build:
	go build -o ${BINARY_NAME} ./cmd/blog/blog.go

run:
	go run ./cmd/blog/blog.go

migrateup:
	migrate -path db/migrations -database "${POSTGRES_URL}" -verbose up

migratedown:
	migrate -path db/migrations -database "${POSTGRES_URL}" -verbose down