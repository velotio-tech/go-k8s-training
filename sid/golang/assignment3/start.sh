#!/bin/bash

docker build -t dependencies -f ./deps.Dockerfile .
docker-compose up
