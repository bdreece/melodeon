# ================================ #
# Build Stage
# ================================ #
FROM golang:1.21-alpine as melodeon-build
WORKDIR /usr/src/melodeon

RUN apk upgrade --no-cache && \
    apk add --no-cache nodejs npm

COPY package.json package-lock.json ./
RUN npm ci

COPY go.mod go.sum ./
RUN go mod download && \
    go mod verify

COPY . .
RUN make build

# ================================ #
# Runtime Stage
# ================================ #
FROM alpine:latest as melodeon-runtime

COPY ./configs/production.yml /usr/share/melodeon/config.yml
COPY ./web/app/dist /usr/share/melodeon/assets/
COPY ./web/static /usr/share/melodeon/static/
COPY ./web/templates /usr/share/melodeon/templates/
COPY --from=melodeon-build /usr/src/melodeon/bin/melodeon /usr/bin/

CMD [ "melodeon" ]
