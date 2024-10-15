 # Golang backend environment build container
FROM golang:1.23 AS build
ENV DEBIAN_FRONTEND noninteractive
ENV ENVIRONMENT=production

WORKDIR /app
ADD . .

ENV GOOS linux
ENV GOARCH amd64
RUN go build -o /app/bin/httpopportunity ./cmd/http/main.go
RUN go build -o /app/bin/chef ./cmd/chef/main.go
RUN go build -o /app/openctl ./cmd/openapi/*.go
RUN /app/openctl --path=public/apidocs --internal-dir-path=services

# Deploy container
FROM ubuntu AS deploy
ENV DEBIAN_FRONTEND noninteractive
ENV ENVIRONMENT=production

# Update and install necessary packages
RUN apt-get update && apt-get install -y \
    ca-certificates \
    poppler-utils 

# Copying environment grpc server binary to /usr/src/app
WORKDIR /app
COPY --from=build /app/bin/httpopportunity .
COPY --from=build /app/bin/chef .
COPY --from=build /app/.env.production ./.env.production

COPY --from=build /app/public ./public

EXPOSE 80

# Start The Project
CMD ["/app/httpopportunity"]
