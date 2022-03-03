package acc

import (
	"fmt"
	"time"
)

type AccData struct {
	SessionLength    time.Duration
	SessionTime      time.Duration
	ProgressWithFule float32
	SessionLaps      int
	SessionLiter     int
	RaceProgress     float32
	FuelLevel        float32
	FuelPerLap       float32
}

func Update(accChan chan<- AccData) {

	session := float32(1800)      // 1h
	sessionTimer := float32(1800) // 1h
	lapTime := float32(120)       // seconds
	fuelLevel := float32(30)
	fuelPerLap := float32(2.45)

	ticker := time.NewTicker(time.Second * time.Duration(1))
	for range ticker.C {

		sessionTimer -= 1
		if sessionTimer == 0 {
			sessionTimer = float32(session)
		}

		rp := 100 - (float32(sessionTimer)*float32(100))/float32(session)

		lapsWithFuel := fuelLevel / fuelPerLap
		percentageWithFuel := (float32(sessionTimer+(lapTime*lapsWithFuel))*float32(100))/float32(session) - 100
		fmt.Printf("laps with fuel %v, percentage with fuel %v %v\n", lapsWithFuel, percentageWithFuel, percentageWithFuel+rp)

		fuelPerSec := fuelPerLap / lapTime
		fuelLevel -= fuelPerSec
		//fmt.Printf("lap time: %v laps with fuel: %v additional time: %v\n", time.Duration(lapTime*float32(time.Second)), lapsWithFuel, time.Duration(additionTimeWithFuel*float32(time.Second)))

		accChan <- AccData{
			SessionLength:    time.Duration(session * float32(time.Second)),
			SessionTime:      time.Duration(sessionTimer * float32(time.Second)),
			SessionLiter:     94,
			SessionLaps:      74,
			RaceProgress:     rp,
			FuelLevel:        fuelLevel,
			FuelPerLap:       fuelPerLap,
			ProgressWithFule: rp + percentageWithFuel,
		}
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
