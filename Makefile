mysql:
	docker run --name go-restaurant-mysql -p 3308:3306 -e MYSQL_ROOT_PASSWORD=be123@ -d mysql:latest

createdb:
	docker exec -it go-restaurant-mysql mysql -uroot -pbe123@ -e "CREATE DATABASE IF NOT EXISTS \`go-restaurant-data\`"

dropdb:
	docker exec -it go-restaurant-mysql mysql -uroot -pbe123@ -e "DROP DATABASE IF EXISTS \`go-restaurant-data\`"

migrateup:
	migrate -path db/migration -database "mysql://root:be123@@tcp(127.0.0.1:3308)/go-restaurant-data?charset=utf8mb4&parseTime=True&loc=Local" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://root:be123@@tcp(127.0.0.1:3308)/go-restaurant-data?charset=utf8mb4&parseTime=True&loc=Local" -verbose down 1

rollback:
	migrate -path db/migration -database "mysql://root:be123@@tcp(127.0.0.1:3308)/go-restaurant-data?charset=utf8mb4&parseTime=True&loc=Local" -verbose down 1

forceup:
	@read -p "Enter version to force migrate up to: " version; \
	migrate -path db/migration -database "mysql://root:be123@@tcp(127.0.0.1:3308)/go-restaurant-data?charset=utf8mb4&parseTime=True&loc=Local" force $$version; \
	migrate -path db/migration -database "mysql://root:be123@@tcp(127.0.0.1:3308)/go-restaurant-data?charset=utf8mb4&parseTime=True&loc=Local" -verbose up $$version

forcedown:
	@read -p "Enter version to force migrate down to: " version; \
	migrate -path db/migration -database "mysql://root:be123@@tcp(127.0.0.1:3308)/go-restaurant-data?charset=utf8mb4&parseTime=True&loc=Local" force $$version; \
	migrate -path db/migration -database "mysql://root:be123@@tcp(127.0.0.1:3308)/go-restaurant-data?charset=utf8mb4&parseTime=True&loc=Local" -verbose down $$version

.PHONY: mysql createdb dropdb migrateup migratedown rollback forceup forcedown
