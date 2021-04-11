FROM golang:1.16.3-buster  AS BUILD
# RUN apt-get install -y build-essential
RUN curl -fsSL https://deb.nodesource.com/setup_lts.x |  bash -
RUN  apt-get install -y nodejs
WORKDIR /build
COPY webapp /build
RUN npm install 
RUN npm install --D
RUN npm run build 

WORKDIR /app
COPY app /app
# RUN go mod download
RUN  go get github.com/prisma/prisma-client-go/generator/types@v0.7.0
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN npm i -g prisma
RUN go run github.com/prisma/prisma-client-go db push --preview-feature
RUN go run github.com/prisma/prisma-client-go generate --schema=/app/schema.prisma
RUN go run github.com/prisma/prisma-client-go prefetch

RUN go build -ldflags="-s -w" -o app .
# RUN ["chmod", "+x", "./app"]
# CMD ./app -prefork


FROM golang:1.16.3-alpine3.13
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64
COPY --from=BUILD /webapp/dist* ./webapp/
COPY --from=BUILD app .
RUN ["chmod", "+x", "./app"]
CMD ./app -prefork