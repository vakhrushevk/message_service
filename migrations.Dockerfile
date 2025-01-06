FROM alpine:3.13

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*
ADD http://github.com/pressly/goose/releases/download/v3.24.0/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /root

ADD migrations/*.sql migrations_message/migrations/
ADD migration_message.sh migrations_message/
ADD .env .

RUN ls
RUN chmod +x migrations_message/migration_message.sh

ENTRYPOINT ["bash","migrations_message/migration_message.sh"]