FROM golang:1.20.8-alpine

WORKDIR /app

COPY . .

RUN apk add --no-cache git && \
    go mod download

ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgres
ENV POSTGRES_DB payment_user

RUN apk add --no-cache postgresql-client && \
    apk add --no-cache --virtual .build-deps gcc musl-dev postgresql-dev && \
    apk --purge del .build-deps

RUN cd cmd && go build -o main .

EXPOSE 5000

CMD ["./cmd/main"]
