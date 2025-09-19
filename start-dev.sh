#!/bin/bash

set -e

echo "--------------------------------- Dev Script ---------------------------------"
echo "Starting docker container of database..."
cd ..
docker-compose up -d

echo "Waiting for database..."
sleep 5

# Migrations:
# The Migrations tool has to be installed, otherwise this will throw an error --> check ReadMe
echo "Database is being set up using migration files..."
migrate -database "postgres://postgres:12345@localhost:5432/kinoprojekt_db?sslmode=disable" -path Backend/migrations up 

echo "--------------------------------- Database is ready ---------------------------------"

# Go Backend Sever start 
echo "Starting Go Backend Server..."
cd Backend
go run ./cmd/api/

echo "--------------------------------- Everything ready! ---------------------------------"
