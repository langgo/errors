package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	{
		err := a()
		// fmt.Println(err)
		fmt.Printf("%+v\n", err)

		e := err.(*Error)
		fmt.Println(e.isTip)
		fmt.Println(e.tipMsg)
	}
}

func a() error {
	return b()
}

func b() error {
	return Wrap(errors.New("xx")).WithTip("参数错误").E()
}
