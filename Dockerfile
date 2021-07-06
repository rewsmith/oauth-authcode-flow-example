FROM golang:1.12-alpine
RUN mkdir /app
ADD . /app
RUN apk add --no-cache git \
    && go get github.com/Sirupsen/logrus \
    && go get github.com/joho/godotenv \
    && apk del git
WORKDIR /app
# Our project will now successfully build with the necessary go libraries included.
RUN go build -o main .
# Our start command which kicks off
# our newly created binary executable
CMD ["/app/main"]