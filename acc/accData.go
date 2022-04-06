package acc

import (
	"goAccFuel/acc/shm"
	"log"
	"syscall"
	"time"
)

type accShmData struct {
	sData shm.ACCStatic
	gData shm.ACCGraphics
	pData shm.ACCPhysics
}

//type accStatus int32

//func (s accStatus) String() string {
//	return [...]string{"off", "replay", "live", "pause"}[s]
//}

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

	if err := shm.ReadStatic(&cData.sData); err != nil {
		log.Fatalf("no acc shm available, retry ...")
		return
	}

	//empty := [15]uint16{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//if empty == cData.sData.ACVersion {
	//	log.Printf("no acc available ...\n")
	//}
	accVersion := syscall.UTF16ToString(cData.sData.ACVersion[:15])
	carModel := syscall.UTF16ToString(cData.sData.CarModel[:33])

	if err := shm.ReadGraphics(&cData.gData); err != nil {
		log.Fatalf("read physics error %v", err)
		return
	}

	if err := shm.ReadPhysics(&cData.pData); err != nil {
		log.Fatalf("read physics error %v", err)
		return
	}

	//status := AccStatus(cData.gData.Status).String()
	sessionType := AccSessionType(cData.gData.SessionType).String()

	lapTime := time.Duration(0 * float32(time.Second)) // set an initial default lap time
	if cData.gData.ILastTime != 2147483647 {
		lapTime = time.Duration(cData.gData.ILastTime) * time.Millisecond
	}

	sessionLaps := 0
	sessionLiter := 0
	lapsToGo := float32(0)
	fuelNeeded := float32(0)
	fuelLevel := cData.pData.Fuel
	raceProgress := float32(0)
	fuelLap := cData.gData.FuelXLap
	lapsWithFuel := float32(0)
	percentageWithFuel := float32(0)

	// car is moving, save the session time
	sessionTimeLeft := time.Duration(cData.gData.SessionTimeLeft) * time.Millisecond
	if sessionLength == 0 || sessionLength < sessionTimeLeft {
		sessionLength = sessionTimeLeft
	}

	// pit window
	//pitWindowLength := uint32(cData.sData.PitWindowEnd - cData.sData.PitWindowStart)
	pitWindowOpenTime := time.Duration(cData.sData.PitWindowStart) * time.Millisecond
	pitWindowCloseTime := time.Duration(cData.sData.PitWindowEnd) * time.Millisecond
	pitWindowOpenPercentage := (float32(pitWindowOpenTime) * float32(100)) / float32(sessionLength)
	pitWindowClosePercentage := (float32(pitWindowCloseTime) * float32(100)) / float32(sessionLength)
	//fmt.Printf("pit window start: %v close: %v\n", pitWindowOpenTime, pitWindowCloseTime)
	//fmt.Printf("car skin %v\n", syscall.UTF16ToString(cData.sData.CarSkin[:33]))
	//fmt.Printf("dry tyre name %v\n", syscall.UTF16ToString(cData.sData.DryTyreName[:33]))
	//fmt.Printf("wet tyre name %v\n", syscall.UTF16ToString(cData.sData.WetTyreName[:33]))

	if lapTime > time.Duration(0*float32(time.Second)) {
		sessionLaps = int(sessionLength.Round(time.Millisecond) / lapTime.Round(time.Millisecond))
		lapsToGo = float32(sessionTimeLeft.Round(time.Millisecond)) / float32(lapTime.Round(time.Millisecond))
		fuelNeeded = fuelLap * float32(lapsToGo+1) // add one lap for safty reasons :-)
		sessionLiter = int((float32(sessionLength) / float32(lapTime)) * float32(fuelLap))

		//fuelUsed := cData.gData.UsedFuel
		lapsWithFuel = fuelLevel / fuelLap
		raceProgress = 100 - (float32(sessionTimeLeft)*float32(100))/float32(sessionLength)
		percentageWithFuel = (float32(float32(lapTime)*lapsWithFuel) * float32(100)) / float32(sessionLength)
		percentageWithFuel += raceProgress

	}

	accChan <- AccData{
		AccVersion:         accVersion,
		CarModel:           carModel,
		SessionLength:      sessionLength,
		SessionTime:        sessionTimeLeft,
		SessionLaps:        sessionLaps,
		SessionLiter:       sessionLiter,
		LapsDone:           int(cData.gData.CompletedLaps),
		RaceProgress:       raceProgress,
		ProgressWithFuel:   percentageWithFuel,
		LapTime:            lapTime,
		Status:             sessionType,
		FuelLevel:          fuelLevel,
		FuelPerLap:         fuelLap,
		RefuelLevel:        fuelNeeded,
		LapsToGo:           lapsToGo,
		LapsWithFuel:       lapsWithFuel,
		PitWindowStartTime: pitWindowOpenTime,
		PitWindowCloseTime: pitWindowCloseTime,
		PitWindowStart:     pitWindowOpenPercentage,
		PitWindowEnd:       pitWindowClosePercentage,
	}

}
