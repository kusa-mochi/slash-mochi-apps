FROM buildpack-deps:bullseye-curl
ARG NODE_VERSION

# install node.js
RUN curl -L https://nodejs.org/dist/v${NODE_VERSION}/node-v${NODE_VERSION}-linux-x64.tar.gz | tar xz -C /usr/local --strip-components=1 && \
    unlink /usr/local/CHANGELOG.md && unlink /usr/local/LICENSE && unlink /usr/local/README.md

# install npm packages
RUN npm i -g @bufbuild/buf @bufbuild/protoc-gen-connect-web @bufbuild/protoc-gen-es @bufbuild/connect @bufbuild/connect-web @bufbuild/protobuf
ENV PATH $PATH:/usr/local/bin/

ENV LANG C.UTF-8
ENV LANGUAGE C.UTF-8
ENV LC_ALL C.UTF-8
ENV DEBUG_COLORS true
ENV FORCE_COLOR true

WORKDIR /project