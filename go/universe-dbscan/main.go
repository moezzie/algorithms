package main

import (
	"app/pkg/cluster/dbscan"
	"app/pkg/display"
	"app/pkg/util"
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

const (
	epsilon        = 2 // max distance
	minSamples     = 2 // minimum neighbors to be concidered a 'core point'
	minLuminecense = 50
)

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	defaultImageFilePath := "data/heic0611b_6200.tif"
	imageFilePath := flag.String("image", defaultImageFilePath, "Path of the image to cluster")
	displayImage := flag.Bool("display", false, "Wether to display the image after finishing")

	flag.Parse()

	// Start the pprof server in a separate goroutine

	img := util.LoadImage(*imageFilePath)

	bounds := img.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y

	fmt.Printf("Image dimensions: %dx%d\n", width, height)

	// Run DBScan over the image to identify separate galaxies (clusters)
	// Returns an 2d array, with one cluster id corresponding to each pixel
	labels := dbscan.DBScan(img, epsilon, minSamples, minLuminecense)

	// Convert the 2d array of cluster ids into
	// a picture where each cluster has its own color
	clusters := util.ConvertLabelsToImage(labels)

	// Layer the new cluster image on top of the original
	// with a 50% opacity
	output := util.Blend(img, clusters)

	if *displayImage {
		display.Display(output)
	}
}
