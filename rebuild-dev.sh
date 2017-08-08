#!/usr/bin/env bash
docker-compose down --rmi all
docker-compose up -d
watchexec --restart --exts "go" --watch . "docker-compose restart app"
