package acc

import (
	"fmt"
	"log"
	"syscall"
	"time"

	"github.com/Dekadee/accshm"
)

type accShmData struct {
	sData accshm.ACCStatic
	gData accshm.ACCGraphics
	pData accshm.ACCPhysics
}

type AccStatus int32

func (s AccStatus) String() string {
	return [...]string{"off", "replay", "live", "pause"}[s]
}

type AccSessionType int32

const (
	ACC_UNKNOWN AccSessionType = iota - 1
	ACC_PRACTICE
	ACC_QUALIFY
	ACC_RACE
	ACC_HOTLAP
	ACC_TIMEATTACK
	ACC_DRIFT
	ACC_DRAG
	ACC_HOTSTINT
	ACC_HOTSTINTSUPERPOLE
)

func (s AccSessionType) String() string {
	if s == -1 {
		return "unknown"
	}

	return [...]string{
		"practice",
		"qualify",
		"race",
		"hotlap",
		"timeattack",
		"drift",
		"drag",
		"hotstint",
		"hotstintsuperpole"}[s]
}

var sessionLength time.Duration

func updateShm(accChan chan<- AccData) {

	var cData accShmData // all shm data

	if err := accshm.ReadStatic(&cData.sData); err != nil {
		log.Fatalf("no acc shm available, retry ...")
		return
	}

	//empty := [15]uint16{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//if empty == cData.sData.ACVersion {
	//	log.Printf("no acc available ...\n")
	//}
	accVersion := syscall.UTF16ToString(cData.sData.ACVersion[:15])

	if err := accshm.ReadGraphics(&cData.gData); err != nil {
		log.Fatalf("read physics error %v", err)
		return
	}

	if err := accshm.ReadPhysics(&cData.pData); err != nil {
		log.Fatalf("read physics error %v", err)
		return
	}

	status := AccStatus(cData.gData.Status).String()
	sessionType := AccSessionType(cData.gData.SessionType).String()
	if status == "off" {
		sessionLength = time.Duration(0 * float32(time.Second))
		accChan <- AccData{AccVersion: accVersion, Status: status}
		return
	}

	lapTime := time.Duration(0 * float32(time.Second)) // set an initial default lap time
	if cData.gData.ILastTime != 2147483647 {
		lapTime = time.Duration(cData.gData.ILastTime) * time.Millisecond
	}

	fuelLevel := cData.pData.Fuel
	fuelLap := cData.gData.FuelXLap
	//fuelUsed := cData.gData.UsedFuel
	lapsWithFuel := fuelLevel / fuelLap

	// car is moving, save the session time
	sessionTimeLeft := time.Duration(cData.gData.SessionTimeLeft) * time.Millisecond
	if sessionLength == 0 {
		fmt.Printf("++++ save session length")
		sessionLength = sessionTimeLeft
	}
	//if cData.gData.DistanceTraveled > 5 && sessionLength == 0 {
	//	sessionLength = sessionTimeLeft
	//}

	lapsToGo := float32(sessionTimeLeft.Round(time.Millisecond)) / float32(lapTime.Round(time.Millisecond))
	fuelNeeded := fuelLap * float32(lapsToGo+1) // add one lap for safty reasons :-)

	raceProgress := 100 - (float32(sessionTimeLeft)*float32(100))/float32(sessionLength)
	percentageWithFuel := (float32(float32(lapTime)*lapsWithFuel) * float32(100)) / float32(sessionLength)
	percentageWithFuel += raceProgress

	if lapTime == time.Duration(0*float32(time.Second)) {
		accChan <- AccData{
			AccVersion:    accVersion,
			SessionLength: sessionLength,
			FuelLevel:     fuelLevel,
			FuelPerLap:    fuelLap,
		}
	} else {
		accChan <- AccData{
			AccVersion:       accVersion,
			SessionLength:    sessionLength,
			SessionTime:      sessionTimeLeft,
			SessionLaps:      int(sessionLength.Round(time.Millisecond) / lapTime.Round(time.Millisecond)),
			RaceProgress:     raceProgress,
			ProgressWithFuel: percentageWithFuel,
			LapTime:          lapTime,
			Status:           sessionType,
			SessionLiter:     int((float32(sessionLength) / float32(lapTime)) * float32(fuelLap)),
			FuelLevel:        fuelLevel,
			FuelPerLap:       fuelLap,
			RefuelLevel:      fuelNeeded,
			LapsToGo:         lapsToGo,
			LapsWithFuel:     lapsWithFuel,
		}

	}

}
