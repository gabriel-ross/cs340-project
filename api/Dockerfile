FROM golang

WORKDIR /app/api

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN mkdir -p bin
RUN go build -o ./bin/server cmd/main.go

EXPOSE 80

CMD ./bin/server