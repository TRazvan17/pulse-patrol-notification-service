# Pulse Patrol - Notification Service (Go)

This repository contains a basic Notification Service implemented in Go.

## REST API

### Run locally
From the repository root:

```bash
go run ./cmd/rest

### Endpoints

## Health Check
curl http://localhost:8080/health
#Expected response 
ok

## Create notifications
curl -X POST http://localhost:8080/notifications \
  -H "Content-Type: application/json" \
  -d '{"to":"lector","message":"hello from pulse patrol"}'
#Expected Response
{"id":"notif-1","status":"queued","createdAt":"2026-02-09T17:30:37Z"}

## gRPC API

### Generate code from proto
From the repository root:

```bash
protoc \
  --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  proto/notification.proto

go run ./cmd/grpc

## Health Check
grpcurl -plaintext localhost:9090 notification.v1.NotificationService/HealthCheck
#SendNotification
grpcurl -plaintext -d '{"to":"lector","message":"hello from grpc"}' \ localhost:9090 notification.v1.NotificationService/SendNotification
