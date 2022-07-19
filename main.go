package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorReset = "\033[0m"

var gpsDataChannel = make(chan startServerChannel)

func main() {
	cfg := loadConfig()
	cfgList := saveCfg() //get list with message formats

	sControl := OpenPort(cfg.Control.TtyPorts, cfg.Control.Baud)
	sServer := OpenPort(cfg.Server.TtyPorts, cfg.Server.Baud)

	//configure NTRIP server
	log.Println("writing settings to port: ", cfg.Server.TtyPorts)
	config(sServer, cfg.Options, false)
	configServer(sServer, cfg.Server, true)
	//save settings to flash
	for i := 0; i < len(cfgList); i++ {
		log.Print("saving settings to flash. Number ", i, "of ", len(cfgList)-1)
		write(cfgList[i], sServer)
	}

	//configure zed-f9p control port
	log.Println("writing settigns to port: ", cfg.Control.TtyPorts)
	config(sControl, cfg.Options, false)
	configControl(sControl, cfg.Control, true)
	//save settings to flash
	for i := 0; i < len(cfgList); i++ {
		log.Print("saving settings to flash. Number ", i, " of ", len(cfgList)-1)
		write(cfgList[i], sControl)
	}

	read(sControl, cfg.Control.TtyPorts)
	go startServer(cfg.Server, gpsDataChannel)

	fmt.Scanln()
	log.Println("end of program")
}

/*
info: start the ntrip server
requires: Config for the tty port
*/
func startServer(cfg Server, channel chan startServerChannel) {
	for content := range channel {
		if content.fix {
			if runtime.GOOS == "windows" {
				fmt.Println("cannot execute this on windows")
			} else {
				file := cfg.Str2Str
				sPort := "serial://" + cfg.TtyPorts + ":" + string(cfg.Baud) + ":8:n:1:off#ubx"
				casterAddress := "ntrips://:hello@" + cfg.Caster + cfg.CasterPort + cfg.Mountpoint
				fmt.Println(sPort, casterAddress)
				arguments := []string{"-in", sPort, "-out", casterAddress, "-msg", "\"1006,1074,1084,1094\"", "-s", "10000"}
				log.Println(arguments)
				cmd := exec.Command(file, arguments[0], arguments[1], arguments[2], arguments[3], arguments[4], arguments[5], arguments[6], arguments[7])

				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				output := cmd.Run()
				log.Println(output)
			}
		}
	}
}
