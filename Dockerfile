FROM golang:1.13

ADD . /go/src/github.com/marccampbell/yaml-toolbox
WORKDIR /go/src/github.com/marccampbell/yaml-toolbox

RUN make yaml-toolbox

FROM debian:stretch-slim
COPY --from=0 /go/src/github.com/marccampbell/yaml-toolbox/bin/yaml-toolbox /yaml-toolbox
ENTRYPOINT ["/bin/sh", "-c"]