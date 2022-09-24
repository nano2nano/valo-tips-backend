FROM golang:1.18 as builder

ARG PORT=8080
ENV PORT=${PORT}
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/nano2nano/valo-tips-backend
COPY . .
RUN go build -o /go/bin/valo-tips-backend ./cmd/valo_tips/main.go


# runtime image
FROM alpine
COPY --from=builder /go/bin/valo-tips-backend /app/main

CMD /app/main
