package acc

import (
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
			updateSim(accChan)
		} else {
			updateShm(accChan)
		}
	}
}
