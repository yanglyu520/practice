package main

//func main() {
//	reverseArray1(a)
//	reverseArrayInplace(a)
//}

func reverseArray1(a []int32) []int32 {
	res := make([]int32, len(a))
	for i, _ := range a {
		res[len(a)-1-i] = a[i]
	}
	return res
}

func reverseArrayInplace(a []int32) []int32 {
	for i, _ := range a {
		a[i], a[len(a)-1] = a[len(a)-1], a[i]
	}
	return a
}
