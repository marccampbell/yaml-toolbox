FROM golang:1.13

ADD . /go/src/github.com/marccampbell/yaml-remarshaler
WORKDIR /go/src/github.com/marccampbell/yaml-remarshaler

RUN make yaml-remarshaler

FROM debian:stretch-slim
COPY --from=0 /go/src/github.com/marccampbell/yaml-remarshaler/bin/yaml-remarshaler /yaml-remarshaler
ENTRYPOINT ["/bin/sh", "-c"]