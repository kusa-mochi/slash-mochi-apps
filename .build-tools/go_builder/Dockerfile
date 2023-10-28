FROM buildpack-deps:bullseye-curl
ARG GO_VERSION

# install go
RUN curl -L https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz | tar xz -C /usr/local
ENV PATH $PATH:/usr/local/go/bin
RUN go version

# install go modules
RUN go install github.com/bufbuild/buf/cmd/buf@latest &&\
    go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest &&\
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest &&\
    go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
ENV PATH $PATH:/root/go/bin/

ENV LANG C.UTF-8
ENV LANGUAGE C.UTF-8
ENV LC_ALL C.UTF-8
ENV DEBUG_COLORS true
ENV FORCE_COLOR true

WORKDIR /project