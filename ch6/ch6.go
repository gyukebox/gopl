package main

import (
	"fmt"
)

type studentExamScore struct {
	math    int
	english int
	science int
}

func main() {
	gyu := &studentExamScore{80, 45, 90}
	fmt.Println(gyu.avg())
	fmt.Println(gyu.max())
}

func (score *studentExamScore) avg() float64 {
	return float64((score.math + score.english + score.science)) / 3
}

func (score *studentExamScore) max() int {
	if score.math > score.english {
		if score.math > score.science {
			return score.math
		}
		return score.science
	} else if score.english > score.science {
		return score.english
	} else {
		return score.science
	}
}
