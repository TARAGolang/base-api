#!/bin/bash
docker stop iothings
docker rm iothings
docker stop mongo-db
docker rm mongo-db
docker stop redis-db
docker rm redis-db