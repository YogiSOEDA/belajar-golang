package main

import "fmt"

func main()  {
	var i int = 0

	// for i := 0; i<5; i++{
	// 	fmt.Println(i)
	// }

	//ganti style perulangan

	for {
		fmt.Println(i)
		i++

		if i>5 {
			break
		}
		continue
	}
}