package main

import (
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/daedaleanai/ublox"
	"github.com/daedaleanai/ublox/ubx"
	"github.com/tarm/serial"
)

/*
info: decode each message recieved in the serial port. After decodeing all the data is stored in a map. See gpsData.go for possible content
requires: pointer to opened port where ubx data is recieved.
*/
func decodeUbxStream(stream *serial.Port, channel chan startServerChannel) {
	stop := false
	navMap := map[string]*GpsData{}
	var timer string

	reader := io.Reader(stream)
	data := ublox.NewDecoder(reader)

	for {
		data2, err := data.Decode()
		if err != nil {
			fmtErrorDisplay(err)
		} else {
			log.Println(reflect.TypeOf(data2).String())
			timer = (time.Now().Format("01-02 15:04:05"))
			switch reflect.TypeOf(data2).String() {
			case "*ubx.NavPvt":
				navPvt := data2.(*ubx.NavPvt)
				var flags [4]bool //fix, avai, date, time

				switch navPvt.FixType {
				case 0:
					navMap[timer] = &GpsData{&Pvt{fixType: "not fixed"}, &Svin{}}
				case 1:
					navMap[timer] = &GpsData{&Pvt{fixType: "Dead reckoning only"}, &Svin{}}
				case 2:
					navMap[timer] = &GpsData{&Pvt{fixType: "n2D-fix"}, &Svin{}}
				case 3:
					navMap[timer] = &GpsData{&Pvt{fixType: "3D-fix"}, &Svin{}}
				case 4:
					navMap[timer] = &GpsData{&Pvt{fixType: "GNSS + Dead reckoning combinded"}, &Svin{}}
				case 5:
					navMap[timer] = &GpsData{&Pvt{fixType: "Time only fix"}, &Svin{}}
				default:
					navMap[timer] = &GpsData{&Pvt{fixType: "error non valid fixtype"}, &Svin{}}
				}

				if strings.Contains(navPvt.Flags.String(), "ConfirmedData") {
					flags[1] = true
				} else {
					flags[1] = false
				}
				if strings.Contains(navPvt.Flags.String(), "confirmedTime") {
					flags[2] = true
				} else {
					flags[2] = false
				}
				if strings.Contains(navPvt.Flags.String(), "ConfirmedAvai") {
					flags[3] = true
				} else {
					flags[3] = false
				}
				navMap[timer] = &GpsData{&Pvt{long: navPvt.Lon_dege7, lat: navPvt.Lat_dege7, height: navPvt.Height_mm,
					msl: navPvt.HMSL_mm, gnssFixOk: flags[0], confirmedDate: flags[1], confirmedTime: flags[2],
					confirmedAvai: flags[3]}, &Svin{}}
			case "*ubx.NavSvin":
				navSvin := data2.(*ubx.NavSvin)
				acc := strconv.FormatFloat((float64(navSvin.MeanAcc)/10000), 'f', -1, 64) + "m"
				navMap[timer] = &GpsData{&Pvt{}, &Svin{accuracy: acc, valid: navSvin.Valid,
					dataPoints: navSvin.Obs}}

				if navMap[timer] != nil {
					if navMap[timer].valid == 1 {
						fmt.Println("Base accuracy acquired. Accuracy is: ", navMap[timer].accuracy, "mm")
						fmt.Println("starting server")
						channel <- startServerChannel{fix: true}
						stop = true
					} else {
						fmt.Println("Base accuracy is: ", navMap[timer].accuracy, "mm")
						fmt.Println("Datapoints used: ", navMap[timer].dataPoints)
						log.Println("Base is not sure of its location yet")
					}
				}
			default:
				printNavMap(navMap, timer)
				fmt.Println("recieved message is of type: ", reflect.TypeOf(data2).String(), "support for this message is not added (yet)")
			}

			fmt.Println("")
		}
		if stop {
			break
		}
	}
}

/*
info: decode a single ubx message.
requires: a byte array containing an ubx message
*/
func decodeUbx(buff []byte) {
	data, err := ubx.Decode(buff)
	if err != nil {
		//log.Fatal(err)
		fmt.Println("error: no data to be decoded")

	} else {
		fmt.Println("data: ", data)
	}
}

/*
info: write a ubx request to the serial port. ubx module will respond with the data linked to the request.
requires: a message formate from the ubx package and a port to send the request to.
*/
func writeUbx(stream *serial.Port, content ubx.Message) {
	fmt.Println("requesting: ", content)

	buff := make([]byte, 1024)

	encoded, err := ubx.Encode(ubx.NavPvt{})
	fatalErrorDisplay(err)
	fmt.Println("encoded: ", reflect.ValueOf(encoded))
	n, err := stream.Write(encoded)
	fatalErrorDisplay(err)
	fmt.Println("n:", n)

	n, err = stream.Read(buff)
	fmtErrorDisplay(err)
	fmt.Println("n:", n)

	decoded, err := ubx.Decode(buff[:n])
	fmtErrorDisplay(err)
	fmt.Println("decoded:", decoded)

	data := reflect.TypeOf(decoded)
	fmt.Println(data, decoded)

	for _, item := range buff[:n] {
		fmt.Println(item)
	}
	fmt.Println("end of read and write")

	fmt.Println(ubx.Decode(encoded))
	closePort(stream)
}
