package beater

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"io/ioutil"
	"encoding/json"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
	"github.com/verticle-io/apexbeat/config"
)

var abt *Apexbeat

type Apexbeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
}


// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {

	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Apexbeat{
		done: make(chan struct{}),
		config: config,
	}
	abt = bt;

	return bt, nil
}


// Runs the beater
func (bt *Apexbeat) Run(b *beat.Beat) error {
	logp.Info("apexbeat is running (port " + bt.config.Port + ")! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()

	go func() {
		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/collector/metrics", CollectorMetrics)
		log.Fatal(http.ListenAndServe(":" + bt.config.Port, router))
	}()

	// loop
	for {
		select {
		case <-bt.done:
			return nil
		}
	}
}


// Stops the beater
func (bt *Apexbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}


// process collector/metrics URL
func CollectorMetrics(w http.ResponseWriter, r *http.Request) {

	type ApexMetricMessage struct {
		meta    map[string]string
		metrics map[string]string
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var f interface{}
	json.Unmarshal([]byte(body), &f)
	fmt.Printf("%+v\n", f)

	event := common.MapStr{
		"@timestamp": common.Time(time.Now()),
		"type":       "metricmessage",
		"meta":       f.(map[string]interface{})["meta"],
		"metrics":    f.(map[string]interface{})["metrics"],
	}
	abt.client.PublishEvent(event)

	fmt.Fprintf(w, "OK")
	logp.Info("Event sent")

}


// function returns a channel
func getChannel() chan string {
	b := make(chan string)
	return b
}
