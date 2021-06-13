postgres:
	docker run --rm --name postgres12 -p 5432:5432 -e POSTGRES_USER=superAdmin -e POSTGRES_PASSWORD=mysecretpassword -d postgres:alpine
createdb:
	docker exec -it postgres12 createdb --username superAdmin --owner superAdmin MyBank
dropdb:
	docker exec -it postgres12 dropdb --username superAdmin MyBank
initdb:
	migrate -path db/migration -database "postgresql://superAdmin:mysecretpassword@localhost:5432/MyBank?sslmode=disable" -verbose up
sqlc: 
	sqlc generate
test: 
	go test -v -cover ./...
start: 
	go run main.go
protoc:
	protoc --go_out=. --go-grpc_out=. proto/*.proto

.PHONY: postgres createdb dropdb initdb sqlc test start proteus protoc