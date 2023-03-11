package main

import "fmt"

// 4, 9, 16, 25, 36, 64, 100
func main() {
	fmt.Println(mySqrt(0))
}

func mySqrt(x int) int {
	var sqrt int

	if x >= 1 && 3 >= x {
		return 1
	}

	for i := 1; i <= x/2; i++ {
		if i*i > x {
			break
		}
		if i*i == x {
			sqrt = i
			break
		}
		sqrt = i
	}
	return sqrt
}

/*
func mySqrt(x int) int {
    i, j := 0, x
    for i <= j {
        mid := i + (j-i)/2
        squared := mid*mid
        if squared == x {
            return mid
        } else if squared < x {
            i = mid + 1
        } else {
            j = mid - 1
        }
    }
    return j
}
*/
/*
func mySqrt(x int) int {
	if x == 1 {
		return 1
	}

	s, l := 0, x
	for {
		m := (s + l) / 2
		if m*m > x {
			l = m
		} else {
			if (m+1)*(m+1) > x {
				return m
			}
			s = m
		}
	}
}
*/
