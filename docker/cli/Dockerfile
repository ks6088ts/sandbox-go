FROM sandbox-go_go:latest AS builder

ARG WORK_DIR=""
ARG HTTP_PROXY=""
ARG HTTPS_PROXY=""

WORKDIR ${WORK_DIR}

COPY ./ ./

RUN make build GOBUILD="CGO_ENABLED=0 go build"

# ---

FROM alpine:3.12.0

ARG WORK_DIR=""

COPY --from=builder ${WORK_DIR}/outputs/sandbox-go ./cli
