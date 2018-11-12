package main

import (
	"github.com/Whitewolf04/Golangstudy/UsefulTools"
	"fmt"
)

func main(){
	for _,k := range UsefulTools.NumGen(){
		for _,v := range UsefulTools.NumGen(){
			for c := range UsefulTools.NumGen(){
				if UsefulTools.Square(k) + UsefulTools.Square(v) == UsefulTools.Square(c){
					if k + v + c == 1000{
						fmt.Println(k*v*c)
					}
				}
			}
		}
	}
}

