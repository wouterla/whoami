FROM golang:alpine

ADD . /app
WORKDIR /app
RUN go build -o http
ENV PORT 8000
EXPOSE 8000

CMD ["/app/http"]
