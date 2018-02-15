FROM golang:1.9.4

RUN mkdir -p /quoteGenerator

WORKDIR /quoteGenerator

ADD . /quoteGenerator

RUN go build ./quoteGenerator.go

CMD ["./quoteGenerator"]