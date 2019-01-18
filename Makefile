.PHONY: build dev bin local-db local-env clean-local-env unit-test integration-test test clean-env-force

local-db: bin
	eval "docker-compose down"
	eval "docker-compose up -d"

cln-bin:
	rm -rf bin
bin-migrator:
	go build -o bin/migrator ./cmd/migrator
bin: cln-bin bin-migrator

env-file:
	@if ! [ -e ".env_migrator.yaml" ] ; then cat .env_migrator.yaml.example > .env_migrator.yaml; fi
	@if ! [ -e "sqlboiler.toml" ] ; then cat sqlboiler.toml.example > sqlboiler.toml; fi
	@if ! [ -e ".env" ] ; then cat .env.example > .env ; fi

local-env: local-db env-file
	@echo "Waiting for database connection..."
	@while ! docker exec vote-now_db pg_isready -h localhost -p 5432 > /dev/null; do \
		sleep 1; \
	done
	bin/migrator up
# SQL_BOILER
gen-sqlboiler-model:
	@sqlboiler --wipe psql

run:
	go build -o main cmd/vote-now/*
	./main