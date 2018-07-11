FROM golang:latest as builder
WORKDIR /go/src/github.com/yuki-toida/knowme/
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep && \
    dep ensure -v && \
    CGO_ENABLED=0 GOOS=linux ENV=dev go build -o app .

FROM alpine:latest
EXPOSE 8080
ENV ENV=dev
RUN apk update && \
    apk upgrade && \
    apk add --no-cache ca-certificates
WORKDIR /opt/app
COPY --from=builder /go/src/github.com/yuki-toida/knowme/app .
COPY --from=builder /go/src/github.com/yuki-toida/knowme/index.html .
COPY --from=builder /go/src/github.com/yuki-toida/knowme/config ./config
COPY --from=builder /go/src/github.com/yuki-toida/knowme/static ./static
CMD ["./app"]
