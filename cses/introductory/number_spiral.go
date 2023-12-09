func NumberSpiral(y, x int) int {
	max := max(x, y)
	startPoint := max * max
	if startPoint%2 == 0 { // even ring
		if max == x { // start from top
			return startPoint - 2*x + 1 + y
		} else {
			return startPoint - x + 1
		}
	} else { // odd ring
		if max == x { // start from top
			return startPoint - y + 1
		} else {
			return startPoint - 2*y + 1 + x
		}
	}
}

// OR

// func main() {
// 	var t int
// 	fmt.Scan(&t)
// 	for i := 0; i < t; i++ {
// 		var x, y int
// 		fmt.Scan(&y)
// 		fmt.Scan(&x)
// 		max := max(x, y)
// 		corner := 0
// 		for j := 0; j < max; j++ {
// 			corner = corner + 2*j
// 		}
// 		var a int
// 		if max%2 == 0 {
// 			if max == x {
// 				a = corner - (x - y) + 1
// 			} else {
// 				a = corner + (y - x) + 1
// 			}
// 		} else {
// 			if max == x {
// 				a = corner + (x - y) + 1
// 			} else {
// 				a = corner - (y - x) + 1
// 			}
// 		}
// 		fmt.Println(a)
// 	}
// }
