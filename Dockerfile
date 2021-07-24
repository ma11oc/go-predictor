FROM golang:alpine as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o predictor cmd/server/main.go


FROM scratch
COPY --from=builder /build/predictor /app/
COPY --from=builder /build/locales/ru-RU.yaml /etc/predictor/locales/

WORKDIR /app
EXPOSE "50051"
EXPOSE "8080"

CMD ["./predictor", "-grpc-port", "50051", "-http-port", "8080", "-locale", "/etc/predictor/locales/ru-RU.yaml"]
