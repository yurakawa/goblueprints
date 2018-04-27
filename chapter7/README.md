# 第7章で学ぶこと

- アジャイルの考え方に基づいて、短くシンプルなユーザーストーリーを通じてプロジェクトの目標を表現する方法
- API設計について意見の一致を得て、多くの人々が同時進行で作業を行うという手順
- 初期バージョンのコードにデータ（フィクスチャーとも呼ばれます）を埋め込んでコンパイルし、後で実装の変更が必要になってもAPIに影響を与えないようにするための方法
- 構造体などの型を公開し、内部的な表現については隠蔽または変形するという設計方針
- 入れ子状のデータを埋め込みの構造体として表現し、同時に型のインタフェースをシンプルに保つ方法
- 外部のAPIにリクエストを行うためのhttp.Get。具体的には、コードを肥大化させずにGoogle Places APIにアクセスする方法
- Goでは定義されていない列挙型を、効率的に実装する方法
- TDD（テスト駆動開発）の実際的な例
- math/randパッケージを使い、スライスの中から1つの項目をランダムに選ぶための簡単な方法
- http.Request型の値の中からURLパラメーターを簡単に取り出す方法

# 開発するもの
ユーザがしたいことについて、お出かけの種類と位置情報に基づいてランダムにおすすめを提示する

### ユーザへ提供するストーリー
- 1つ目のストーリー

|  ユーザーの立場 | 旅行者 |
|  ------ | ------ |
|  行いたいこと | 複数の種類のお出かけが提示される |
|  目的 | どのような種類のお出かけに仲間を連れてゆくか決める |

- 2つ目のストーリー

|  ユーザーの立場 | 旅行者 |
|  ------ | ------ |
|  行いたいこと | 自分が選んだ種類のお出かけについて、ランダムなおすすめが提示される |
|  目的 | 向かうべき場所と、イベントの内容を知る |

### エンドポイント
- Google Place APIへのリクエストURL例
` GET /recommendations?lat=1&lng=2&journey=bar|cafe&radius=10&cost=$...$$$$$ `

lat(latitude): 緯度
lng()longitude): 軽度
radius: おすすめを取得する範囲を中心からの半径としてメートル単位で指定
const: 施設での費用価格帯($$$$が最も高価)

- レスポンス

```json
[
  {
    name: "Romantic",
    journey: "park|bar|movie_theater|restaurant|florist"
  },
  {
    name: "Shopping",
    journey: "department_store|clothing_store|jewelry_store"
  }
]
```

### meander リクエスト例
http://localhost:8080/recommendations?lat=51.520707&lng=-0.153809&radius=5000&journey=cafe|bar|casino|restaurant&cost=$...$$$

### GooGle Places API ドキュメント
https://developers.google.com/places/documentation/supported_types

### geo 確認
http://mygeoposition.com


