FROM golang:1.23

WORKDIR /app

COPY go.mod .
COPY . .


RUN go get
RUN go build -o bin . 

ENTRYPOINT [ "/app/bin" ]