FROM golang:alpine AS builder

# Move to working directory /build
WORKDIR /go/src/github.com/catena-x/gh-org-checks

# Copy and download dependency using go mod
COPY ./pkg ./pkg/
COPY go.mod ./
COPY go.sum ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/main main.go

#use a small image to run
FROM alpine:3.10 as runner

RUN apk update && apk upgrade
RUN apk add --no-cache tzdata

WORKDIR /home/appuser

COPY --from=builder /go/bin/main .

EXPOSE 8000
ENTRYPOINT ["./main"]
