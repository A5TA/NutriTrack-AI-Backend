FROM golang:1.23

# Step 1: Set the working directory inside the container
WORKDIR /app

# Step 2: Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Step 3: Copy the entire application code into the container
COPY . .

# Step 4: Build the Go application
RUN go build -o bin ./cmd/api-server

# Expose the application port - optional
EXPOSE 8050

ENTRYPOINT [ "/app/bin" ]