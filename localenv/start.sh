#!/bin/bash

echo Bem vindo ao meu ambiente local!

echo Criando network...
docker network create --driver=bridge local-environment-network

sleep 10

echo Executando docker-compose para o serviço do MySQL...
docker-compose -f mysql/docker-compose.yaml up -d --remove-orphans