FROM golang:1.24.2

COPY . .

RUN go test -v > test.log

CMD ["cat", "test.log"]
