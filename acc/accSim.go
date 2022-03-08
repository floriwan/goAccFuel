package acc

import "time"

func updateSim(accChan chan<- AccData) {
	sessionTimer -= 1
	if sessionTimer == 0 {
		sessionTimer = float32(session)
	}

	rp := 100 - (float32(sessionTimer)*float32(100))/float32(session)

	fuelPerSec := fuelPerLap / lapTime
	fuelLevel -= fuelPerSec
	if fuelLevel <= 0 {
		fuelLevel = 30
	}

	lapsWithFuel := fuelLevel / fuelPerLap
	percentageWithFuel := (float32(lapTime*lapsWithFuel) * float32(100)) / float32(session)
	percentageWithFuel += rp

	sessionLaps := int(session / lapTime)

	lapsDone := int((session - sessionTimer) / lapTime)
	refuelLevel := float32((sessionLaps - lapsDone - int(lapsWithFuel))) * fuelPerLap

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
		LapsToGo:         lapsWithFuel,
		BoxLap:           lapsDone + int(lapsWithFuel),
		RefuelLevel:      refuelLevel,
	}
}
