# Single-stage: the gitops-platform compile stage already runs `go build -o
# ./bin/ ./...` and downloads modules (with any private-module credentials
# this app declares) before Kaniko ever runs. This Dockerfile only packages
# that already-built binary -- it does not rebuild from source.
FROM alpine:3.19
RUN adduser -D appuser
USER appuser
WORKDIR /app
COPY bin/ ./bin/
EXPOSE 8081
CMD ["./bin/products-service"]
