FROM golang:1.14-alpine3.11 AS builder

LABEL maintainer="Nanang Suryadi <nanang.suryadi@aegis.gg>"

ARG BINARY_NAME="{{cookiecutter.repo_name}}"
ARG SSH_PRIVATE_KEY

#-----------------------------------------------------------
#   SET SSH REPO
#-----------------------------------------------------------
RUN set -eux && apk --update --no-cache add ca-certificates git binutils openssh bash upx
RUN set -eux && mkdir ~/.ssh/ \
  && touch ~/.ssh/config \
  && echo -e "${SSH_PRIVATE_KEY}" > ~/.ssh/id_rsa \
  && chmod 0600 ~/.ssh/id_rsa \
  && echo "IdentityFile $HOME/.ssh/id_rsa" >> ~/.ssh/config \
  && echo -e "Host *\n\tStrictHostKeyChecking no\n\n" > ~/.ssh/config \
  && git config --global url."git@gitlab.com:".insteadOf "https://gitlab.com/"

#-----------------------------------------------------------
#   COPY AND COMPILE APPS
#-----------------------------------------------------------
WORKDIR "$(echo $GOPATH)/src/{{cookiecutter.repo_path}}/{{cookiecutter.repo_name}}"
COPY . .

# Using go mod.
RUN go mod download
RUN go mod verify

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o "/go/bin/$(echo $BINARY_NAME)" ./cmd/{{cookiecutter.repo_name}}/main.go
RUN strip /go/bin/*

RUN mkdir -p /etc/config/
COPY ./config/app.* /etc/config/

# Compress binary
RUN set -eux && upx "/go/bin/$(echo $BINARY_NAME)" && \
  upx -t "/go/bin/$(echo $BINARY_NAME)" && \
  chmod +x "/go/bin/$(echo $BINARY_NAME)"

# Clear cache and app files
RUN set -eux && apk del ca-certificates git binutils openssh bash upx && \
  rm -rf /var/cache/apk/* $GOPATH/src/{{cookiecutter.repo_path}}

FROM alpine:3.8

ARG jaeger_agent_host
ENV JAEGER_AGENT_HOST=${jaeger_agent_host}

RUN mkdir -p /go/bin/config

COPY --from=builder /go/bin/{{cookiecutter.repo_name}}  /{{cookiecutter.repo_name}}
COPY --from=builder /etc/config/app.* /go/bin/config

WORKDIR /go/bin/

EXPOSE 8009

ENTRYPOINT ["/{{cookiecutter.repo_name}}"]
