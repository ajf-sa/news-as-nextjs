FROM node:15.12 AS JS_BUILD
COPY webapp /webapp
WORKDIR webapp
RUN npm install 
RUN npm run build 

FROM golang:1.16.2-alpine3.13 AS GO_BUILD
# RUN apk add build-base
RUN apk update 
WORKDIR /app
COPY app /app
COPY app/.env ./.env
ENV DATABASE_URL = postgresql://admin:secret@db01:5432/newsdemo

RUN go run github.com/prisma/prisma-client-go prefetch
RUN go run github.com/prisma/prisma-client-go generate
RUN go build -ldflags="-s -w" -o app .
COPY --from=JS_BUILD /webapp/build* ./webapp/
CMD ./app -prefork
