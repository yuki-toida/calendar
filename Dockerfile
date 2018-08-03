FROM golang:latest as builder
WORKDIR /go/src/github.com/yuki-toida/knowme/
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep && \
    dep ensure && \
    CGO_ENABLED=0 GOOS=linux GIN_MODE=release ENV=dev go build -o app .

FROM alpine:latest
EXPOSE 8080
ENV ENV=dev \
    GOOGLE_APPLICATION_CREDENTIALS="./cred/gcs.json"
RUN apk update && \
    apk upgrade && \
    apk add --no-cache ca-certificates && \
    apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata
WORKDIR /opt/app
COPY --from=builder /go/src/github.com/yuki-toida/knowme/app .
COPY --from=builder /go/src/github.com/yuki-toida/knowme/config ./config
COPY --from=builder /go/src/github.com/yuki-toida/knowme/cred ./cred
COPY --from=builder /go/src/github.com/yuki-toida/knowme/interface/template ./interface/template
COPY --from=builder /go/src/github.com/yuki-toida/knowme/assets/dist ./static
CMD ["./app"]
