FROM golang:1.21-alpine AS builder
WORKDIR /build

COPY go.mod .
RUN go mod download
COPY . .
RUN ls -l
RUN go build -o apix main.go

FROM alpine:latest
COPY --from=builder /build/apix .

EXPOSE 8080

CMD ["./apix"]