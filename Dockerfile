
FROM golang:alpine AS base_alpine

RUN apk add bash ca-certificates git gcc g++ libc-dev

WORKDIR /temp

ADD . /temp

RUN go build -v -o ssoService && mkdir /final \
    && cp -r /temp/ssoService /final \
    && cp -r /temp/envs /final \
    && cp -r /temp/docs/ /final

FROM alpine

# ARG GLOBAL_ENVIRONMENT

# ENV GLOBAL_ENV=$GLOBAL_ENVIRONMENT

RUN apk update && apk add ca-certificates \
    && rm -rf /var/cache/apk/*

WORKDIR /usr/src/app

COPY --from=base_alpine /final /usr/src/app/

EXPOSE 8181

CMD [ "./ssoService" ]