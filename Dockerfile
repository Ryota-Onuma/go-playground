FROM golang:1.22.0

RUN useradd -m -s /bin/bash app && \
    apt-get update && \
    apt-get install -y git && \
    apt-get clean

USER app

WORKDIR /home/app

COPY . .

CMD ["go", "run", "src/main.go"]