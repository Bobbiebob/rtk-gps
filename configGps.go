package main

import (
	"log"

	"github.com/daedaleanai/ublox/ubx"
	"github.com/tarm/serial"
)

/*
info: Turinging on and/or off nav, NMEA and RXM messages. This is specific for the ntrip server.
requires: pointer to serial.port and a Config struct
*/
func configServer(s *serial.Port, configs Server, on bool) {
	if on {
		//NAV messages
		for _, item := range configs.Nav {
			log.Println("turning on: ", item)
			temp, err := turnOnNav(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}

		//RXM messages
		for _, item := range configs.Rxm {
			log.Println("turning on: ", item)
			temp, err := turnOnRxm(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}

		// RTCM messages
		for _, item := range configs.Rtcm {
			log.Println("turning on: ", item)
			temp, err := turnOnRTCM(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}

		// NMEA messages
		for _, item := range configs.Nmea {
			log.Println("turning on: ", item)
			temp, err := turnOnNmea(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}
	} else {
		//NAV Messages
		for _, item := range configs.Nav {
			log.Println("turning off: ", item)
			temp, err := turnOffNav(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}

		//RXM Messages
		for _, item := range configs.Rxm {
			log.Println("turning off: ", item)
			temp, err := turnOffRxm(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}

		//RTCM Messages
		for _, item := range configs.Rtcm {
			log.Println("turning off: ", item)
			temp, err := turnOffRTCM(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}

		//NMEA Messages
		for _, item := range configs.Nmea {
			log.Println("turning off: ", item)
			temp, err := turnOffNmea(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}
	}
}

/*
info: Turinging on and/or off nav, NMEA and RXM messages. This is specific for the control port to keep check of the gps module.
requires: pointer to serial.port and a Config struct
*/
func configControl(s *serial.Port, configs Control, on bool) {
	if on {
		//NAV messages
		for _, item := range configs.Nav {
			log.Println("turning on: ", item)
			temp, err := turnOnNav(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}

		//RXM messages
		for _, item := range configs.Rxm {
			log.Println("turning on: ", item)
			temp, err := turnOnRxm(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}

		// RTCM messages
		for _, item := range configs.Rtcm {
			log.Println("turning on: ", item)
			temp, err := turnOnRTCM(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}

		// NMEA messages
		for _, item := range configs.Nmea {
			log.Println("turning on: ", item)
			temp, err := turnOnNmea(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}

		if configs.Svin != nil {
			sec, accuracy := configs.Svin[0], configs.Svin[1]
			log.Println("setting svin. sec: ", sec, " accurracy ", accuracy)
			write(timeMode3(sec, accuracy), s)
		}

	} else {
		//NAV Messages
		for _, item := range configs.Nav {
			log.Println("turning off: ", item)
			temp, err := turnOffNav(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}

		//RXM Messages
		for _, item := range configs.Rxm {
			log.Println("turning off: ", item)
			temp, err := turnOffRxm(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}

		//RTCM Messages
		for _, item := range configs.Rtcm {
			log.Println("turning off: ", item)
			temp, err := turnOffRTCM(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}

		//NMEA Messages
		for _, item := range configs.Nmea {
			log.Println("turning off: ", item)
			temp, err := turnOffNmea(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}
	}
}

/*
info: Turinging on and/or off nav, NMEA and RXM messages. This is to "reset" the gps module. Normally, all settings would be turned of to start with an clean slate.
requires: pointer to serial.port and a Config struct
*/
func config(s *serial.Port, configs Options, on bool) {

	if on {
		//NAV messages
		for _, item := range configs.Nav {
			log.Println("turning on: ", item)
			temp, err := turnOnNav(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}

		//RXM messages
		for _, item := range configs.Rxm {
			log.Println("turning on: ", item)
			temp, err := turnOnRxm(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}

		// RTCM messages
		for _, item := range configs.Rtcm {
			log.Println("turning on: ", item)
			temp, err := turnOnRTCM(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}

		// NMEA messages
		for _, item := range configs.Nmea {
			log.Println("turning on: ", item)
			temp, err := turnOnNmea(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}
	} else {
		//NAV Messages
		for _, item := range configs.Nav {
			log.Println("turning off: ", item)
			temp, err := turnOffNav(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}

		//RXM Messages
		for _, item := range configs.Rxm {
			log.Println("turning off: ", item)
			temp, err := turnOffRxm(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}

		//RTCM Messages
		for _, item := range configs.Rtcm {
			log.Println("turning off: ", item)
			temp, err := turnOffRTCM(item)
			fmtErrorDisplay(err)
			write(temp, s)
		}

		//NMEA Messages
		for _, item := range configs.Nmea {
			log.Println("turning off: ", item)
			temp, err := turnOffNmea(item)
			fmtErrorDisplay(err)
			write(temp, s)

		}
	}
}

var cfgItems = []ubx.CfgCfgClearMask{ubx.CfgCfgIoPort, ubx.CfgCfgMsgConf, ubx.CfgCfgInfMsg, ubx.CfgCfgNavConf,
	ubx.CfgCfgRxmConf, ubx.CfgCfgSenConf, ubx.CfgCfgRinvConf, ubx.CfgCfgAntConf, ubx.CfgCfgLogConf,
	ubx.CfgCfgFtsConf}

/*
info: save current gps settings on the module to flash
return: list of ubx.Message
*/
func saveCfg() []ubx.Message {
	var listToReturn = []ubx.Message{}

	for index := 0; index < len(cfgItems); index++ {
		toReturn := ubx.CfgCfg{}
		toReturn.DeviceMask = ubx.CfgCfgDeviceMask(ubx.CfgCfgDevFlash)
		toReturn.SaveMask = cfgItems[index]

		listToReturn = append(listToReturn, toReturn)
	}
	return listToReturn
}

/*
info: clear all settings from the module
return: list of ubx.Message
*/
func clearCfg() []ubx.Message {
	var listToReturn = []ubx.Message{}

	for index := 0; index < len(cfgItems); index++ {
		toReturn := ubx.CfgCfg{}
		toReturn.DeviceMask = ubx.CfgCfgDeviceMask(ubx.CfgCfgDevFlash)
		toReturn.ClearMask = cfgItems[index]

		listToReturn = append(listToReturn, toReturn)
	}
	return listToReturn
}

/*
info: load saved settings from the module in to active settings
return: list of ubx.Message
*/
func loadCfg() []ubx.Message {
	var listToReturn = []ubx.Message{}

	for index := 0; index < len(cfgItems); index++ {
		toReturn := ubx.CfgCfg{}
		toReturn.DeviceMask = ubx.CfgCfgDeviceMask(ubx.CfgCfgDevFlash)
		toReturn.LoadMask = cfgItems[index]

		listToReturn = append(listToReturn, toReturn)
	}
	return listToReturn
}
