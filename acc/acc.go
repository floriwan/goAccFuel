package acc

import "time"

type AccData struct {
	RaceProgress float32
}

func Update(accChan chan<- AccData) {

	rp := float32(0.)

	ticker := time.NewTicker(time.Second * time.Duration(1))
	for range ticker.C {

		rp += 1

		if rp > 100 {
			rp = 0
		}

		accChan <- AccData{RaceProgress: rp}
	}
}
