package main

import (
	"errors"

	"github.com/daedaleanai/ublox/ubx"
)

func turnOnNmea(topic string) (ubx.Message, error) {
	switch topic {
	case "DTM":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x0a, Rate: 1}
		return toReturn, nil

	case "GBS":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x09, Rate: 1}
		return toReturn, nil
	case "GGA":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x00, Rate: 1}
		return toReturn, nil
	case "GLL":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x01, Rate: 1}
		return toReturn, nil

	case "GNS":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x0d, Rate: 1}
		return toReturn, nil

	case "GRS":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x06, Rate: 1}
		return toReturn, nil
	case "GSA":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x02, Rate: 1}
		return toReturn, nil
	case "GST":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x07, Rate: 1}
		return toReturn, nil
	case "GSV":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x03, Rate: 1}
		return toReturn, nil
	case "RMC":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x04, Rate: 1}
		return toReturn, nil

	case "VLW":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x0f, Rate: 1}
		return toReturn, nil
	case "VTG":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x05, Rate: 1}
		return toReturn, nil
	case "ZDA":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x08, Rate: 1}
		return toReturn, nil
	default:
		return nil, errors.New("wrong input type")
	}
}

func turnOffNmea(topic string) (ubx.Message, error) {
	switch topic {
	case "DTM":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x0a, Rate: 0}
		return toReturn, nil

	case "GBS":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x09, Rate: 0}
		return toReturn, nil
	case "GGA":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x00, Rate: 0}
		return toReturn, nil
	case "GLL":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x01, Rate: 0}
		return toReturn, nil
	case "GNS":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x0d, Rate: 0}
		return toReturn, nil
	case "GRS":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x06, Rate: 0}
		return toReturn, nil
	case "GSA":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x02, Rate: 0}
		return toReturn, nil
	case "GST":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x07, Rate: 0}
		return toReturn, nil
	case "GSV":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x03, Rate: 0}
		return toReturn, nil
	case "RMC":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x04, Rate: 0}
		return toReturn, nil
	case "VLW":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x0f, Rate: 0}
		return toReturn, nil
	case "VTG":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x05, Rate: 0}
		return toReturn, nil
	case "ZDA":
		toReturn := ubx.CfgMsg1{MsgClass: 0xf0, MsgID: 0x08, Rate: 0}
		return toReturn, nil
	default:
		return nil, errors.New("wrong input type")
	}
}
