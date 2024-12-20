FROM nvcr.io/nvidia/cuda:12.6.2-cudnn-devel-ubuntu20.04

RUN apt-get update && apt-get install -y wget

RUN wget https://golang.org/dl/go1.23.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.23.linux-amd64.tar.gz \
    && rm go1.23.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin

RUN go install github.com/ei-sugimoto/cudair

CMD ["cudair", "run", "-c", ".cudair.toml"]