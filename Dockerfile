FROM golang:1.23-alpine AS builder
RUN apk add --no-cache git
WORKDIR /build
COPY go.mod .
COPY . .
RUN go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -v -o /build/resume ./cmd/webserver/main.go

# final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /build/resume /app/resume
# copy the templates directory
COPY --from=builder /build/internal/templates /app/internal/templates
COPY --from=builder /build/static /app/static
RUN chmod +x /app/resume
CMD ["/app/resume"]
