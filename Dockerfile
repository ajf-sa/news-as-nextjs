FROM golang:1.14

WORKDIR /fiber

COPY ./src /fiber

RUN go get github.com/valyala/quicktemplate/qtc

RUN go generate ./templates
RUN go build -ldflags="-s -w" -o app .

CMD ./app -prefork

# FROM golang:1.15.3-alpine AS GO_BUILD
# RUN apk add build-base
# COPY app /app
# WORKDIR /app
# RUN go build  -ldflags="-s -w" -o /go/bin/app

# FROM alpine:3.12.0
# COPY app/prod.env ./.env
# COPY --from=GO_BUILD /go/bin/app ./
# CMD ./app -prefork