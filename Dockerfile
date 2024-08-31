ARG GO_VERSION=1.22-alpine

FROM golang:${GO_VERSION} AS build

WORKDIR /go/src/app

COPY . .
RUN go build -v -o api .

# Final Stage
FROM alpine:latest
COPY --from=build /go/src/app/api /api
# Copy other necessary files
ENV Origins="*"
CMD ["/api"]
