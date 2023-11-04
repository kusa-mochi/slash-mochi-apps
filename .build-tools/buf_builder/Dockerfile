FROM buildpack-deps:bullseye-curl
ARG GO_VERSION
ARG GO_BUILDER_OS
ARG GO_BUILDER_ARCH
ARG NODE_VERSION
ARG NODE_BUILDER_OS
ARG NODE_BUILDER_ARCH

# install go
RUN curl -L https://go.dev/dl/go${GO_VERSION}.${GO_BUILDER_OS}-${GO_BUILDER_ARCH}.tar.gz | tar xz -C /usr/local
ENV PATH $PATH:/usr/local/go/bin
RUN go version

# install node.js
RUN curl -L https://nodejs.org/dist/v${NODE_VERSION}/node-v${NODE_VERSION}-${NODE_BUILDER_OS}-${NODE_BUILDER_ARCH}.tar.gz | tar xz -C /usr/local --strip-components=1 && \
    unlink /usr/local/CHANGELOG.md && unlink /usr/local/LICENSE && unlink /usr/local/README.md

# install go modules
RUN go install github.com/bufbuild/buf/cmd/buf@latest &&\
    go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest &&\
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest &&\
    go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
ENV PATH $PATH:/root/go/bin/

# install npm packages
RUN npm i -g @bufbuild/buf @connectrpc/protoc-gen-connect-es @bufbuild/protoc-gen-es && \
    npm i -g @connectrpc/connect @connectrpc/connect-web @bufbuild/protobuf
ENV PATH $PATH:/usr/local/bin/

ENV LANG C.UTF-8
ENV LANGUAGE C.UTF-8
ENV LC_ALL C.UTF-8
ENV DEBUG_COLORS true
ENV FORCE_COLOR true

WORKDIR /project