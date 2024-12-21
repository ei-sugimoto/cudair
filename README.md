# cudair
cudair enable live-reloading for developing CUDA applications like golang-air.
I recommend using docker.



## Feature
- [x] Live-reloading for CUDA applications
- [x] Docker support
- [x] Easy configuration

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

ENV PATH=$PATH:/usr/local/go/bin

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

