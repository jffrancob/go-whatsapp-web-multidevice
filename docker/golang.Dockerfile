############################
# STEP 1 build executable binary
############################
FROM golang:1.21.5-alpine3.19 AS builder
RUN apk update && apk add --no-cache vips-dev gcc musl-dev gcompat ffmpeg
WORKDIR /whatsapp
COPY ./src .

# Fetch dependencies.
RUN go mod download
# Install pkger
RUN go install github.com/markbates/pkger/cmd/pkger@latest
# Build the binary.
RUN pkger && go build -o /app/whatsapp

#############################
## STEP 2 build a smaller image
#############################
FROM alpine:3.19
RUN apk update && apk add --no-cache vips-dev ffmpeg
COPY ./docs /docs
WORKDIR /app
# Copy compiled from builder.
COPY --from=builder /app/whatsapp /app/whatsapp
# Run the binary.
ENTRYPOINT ["/app/whatsapp"]