# ベースイメージ
FROM golang:1.23.4-alpine3.20

# 必要なツールをインストール
RUN apk add --no-cache git curl bash

# ホットリロードツール(air)をインストール
RUN curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b /go/bin

# PATHに追加
ENV PATH="/go/bin:$PATH"

# 作業ディレクトリを設定
WORKDIR /app

# go.modとgo.sumをコピーし、依存関係を解決
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをすべてコピー
COPY . .

# ホットリロードの実行
CMD ["air"]
