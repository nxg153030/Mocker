FROM golang:1.21-bullseye

WORKDIR /app

# Copy your source code into the container
COPY . .

# We use a shell as the default entrypoint so we can run commands manually
CMD ["/bin/bash"]