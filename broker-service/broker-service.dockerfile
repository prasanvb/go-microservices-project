#base go image
FROM golang:1.18-alpine as builder

#create app directory
RUN mkdir /app

#copy contents from current to newly created app folder
COPY . /app

# nominate workdir
WORKDIR /app

# run go build command
RUN CDO_ENABLED=0 go build -o brokerAPP ./src/api

# specifies that executable permissions should be added
RUN chmod +x /app/brokerAPP

# build tiny docker image from the above
FROM alphine:latest

RUN mkdir /app

COPY --from=builder /app/brokerAPP /app

CMD [ "/app/brokerAPP" ]