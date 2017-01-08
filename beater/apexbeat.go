package beater

import (
	"fmt"
  "html"
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
//var messages chan string

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

func (bt *Apexbeat) Run(b *beat.Beat) error {
	logp.Info("apexbeat is running! Hit CTRL-C to stop it.")

	//messages = getChannel()
	bt.client = b.Publisher.Connect()

	//go Publish(bt)

	go func(){
		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/", Index)
		log.Fatal(http.ListenAndServe(":8080", router))
	}()
/*
	go func(){
		msg := <-messages
		fmt.Println("MSG: " + msg)
		}()

	ticker := time.NewTicker(bt.config.Period)
	counter := 1
*/
	for {
		select {
		case <-bt.done:
			return nil
		}
	}


	//return nil
}


func (bt *Apexbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

/*

func Publish(bt *Apexbeat){
	logp.Info("Publish() called")

	event := common.MapStr{
		"@timestamp": common.Time(time.Now()),
		"type":       "test",
		"message":    "trsasdasdas",
	}
	bt.client.PublishEvent(event)
	logp.Info("Event sent")



}
*/



func Index(w http.ResponseWriter, r *http.Request) {
	// messages <- "channel message"





	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

	body,err := ioutil.ReadAll(r.Body)
  fmt.Fprintf(w, "Body:, %q", body,err)
	type ApexMetricMessage struct {
	    meta map[string]string
			metrics map[string]string
	}

	var f interface{}
	//metricsMessage := &ApexMetricMessage{}
  json.Unmarshal([]byte(body), &f)
	fmt.Printf("%+v\n", f)
	//logp.Info(fmt.Printf("%+v\n", stats))
	event := common.MapStr{
		"@timestamp": common.Time(time.Now()),
		"type":       "test2",
		"meta":    f.(map[string]interface{})["meta"],
		"metrics":	 f.(map[string]interface{})["metrics"],
	}
	abt.client.PublishEvent(event)
	logp.Info("Event sent")


}


// function returns a channel
func getChannel() chan string {
     b := make(chan string)
     return b
}
