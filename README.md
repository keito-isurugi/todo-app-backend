# todo-app-backend

## 開発環境の構築

### プロジェクトのアーキテクチャ図
[README.md](docs%2FREADME.md)

### goのインストール
```shell
# 実行時点の最新版がインストールされます
# brewがインストールされていない場合はインストールしてください
$ brew install go 
```

### 起動
dockerを立ち上げて、必要なライブラリのインストールを行います。
```shell
$ make init
```

### APIをpostmanで実行する方法
[README.md](postman%2FREADME.md)

### DB定義書
[README.md](schema%2FREADME.md)

### linterで静的解析
```shell
$ make lint
```

### 自動成型
```shell
$ make fmt
```

### mock生成
```shell
$ make mockgen
```

### test
```shell
$ make test
```

### カバレッジ計測してHTMLで出力
```shell
$ make coverage
```

### swaggerを生成
```shell
$ make swag
```

### 全テーブルをDROPさせて再度DDLを実行したい時
```shell
$ make refresh-schema
```

### ダミーデータを再度挿入したい時
```shell
$ make exec-dummy
```

### 予約のダミーデータ挿入
```shell
$ make exec-dummy-appointments
```

### tblsでDB定義書を生成する方法
```shell
$ make generate-schema
```

DB定義に変更がある場合、再度 `make generate-schema`をすると強制的に上書きされます