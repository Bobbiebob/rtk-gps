package main

import (
	"errors"

	"github.com/daedaleanai/ublox/ubx"
)

/*
info: turning on given RTCM message
requires: Message wich needs to be turned on
*/
func turnOnRTCM(topic string) (ubx.Message, error) {
	switch topic {
	case "1005":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x05, Rate: 1}, nil
	case "1006":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x06, Rate: 1}, nil

	case "1009":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x09, Rate: 1}, nil
	case "1010":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x0a, Rate: 1}, nil

	case "1074":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x4a, Rate: 1}, nil
	case "1077":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x4d, Rate: 1}, nil

	case "1084":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x54, Rate: 1}, nil
	case "1087":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x57, Rate: 1}, nil

	case "1094":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x5e, Rate: 1}, nil
	case "1097":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x61, Rate: 1}, nil

	case "1124":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x7c, Rate: 1}, nil
	case "1127":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x7f, Rate: 1}, nil

	case "1230":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0xe6, Rate: 5}, nil

	case "4072_0":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0xfe, Rate: 1}, nil
	case "4072_1":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0xfd, Rate: 1}, nil
	default:
		return nil, errors.New("wrong input type")
	}

}

/*
info: turning of given RTCM message
requires: Message wich needs to be turned off
*/
func turnOffRTCM(topic string) (ubx.Message, error) {
	switch topic {
	case "1005":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x05, Rate: 0}, nil
	case "1006":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x06, Rate: 0}, nil

	case "1009":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x09, Rate: 0}, nil
	case "1010":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x0a, Rate: 0}, nil

	case "1074":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x4a, Rate: 0}, nil
	case "1077":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x4d, Rate: 0}, nil

	case "1084":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x54, Rate: 0}, nil
	case "1087":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x57, Rate: 0}, nil

	case "1094":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x5e, Rate: 0}, nil
	case "1097":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x61, Rate: 0}, nil

	case "1124":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x7c, Rate: 0}, nil
	case "1127":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0x7f, Rate: 0}, nil

	case "1230":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0xe6, Rate: 0}, nil

	case "4072_0":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0xfe, Rate: 0}, nil
	case "4072_1":
		return ubx.CfgMsg1{MsgClass: 0xf5, MsgID: 0xfd, Rate: 0}, nil
	default:
		return nil, errors.New("wrong input type")
	}

}
