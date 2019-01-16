.PHONY: build dev bin local-db local-env clean-local-env unit-test integration-test test clean-env-force

local-db: bin
	eval "docker-compose -f localdb-docker-compose.yaml down"
	eval "docker-compose -f localdb-docker-compose.yaml up -d"
