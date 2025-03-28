# Build documentation and run binary
FROM golang:1.23-alpine AS builder

ARG CGO_ENABLED=1
ARG APP_ENV
ARG SERVER_ADDRESS
ARG PORT
ARG CONTEXT_TIMEOUT
ARG DB_TYPE
ARG DB_HOST
ARG DB_PORT
ARG DB_USER
ARG DB_PASS
ARG DB_NAME
ARG ACCESS_TOKEN_EXPIRY_HOUR
ARG REFRESH_TOKEN_EXPIRY_HOUR
ARG ACCESS_TOKEN_SECRET
ARG REFRESH_TOKEN_SECRET
ARG APP_BINARY_NAME

WORKDIR /app
# !! for sqlite3 dependency
RUN apk add --no-cache gcc musl-dev 
COPY go.mod go.sum ./
RUN go mod download
# !! MAKE SWAGGER DOCS WITH swag
RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY . .
RUN swag init -g cmd/main.go
RUN go build -ldflags='-s -w -extldflags "-static"' -o /app/platform-core cmd/main.go

# Run Binary
FROM scratch AS runner

ENV APP_BINARY_NAME=${APP_BINARY_NAME}
ENV APP_ENV=${APP_ENV}
ENV SERVER_ADDRESS=${SERVER_ADDRESS}
ENV PORT=${PORT}
ENV CONTEXT_TIMEOUT=${CONTEXT_TIMEOUT}
ENV DB_TYPE=${DB_TYPE}
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_USER=${DB_USER}
ENV DB_PASS=${DB_PASS}
ENV DB_NAME=${DB_NAME}
ENV ACCESS_TOKEN_EXPIRY_HOUR=${ACCESS_TOKEN_EXPIRY_HOUR}
ENV REFRESH_TOKEN_EXPIRY_HOUR=${REFRESH_TOKEN_EXPIRY_HOUR}
ENV ACCESS_TOKEN_SECRET=${ACCESS_TOKEN_SECRET}
ENV REFRESH_TOKEN_SECRET=${REFRESH_TOKEN_SECRET}

COPY --from=builder /app/platform-core /platform-core
COPY --from=builder /app/docs/swagger.json /docs/swagger.json
COPY --from=builder /app/docs/swagger.yaml /docs/swagger.yaml

EXPOSE 8080

CMD ["/platform-core"]          
