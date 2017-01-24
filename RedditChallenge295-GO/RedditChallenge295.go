package main

import "fmt"
import "os"
import "bufio"

func replaceAtIndex(in1 string, in2 string, i int) string {
    //mylength = len(in2)
	returnstring := []rune(in1)
	compstring := []rune(in2)
	a := 0
	for a <= i {
		returnstring[a] = compstring[a]
		a += 1
	}
	return string(returnstring)
}

func main(){
	//get input text to change
	reader1 := bufio.NewReader(os.Stdin)
	fmt.Print("This program will rearrange a word or sentence letter by letter.\nInput and output must have equal length.\n")	
	fmt.Print("Enter text to transofrm: ")
	starttext, _ := reader1.ReadString('\n')
	fmt.Println(starttext)
	
	//get desired output
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter desired output: ")
	endtext, _ := reader2.ReadString('\n')
	fmt.Println(endtext)
	
	//Call our runes replace at index function to change string compare string length. length(input) must equal length(output)
	lengthstart := len(starttext)
	lengthend := len(endtext)
	if lengthstart == lengthend {
		fmt.Print("both strings are equal in length, we will now proceed\n")
		ctr := 0
		for ctr < lengthend-2 {   //I think the need to subtract two is driven by '\n' addition in reader.
			output := replaceAtIndex(starttext, endtext, ctr)
			fmt.Print(output)
			ctr += 1
		}
	}else{
		fmt.Print("strings are not equal")
	}
	
}