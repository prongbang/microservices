FROM golang:alpine as builder
COPY web.go .
RUN go build -o /web .

FROM alpine:3.7
# Define GOTRACEBACK to mark this container as using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=single
CMD ["./web"]
COPY --from=builder /web .