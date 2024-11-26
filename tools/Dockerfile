FROM bash:5.0.18

WORKDIR /app

COPY ./merger.sh /app/merger.sh
COPY ./csv_export.py /app/csv_export.py

RUN apk update && apk upgrade

RUN apk add --no-cache sqlite python3 bash

RUN chmod +x /app/merger.sh

CMD ["/usr/local/bin/bash", "/app/merger.sh"]

