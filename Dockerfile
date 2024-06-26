FROM golang:latest as BUILDER
LABEL maintainer="shishupei"

# build binary
ARG USER
ARG PASS
RUN echo "machine github.com login $USER password $PASS" >/root/.netrc
RUN mkdir -p /go/src/github.com/opensourceways/message-collect
COPY . /go/src/github.com/opensourceways/message-collect
RUN cd /go/src/github.com/opensourceways/message-collect && CGO_ENABLED=1 go build -v -o ./message-collect main.go

# copy binary config and utils
FROM openeuler/openeuler:22.03
RUN dnf -y update && \
    dnf in -y shadow && \
    groupadd -g 1000 message-center && \
    useradd -u 1000 -g message-center -s /bin/bash -m message-center

USER message-center
COPY  --chown=message-center --from=BUILDER /go/src/github.com/opensourceways/message-collect /opt/app
WORKDIR /opt/app/
ENTRYPOINT ["/opt/app/message-collect"]