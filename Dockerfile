FROM golang:1.21.2-alpine

WORKDIR /gameback

# Download Go modules
COPY ["go.mod" , "go.sum" , "./"]
#RUN go mod init gameback_v1
#RUN go mod tidy
RUN go mod download

COPY app ./
COPY config.json ./
# Build
RUN go build -o gameback_v1 ./cmd/

EXPOSE 3030

# Run
CMD ["./gameback_v1"]