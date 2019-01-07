# note: call scripts from /scripts
# run:
# 	go run cmd/server/main.go
build: local-env 
	eval "docker-compose -f docker-compose.yml down"
	eval "docker-compose -f docker-compose.yml up -d"

local-env: local-db
	@cat .env_migrator.yaml.example > .env_migrator.yaml
	@cat .env.example > .env
	@echo "Waiting for database connection..."
	@while ! docker exec vote-now_db_1 pg_isready -h localhost -p 5432 > /dev/null; do \
		sleep 1; \
	done
	@echo "Create database connection succesfully!"
local-db: 
	eval "docker-compose -f localdb-docker-compose.yaml down"
	eval "docker-compose -f localdb-docker-compose.yaml up -d"