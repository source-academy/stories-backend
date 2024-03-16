FROM golang:1.21.5 AS build_stage

WORKDIR /go/src/github.com/source-academy/

# Limit CPU and memory usage to prevent
# crashes during server build
ENV GOMAXPROCS=1
ENV GOMEMLIMIT=256MiB

RUN git clone https://github.com/source-academy/stories-backend.git
RUN cd stories-backend && make build

FROM ubuntu:latest

RUN apt-get update && apt-get install -y \
    # Required to validate SSL certificates
    ca-certificates \
    # Remove apt package lists to save space
    && rm -rf /var/lib/apt/lists/*

WORKDIR /var/lib/sa_stories
# Output entrypoint from "make build" is ./app
COPY --from=build_stage /go/src/github.com/source-academy/stories-backend/app ./server

ENTRYPOINT [ "./server" ]
