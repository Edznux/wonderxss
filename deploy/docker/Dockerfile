FROM golang:1.13-alpine AS BUILDER
RUN apk add --update ca-certificates git build-base
RUN set -ex && \
    git clone https://github.com/edznux/wonderxss && \
    cd wonderxss && \
    go build -v && \
    mv wonderxss /usr/bin/wonderxss && \
    mkdir /etc/wonderxss && \
    mv wonderxss.toml.example /etc/wonderxss/wonderxss.toml

FROM alpine
RUN apk add --update ca-certificates && \
    mkdir /etc/wonderxss

# Retrieve the binary from the previous stage
COPY --from=BUILDER /usr/bin/wonderxss /usr/bin/wonderxss
COPY --from=BUILDER /etc/wonderxss.toml /etc/wonderxss/wonderxss.toml

# Set the binary as the entrypoint of the container
ENTRYPOINT [ "/usr/bin/wonderxss"]
CMD ["/usr/bin/wonderxss", "serve"]