FROM golang:1.14.6-buster

ARG work_dir=""
ARG http_proxy=""
ARG https_proxy=""

WORKDIR ${work_dir}

RUN apt-get -y update && apt-get install -y \
    make

COPY Makefile Makefile
RUN make grpc-install && \
    make tools-install && \
    make lint-install && \
    make cobra-install
