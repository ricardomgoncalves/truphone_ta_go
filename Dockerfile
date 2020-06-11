FROM golang:1.14-alpine as builder
RUN apk --no-cache add ca-certificates git
WORKDIR /build/truphone

# Fetch dependencies
COPY go.mod ./
RUN go mod download

# Build
COPY . ./
RUN CGO_ENABLED=0 go build -o truphone cmd/service/main.go

# Create final image
FROM alpine
WORKDIR /root
COPY --from=builder /build/truphone/truphone .
EXPOSE 8080
CMD ["./truphone"]