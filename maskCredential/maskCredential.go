package maskcredential

import "fmt"

type ConfidentialCustomer struct {
	CustomerID int64
	CreditCard CreditCard
}

type CreditCard string

func (c CreditCard) String() string {
	return "xxxx-xxxx-xxxx-xxxx"
}

func (c CreditCard) GoString() string {
	return "xxxx-xxxx-xxxx-xxxx"
}

func Main() {
	c := ConfidentialCustomer{
		CustomerID: 123,
		CreditCard: "1234-5678-9123-4567",
	}

	fmt.Println(c)
	fmt.Printf("%#v", c)
}
