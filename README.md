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




