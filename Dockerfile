FROM golang:1.21

WORKDIR /app

COPY . .

RUN apt-get update && \
    apt-get install -y libpcap-dev

RUN go build -o pegasus

CMD [ "./pegasus" ]