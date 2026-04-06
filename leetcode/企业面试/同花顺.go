package main

import (
	"fmt"
	// "time"
)

func main(){
	var m Model
	if m == nil { fmt.Println("nil") } 
	var l *LLM
	m = l // l is nil
	if m == nil { fmt.Println("nil") } else { fmt.Println("not nil") }
}

type Model interface { Generate() string }
type LLM struct {}
func (l *LLM) Generate() string { return "Hi" }

// func main() {  
//     ch := make(chan int)  
//     go func() {  
//         ch <- 1  
//     }()  
//     x := <-ch  
//     fmt.Println(x)  
// }