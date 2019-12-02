FROM golang:1.13-alpine AS BUILDER
RUN apk add --update ca-certificates git build-base
RUN set -ex && \
    git clone https://github.com/edznux/wonderxss && \
    cd wonderxss && \
    go build -v && \
    mv wonderxss /usr/bin/wonderxss

FROM alpine
RUN apk add --update ca-certificates

# Retrieve the binary from the previous stage
COPY --from=BUILDER /usr/bin/wonderxss /usr/bin/wonderxss

# Set the binary as the entrypoint of the container
ENTRYPOINT [ "/usr/bin/wonderxss" ]