package acc

import "time"

var (
	session      = float32(3600) // 1h
	sessionTimer = float32(3600)
	lapTime      = float32(120) // seconds
	fuelLevel    = float32(6)
	fuelPerLap   = float32(2.45)
)

func updateAccSim() (AccData, error) {
	sessionTimer -= 1
	if sessionTimer == 0 {
		sessionTimer = float32(session)
	}

	rp := 100 - (float32(sessionTimer)*float32(100))/float32(session)

	fuelPerSec := fuelPerLap / lapTime
	fuelLevel -= fuelPerSec
	if fuelLevel <= 0 {
		fuelLevel = 40
	}

	lapsWithFuel := fuelLevel / fuelPerLap
	percentageWithFuel := (float32(lapTime*lapsWithFuel) * float32(100)) / float32(session)
	percentageWithFuel += rp

	sessionLaps := int(session / lapTime)

	lapsDone := int((session - sessionTimer) / lapTime)
	refuelLevel := float32((sessionLaps - lapsDone - int(lapsWithFuel))) * fuelPerLap

	return AccData{SessionLength: time.Duration(session * float32(time.Second)),
		SessionTime:        time.Duration(sessionTimer * float32(time.Second)),
		LapTime:            time.Duration(lapTime * float32(time.Second)),
		SessionLiter:       int((session / lapTime) * fuelPerLap),
		SessionLaps:        sessionLaps,
		RaceProgress:       rp,
		FuelLevel:          fuelLevel,
		FuelPerLap:         fuelPerLap,
		ProgressWithFuel:   percentageWithFuel,
		LapsWithFuel:       lapsWithFuel,
		LapsDone:           lapsDone,
		LapsToGo:           lapsWithFuel,
		BoxLap:             lapsDone + int(lapsWithFuel),
		RefuelLevel:        refuelLevel,
		PitWindowStartTime: time.Duration(75 * float32(time.Minute)),
		PitWindowCloseTime: time.Duration(75 * float32(time.Minute)),
		PitWindowStart:     float32(25),
		PitWindowEnd:       float32(75),
		Status:             "race",
		AccVersion:         "sim"}, nil

}
