FROM golang:1.21.5-alpine3.19
WORKDIR /ordersapi
COPY . /ordersapi
RUN go build /ordersapi
EXPOSE 3000
RUN go build -o api main.go
ENTRYPOINT [ "./api" ]