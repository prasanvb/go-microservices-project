#base go image
FROM golang:1.23-alpine as builder

#create app directory
RUN mkdir /app

#copy contents from current to newly created app folder
COPY . /app

# nominate workdir
WORKDIR /app

# run go build command
RUN CDO_ENABLED=0 go build -o brokerApp ./src/api

# specifies that executable permissions should be added
RUN chmod +x /app/brokerApp

# build tiny docker image from the above
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/brokerApp /app

CMD [ "/app/brokerApp" ]