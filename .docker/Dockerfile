FROM alpine:3.15.1

RUN addgroup -S ory; \
    adduser -S ory -G ory -D -H -s /bin/nologin
RUN apk add -U --no-cache ca-certificates

COPY kratos /usr/bin/kratos
COPY embedx/config.schema.json /embedx/config.schema.json
COPY identity/stub/identity.schema.json /contrib/extension/identity.schema.json
COPY persistence/sql/migrations/sql /persistence/sql/migrations/sql
USER ory

ENTRYPOINT ["kratos"]
CMD ["serve"]