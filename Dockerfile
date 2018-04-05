FROM golang:1.9.5

WORKDIR /go/src/github.com/timwebster9/encryptor-api-go/
RUN go get -d -v github.com/timwebster9/encryptor
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o encryptor-api .

FROM alpine:3.7

WORKDIR /encryptor/
COPY --from=0 /go/src/github.com/timwebster9/encryptor-api-go/encryptor-api .
COPY static ./static

CMD ["/encryptor/encryptor-api"]