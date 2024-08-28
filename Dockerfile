FROM golang:1.22.6 as BUILDER

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o api main.go

FROM scratch

WORKDIR /app

COPY --from=BUILDER /app/api ./
COPY --from=BUILDER /app/.env ./


EXPOSE 5000

CMD [ "./api" ]