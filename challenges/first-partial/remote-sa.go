package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"math"
)

type Point struct {
	X, Y float64
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//generatePoints array
func generatePoints(s string) ([]Point, error) {

	points := []Point{}

	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	vals := strings.Split(s, ",")
	if len(vals) < 2 {
		return []Point{}, fmt.Errorf("Point [%v] was not well defined", s)
	}

	var x, y float64

	for idx, val := range vals {

		if idx%2 == 0 {
			x, _ = strconv.ParseFloat(val, 64)
		} else {
			y, _ = strconv.ParseFloat(val, 64)
			points = append(points, Point{x, y})
		}
	}
	return points, nil
}

// getArea gets the area inside from a given shape
func getArea(points []Point) float64 {
	// Your code goes here
	var caso =len(points)
	var areaShape float64 =0

	//Depending on the number of points the area that it will have
	switch caso {
    case 3:
        for i := 0; i <1; i++ {
          areaShape = 0.5*(math.Abs(float64(
			  //Gaussian area formula for 3 coordinates
			  ((points[i].X*points[i+1].Y)+(points[i+1].X*points[i+2].Y)+(points[i+2].X*points[i].Y)-
			  (points[i+1].X*points[i].Y)-(points[i+2].X*points[i+1].Y)-(points[i].X*points[i+2].Y)))))
		}
		fmt.Println(areaShape)

	case 4:
		for i := 0; i <1; i++ {
			areaShape = 0.5*(math.Abs(float64(
				//Gaussian area formula for 4 coordinates
				((points[i].X*points[i+1].Y)+(points[i+1].X*points[i+2].Y)+(points[i+2].X*points[i+3].Y)+(points[i+3].X*points[i+0].Y)-
				(points[i+1].X*points[i].Y)-(points[i+2].X*points[i+1].Y)-(points[i+3].X*points[i+2].Y)-(points[i].X*points[i+3].Y)))))
		}
		fmt.Println(areaShape)

	case 5:
        for i := 0; i <1; i++ {
			areaShape = 0.5*(math.Abs(float64(
				//Gaussian area formula for 5 coordinates
				((points[i].X*points[i+1].Y)+(points[i+1].X*points[i+2].Y)+(points[i+2].X*points[i+3].Y)+(points[i+3].X*points[i+4].Y)+(points[i+4].X*points[i].Y)-
				(points[i+1].X*points[i].Y)-(points[i+2].X*points[i+1].Y)-(points[i+3].X*points[i+2].Y)-(points[i+4].X*points[i+3].Y)-(points[i].X*points[i+4].Y)))))
		}
		fmt.Println(areaShape)

  	}
	return areaShape
	
	
}

// getPerimeter gets the perimeter from a given array of connected points
func getPerimeter(points []Point) float64 {
	// Your code goes here
	var perimeterShape float64 =0 
	//Get all of the points and calculate the distance from one to another
	for i := 0; i < len(points)-1; i++ {
		perimeterShape +=  math.Sqrt(math.Pow((points[i+1].X-points[i].X),2) + math.Pow((points[i+1].Y-points[i].Y),2))
	}
	perimeterShape +=  math.Sqrt(math.Pow((points[len(points)-1].X-points[0].X),2) + math.Pow((points[len(points)-1].Y-points[0].Y),2))
	return perimeterShape
}

// handler handles the web request and reponds it
func handler(w http.ResponseWriter, r *http.Request) {

	var vertices []Point
	for k, v := range r.URL.Query() {
		if k == "vertices" {
			points, err := generatePoints(v[0])
			if err != nil {
				fmt.Fprintf(w, fmt.Sprintf("error: %v", err))
				return
			}
			vertices = points
			break
		}
	}

	// Results gathering
	area := getArea(vertices)
	perimeter := getPerimeter(vertices)

	// Logging in the server side
	log.Printf("Received vertices array: %v", vertices)

	// Response construction
	response := fmt.Sprintf("Welcome Friend to the Remote Shapes Analyzer\n")
	response += fmt.Sprintf(" - Your figure has : [%v] vertices\n", len(vertices))
	if len(vertices)>2{
		response += fmt.Sprintf(" - Vertices        : %v\n", vertices)
		response += fmt.Sprintf(" - Perimeter       : %v\n", perimeter)
		response += fmt.Sprintf(" - Area            : %v\n", area)
	}else{
		response += fmt.Sprintf("ERROR - Your shape is not compliying with the minimum number of vertices.")
	}

	// Send response to client
	fmt.Fprintf(w, response)
}
