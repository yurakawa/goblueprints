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
以下をそれぞれ別のウインドウから実行する

    nsqlookupd
    nsqd  --lookupd-tcp-address=127.0.0.1:4160
    mongod --dbpath ./db
    
    # mongoコマンドを実行後
    ballotsデータベースのpollsコレクションに新たな項目を追加する。
    > use ballots
    > db.polls.insert({"title":"今の気分は?","options":["happy","sad","fail","win"]})
    Ctrl + C

    // 不要? nsq_tail  --topic="votes" --lookupd-http-address=localhost:4161
    
    ./twittervotes
    ./counter

