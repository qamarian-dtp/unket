package unket

import (
	"crypto/rand"
	"github.com/qamarian-dtp/err"
	"math/big"
	"strconv"
)

// Function New () creates a new unket. Its argument should be the number of numbers you
// want the unket to contain.
func New (n int64) (*Unket, error) {
	if n < 1 {
		return nil, err.New ("Invalid argument (number).", nil, nil)
	}

	return &Unket {map[string]string {}, big.NewInt (n)}, nil
}

type Unket struct {
	repositioned map[string]string
	noOfElement *big.Int
}

// Function Pick () helps pick a number at random. If there is no more number to pick,
// nil would be returned.
func (n Unket) Pick () (o *big.Int, e error) {
	// ..1.. {
	if n.noOfElement.String () == "0" {
		return
	}
	// ..1.. }

	// ..1.. {
	elements, okX := big.NewInt (0).SetString (n.noOfElement.String (), 0)
	if okX == false {
		e = err.New ("Bug detected. Ref: 0.", nil, nil)
		return
	}

	noOfTimes := elements.Div (elements, big.NewInt (256))

	if noOfTimes.String () == "0" {
		noOfTimes = big.NewInt (1)
	} else {
		z := big.NewInt (0)
		z = z.Mul (z, noOfTimes)

		if z.String () != n.noOfElement.String () {
			noOfTimes.Add (noOfTimes, big.NewInt (1))
		}
	}
	// ..1.. }

	totalRandNo := big.NewInt (0)

	// ..1.. {
	for i := big.NewInt (1); i.Cmp (noOfTimes) <= 0; i.Add (i, big.NewInt (1)) {
		randByte := make ([]byte, 1)
		_, errA := rand.Read (randByte)
		if errA != nil {
			e = err.New ("Unable to source rand byte.", nil, nil, errA)
			return
		}

		randNo, okB := big.NewInt (0).SetString (strconv.Itoa (int (randByte [0])),
			0)
		if okB == false {
			e = err.New ("Bug detected. Ref: 1.", nil, nil)
			return
		}

		randNo = randNo.Add (randNo, big.NewInt (1))

		totalRandNo = totalRandNo.Add (totalRandNo, randNo)
	}

	if totalRandNo.Cmp (n.noOfElement) == 1 {
		totalRandNo = totalRandNo.Mod (totalRandNo, n.noOfElement)
		totalRandNo = totalRandNo.Add (totalRandNo, big.NewInt (1))
	}
	// ..1.. }

	// ..1.. {
	no, okC := n.repositioned [totalRandNo.String ()]

	if okC == false {
		okD := true
		o, okD = big.NewInt (0).SetString (totalRandNo.String (), 0)
		if okD == false {
			e = err.New ("Bug detected. Ref: 2.", nil, nil)
			return
		}
	} else {
		okE := true
		o, okE = big.NewInt (0).SetString (no, 0)
		if okE == false {
			e = err.New ("Bug detected. Ref: 3.", nil, nil)
			return
		}
	}
	// ..1.. }

	// ..1.. {
	if totalRandNo.Cmp (n.noOfElement) != 0 {
		replacingElement := n.noOfElement.String ()

		if elem, _ := n.repositioned [replacingElement]; elem != "" {
			replacingElement = elem
		}

		n.repositioned [totalRandNo.String ()] = replacingElement
	}
	// ..1.. }

	n.noOfElement.Sub (n.noOfElement, big.NewInt (1))

	return
}
