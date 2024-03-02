アーキテクチャ図
<pre>
.
├── DDL
├── cmd // APIやバッチなどの実行ファイルを置く場所
│   ├── api // api
│   └── seeder // ダミーデータ生成
├── docker // dockerfileやdocker-compose.ymlを置く場所
│   └── localstack // localstackの設定ファイル
├── docs // ドキュメントを置く場所
├── internal // 内部のAPIのプログラムのソースコードを置く場所。外部との接続についてはここには置かない
│   ├── domain
│   │   ├── entity // ドメインオブジェクト いわゆるmodel
│   │   └── repository // データストアとのやり取りを抽象化するためのリポジトリインターフェイス
│   │       └── mock // テストのためのモック
│   ├── handler // APIリクエストを処理するディレクトリ
│   ├── helper // 共通の補助関数
│   ├── infra // プロジェクトのインフラストラクチャに関連するコード
│   │   ├── aws // AWS
│   │   ├── db // DBとの接続やロガーの設定
│   │   ├── env // 環境変数
│   │   ├── logger // ロガー。DBのロガーとは別で、HTTPリクエストやレスポンスで使用する
│   │   ├── postgres // PostgreSQLからデータを取得する処理
│   │   └── storage // S3などのファイルストレージとのやり取り
│   ├── server // APIサーバーの設定。middlewareやrouterの設定
│   │   └── middleware // middleware
│   └── usecase // アプリケーションのユースケース（ビジネスロジック）を格納。メソッドごとにUsecaseを作成する
│       ├── appointment
│       ├── banner
│       ├── branch
│       ├── daily_hour
│       ├── holiday
│       ├── menu
│       ├── menu_group
│       └── weekly_hour
├── persist // データの永続化に関連するコードや設定
│   ├── keycloak // keycloakの設定ファイル
│   ├── localstack // localstackの設定ファイル
│   ├── postgres // postgresの設定ファイル
│   ├── postgres-keycloak // keycloakで使うpostgresの設定ファイル
│   └── postgres-test // テスト用のpostgresの設定ファイル
├── postman // postmanの設定ファイル
├── schema // DB定義書
├── swagger // swaggerの設定ファイル。API仕様書をGoのコードから自動生成する
└── tmp // airで起動するときに生成されるファイル
</pre>