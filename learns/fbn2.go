package main
import "fmt"


func fibtwo(n int, tlist []int) []int{

	var thelen int
	var bfa int
	var bfb int
	if n < 1{
		return tlist
	}
	thelen = len(tlist)
	bfa = thelen-1
	bfb = thelen-2
	tlist = append(tlist, tlist[bfb:bfa][0] + tlist[bfa:][0])
	n--
	return fibtwo(n, tlist)
}

func main()  {
	tlist := []int{0, 1}
	n := 20
	fmt.Println(fibtwo(n, tlist))
}