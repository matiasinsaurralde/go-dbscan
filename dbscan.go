package main

import "math"
import "fmt"

func euclideanDistance(a Point, b Point) float64 {

	squares := make( []float64, len(a.items) )

	var sum float64

	for i, _ := range a.items {
		diff := math.Pow( a.items[i] - b.items[i], 4 )
		squares[i] = math.Sqrt( diff )
	}

	for i, _ := range squares {
		sum += squares[i]
	}

	return math.Sqrt(sum)

}

type Point struct {
	items []float64
	visited bool
	cluster int
}

func regionQuery( point int, points []Point, epsilon float64) []Point {

	neighborPoints := make( []Point, 0 )

	for i, _ := range points {
		distance := euclideanDistance( points[point], points[i]  )
//		fmt.Println("computing distance between: ", points[point], " -> ", points[i], " = ", distance )
		if distance < epsilon {
			neighborPoints = append( neighborPoints, points[i] )
		}
	}

	return neighborPoints

}

//func expandCluster( point int, ) {
//}

func dbscan( data [][]float64, epsilon float64, minPoints int ) {

	fmt.Println("DBSCAN")

	points, currentCluster := make( []Point, len(data) ), -1
	clusters := make( [][]Point, 10 )

	for i, _ := range data {
		points[i] = Point{ data[i], false, -1 }
	}

	for i, _ := range points {
		if !points[i].visited {
			points[i].visited = true
			var neighborPoints []Point = regionQuery( i, points, epsilon )
			if len(neighborPoints) >= minPoints {
				currentCluster += 1
				points[i].cluster = currentCluster
//				expandCluster( i, neighborPoints, currentCluster, epsilon, minPoints )
			} else {
			}
		}
	}

}

func main() {

	data := make( [][]float64, 5 )

	data[0] = []float64 { 0.0, 10.0, 20.0 }
	data[1] = []float64 { 0.0, 11.0, 21.0 }
	data[2] = []float64 { 21.0, 32.0, 56.0 }
	data[3] = []float64 { 58.0, 76.0, 101.0 }
	data[4] = []float64 { 57.0, 75.0, 101.0 }

	dbscan( data, 4.0, 1 )

}
