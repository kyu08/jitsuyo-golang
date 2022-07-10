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
