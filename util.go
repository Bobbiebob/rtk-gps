package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/daedaleanai/ublox"
	"github.com/daedaleanai/ublox/ubx"
	"github.com/tarm/serial"
)

/*
info: read the data from the port, assuming it is ubx, decode it
requires: a pointer to the opened port and a string containting the port which has been opened
*/
func read(port *serial.Port, tty string) {
	//goRoutineChannels <- +1
	log.Println("reading port: " + tty)
	go decodeUbxStream(port, gpsDataChannel)
}

/*
info: Read the reply that is expected after sendin an message. The repley will be displayed
requirs: pointer to opened serial port
*/
func readconcurrent(s *serial.Port) {
	reader := io.Reader(s)
	decoder := ublox.NewDecoder(reader)
	data2, _ := decoder.Decode()

	if data2 != nil {
		reflected := reflect.TypeOf(data2).String()
		switch reflected {
		case "*ubx.CfgMsg2":
			cfgMsg2 := data2.(*ubx.CfgMsg2)
			log.Println("message class = ", cfgMsg2.MsgClass)
			log.Println("message id = ", cfgMsg2.MsgID)
			log.Println("message rate = ", cfgMsg2.Rate)
		case "*ubx.CfgMsg1":
			fmt.Printf("%#v", data2.(*ubx.CfgMsg1))
		case "*ubx.AckAck":
			log.Println("setting acknowledged")
		case "*ubx.AckNak":
			log.Println(string(colorRed), "setting not set", "\n", string(colorReset))
		default:
			log.Println(reflected)
			log.Println("non dedocded message recieved")
		}
		fmt.Println("")
	} else {
		log.Println("unkown data recieved")
	}
}

/*
info: Writes a request message for info to the given port
requires: a pointer to the opened port and a message.
*/
func write(message ubx.Message, s *serial.Port) {
	s.Flush()

	//write a message and expact a reply
	encode, err := ubx.Encode(message)
	fatalErrorDisplay(err)
	n, err := s.Write(encode)
	_ = n
	fatalErrorDisplay(err)

	readconcurrent(s)
	time.Sleep(time.Millisecond * 250)
}

/*
info: prints al the contend from the navMap in an overview
requires: a map where key = string and data = *GpsData, allong with the currend time to only print the disered message
*/
func printNavMap(item map[string]*GpsData, time string) {
	for key, val := range item {
		if key == time {
			fmt.Println("key: ", key,
				//"val: ", *val.Pvt, "; ", *val.Svin,
				", accuracy: ", val.Svin.accuracy,
				", fixtype: ", val.fixType,
				", dataPoints", val.dataPoints,
				", long", val.long,
				", lat", val.lat,
				", mena sea level", val.msl,
				", height", val.height,

				//*val.Svin
			)
		}
	}
}

/*
info: load the config file
return: struct Config
*/
func loadConfig() Config {
	jsonFile, err := os.Open("config.json")
	fatalErrorDisplay(err)
	defer jsonFile.Close()

	content, err := ioutil.ReadAll(jsonFile)
	fatalErrorDisplay(err)
	var config Config
	err = json.Unmarshal(content, &config)
	fatalErrorDisplay(err)

	return config
}

/*
info: standard out for a log.Fatal message. Makes code easier to read.
requires: an error message to be printed.
*/
func fatalErrorDisplay(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

/*
info: standard out for a log.Fatal message but with fmt so the code will go on. Makes code easier to read
requires: an error message to be printed
*/
func fmtErrorDisplay(err error) {
	if err != nil {
		log.Println(err)
	}
}
