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

