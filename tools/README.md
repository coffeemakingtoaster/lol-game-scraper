# Merger

Merge multiple sqlite dbs resulting from scraper run into one and export as csv in format that contains winner and loser champions for every role.
Simply run:

```sh
docker run -v $(pwd)/in:/app/in -v $(pwd)/export:/app/export $(docker build -q .)
```
