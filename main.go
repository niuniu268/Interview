package main

import (
	"fmt"
	"math"
)

type Glass struct {
	amount float64
}

func fillWater(time float64, row, col int) float64 {

	fillTime := 10.0

	rows := row

	pyramid := make([][]*Glass, rows+1)
	for i := 0; i <= rows; i++ {
		pyramid[i] = make([]*Glass, i+1)
		for j := 0; j <= i; j++ {
			pyramid[i][j] = &Glass{}
		}
	}

	pyramid[0][0].amount = time / fillTime

	for i := 0; i < rows; i++ {
		for j := 0; j <= i; j++ {
			if pyramid[i][j].amount > 1 {
				overflow := pyramid[i][j].amount - 1
				pyramid[i][j].amount = 1
				if i+1 <= rows {
					pyramid[i+1][j].amount += overflow / 2
					if j+1 <= i+1 {
						pyramid[i+1][j+1].amount += overflow / 2
					}
				}
			}
		}
	}

	return math.Min(pyramid[row-1][col-1].amount, 1.0)
}

func main() {
	time := 136.66666
	row := 4
	col := 1

	amount := fillWater(time, row, col)
	fmt.Printf("The glass at row %d, column %d is filled: %.3f\n", row, col, amount)
}
