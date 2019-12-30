package main

import (
	"encoding/json"
	"fmt"
	"os"

	pb "github.com/hexagram30/raster/api"
	"github.com/hexagram30/raster/client"
	"github.com/hexagram30/raster/components/config"
	"github.com/hexagram30/raster/components/logging"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Create the application object and assign components to it
	c := client.New()
	c.Config = config.New()
	c.Logger = logging.LoadClient(c.Config)
	c.SetupConnection()
	cmd := os.Args[1]
	switch cmd {
	case "ping":
		println(c.Ping())
	case "version":
		version := gRPCVersionToJSON(c.Version())
		fmt.Println(version)
	default:
		log.Errorf("Unknown command '%s'", cmd)
	}
}

func gRPCVersionToJSON(structData *pb.VersionReply) string {
	jsonData, err := json.Marshal(structData)
	if err != nil {
		log.Error(err)
		log.Fatalf("Couldn't marshal: %v", structData)
	}
	return string(jsonData)
}
