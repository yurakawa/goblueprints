# 第8章で学ぶこと

- 複数のパッケージやコマンドラインツールを含むプロジェクトの構成
- シンプルなデータを永続化し、ツールの実行のたびに参照できるようにするための現実的なアプローチ
- osパッケージを使ったファイルシステムとのインタラクション
- コードを実行し続け、Ctrl＋Cが押されたら終了するための方法
- filepath.Walkを使った、すべてのファイルとフォルダーへのアクセス
- フォルダー内のファイルの内容が変化したことを迅速に検出する方法
- archive/zipパッケージによるファイルの圧縮
- コマンドラインフラグと通常の引数の組み合わせを考慮したツールの作成

# システムの設計

- ファイルのスナップショットを定期的に作成し、ソースコードを含むプロジェクトへの変更を記録します。
- 変更の有無をチェックする間隔を変更できます。
- 主にテキストベースのプロジェクトをZIP圧縮するので、アーカイブされたファイルのサイズはとても小さくなります。
- ビルドは早期に行いつつ、将来的な改善の可能性を検討します。
- 実装上の判断は容易に修正できるようにし、今後の変更に備えます。
- 2つのコマンドラインツールを作成します。1つは実際の処理を行うバックエンドのデーモンで、もう1つはバックアップ対象のパスの一覧表示や追加と削除を行うユーザー向けのユーティリティです。

# プロジェクト構造

```
backup/   (パッケージ)
 `-- cmds /
      |-- backup/ (ユーザー向けツール)
      `-- backupd/ (デーモン)
```

# backupd
### ビルド

        cd ./chapter8/backup/cmds/backup
        go build -o backup

### 実行例
- 追加

        ./backup -db=./backupdata add ./test ./test2
        ./backup -db=./backupdata add ./test3

- リスト

        ./backup -db=./backupdata list

- 削除

        ./backup -db=./backupdata remove ./test3
        ./backup -db=./backupdata list


# backupd
### ビルド

        cd ./chapter8/backup/cmds/backupd
        go build -o backupd

### 実行例

        ./backupd -db="../backup/backupdata/" -archive="./archive" -interval=5

