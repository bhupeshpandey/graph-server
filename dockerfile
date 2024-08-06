# Use a minimal base image
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary file from the host
COPY ./build/linux/amd64/main .

# Expose port 8080 to the outside world
EXPOSE 2007

# Command to run the executable
CMD ["./main"]