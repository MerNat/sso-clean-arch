
FROM golang:alpine AS base_alpine

RUN apk add bash ca-certificates git gcc g++ libc-dev

WORKDIR /temp

ADD . /temp

RUN go build -v -o ssoService && mkdir /final \
    && cp -r /temp/ssoService /final \
    && cp -r /temp/envs /final \
    && cp -r /temp/docs/ /final \
    && cp -r /temp/sso.db /final

FROM alpine

ENV ELASTIC_APM_LOG_FILE=stderr
ENV ELASTIC_APM_LOG_LEVEL=debug
ENV ELASTIC_APM_SERVICE_NAME=sso-service-docker
ENV ELASTIC_APM_SERVER_URL=http://apm-server:8200
# ARG GLOBAL_ENVIRONMENT
# ENV GLOBAL_ENV=$GLOBAL_ENVIRONMENT


RUN apk update && apk add ca-certificates \
    && rm -rf /var/cache/apk/*

WORKDIR /usr/src/app

COPY --from=base_alpine /final /usr/src/app/

EXPOSE 8181

CMD [ "./ssoService" ]