FROM golang:1.13-alpine AS build-env

LABEL maintainer="bondhan novandy<bondhan.novandy@gmail.com>"

ENV GO111MODULE=on
ENV PROJECT=src/github.com/bondhan/godddnews

COPY . $GOPATH/$PROJECT
WORKDIR $GOPATH/$PROJECT

RUN cd . && go mod download && go build -o $GOPATH/$PROJECT/godddnews main.go


# clean container
FROM alpine

ENV TZ=Asia/Jakarta

RUN apk update && apk upgrade
RUN apk add --no-cache --virtual .build-deps --no-progress -q \
    bash \
    curl \
    busybox-extras \
    tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apk update && apk add --no-cache coreutils

RUN mkdir /app
WORKDIR /app

COPY --from=build-env /go/src/github.com/bondhan/godddnews/godddnews /app/godddnews
COPY --from=build-env /go/src/github.com/bondhan/godddnews/.env /app/.env
COPY --from=build-env /go/src/github.com/bondhan/godddnews/wait-for-it.sh /app/wait-for-it.sh
COPY --from=build-env /go/src/github.com/bondhan/godddnews/doit.sh /app/doit.sh
COPY --from=build-env /go/src/github.com/bondhan/godddnews/godddnews /app/godddnews

RUN chmod +x /app/wait-for-it.sh
RUN chmod +x /app/doit.sh

# RUN /app/wait-for-it.sh --host=postgres-svc --port=5432 --timeout=30 -- /app/godddnews

EXPOSE 8080

# CMD ["./godddnews"]
CMD ["./doit.sh"]

