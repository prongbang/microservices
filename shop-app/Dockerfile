FROM golang:alpine as builder
COPY app.go .
RUN go build -o /app .

FROM alpine:3.7
# Define GOTRACEBACK to mark this container as using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=single
CMD ["./app"]
COPY --from=builder /app .