FROM golang:1.18-alpine AS build

WORKDIR /out

RUN apk add make gcc musl-dev

COPY . .

RUN CGO_ENABLED=1 make build

RUN apk --no-cache add ca-certificates \
  && update-ca-certificates

ENTRYPOINT [ "./main.out" ]
CMD ["run", "--kafka-env", "dev"]
