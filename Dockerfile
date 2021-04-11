FROM node:15.14.0-buster AS JS_BUILD
COPY webapp /webapp
WORKDIR webapp
RUN npm install -g npm@7.9.0
RUN npm install 
RUN npm install --D
RUN npm run build 



FROM  golang:1.16.3-buster AS GO_BUILD
# RUN apk add build-base
WORKDIR /app
COPY app /app
RUN go mod download
# RUN go get -d -v
# Node PPA
RUN curl -sL https://deb.nodesource.com/setup_13.x | bash -
# Downloading Node
RUN apt-get install -y
RUN apt-get install nodejs
# Prisma CLI
RUN npm i -g prisma
RUN go run github.com/prisma/prisma-client-go db push --preview-feature
RUN go run github.com/prisma/prisma-client-go generate --schema=/app/schema.prisma
RUN go run github.com/prisma/prisma-client-go prefetch
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64
COPY --from=JS_BUILD /webapp/dist* ./webapp/
RUN go build -ldflags="-s -w" -o app .
RUN ["chmod", "+x", "./app"]
CMD ./app -prefork





# FROM golang:1.16.3-alpine3.13
# RUN apk add build-base
# WORKDIR /app
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64
# COPY --from=JS_BUILD /webapp/dist* ./webapp/
# COPY --from=GO_BUILD app .
# RUN go get -a
# RUN go build -ldflags="-s -w" -o app .
# RUN ["chmod", "+x", "./app"]
# CMD ./app -prefork