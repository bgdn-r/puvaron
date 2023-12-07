FROM golang:1.21.3-alpine as build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o puvaron cmd/puvaron/main.go

FROM alpine
WORKDIR /app
COPY --from=build /app/puvaron .
COPY .env .
EXPOSE 9090
CMD [ "/app/puvaron" ]

