#ドッカーハブにあるnodeのイメージを持ってくる
FROM node:18-alpine AS base
#コンテナ内のディレクトリを/appに移動する
WORKDIR /app
#現在のディレクトリのpackage.jsonをコンテナ内のディレクトリ（現在は/appディレクトリにコピー）
COPY ./package.json ./
#コンテナ内のpackage.jsonをもとに必要なライブラリをインストール
RUN npm install
