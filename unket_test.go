package unket

import (
	"fmt"
	"github.com/qamarian-dtp/err"
	errLib "github.com/qamarian-lib/err"
	"gopkg.in/qamarian-lib/str.v3"
	"strconv"
	"testing"
)

func TestUnket_Pick (t *testing.T) {
	str.PrintEtr ("Test started...", "std", "TestUnket_Pick ()")
	fmt.Println ()

	// ..1.. {
	n := 8
	u, errX := New (int64 (n))
	if errX != nil {
		e := err.New ("Test failed. Ref: 0.", nil, nil, errX)
		str.PrintEtr (errLib.Fup (e), "err", "TestUnket_Pick ()")
		t.FailNow ()
	}
	// ..1.. }

	// ..1.. {
	for i := 1; i <= n; i ++ {
		randNo, errY := u.Pick ()
		if errY != nil {
			e := err.New ("Test failed. Ref: 1; Iteration: " +
				strconv.Itoa (i), nil, nil, errY)
			str.PrintEtr (errLib.Fup (e), "err", "TestUnket_Pick ()")
			t.FailNow ()
		}

		if randNo == nil {
			str.PrintEtr ("Nil data at iteration: " + strconv.Itoa (i), "std",
				"TestUnket_Pick ()")
			continue
		}

		o := fmt.Sprintf ("%s; ", randNo.String ())
		fmt.Print (o)
	}
	// ..1.. }

	fmt.Println ()
	str.PrintEtr ("Test passed!", "std", "TestUnket_Pick ()")
}
