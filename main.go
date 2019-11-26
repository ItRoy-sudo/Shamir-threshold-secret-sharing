package main

import "fmt"

func CreateMiyao(n int,number int,threshold int,note int,sushu int) int {
	var y =0
	var s int
	for i:=1;i<=threshold-1;i++{
		s=number-i
		for j:=1;j<=i;j++  {
			s=s*n
		}
		y=y+s
	}
	y=(y+note)%sushu
	return y
}

func main()  {
	var number int
	var threshold int
	var note int
	var sushu =23
	var n = 1
	fmt.Print("请输入份数：")
	fmt.Scanln(&number)
	fmt.Print("请输入门限：")
	fmt.Scanln(&threshold)
	fmt.Print("请输入秘密：")
	fmt.Scanln(&note)
	m1:=make(map[int]int,5)
	for i:=1;i<=number;i++ {
		m1[i]=CreateMiyao(n,number,threshold,note,sushu)
		n++
	}
	fmt.Println("子密钥为：",m1)

	var c int
	var x int
	var y int
	var S int
	var sum =0
	m2:=make(map[int]int,5)
	s2:=make([]int,1)
	fmt.Print("请输入子密钥个数：")
	fmt.Scanln(&c)
	for i:=1;i<=c;i++ {
		fmt.Print("输入子密钥x：")
		fmt.Scanln(&x)
		fmt.Print("输入子密钥y：")
		fmt.Scanln(&y)
		s2=append(s2,x)
		m2[x]=y
	}
	for i:=1;i<=c;i++ {
		var m float64=1
		for j:=1;j<=c;j++{
			if j!=i {
				m=m*float64((0-s2[j]))/float64((s2[j]-s2[i]))
			}
		}
		sum=sum+m2[i]*int(m)
	}
	S=sum%sushu
	if S<0 {
		S=S+sushu
	}
	fmt.Println("秘密为：",S)
}
