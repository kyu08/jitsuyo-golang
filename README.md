# 1章 「goらしさ」に触れる
## 1.1.1 変数名
- 変数名が複数の単語で構成されるとき、_を使用せず、MixedCapsとするのが慣習となっている。(`max_length`ではなく`maxLength`) 
- 頭文字は`Url`や`Id`などとせずに`URL`,`ID`とするか`url`,`id`とする。他の単語と組み合わせるときは`ServeHTTP`や`stackID`のように大文字にする
- エラーの変数名、つまりErrorインタフェースを満たすError型としてふるまう型の名前には接尾語Errorをつけるのが慣例。`MarshalerError`, `UnsupportedTypeError`のようにする
- 同様に`errors.New("my error")`と宣言されるようなエラーの変数はErrまたはerrから始まる

## 1.1.2 パッケージ名
- パッケージ名は簡潔でパッケージの内容を想像しやすい名前が適切。また慣例としてパッケージ名は小文字で構成される1つの単語とされる。
- 望ましくないパッケージ名はたとえばutil, common, apiなどである。この名前からはこのパッケージがどのようなものなのか明らかではない。
- 複数の用途で用いられるパッケージの場合はencoding/json, encoding/xmlのようにフォルダを分、末端のパッケージ名はpackage json, package xml のようにするのが一般的
- また、priority_queueのようなスネークケースやArrayListのようなキャメルケースを用いたパッケージ名はgoらしくない。複数の名詞で構成したくなった場合は上記のencodingで例示したようにフォルダを分ける

## 1.1.3 インタフェース名
- 1つのメソッドのみを持つインタフェースでは、`er`という接尾辞がついた名前が用いられることがある。具体的には`Read`メソッドのみを持つ`io.Reader`や`String`メソッドのみを持つ`fmt.Stringer`のようなものである。(英語として正しくない場合もある)

## 1.1.4 レシーバー名
- レシーバーの型を反映した名前をつけるのが望ましい
- 通常は1,2文字で十分


## 1.3.2 iotaの挙動
- iotaは最初に使われた行から行が変わるたびにインクリメントされる(空行は除く)
- constブロックが複数ある場合はブロックごとにiotaのカウンターがリセットされて0に戻る

```go
const (
	a = iota + 1 // 1
	b            // 2
	c            // 3
	_            // 4だが使われない
	e            // 5
	f = iota     // 6 iotaの値はリセットされない
)

const (
	g = iota // 再び0になる
	h        // 1
	i        // 2
)
```

## 1.6 関数のオプション引数
- うどん屋のオプションを例に以下の`NewUdon`関数をいかに使いやすくできるかがここでのテーマ
- 現状の実装だとオプションが増えた時に全利用箇所を修正する必要がある

```go
type Portion int

const (
	Regular Portion = iota
	Small
	Large
)

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   int
}

func NewUdon(p Portion, aburaage bool, ebiten int) Udon {
	return Udon{
		men:      p,
		aburaage: aburaage,
		ebiten:   ebiten,
	}
}
```

### 1.6.1 別名の関数によるオプション引数
- よく使われる呼び出し方法があらかじめわかっている場合には一番簡単に使うことができ、コードも読みやすいのがこの方法

```go
func NewKakeudon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: false,
		ebiten:   0,
	}
}

func NewKitsuneUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: true,
		ebiten:   0,
	}
}

func NewTempuraUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: false,
		ebiten:   3,
	}
}
```
- pros
  - Udon構造体のフィールドに変更があっても変更量は関数の数に比例する(利用側が多くても関数が少なければ修正は大変じゃない)
  - 利用側の記述量が少なくて済む
- cons
  - パターンが多いと定義する関数の数が爆増してメンテコストが増える

### 1.6.2 構造体を利用したオプション引数
- Goの標準ライブラリでよくみられるオプション指定用の構造体を用いるパターン

```go
type Option struct {
	men      Portion
	aburaage bool
	ebiten   int
}

func NewUdonUsingStruct(o Option) *Udon {
  // 朝食時間は海老天1本無料
  if o.ebiten == 0 && time.Now().Hour() < 10 {
    o.ebiten = 1
  }
	return &Udon{
		men:      o.men,
		aburaage: o.aburaage,
		ebiten:   o.ebiten,
	}
}
```

- pros
  - オプションが大量にある機能を(定義側は)比較的少ない記述量で記述できる
  - ロジックの追加・変更が容易にできる
- cons
  - ゼロ値やデフォルト引数の実装がやや面倒臭い(とはいえ全然面倒臭くないきはする)
  - Option のフィールドの増減があった場合、利用側のコードを全箇所修正する必要がある

### 1.6.3 ビルダーを利用したオプション引数
- ビルダーパターン

```go
type fluentOpt struct {
	men      Portion
	aburaage bool
	ebiten   int
}

func NewUdonUsingBuilder(p Portion) *fluentOpt {
	// デフォルトはコンストラクタで定義
	return &fluentOpt{
		men:      p,
		aburaage: false,
		ebiten:   1,
	}
}

func (o *fluentOpt) Aburaage() *fluentOpt {
	o.aburaage = true
	return o
}

func (o *fluentOpt) Eviten(n int) *fluentOpt {
	o.ebiten = n
	return o
}

func (o *fluentOpt) Order() *Udon {
	return &Udon{
		men:      o.men,
		aburaage: false,
		ebiten:   0,
	}
}
```
- pros
  - オプションを追加してもそのオプションを利用したい箇所だけを変更すればよいためたくさんある利用側のコードのうち一部だけを変更したいときは変更量が少なくてすむ
- cons
  - pros の裏返しで、変更を利用側のコード全箇所に適用する場合は変更量が多くなる

### 1.6.4 Functional Optionパターンを使ったオプション引数
- ビルダーパターンはオプションごとに関数を作成するが、こちらのパターンでは構造体のメソッドにせず、独立した関数にするのが大きな違い
```go
type OptFunc func(r *Udon)

func NewUdonUsingFunctionalOption(opts ...OptFunc) *Udon {
	r := &Udon{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func OptMen(p Portion) OptFunc {
	return func(r *Udon) { r.men = p }
}

func OptAbura() OptFunc {
	return func(r *Udon) { r.aburaage = true }
}

func OptEbiten(e int) OptFunc {
	return func(r *Udon) { r.ebiten = e }
}
```

- pros
  - Builder パターンは NewUdonUsingBuilder, order の2つのメソッドを呼ぶ必要があるのに比べて Functional Option パターンでは NewUdonUsingFunctionalOption 関数だけ呼べば同じことができるのでスッキリしていいかも
  - オプションを追加してもそのオプションを利用したい箇所だけを変更すればよいため変更量が少なくなる場合がある
- cons
  - pros の裏返しで、変更を利用側のコード全箇所に適用する場合は変更量が多くなる

### 利用側
```go
// 利用側
func Main() {
  udon := NewUdon(Large, false, 2)

	// 関数
	tempuraUdon := NewKitsuneUdon(2)

	// 構造体
	udonUsingStruct := NewUdonUsingStruct(Option{
		men:      0,
		aburaage: false,
		ebiten:   0,
	})

	// ビルダー
	udonUsingBuilder := NewUdonUsingBuilder(1).Aburaage().Eviten(5).Order()

	// Functional Option パターン
	udonUsingFunctionalOption := NewUdonUsingFunctionalOption(
		OptMen(1),
		OptAbura(),
		OptEbiten(3),
	)
}
```

### 1.6.5 どの実装方法を洗濯すべきか
- おすすめはコード量の少ない構造体パターンをまず実装して提供すること
- ビルダーパターンや Functional Optionパターンはそれを土台に必要になってから実装すればよい
- 「比較的シンプルな方法がおすすめなら他のものを紹介する必要はなかったのでは」と思われるかもしれないが、ライブラリを利用する立場だといろんなパターンを使うことがあるのでそれぞれ理解しておくのは大事

## defer の注意点
`Close()` など、エラー処理のメソッドによってはエラーを返すケースもある
この場合、普通に defer で呼ぶだけではエラーを取りこぼしてしまうため、無名関数で括ってそのエラーを名前付き返り値に代入すると呼び出しもとに返すことができる

```go
func deferReturnSample(fname string) (err error) {
  var f *os.File
  f, err = os.Create(fname)
  if err != nil {
    return fmt.Errorf("file open error %w", err)
  }
  
  defer func() {
    err = f.Close()
  }()

  io.WriteString(f, "defer error sample")
  return
}
```

## 型
go だとこれはちゃんとコンパイル通らない

```go
type MyInt int64
var int1 int64 = 1
var myint MyInt = int1
```

### 型変換
```go

var i int

type ErrorCode int

var e ErrorCode

i = e // error
e = i // error


e = ErrorCode(i)
```

# 構造体
以下のように書くと構造体を埋め込んで共通部分を使いまわすことができる
```go

type struct Book {
  ID string
}

type struct AmazonBook {
  Book
  isPrimeBook bool
}
```

これは型同名のフィールドが宣言されたのと同じように振る舞う

```go
type struct AmazonBook {
  Book Book
  isPrimeBook bool
}

amazonBook := AmazonBook{
  Book: Book{
    ID: "id",
  },
  isPrimeBook: true,
}

// 呼び出すときはこんなかんじ
id := amazonBook.Book.ID
```

## タグを使って構造体にメタデータを埋め込む
タグの記法はこう
フィールド名 型 `json:"field"`

文法的にはどのような文字列もタグとして扱えるが、監修的には上記のように書くことがおおい

## 構造体を設計する際のポイント
###  ポインタ型として扱うかどうか
内部にスライスやmap,ポインタなどの参照型の要素を持っている場合には基本的にポインタ型で扱うようにする
なぜなら構造体をコピーすると複数の構造体が同一のポインタを参照する状態ができてしまい、1箇所での変更が全体に波及してしまいバグの温床になりえるから。

# インターフェース
## 型アサーション

`io.closer` への型アサーションが成功した場合のみ`io.closer` のメソッドである `Close()` を呼び出す実装

```go
if c, ok := r.(io.Closer); ok {
  c.Close()
}
```

# エラーハンドリング
エラーに含まれる文字列を特定の文字列と比較する方法はアンチパターンである
特定のエラーに応じたハンドリングが必要な場合は `errors.Is()`や`errors.As()` を使ってハンドリングする


エラーをログ出力する際はどのような処理なのか、どのような引数を元に動いて、どんなエラーが発生したのか明確にわかるようんいエラーメッセージを記述するようにしよう
```go
user, err := getInvitedUserWithEmail(ctx, email) if err != nil {
// 呼び出し先で発生したエラーをラップし、付加情報を付与して呼び出し元に返却
return fmt.Errorf("fail to get invited user with email(%s): %w", email, err) }
```
