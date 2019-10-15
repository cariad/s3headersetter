package main

import (
	"flag"
	"log"

	"gopkg.in/yaml.v2"
	"io/ioutil"

	"github.com/cariad/gos3headersetter"
)

type config struct {
	Rules []gos3headersetter.Rule `yaml:"rules"`
}

func loadConfig(filename string) (*config, error) {
	log.Printf("Loading configuration from \"%s\"...", filename)

	c := &config{}

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return c, err
	}
	err = yaml.Unmarshal(f, c)
	return c, err
}

func main() {
	bucketFlag := flag.String("bucket", "", "name of the bucket to update")
	prefixFlag := flag.String("key-prefix", "", "(optional) key prefix of objects to update")
	configFlag := flag.String("config", "", "configuration file name")
	flag.Parse()

	if *configFlag == "" {
		log.Fatal("-config must be specified")
	}

	if *bucketFlag == "" {
		log.Fatal("-bucket must be specified")
	}

	config, err := loadConfig(*configFlag)
	if err != nil {
		log.Fatal(err)
	}

	bucket := gos3headersetter.Bucket{
		Bucket:    *bucketFlag,
		KeyPrefix: *prefixFlag,
	}

	err = bucket.Apply(config.Rules)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Update complete.")
}
