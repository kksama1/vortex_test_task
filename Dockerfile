FROM golang:1.22-alpine AS builder

WORKDIR /usr/loclal/src

RUN apk --no-cache add bash gcc gettext

# dep
COPY /go.mod /go.sum ./
RUN go mod download

COPY . .



RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/app ./cmd/main.go

FROM alpine

COPY --from=builder /usr/loclal/src/bin/app /

COPY src/sql/algorithm.sql /usr/local/src/db/sql/
COPY src/sql/client.sql /usr/local/src/db/sql/
COPY .env  /

CMD ["/app"]