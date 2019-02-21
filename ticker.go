package dateticker

import (
	"log"
	"math"
	"time"
)

type TickerFromDateTime struct {
	DateTime time.Time
	Tick     time.Duration
}

func (t *TickerFromDateTime) getNextStartDate() (time.Time, time.Duration) {
	now := time.Now().UTC()
	dateDuration := now.Sub(t.DateTime)

	steps := uint64(math.Ceil(float64(dateDuration) / float64(t.Tick)))

	nextStartDate := t.DateTime.Add(time.Duration(uint64(t.Tick) * steps))

	return nextStartDate, nextStartDate.Sub(now)
}

func (t *TickerFromDateTime) Run(f func()) {
	_, startAfter := t.getNextStartDate()
	log.Printf("TickerFromDateTime start after : %v\n", startAfter)
	<-time.After(startAfter)
	for c := time.Tick(t.Tick); ; <-c {
		f()
	}
}
