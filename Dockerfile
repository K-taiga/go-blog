FROM golang:latest

RUN apt-get update \
  && apt-get install -y \
  tree

RUN curl -sfL https://direnv.net/install.sh | bash 

WORKDIR /app
