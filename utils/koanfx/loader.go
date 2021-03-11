package koanfx

import (
	"github.com/knadh/koanf/parsers/json"
	log "github.com/sirupsen/logrus"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

var K *koanf.Koanf

func LoadConfig (){
	f := file.Provider("mock/mock.yaml")
	// Load YAML config
	if err := K.Load(f, yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	// Watch the file and get a callback on change. The callback can do whatever,
	// like re-load the configuration.
	// File provider always returns a nil `event`.
	go func() {
		f.Watch(func(event interface{}, err error) {
			if err != nil {
				log.Printf("watch error: %v", err)
				return
			}

			log.Println("config changed. Reloading ...")
			K.Load(f, json.Parser())
			K.Print()
		})

		// Block forever (and manually make a change to mock/mock.json) to
		// reload the config.
		log.Println("waiting forever. Try making a change to mock/mock.json to live reload")
		<-make(chan bool)
	}()

}

func init(){
	K = koanf.New(".")
}