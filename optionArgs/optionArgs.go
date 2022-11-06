package optionArgs

import "fmt"

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

// 利用サンプル
// 使い分けの観点となりそうなこと
// - 記述量
// - 変更時の作業量
// - 読みやすさ
// - 書き足す時のミスの起きづらさ
func Main() {
	// 関数
	tempuraUdon := NewKitsuneUdon(2)
	fmt.Printf("udon=%#v\n", tempuraUdon)

	// 構造体
	udonUsingStruct := NewUdonUsingStruct(Option{
		men:      0,
		aburaage: false,
		ebiten:   0,
	})
	fmt.Printf("udonUsingStruct=%#v\n", udonUsingStruct)

	// ビルダー
	udonUsingBuilder := NewUdonUsingBuilder(1).Aburaage().Eviten(5).Order()
	fmt.Printf("udonUsingBuilder=%#v\n", udonUsingBuilder)

	// Functional Option パターン
	udonUsingFunctionalOption := NewUdonUsingFunctionalOption(
		OptMen(1),
		OptAbura(),
		OptEbiten(3),
	)
	fmt.Printf("udonUsingFunctionalOption=%#v\n", udonUsingFunctionalOption)
}

// 別名の関数によるオプション引数
// 疑問: NewKakeudon(100) みたいに iota の値を超過した値を渡してもコンパイルエラーにはならないが別途バリデーションを追加する以外に防ぐ方法はないか、、、
// pro(s):
//   利用側の記述量がすくない
//   Udon の構造体のフィールドに変更があっても変更量は関数の数に比例する(利用側が多くても関数が少なければ修正は大変じゃない)
// con(s):
//  パターンが多いと定義する関数の数が爆増してメンテコストが増える
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

// 構造体を利用したオプション引数
// pro(s):
//   オプションが大量にある機能を(定義側は)比較的少ない記述量で記述できる
// con(s):
//   ゼロ値やデフォルト引数の実装がやや面倒臭い(とはいえ全然面倒臭くないきはする)
//   Option のフィールドの増減があった場合、利用側のコードを全箇所修正する必要がある
type Option struct {
	men      Portion
	aburaage bool
	ebiten   int
}

func NewUdonUsingStruct(o Option) *Udon {
	return &Udon{
		men:      o.men,
		aburaage: o.aburaage,
		ebiten:   o.ebiten,
	}
}

// ビルダーを利用したオプション引数
// pro(s
//   オプションを追加してもそのオプションを利用したい箇所だけを変更すればよいためたくさんある利用側のコードのうち一部だけを変更したいときは変更量が少なくてすむ
// con(s):
//   pros の裏返しで、変更を利用側のコード全箇所に適用する場合は変更量が多くなる
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

// Functional Option パターンを使ったオプション引数
// pro(s): Builder パターンは NewUdonUsingBuilder, order の2つのメソッドを呼ぶ必要があるのに比べて Functional Option パターンでは NewUdonUsingFunctionalOption 関数だけ呼べば同じことができるのでスッキリしていいかも
//   オプションを追加してもそのオプションを利用したい箇所だけを変更すればよいため変更量が少なくなる場合がある
// con(s):
//   pros の裏返しで、変更を利用側のコード全箇所に適用する場合は変更量が多くなる
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
