package acc

import (
	"fmt"
	"time"
)

type AccData struct {
	AccVersion         string
	Status             string
	CarModel           string
	SessionLength      time.Duration
	SessionTime        time.Duration
	LapTime            time.Duration
	ProgressWithFuel   float32
	SessionLaps        int
	SessionLiter       int
	RaceProgress       float32
	FuelLevel          float32
	FuelPerLap         float32
	CompletedLaps      int
	LapsWithFuel       float32
	LapsDone           int
	BoxLap             int
	LapsToGo           float32
	RefuelLevel        float32
	PitWindowStartTime time.Duration
	PitWindowCloseTime time.Duration
	PitWindowStart     float32
	PitWindowEnd       float32
}

type AccUpdater interface {
	update() (AccData, error)
}

func Update(sim bool, accChan chan<- AccData) {

	ticker := time.NewTicker(time.Second * time.Duration(1))
	for range ticker.C {

		if sim {
			data, err := updateAccSim()
			if err != nil {
				fmt.Printf("error: %e\n", err)
			}
			accChan <- data

		} else {
			data, err := updateAccShm()
			if err != nil {
				fmt.Printf("error: %e\n", err)
			}
			accChan <- data
		}

	}
}

func saveLapData() {

}
