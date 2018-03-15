# 第1章で学ぶこと

- net/httpパッケージを使い、HTTPリクエストに応答します。
- テンプレートを使ったコンテンツをクライアントに提供します。
- Goのhttp.Handler型のインタフェースに適合させます。
- Goのgoroutineを使い、アプリケーションが複数の作業を同時に行えるようにします。
- チャネルを使い、実行中の各goroutine間で情報を共有します。
- HTTPリクエストをアップグレードし、WebSocketなどの新しい機能を使えるようにします。
- アプリケーション内部でのはたらきをよりよく理解するために、ログの記録を行います。
- テスト駆動型の手順に従って、Goの完全なパッケージ構造を作成します。
- 非公開の型を、公開されているインタフェース型で返します。

# 実行方法

       go build -o chat
       ./chat
       