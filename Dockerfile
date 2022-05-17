FROM golang:alpine

RUN mkdir /app
COPY . /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

COPY . .


RUN go build -o aplikasi .

EXPOSE 8080

CMD ["./aplikasi"]