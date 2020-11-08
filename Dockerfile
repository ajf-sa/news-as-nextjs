FROM node:12.19 AS JS_BUILD
COPY webapp /webapp
WORKDIR webapp
RUN npm install 
RUN npm run build 

FROM golang:1.15.3 AS GO_BUILD
# RUN apk add build-base
WORKDIR /app
COPY app /app
COPY app/prod.env ./.env
RUN go build -ldflags="-s -w" -o app .
COPY --from=JS_BUILD /webapp/build* ./webapp/
CMD ./app 
#-prefork
