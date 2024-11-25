# lol game scraper

Scraper for lol game data via the riot api. Supports multiple instances that try to avoid duplicated fetches by using a centralized mongodb.

To run just do 
```sh
docker build -t lol-scraper:latest .

docker run \                                                                                                                                                                                    [±main ●▴]
-e RIOT_API_KEY= \
-e ENTRY_USER_NAME= \
-e ENTRY_USER_TAG= \
-e MONGO_CONNECTION_STRING= \
-v $(pwd)/data:/app/data lol-scraper:latest
```
