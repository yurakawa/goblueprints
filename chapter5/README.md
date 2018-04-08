# 第5章で学ぶこと

- NoSQLの分散データベース。特に、MongoDBとのインタラクション
- 分散メッセージキュー。特に、Bit.lyのNSQとgo-nsqパッケージを使ったイベントのパブリッシュとサブスクライブ
- TwitterのストリーミングAPIを使ったツイートのライブ表示や、長期間のネットワーク接続の管理
- 内部で多くのgoroutineを実行しているプログラムを、適切に終了させる方法
- メモリ消費量の少ないチャネルを使ってシグナルを送受信する方法


# 環境構築
https://github.com/matryer/goblueprints/blob/master/chapter5/README.md

## nsq
### インストール
    brew install nsq
### 起動確認
    nsqd
### go 用のnsqドライバ
    go get github.com/bitly/go-nsq
    
## mongodb
### インストール
    brew install mongodb
### 起動確認
    mongod --dbpath /tmp
### go 用のMongoDBドライバ
    go get gopkg.in/mgo.v2

## 実行環境の起動
nsqlookupdを起動し、nsqdのTCPポート(デフォルトは4160)を確認する

    nsqlookupd
    
nsqdを起動し、どのnsqlookupdを利用するか指定する

    nsqd --lookupd-tcp-address=localhost:4160

mongodを起動してデータ関連のサービスを実行する

    mongod --dbpath ./db
