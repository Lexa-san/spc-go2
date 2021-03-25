package models

import "math"

//Quadratic model...
type Quadratic struct {
	A      int `json:"A"`
	B      int `json:"B"`
	C      int `json:"C"`
	Nroots int `json:"Nroots"`
}

//Solve quadratic and return quantity of roots.
func (q *Quadratic) Solve() (int, error) {
	var (
		a  float64
		b  float64
		c  float64
		D  float64
		x1 float64
		x2 float64
	)

	a = float64(q.A)
	b = float64(q.B)
	c = float64(q.C)

	if a == 0 && b == 0 {
		//fmt.Println("корней нет")
		return 0, nil
	}

	if a == 0 {
		//fmt.Println("один корень")
		return 1, nil
	}

	D = math.Pow(b, 2) - 4*a*c

	if D < 0 {
		//fmt.Println("корней нет")
		return 0, nil
	}

	if D == 0 {
		//fmt.Println("один корень")
		return 1, nil
	}

	if D > 0 {
		x1 = (-b + math.Sqrt(D)) / 2 / a
		x2 = (-b - math.Sqrt(D)) / 2 / a
		if x1 == x2 {
			//fmt.Println("один корень")
			return 1, nil
		}

		//fmt.Println("два корня")
		return 2, nil
	}

	return 0, nil
}
