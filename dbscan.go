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

func regionQuery( point Point, points []Point, epsilon float64 ) []Point {
	neighborPoints := make( []Point, 0 )
	for i, _ := range points {
		distance := euclideanDistance( point, points[i] )
		if ( distance < epsilon && distance != 0 ) {
			neighborPoints = append( neighborPoints, points[i] )
		}
	}
	return neighborPoints
}

func expandCluster( point Point, neighborPoints []Point, points []Point, currentClusterN int, epsilon float64,  minPoints int ) []Point {
	clusterPoints := make( []Point, 0 )

	for i, _ := range neighborPoints {
		if( !neighborPoints[i].visited ) {
			neighborPoints[i].visited = true
			newPoints := make( []Point, 0 )
			newPoints = regionQuery( point, points, epsilon )
			if( len( newPoints) >= minPoints ) {
				for n, _ := range newPoints {
					neighborPoints = append( neighborPoints, newPoints[n] )
				}
			}
		}

		if( neighborPoints[i].cluster == 0 ) {
			clusterPoints = append( clusterPoints, neighborPoints[i] )
			neighborPoints[i].cluster = currentClusterN
		}
	}

	return clusterPoints
}

func dbscan( data [][]float64, epsilon float64, minPoints int ) [][]Point {

	fmt.Println("DBSCAN")

	points := make( []Point, len(data) )
        currentClusterN := -1
	clusters := make( [][]Point, 1 )

	for i, _ := range data {
		points[i] = Point{ data[i], false, 0 }
	}

        for i, _ := range points {
		if( !points[i].visited ) {

			points[i].visited = true

			var neighborPoints []Point = regionQuery( points[i], points, epsilon )

			if len(neighborPoints) >= minPoints {
				currentClusterN += 1
				points[i].cluster = currentClusterN

				//cluster := make( []Point, 0 )
				//cluster = append( cluster, points[i] )
				//clusters[ currentClusterN ] = cluster

				var cluster = make( []Point, 0 )
				cluster = append( cluster, points[i] )
				cluster = expandCluster( points[i], neighborPoints, points, currentClusterN, epsilon, minPoints )

				clusters = append( clusters, cluster )

				//fmt.Println( clusters )
			} else {
				clusters[ 0 ] = append( clusters[0], points[i] )
			}
		}

//		fmt.Println( regionQuery( points[i], points, epsilon ) )
	}

        //fmt.Println( clusters )
	//fmt.Println( regionQuery( points[3], points, epsilon ) )

	return clusters

}

func main() {

	data := make( [][]float64, 11 )


	data[0] = []float64 { 0.0, 10.0, 20.0 }
	data[1] = []float64 { 0.0, 11.0, 21.0 }
	data[2] = []float64 { 0.0, 12.0, 20.0 }
	data[3] = []float64 { 20.0, 33.0, 59.0 }
	data[4] = []float64 { 21.0, 32.0, 56.0 }
	data[5] = []float64 { 59.0, 77.0, 101.0 }
	data[6] = []float64 { 58.0, 79.0, 100.0 }
	data[7] = []float64 { 58.0, 76.0, 102.0 }
	data[8] = []float64 { 300.0, 70.0, 20.0 }
        data[9] = []float64 { 500.0, 300.0, 202.0 }
        data[10] = []float64 {  500.0, 302.0, 204.0 }

	var output [][]Point = dbscan( data, 4.0, 2 )

	for i, _ := range output {
		fmt.Println( "Cluster #", i )
		for e, _ := range output[i] {
			fmt.Println( output[i][e] )
		}
	}

}
