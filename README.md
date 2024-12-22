# cudair

cudair enable live-reloading for developing CUDA applications like golang-air.
I recommend using docker.

![マイ ムービー](https://github.com/user-attachments/assets/bf339914-acb7-4787-93cb-e8a9d9934ee5)

## Feature

- [x] Live-reloading for CUDA applications
- [x] Docker support
- [x] Easy configuration
- [x] Generate log file on build errors

## Go version

I used only v1.23~ . So I recommend v1.23~

## Install

### Via go install(recommended)

```bash
go install github.com/ei-sugimoto/cudair@latest
```

## Usage

First enter into your project

```bash
cd /path/to/your_projects
```

Next, execute `cudair`

```bash
cudair run -c .cudair.toml
```

You can initialize the configuration file.

```bash
cudair init
```

## configuration

config file must ".toml" .

```toml
# you want to watch root dir path
root = "."
# tmp dir name
tmp_dir = "tmp"

[build]
# execute command
bin = "./tmp/main"
# complie command
cmd = "nvcc --std=c++17 -o ./tmp/main main.cu"
# Generate log file on build errors
log = "build-errors.log"

# Specify the dirs you want to exclude from monitoring
exclude_dir = ["tmp"]
```

## In Docker

How to use docker is stored in the sample directory.

### required

- nvidia driver
- nvidia-container-toolkit

```Dockerfile
FROM nvcr.io/nvidia/cuda:12.6.2-cudnn-devel-ubuntu20.04

RUN apt-get update && apt-get install -y wget

# installing Golang
RUN rm -rf /usr/local/go && wget https://go.dev/dl/go1.23.4.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin:/root/go/bin

RUN go install github.com/ei-sugimoto/cudair@latest

WORKDIR /app

COPY . .

CMD ["cudair", "run", "-c", ".cudair.toml"]
```

building Image...

```bash
docker build -t cudair-sample
```

creating container...

```bash
docker run --rm -it --gpus all -v path/to/your_project:/app cudair-sample
```
