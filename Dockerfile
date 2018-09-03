FROM alpine:latest

RUN apk add --no-cache bash curl wget
RUN addgroup -g 1000 -S www-user && \
    adduser -u 1000 -S www-user -G www-user

ENV MY_VERSION=v.0.1

ADD vpncentralmanager /
USER www-user
CMD ["./vpncentralmanager", "serve"]