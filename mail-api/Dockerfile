FROM golang:1.19.3-alpine
RUN mkdir api
COPY . /api
WORKDIR /api
RUN go mod tidy
RUN go build -o main cmd/main.go 
CMD ./main