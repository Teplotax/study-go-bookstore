#!/bin/bash

docker-compose -f ./mysql/docker-compose.yaml down

docker network rm local-environment-network

echo Ambiente local encerrado com sucesso. Até logo!