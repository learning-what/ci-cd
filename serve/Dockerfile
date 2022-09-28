# Docker official guides about Go: https://docs.docker.com/language/golang/

# golang official images
# You can see more at https://hub.docker.com/_/golang
FROM golang:1.19-alpine

# Your code directory
WORKDIR /src

# Set GOPROXY to fast download speed in China
ENV GOPROXY https://proxy.golang.com.cn,direct

# Copy core code files
COPY . .

# Download the project dependencies
RUN go mod download

# Build the project
# You can rename the built file you like
RUN go build -o app

# Expose container port
EXPOSE 9090

# Run the built excutable file
ENTRYPOINT ["./app"]
