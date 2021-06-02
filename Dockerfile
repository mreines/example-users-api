FROM golang:1.10 AS build
WORKDIR /go/src
COPY api ./api
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o usersapi .

FROM scratch AS runtime
COPY --from=build /go/src/usersapi ./
EXPOSE 8080/tcp
ENTRYPOINT ["./usersapi"]
