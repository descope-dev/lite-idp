FROM golang:1-alpine

COPY . /go/src/github.com/descope-dev/lite-idp
WORKDIR /go/src/github.com/descope-dev/lite-idp
RUN go build .

FROM alpine

COPY --from=0 /go/src/github.com/descope-dev/lite-idp/lite-idp /usr/bin/lite-idp

ENTRYPOINT ["lite-idp"]
CMD ["serve"]
