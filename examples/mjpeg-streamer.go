// What it does:
//
// This example opens a video capture device, then streams MJPEG from it.
// Once running point your browser to the hostname/port you passed in the
// command line (for example http://localhost:8080) and you should see 
// the live video stream.
//
// How to run:
//
// mjpeg-streamer [camera ID] [host:port]
//
// 		go run ./examples/mjpeg-streamer.go 1 0.0.0.0:8080
//
// +build example

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	opencv3 ".."
	"github.com/saljam/mjpeg"
)

var (
	deviceID int
	err      error
	webcam   opencv3.VideoCapture
	img      opencv3.Mat

	stream *mjpeg.Stream
)

func capture() {
	for {
		if ok := webcam.Read(img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		buf, _ := opencv3.IMEncode(".jpg", img)
		stream.UpdateJPEG(buf)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("How to run:\n\tmjpeg-streamer [camera ID] [host:port]")
		return
	}

	// parse args
	deviceID, _ = strconv.Atoi(os.Args[1])
	host := os.Args[2]

	// open webcam
	webcam, err = opencv3.VideoCaptureDevice(int(deviceID))
	if err != nil {
		fmt.Printf("error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	// prepare image matrix
	img = opencv3.NewMat()
	defer img.Close()

	// create the mjpeg stream
	stream = mjpeg.NewStream()

	// start capturing
	go capture()

	// start http server
	http.Handle("/", stream)
	log.Fatal(http.ListenAndServe(host, nil))
}
