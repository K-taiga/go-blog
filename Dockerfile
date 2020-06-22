FROM golang:latest

RUN apt-get update \
  && apt-get install -y \
  tree

RUN curl -sfL https://direnv.net/install.sh | bash \
  # bashrcにdirenvをhook evalは文字列をコマンドとして実行する　
  && echo 'eval "$(direnv hook bash)"' >> ~/.bashrc 

WORKDIR /app
