package main

import "fmt"

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	
	fmt.Println(a, "|", b)
	fmt.Println(len(a), cap(a), "|", len(b), cap(b))
	
	b = append(b, 100)
	b = append(b, 102)
	
	fmt.Println(a, "|", b)
	fmt.Println(len(a), cap(a), "|", len(b), cap(b))
}

// будет создан срез [77, 78, 79]
// len = 3
// cap = 4

// Значения в 'a' тоже будут меняться если менять значения в 'b', 
// так как они ссылаются на одну и ту же область памяти.

// Но если добавить элементы через append, превышающие capacity (4), 
// то создастся новый базовый массив, capacity увеличится 
// (×2 если cap < 256, или +25-30% если cap ≥ 256), 
// и после этого 'a' и 'b' перестанут быть связаны.