package optionArgs

import "fmt"

type Portion int

const (
	Regular Portion = iota
	Small
	Large
)

func Main() {
	var tempuraUdon = NewKitsuneUdon(2)
	fmt.Printf("udon=%#v\n", tempuraUdon)

	var udonUsingStruct = NewUdonUsingStruct(Option{
		men:      0,
		aburaage: false,
		ebiten:   0,
	})
	fmt.Printf("udonUsingStruct=%#v\n", udonUsingStruct)

	var udonUsingBuilder = NewUdonUsingBuilder(1).Aburaage().Eviten(5).Order()
	fmt.Printf("udonUsingBuilder=%#v\n", udonUsingBuilder)
}

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

// 別名の関数によるオプション引数
// 疑問: NewKakeudon(100) みたいに iota の値を超過した値を渡してもコンパイルエラーにはならないが別途バリデーションを追加する以外に防ぐ方法はないか、、、
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
//   比較的少ない記述量でオプションが大量にある機能を記述できる
// con(s):
//   ゼロ値やデフォルト引数の実装がやや面倒臭い(とはいえ全然面倒臭くないきはする)
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
// pro(s):
//   オプションを追加してもそのオプションを利用したい箇所だけを変更すればよいため変更量が少なくなる場合がある
// con(s):
//
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
