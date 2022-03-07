package acc

import (
	"fmt"
	"log"
	"time"
)

type AccData struct {
	SessionLength    time.Duration
	SessionTime      time.Duration
	LapTime          time.Duration
	ProgressWithFuel float32
	SessionLaps      int
	SessionLiter     int
	RaceProgress     float32
	FuelLevel        float32
	FuelPerLap       float32
	LapsWithFuel     float32
	LapsDone         int
	BoxLap           int
	RefuelLevel      float32
}

var (
	session      = float32(1800)
	sessionTimer = float32(1800) // 1h
	lapTime      = float32(120)  // seconds
	fuelLevel    = float32(30)
	fuelPerLap   = float32(2.45)
)

func Update(sim bool, accChan chan<- AccData) {

	ticker := time.NewTicker(time.Second * time.Duration(1))
	for range ticker.C {
		if sim {
			fillSimValues(accChan)
		} else {
			log.Printf("acc update not implemented ...\n")
		}
	}
}

func fillSimValues(accChan chan<- AccData) {
	sessionTimer -= 1
	if sessionTimer == 0 {
		sessionTimer = float32(session)
	}

	rp := 100 - (float32(sessionTimer)*float32(100))/float32(session)

	fuelPerSec := fuelPerLap / lapTime
	fuelLevel -= fuelPerSec

	lapsWithFuel := fuelLevel / fuelPerLap
	percentageWithFuel := (float32(lapTime*lapsWithFuel) * float32(100)) / float32(session)
	percentageWithFuel += rp

	sessionLaps := int(session / lapTime)

	lapsDone := int((session - sessionTimer) / lapTime)
	refuelLevel := float32((sessionLaps - lapsDone - int(lapsWithFuel))) * fuelPerLap
	//fmt.Printf("fuel per second %v, laps with fuel %v, percentage with fuel %v %v \n", fuelPerSec, lapsWithFuel, percentageWithFuel, percentageWithFuel+rp)

	accChan <- AccData{
		SessionLength:    time.Duration(session * float32(time.Second)),
		SessionTime:      time.Duration(sessionTimer * float32(time.Second)),
		LapTime:          time.Duration(lapTime * float32(time.Second)),
		SessionLiter:     int((session / lapTime) * fuelPerLap),
		SessionLaps:      sessionLaps,
		RaceProgress:     rp,
		FuelLevel:        fuelLevel,
		FuelPerLap:       fuelPerLap,
		ProgressWithFuel: percentageWithFuel,
		LapsWithFuel:     lapsWithFuel,
		LapsDone:         lapsDone,
		BoxLap:           lapsDone + int(lapsWithFuel),
		RefuelLevel:      refuelLevel,
	}
}

func FmtDuration(d time.Duration) string {
	d = d.Round(time.Second)

	h := d / time.Hour
	d -= h * time.Hour

	m := d / time.Minute
	d -= m * time.Minute

	s := d / time.Second

	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
