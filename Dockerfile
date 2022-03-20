FROM alpine:latest
WORKDIR /app

COPY ./*.env .

COPY gomonserver .

CMD ["./gomonserver"]
