FROM alpine:3.14

# Update package lists and install required dependencies
RUN apk update && apk upgrade && apk --no-cache add wget bash git openssh gcc libc-dev tzdata
RUN apk add --no-cache tzdata

# Set the working directory inside the container to /app
WORKDIR /app

# Copy the binary executable "app" (presumably built from Go source code) into the container
COPY app /app

# Expose port 8888 on the container, allowing connections to the app
EXPOSE 8080

# Set the default command to run when the container starts
CMD ["./app"]