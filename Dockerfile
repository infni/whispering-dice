FROM golang:1.19 as GOBASE

WORKDIR /app

COPY . .

RUN go vet ./... \
    && result=`go test -timeout 2s -v ./test/... 2>&1` ; rc=$? \
    && echo "$result" | tee testResults.txt \
    && [ "$rc" -eq 0 ] \
    && go build -v -o "./bin/whisperingdice" -ldflags="-X main.version=$(git describe --always --long)" ./cmd/whisperingdice

RUN ./bin/whisperingdice -version > version && cat version

ARG PUB="/pub"
RUN mkdir -p $PUB \
    && cp ./bin/whisperingdice $PUB \
    && cp testResults.txt $PUB  \
    && cp version $PUB

# docker-in-docker
#FROM gcr.io/distroless/base
FROM ubuntu:23.04

RUN apt-get update && \
    apt-get install -y ca-certificates

WORKDIR /pub

COPY --from=GOBASE /pub .

ARG BOT_TOKEN
ENV TOKEN=${BOT_TOKEN}
ENV APPID=${APP_ID}
ENV GUILDID=${GUILD_ID}

CMD ["/pub/whisperingdice"]
