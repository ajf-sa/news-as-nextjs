FROM node:16.1.0-alpine as build
WORKDIR /app
COPY ./app/package*.json ./
COPY ./app .
RUN yarn install \
  --prefer-offline \
  # --frozen-lockfile \
  --non-interactive \
  --production=false

RUN npm run build

CMD npm run start