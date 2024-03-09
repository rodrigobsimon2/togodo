# First stage: build the application
FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o togodo-cli ./cmd/togodo-cli

# Second stage: create a minimal image
FROM scratch

WORKDIR /app

# Copy only the necessary files from the builder stage
COPY --from=builder /app/togodo-cli .

ENV DB_DATABASE=togodo
ENV DB_USER=togodo
ENV DB_PASSWORD=secret
ENV DB_HOST=db
ENV DB_PORT=3306

ENTRYPOINT [ "./togodo-cli" ]