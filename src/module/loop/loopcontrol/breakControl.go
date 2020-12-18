package loopcontrol

func BreakFor(){
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 6 {
				break //现在终止的是j 循环，而不是i的那个
			}
			//fmt.Println(i)
			print(i)
		}
		print("\n")
	}
}
