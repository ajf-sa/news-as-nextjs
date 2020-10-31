FROM golang:1.15.3-alpine AS GO_BUILD
RUN apk add build-base
COPY app /app
WORKDIR /app
RUN go build -o /go/bin/app

FROM alpine:3.12.0
COPY --from=GO_BUILD /go/bin/app ./
CMD ./app