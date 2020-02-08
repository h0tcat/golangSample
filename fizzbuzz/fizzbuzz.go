package main

import "fmt"

func main() {
  min:= 0
  max:= 9999999
  for ;min<max;min++{
	  go fizzbuzz(&min,&max)
	}
}

func fizzbuzz(min *int, max *int) {
	go fmt.Printf("%d\n",*min);
	
}
