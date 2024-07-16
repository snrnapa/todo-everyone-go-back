#!/bin/bash

# 開発環境では、デバッグを行いたいため、postgresqlのみを起動する。
# sudo docker comopse up -d

# 本番環境用
# 本番環境では、goのインストールを行っていないため、go appの起動自体もdockerで行っているため、
# docker-compose-prod.ymlの内容でコンテナの作成を行う必要がある。
sudo docker-compose -f docker-compose.yml -f docker-compose-prod.yml build
sudo docker-compose -f docker-compose.yml -f docker-compose-prod.yml up -d