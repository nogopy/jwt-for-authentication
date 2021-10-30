#!/usr/bin/env bash

gin --appPort 8080 --port 9000 --build cmd/jwt-for-authentication --immediate run cmd/jwt-for-authentication/main.go