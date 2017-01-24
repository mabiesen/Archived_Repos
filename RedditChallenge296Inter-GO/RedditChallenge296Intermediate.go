package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "sort"
import "strings"
import "math"

func infertwopoints(x1, y1, x2, y2 float64)(float64, float64, float64, float64){
	x3 := x1
	y3 := y2
	x4 := x2
	y4 := y1
	
	return x3, y3, x4, y4
}

type Rectangle struct {
	x1 float64
	y1 float64 
	x2 float64
	y2 float64
	x3 float64
	y3 float64
	x4 float64
	y4 float64
	top float64
	bottom float64
	left float64
	right float64
}

func (r *Rectangle) findtop() float64{
	temparray := []float64{r.y1, r.y2, r.y3, r.y4,}
	sort.Float64s(temparray)
	return temparray[3]
}

func (r *Rectangle) findbottom() float64{
	temparray := []float64{r.y1,r.y2,r.y3,r.y4}
	sort.Float64s(temparray)
	return temparray[0]
}

func (r *Rectangle) findleft() float64{
	temparray := []float64{r.x1,r.x2,r.x3,r.x4}
	sort.Float64s(temparray)
	return temparray[0]
}

func (r *Rectangle) findright() float64{
	temparray := []float64{r.x1,r.x2,r.x3,r.x4}
	sort.Float64s(temparray)
	return temparray[3]
}



func main(){
//user prompts to help understand the program
	fmt.Println("This program will determine the area of two overlapping rectangles")
	fmt.Println("Your input is required.  Please provide the startpoint and endpoint coordinates for each rectangle.\n")
	fmt.Println("The starting point and end point can be any two opposing corners of the rectangle")
	fmt.Println("You will now be prompted to provide inputs for the rectangles.\n\n")
	
//read user input from console	

	z:= "" 
	
	//read the first rectangle
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Rectangle1, starting x: ")
	z, _  = reader.ReadString('\n')
	xcoord1rect1, _ := strconv.ParseFloat(strings.TrimSpace(z), 64)	
	reader = bufio.NewReader(os.Stdin)	
	fmt.Print("Rectangle1, starting y: ")
	z, _ = reader.ReadString('\n')	
	ycoord1rect1, _ := strconv.ParseFloat(strings.TrimSpace(z), 64)
	reader = bufio.NewReader(os.Stdin)		
	fmt.Print("Rectangle1, ending x: ")
	z, _ = reader.ReadString('\n')		
	xcoord2rect1, _ := strconv.ParseFloat(strings.TrimSpace(z), 64)
	reader = bufio.NewReader(os.Stdin)	
	fmt.Print("Rectangle1, ending y: ")
	z, _ = reader.ReadString('\n')
	ycoord2rect1, _ := strconv.ParseFloat(strings.TrimSpace(z), 64)
	
	//read the second rectangle
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Rectangle2, starting x: ")
	z, _ = reader.ReadString('\n')
	xcoord1rect2, _ := strconv.ParseFloat(strings.TrimSpace(z), 64)
	reader = bufio.NewReader(os.Stdin)	
	fmt.Print("Rectangle2, starting y: ")
	z, _ = reader.ReadString('\n')	
	ycoord1rect2, _ := strconv.ParseFloat(strings.TrimSpace(z), 64)
	reader = bufio.NewReader(os.Stdin)		
	fmt.Print("Rectangle2, ending x: ")
	z, _ = reader.ReadString('\n')		
	xcoord2rect2, _ := strconv.ParseFloat(strings.TrimSpace(z), 64)
	reader = bufio.NewReader(os.Stdin)	
	fmt.Print("Rectangle2, ending y: ")
	z, _ = reader.ReadString('\n')
	ycoord2rect2, _ := strconv.ParseFloat(strings.TrimSpace(z), 64)
	
	//gather remaining points for the rectangle
	xcoord3rect1, ycoord3rect1, xcoord4rect1, ycoord4rect1 := infertwopoints(xcoord1rect1, ycoord1rect1, xcoord2rect1, ycoord2rect1)		
	xcoord3rect2, ycoord3rect2, xcoord4rect2, ycoord4rect2 := infertwopoints(xcoord1rect2, ycoord1rect2, xcoord2rect2, ycoord2rect2)
	
	//create the rectangle structures
	r1 := Rectangle{xcoord1rect1, ycoord1rect1, xcoord2rect1, ycoord2rect1, xcoord3rect1, ycoord3rect1, xcoord4rect1, ycoord4rect1, 0, 0, 0, 0}
	r2 := Rectangle{xcoord1rect2, ycoord1rect2, xcoord2rect2, ycoord2rect2, xcoord3rect2, ycoord3rect2, xcoord4rect2, ycoord4rect2, 0, 0, 0, 0}
	
	//find furthest points	
	r1.left = r1.findleft()
	r1.right = r1.findright()
	r1.bottom = r1.findbottom()
	r1.top = r1.findtop()
	
	r2.left = r2.findleft()
	r2.right = r2.findright()
	r2.bottom = r2.findbottom()
	r2.top = r2.findtop()
	
	x_overlap := math.Max(0, math.Min(r2.right, r1.right) - math.Max(r2.left, r1.left))
	y_overlap := math.Max(0, math.Min(r1.top, r2.top) - math.Max(r1.bottom, r2.bottom))
	overlaparea := x_overlap * y_overlap
	
	
	//Test

	fmt.Print("the area of overlap is: ", overlaparea)
}

