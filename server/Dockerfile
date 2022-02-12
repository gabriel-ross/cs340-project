FROM golang

WORKDIR /app/server

COPY . .
RUN go mod download
RUN mkdir -p bin
RUN go build -o bin/server cmd/main.go

EXPOSE 80

CMD /bin/server