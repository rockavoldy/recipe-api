# Recipes API
TLab technical test

## How to run
Make sure [docker](https://docs.docker.com/engine/install/) (and `docker compose` too) already installed
```shell
docker compose up
```

## Postgresql DB CLI
```shell
docker exec -it recipe-api-db-1 psql -d recipes_api -U recipe
```