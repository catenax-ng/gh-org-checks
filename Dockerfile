FROM golang:alpine AS backendBuilder

# Move to working directory /build
WORKDIR /go/src/github.com/catena-x/gh-org-checks

# Copy and download dependency using go mod
COPY ./pkg ./pkg/
COPY go.mod ./
COPY go.sum ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/main main.go

FROM node:18.7.0-alpine As frontendBuilder

WORKDIR /usr/src/app
RUN npm install -g npm@8.18.0
COPY dashboard/package.json dashboard/package-lock.json ./
RUN npm install
COPY ./dashboard ./
RUN npm run build --prod

#use a small image to run
FROM alpine:3.10 as runner

RUN apk update && apk upgrade
RUN apk add --no-cache tzdata

WORKDIR /home/appuser

COPY --from=backendBuilder /go/bin/main .
COPY --from=frontendBuilder /usr/src/app/dist ./dashboard/dist/

EXPOSE 8000
ENTRYPOINT ["./main"]
