FROM golang:1.18-alpine

# Chrome için dependencyler
RUN apk add --no-cache \
    udev \
    ttf-freefont \
    chromium

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN ls -la
RUN go build -o takepic main.go
RUN ls -la

# PATH değişkenine chromium ekle tarayıcı çalışması için
ENV PATH="/usr/lib/chromium:${PATH}"

CMD ["./takepic"]
