package main

import "fmt"

func main (){

	whichday := [12]string{
		"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twefth",
	}
	
	xmaslist := [12]string {
		"partridge", "turtle doves", "french hens", "calling birds", "golden rings", "geese", "swans", "maids", "ladies", "lords", "pipers", "drummers",
	}

	twelvedays := map[string]map[string]string{
		"partridge" : map[string]string{
			"doing" : "in a pear tree",
			"number" : "and 1",
		},
		"turtle doves" : map[string]string{
			"doing" : "",
			"number" : "2",
		},
		"french hens" : map[string]string{
			"doing" : "",
			"number" : "3",
		},
		"calling birds" : map[string]string{
			"doing" : "",
			"number" : "4",
		},
		"golden rings" : map[string]string{
			"doing" : "",
			"number" : "5",
		},
		"geese" : map[string]string{
			"doing" : "a laying",
			"number" : "6",
		},
		"swans" : map[string]string{
			"doing" : "a swimming",
			"number" : "7",
		},
		"maids" : map[string]string{
			"doing" : "a milking",
			"number" : "8",
		},
		"ladies" : map[string]string{
			"doing" : "dancing",
			"number" : "9",
		},
		"lords" : map[string]string{
			"doing" : "a leaping",
			"number" : "10",
		},
		"pipers" : map[string]string{
			"doing" : "piping",
			"number" : "11",
		},
		"drummers" : map[string]string{
			"doing" : "drumming",
			"number" : "12",
		},
	}
	
	ctr := 0
	for ctr < 12 {
		fmt.Print("\nOn the ", whichday[ctr], " day of Christmas \nmy true love sent to me:\n")
		a := ctr
		if ctr == 0{
			fmt.Print("1 partridge in a pear tree\n")
		}else{
		for a >= 0 {
			presents, _ := twelvedays[xmaslist[a]]
			fmt.Print(presents["number"], " ",xmaslist[a], " ", presents["doing"],"\n")
			a -= 1
		}
		}
		ctr += 1
		fmt.Print("\n")
	}
	fmt.Print(twelvedays[xmaslist[1]])
}