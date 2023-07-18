FROM golang:1.20.5

# Needed for pg_isready
RUN apt-get update
RUN apt-get install -y postgresql-client

WORKDIR /go/src/github.com/source-academy/

RUN git clone https://github.com/source-academy/stories-backend
WORKDIR /go/src/github.com/source-academy/stories-backend
COPY entrypoint.sh .

ENTRYPOINT [ "./entrypoint.sh" ]
