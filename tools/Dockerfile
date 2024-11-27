FROM alpine:3.20

WORKDIR /app

COPY ./csv_export.py /app/csv_export.py

RUN apk update && apk upgrade

RUN apk add --no-cache sqlite python3

CMD ["python3", "-u", "/app/csv_export.py"]

