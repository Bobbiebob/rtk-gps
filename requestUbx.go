package main

import (
	"errors"

	"github.com/daedaleanai/ublox/ubx"
)

func cfgInf1() ubx.Message {
	return ubx.CfgInf1{}
}

func logInfo() ubx.Message {
	return ubx.LogInfo{}
}

func cfgMessage() ubx.Message {
	return ubx.CfgMsg{}
}

func cfgNavx5() ubx.Message {
	return ubx.CfgNavx5{}
}

func cfgPortUsb(port string) ubx.Message {
	toreturn := ubx.CfgPrt{}
	if port == "USB" {
		toreturn.PortID = 3
	}
	if port == "I2C" {
		toreturn.PortID = 0
	}
	if port == "UART1" {
		toreturn.PortID = 1
	}
	if port == "UART2" {
		toreturn.PortID = 2
	}
	return toreturn
}

func cfgTp5() ubx.Message {
	return ubx.CfgTp5{}
}

func esfAlgIMUAlignment() ubx.Message {
	return ubx.EsfAlg{}
}

func esfIns() ubx.Message {
	return ubx.EsfIns{}
}

func hnrAtt() ubx.Message {
	return ubx.HnrAtt{}
}

func hnrIns() ubx.Message {
	return ubx.HnrIns{}
}

func hnrPvt() ubx.Message {
	return ubx.HnrPvt{}
}

func logBatch() ubx.Message {
	return ubx.LogRetrievebatch{}
}

func cfgInfo() ubx.Message {
	toReturn := ubx.CfgInf1{}
	toReturn.ProtocolID = 3
	return toReturn
}

func usbInfo() ubx.Message {
	toReturn := ubx.CfgMsg1{MsgClass: 0x06, MsgID: 0x00}
	return toReturn
}

func cfgRst(set string) (ubx.Message, error) {
	toReturn := ubx.CfgRst{}
	switch set {
	case "now":
		toReturn.ResetMode = 0x00
		return toReturn, nil
	case "software":
		toReturn.ResetMode = 0x01
		return toReturn, nil
	case "gnss":
		toReturn.ResetMode = 0x02
		return toReturn, nil
	case "after":
		toReturn.ResetMode = 0x04
		return toReturn, nil
	case "gnssStop":
		toReturn.ResetMode = 0x08
		return toReturn, nil
	case "gnssStart":
		toReturn.ResetMode = 0x09
		return toReturn, nil
	default:
		return nil, errors.New("wrong input type")
	}
}

/*
info: turn on survey-in with given settings
requires: time in seconds and acc in mm
*/
func timeMode3(time uint32, acc uint32) ubx.Message {
	toReturn := ubx.CfgTmode3{}
	toReturn.SvinAccLimit = acc * 10
	var mode = ubx.CfgTmode3Mode
	mode = 1
	toReturn.Flags = mode
	toReturn.SvinMinDur_s = time
	return toReturn
}

/*
info: turning on given NAV message
requires: Message wich needs to be turned on
*/
func turnOnNav(topic string) (ubx.Message, error) {
	switch topic {

	case "CLOCK":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x22, Rate: 1}
		return toReturn, nil
	case "DOP":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x04, Rate: 1}
		return toReturn, nil
	case "EOE":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x61, Rate: 1}
		return toReturn, nil
	case "GEOFENCE":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x39, Rate: 1}
		return toReturn, nil
	case "HPPOSECEF":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x13, Rate: 1}
		return toReturn, nil
	case "HPPOSLLH":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x14, Rate: 1}
		return toReturn, nil
	case "ODO":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x09, Rate: 1}
		return toReturn, nil
	case "ORB":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x34, Rate: 1}
		return toReturn, nil
	case "POSECEF":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x01, Rate: 1}
		return toReturn, nil
	case "POSLLH":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x02, Rate: 1}
		return toReturn, nil
	case "PVT":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x07, Rate: 1}
		return toReturn, nil
	case "RELPOSNED":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x3c, Rate: 1}
		return toReturn, nil
	case "SAT":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x35, Rate: 1}
		return toReturn, nil
	case "SBAS":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x32, Rate: 1}
		return toReturn, nil
	case "SIG":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x43, Rate: 1}
		return toReturn, nil
	case "SLAS":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x42, Rate: 1}
		return toReturn, nil
	case "STATUS":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x03, Rate: 1}
		return toReturn, nil
	case "SVIN":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x3b, Rate: 1}
		return toReturn, nil
	case "TIMEBDS":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x24, Rate: 1}
		return toReturn, nil
	case "TIMEGAL":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x25, Rate: 1}
		return toReturn, nil
	case "TIMEGLO":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x23, Rate: 1}
		return toReturn, nil
	case "TIMELS":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x26, Rate: 1}
		return toReturn, nil
	case "TIMEQZSS":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x27, Rate: 1}
		return toReturn, nil
	case "TIMEUTC":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x21, Rate: 1}
		return toReturn, nil
	case "VELECEF":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x11, Rate: 1}
		return toReturn, nil
	case "VELNED":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x12, Rate: 1}
		return toReturn, nil
	default:
		return nil, errors.New("wrong input type")
	}
}

/*
info: turning off given NAV message
requires: Message wich needs to be turned off
*/
func turnOffNav(topic string) (ubx.Message, error) {
	switch topic {
	case "CLOCK":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x22, Rate: 0}
		return toReturn, nil
	case "DOP":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x04, Rate: 0}
		return toReturn, nil
	case "EOE":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x61, Rate: 0}
		return toReturn, nil
	case "GEOFENCE":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x39, Rate: 0}
		return toReturn, nil
	case "HPPOSECEF":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x13, Rate: 0}
		return toReturn, nil
	case "HPPOSLLH":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x14, Rate: 0}
		return toReturn, nil
	case "ODO":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x09, Rate: 0}
		return toReturn, nil
	case "ORB":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x34, Rate: 0}
		return toReturn, nil
	case "POSECEF":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x01, Rate: 0}
		return toReturn, nil
	case "POSLLH":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x02, Rate: 0}
		return toReturn, nil
	case "PVT":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x07, Rate: 0}
		return toReturn, nil
	case "RELPOSNED":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x3c, Rate: 0}
		return toReturn, nil
	case "SAT":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x35, Rate: 0}
		return toReturn, nil
	case "SBAS":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x32, Rate: 0}
		return toReturn, nil
	case "SIG":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x43, Rate: 0}
		return toReturn, nil
	case "SLAS":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x42, Rate: 0}
		return toReturn, nil
	case "STATUS":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x03, Rate: 0}
		return toReturn, nil
	case "SVIN":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x3b, Rate: 0}
		return toReturn, nil
	case "TIMEBDS":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x24, Rate: 0}
		return toReturn, nil
	case "TIMEGAL":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x25, Rate: 0}
		return toReturn, nil
	case "TIMEGLO":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x23, Rate: 0}
		return toReturn, nil
	case "TIMELS":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x26, Rate: 0}
		return toReturn, nil
	case "TIMEQZSS":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x27, Rate: 0}
		return toReturn, nil
	case "TIMEUTC":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x21, Rate: 0}
		return toReturn, nil
	case "VELECEF":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x11, Rate: 0}
		return toReturn, nil
	case "VELNED":
		toReturn := ubx.CfgMsg1{MsgClass: 0x01, MsgID: 0x12, Rate: 0}
		return toReturn, nil
	default:
		return nil, errors.New("wrong input type")
	}
}

/*
info: turning on given RXM Message
requires: Message wich needs to be turned on
*/
func turnOnRxm(topic string) (ubx.Message, error) {
	switch topic {
	case "MEASX":
		toReturn := ubx.CfgMsg1{MsgClass: 0x02, MsgID: 0x14, Rate: 1}
		return toReturn, nil
	case "PMREQ":
		toReturn := ubx.CfgMsg1{MsgClass: 0x02, MsgID: 0x14, Rate: 1}
		return toReturn, nil
	case "RAWX":
		toReturn := ubx.CfgMsg1{MsgClass: 0x02, MsgID: 0x14, Rate: 1}
		return toReturn, nil
	case "RLM":
		toReturn := ubx.CfgMsg1{MsgClass: 0x02, MsgID: 0x14, Rate: 1}
		return toReturn, nil
	case "RTCM":
		toReturn := ubx.CfgMsg1{MsgClass: 0x02, MsgID: 0x14, Rate: 1}
		return toReturn, nil
	case "SFRBX":
		toReturn := ubx.CfgMsg1{MsgClass: 0x02, MsgID: 0x14, Rate: 1}
		return toReturn, nil
	default:
		return nil, errors.New("wrong input type")
	}
}

/*
info: turning off given RXM Message
requires: Message wich needs to be turned off
*/
func turnOffRxm(topic string) (ubx.Message, error) {
	switch topic {
	case "MEASX":
		toReturn := ubx.CfgMsg1{MsgClass: 0x02, MsgID: 0x14, Rate: 0}
		return toReturn, nil
	case "PMREQ":
		toReturn := ubx.CfgMsg1{MsgClass: 0x02, MsgID: 0x14, Rate: 0}
		return toReturn, nil
	case "RAWX":
		toReturn := ubx.CfgMsg1{MsgClass: 0x02, MsgID: 0x14, Rate: 0}
		return toReturn, nil
	case "RLM":
		toReturn := ubx.CfgMsg1{MsgClass: 0x02, MsgID: 0x14, Rate: 0}
		return toReturn, nil
	case "RTCM":
		toReturn := ubx.CfgMsg1{MsgClass: 0x02, MsgID: 0x14, Rate: 0}
		return toReturn, nil
	case "SFRBX":
		toReturn := ubx.CfgMsg1{MsgClass: 0x02, MsgID: 0x14, Rate: 0}
		return toReturn, nil
	default:
		return nil, errors.New("wrong input type")
	}
}
