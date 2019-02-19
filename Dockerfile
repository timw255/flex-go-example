FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

RUN mkdir /build
ADD . /build/
WORKDIR /build

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM scratch

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 10001/tcp
CMD ["./main"]