package main

import (
	"app/pkg/cluster/dbscan"
	"app/pkg/util"
	"fmt"
	"testing"
	"time"
	"unsafe"
)

func Benchmark_clustering(b *testing.B) {
	defaultImageFilePath := "data/heic0611b_6200.tif"
	img := util.LoadImage(defaultImageFilePath)

	var labels [][]int

	startTime := time.Now().UnixMilli()
	for b.Loop() {
		// Run DBScan over the image to identify separate galaxies (clusters)
		// Returns an 2d array, with one cluster id corresponding to each pixel
		labels = dbscan.DBScan(img, epsilon, minSamples, minLuminecense)
	}
	endTime := time.Now().UnixMilli()
	fmt.Printf("Elapsed: %dms\n", endTime-startTime)

	b.StopTimer()
	clusterMap := make(map[int]int)
	for _, x := range labels {
		for _, c := range x {
			if _, ok := clusterMap[c]; !ok {
				clusterMap[c] = 1
			} else {
				clusterMap[c] += 1
			}
		}
	}

	fmt.Println("Total num clusters: ", len(clusterMap))

	// Layer the new cluster image on top of the original
	// with a 50% opacity
	//output := util.Blend(img, clusters)
}

func Test_abs(t *testing.T) {
	value := 8

	mask := value >> (int(unsafe.Sizeof(int(1)))*8 - 1)
	absolute_value := (value ^ mask) - mask

	if absolute_value != 8 {
		t.Fatalf("Expected absolute value to be %d, got %d", value, absolute_value)
	}

	value = -8
	absolute_value = (value ^ mask) - mask

	if absolute_value != -8 {
		t.Fatalf("Expected absolute value to be %d, got %d", value, absolute_value)
	}
}
