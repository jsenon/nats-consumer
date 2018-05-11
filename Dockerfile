FROM alpine:latest

RUN apk add --no-cache bash curl wget
RUN addgroup -g 1000 -S www-user && \
    adduser -u 1000 -S www-user -G www-user

ENV MY_NATSBOOTSTRAP=127.0.0.1

ENV MY_TOPIC=mytest

ADD nats-consumer /
USER www-user

CMD ["./nats-consumer"]