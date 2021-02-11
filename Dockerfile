FROM golang:1.15.6-alpine3.12 as build-env

RUN mkdir /stategate
RUN apk --update add ca-certificates build-base
RUN apk add make git
WORKDIR /stategate
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go install ./cmd/...

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=build-env /go/bin/ /usr/local/bin/
WORKDIR /workspace
EXPOSE 8080
EXPOSE 8081

ENTRYPOINT ["/usr/local/bin/stategate"]