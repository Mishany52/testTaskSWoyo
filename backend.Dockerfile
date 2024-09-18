FROM golang:1.22-alpine AS backend

WORKDIR /go/src/app
COPY go.* ./
COPY internal ./internal
COPY pkg ./pkg
COPY utils ./utils
RUN go mod download
COPY cmd/main ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-s' -o main .

FROM scratch
COPY --from=backend /go/src/app/main /backend
EXPOSE 8080
CMD ["/backend", "-d"]
COPY .env ./