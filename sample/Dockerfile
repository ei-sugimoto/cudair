FROM nvcr.io/nvidia/cuda:12.6.2-cudnn-devel-ubuntu20.04

RUN apt-get update && apt-get install -y wget

RUN rm -rf /usr/local/go && wget https://go.dev/dl/go1.23.4.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin:/root/go/bin

RUN go install github.com/ei-sugimoto/cudair@latest

WORKDIR /app

COPY . .

CMD ["cudair", "run", "-c", ".cudair.toml"]