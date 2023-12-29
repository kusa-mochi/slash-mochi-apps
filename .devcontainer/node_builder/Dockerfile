FROM buildpack-deps:bullseye-curl
ARG NODE_VERSION
ARG NODE_BUILDER_OS
ARG NODE_BUILDER_ARCH

# install node.js
RUN curl -L https://nodejs.org/dist/v${NODE_VERSION}/node-v${NODE_VERSION}-${NODE_BUILDER_OS}-${NODE_BUILDER_ARCH}.tar.gz | tar xz -C /usr/local --strip-components=1 && \
    unlink /usr/local/CHANGELOG.md && unlink /usr/local/LICENSE && unlink /usr/local/README.md

ENV LANG C.UTF-8
ENV LANGUAGE C.UTF-8
ENV LC_ALL C.UTF-8
ENV DEBUG_COLORS true
ENV FORCE_COLOR true

WORKDIR /project