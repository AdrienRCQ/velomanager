FROM golang:1.23.4

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod tidy

COPY backend/ ./

RUN go build -o server .

ENV MYSQL_DSN=user:password@tcp(db:3306)/velodb?parseTime=true

CMD ["./server"]
