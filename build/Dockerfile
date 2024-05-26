# ================================ #
# Build Stage - Go
# ================================ #
FROM golang:1.22-alpine as melodeon-build-golang
WORKDIR /usr/src/melodeon

COPY go.mod go.sum ./
RUN go mod download && \
    go mod verify

COPY . .
RUN mkdir -p ./bin && \
    go build -v -o ./bin ./...

# ================================ #
# Build Stage - NPM
# ================================ #
FROM node:22-alpine as melodeon-build-npm
WORKDIR /usr/src/melodeon

COPY package.json package-lock.json ./
RUN npm ci

COPY . .
RUN npm run -ws build

# ================================ #
# Runtime Stage
# ================================ #
FROM alpine:latest as melodeon-runtime

COPY ./configs/production.yml /usr/share/melodeon/config.yml
COPY ./web/static /usr/share/melodeon/static/
COPY ./web/templates /usr/share/melodeon/templates/
COPY --from=melodeon-build-npm /usr/src/melodeon/web/app/dist /usr/share/melodeon/assets/
COPY --from=melodeon-build-golang /usr/src/melodeon/bin/melodeon /usr/bin/

CMD [ "melodeon" ]
