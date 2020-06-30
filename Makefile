
run: build
	./bin/mysql-guard --db_user mylxsw --adanos_server http://localhost:19999 --killer --killer_busy_time 10 --deadlock_logger

build: build-orm
	go build -o bin/mysql-guard main.go

build-orm:
	orm models/*.yml
