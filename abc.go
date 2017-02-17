//package main;import("os";"fmt");func main(){b:=[3]int{65,66,67};for _,e:=range os.Args[1]{b[e-65]=1};fmt.Print(string(b[0])+string(b[1])+string(b[2]))}
//http://codegolf.stackexchange.com/questions/110219/as-easy-as-a-b-c/110409#110409

package main

import (
	"os"
	"fmt"
)

func main() {
	//A,B,C are represented as 65,66,67 in int
	b := [3]int{65, 66, 67}
	//Loop through the ABC argument.
	for _,e := range os.Args[1]{
		//Find the letter that exists. If it's A then b[65-65] -> b[0] -> A
		//We replace the integer, with 1 which is equal to an empty string if converted.
		b[e-65] = 1
	}
	//Covert each int to string.
	fmt.Print(string(b[0])+string(b[1])+string(b[2]))
}