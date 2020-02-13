# build stage
FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root

COPY storage/vocabulary.xml storage/vocabulary.xml
COPY --from=builder /app/main .

ENTRYPOINT ./main
