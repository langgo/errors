package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	{
		err := a()
		fmt.Printf("==========================: %%s\n")
		fmt.Printf("%s\n", err)

		fmt.Printf("==========================: %%v\n")
		fmt.Printf("%v\n", err)

		fmt.Printf("==========================: %%+v\n")
		fmt.Printf("%+v\n", err)

		e := err.(*Error)
		fmt.Printf("==========================: isTip\n")
		fmt.Printf("%v\n", e.isTip)

		fmt.Printf("==========================: tipMsg\n")
		fmt.Printf("%v\n", e.tipMsg)
	}
}

func a() error {
	return b()
}

func b() error {
	return Wrap(errors.New("xx")).WithTip("参数错误").E()
}
