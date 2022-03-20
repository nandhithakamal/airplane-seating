package drawer

import (
	"airplane-seating/internal/seat"
	"fmt"
	tm "github.com/buger/goterm"
)

const boxWidth = 6
const boxHeight = 5

func DrawPassengerSeatMap(seats []*seat.Seat) {
	tm.Clear()

	for _, seat := range seats {
		box := tm.NewBox(boxWidth|tm.PCT, boxHeight, 0)
		if seat.PassengerId() != -1 {
			fmt.Fprint(box, seat.PassengerId())
		}
		tm.Print(tm.MoveTo(box.String(), (seat.Column()*(boxWidth))|tm.PCT, seat.Row()*(boxHeight-1)))
		tm.Flush()
	}
}
