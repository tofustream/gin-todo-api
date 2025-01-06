#!/bin/bash
set -e

# マイグレーションの実行
go run ./cmd/internal/migrations/migration.go

# ホットリロードの実行
exec air
