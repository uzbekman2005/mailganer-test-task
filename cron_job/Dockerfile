FROM golang:1.19.3-alpine
RUN mkdir cron_job
COPY . /cron_job
WORKDIR /cron_job
RUN go mod tidy
RUN go build -o main cmd/main.go 
CMD ./main