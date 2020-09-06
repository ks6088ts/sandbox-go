FROM golang:1.14.6-buster

ARG WORK_DIR=""
ARG HTTP_PROXY=""
ARG HTTPS_PROXY=""

WORKDIR ${WORK_DIR}

ENV HTTP_PROXY=${HTTP_PROXY}
ENV HTTPS_PROXY=${HTTPS_PROXY}
ENV http_proxy=${HTTP_PROXY}
ENV https_proxy=${HTTPS_PROXY}

RUN apt-get -y update && apt-get install -y \
    make \
    sudo

COPY Makefile Makefile
RUN make grpc-install && \
    make tools-install && \
    make lint-install && \
    make cobra-install && \
    make gqlgen-install && \
    make xo-install 
