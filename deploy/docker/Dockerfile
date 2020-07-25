# docker build -t edznux/wonderxss:$VERSION .
# docker run -it -p 80:80 -v $(pwd)/wonderxss.toml:/etc/wonderxss/wonderxss.toml edznux/wonderxss:latest

FROM golang:1.13-alpine AS BUILDER
RUN apk add --update ca-certificates git build-base
COPY . /app
WORKDIR /app
RUN ls && go build -v

FROM alpine
RUN apk add --update ca-certificates && \
    mkdir /etc/wonderxss/

# Retrieve the binary from the previous stage
COPY --from=BUILDER /app/wonderxss /app/wonderxss

# Set the binary as the entrypoint of the container
ENTRYPOINT [ "/app/wonderxss"]
CMD ["serve"]