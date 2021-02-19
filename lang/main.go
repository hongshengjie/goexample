package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	f, _ := ioutil.ReadFile("t")
	line := bytes.Split(f, []byte("\n"))
	for _, v := range line {

		k := strings.Split(string(v), ",")
		fmt.Printf(`grpcDebug -addr=172.21.41.229:9000 -data='{"mid":%s,"order_id":%s}' -method=/live.xroomfeed.v1.AnchorCpm/StopSpread`, k[1], k[0])
		fmt.Println()
	}

}
