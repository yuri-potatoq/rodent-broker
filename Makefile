WEB_SERVER_ADDR?=localhost
WEB_SERVER_PORT?=5000

KAFKA_ENV=local

SERVICE_FLAGS=${WEB_SERVER_ADDR} 
SERVICE_FLAGS+=${WEB_SERVER_PORT}
SERVICE_FLAGS+=${KAFKA_ENV}

BINARY_OUT?=main.out
GO_TAGS=musl

tests:
	go test -tags ${GO_TAGS} ./...


build: tests
	go build -o ${BINARY_OUT} -tags ${GO_TAGS} ./...


run: build
	./${BINARY_OUT} run ${SERVICE_FLAGS}


up-kafka: ./docker-compose.yml
	docker-compose up --build -d broker zookeeper


rm: *.out
	rm -rf $^


clean-images:
	docker rmi \
		confluentinc/cp-zookeeper:7.0.1 \
		confluentinc/cp-kafka:7.0.1


