FROM golang:1.21 AS build
WORKDIR /app
COPY main.go .
RUN go mod init blog && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o main .


FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main .
CMD ["./main"]