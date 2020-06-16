FROM golang:latest

RUN apt-get update \
  && apt-get install -y \
  tree

WORKDIR /app
