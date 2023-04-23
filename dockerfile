FROM golang:1.19.5

WORKDIR /home/ubuntu/src/blog

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
#RUN go build -v -o /usr/local/bin/app ./...
#RUN go run ./cmd

CMD ["go", "run", "./cmd"]
EXPOSE 8080