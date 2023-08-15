# Recipes API
TLab technical test

## How to run
Make sure [docker](https://docs.docker.com/engine/install/) (and `docker compose` too) already installed.

Then just use this command
```shell
docker compose up
```

## Postgresql DB CLI
```shell
docker exec -it recipes_db psql -d recipes_api -U recipe
```