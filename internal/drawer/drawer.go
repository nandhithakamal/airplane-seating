package drawer

import (
	"airplane-seating/internal/seat/seatmap"
	"fmt"
	tm "github.com/buger/goterm"
)

const BOX_WIDTH = 4
const BOX_HEIGHT = 3
const GO_TERM_FLAGS = 0

func DrawPassengerSeatMap(m *seatmap.SeatMap) {
	tm.Clear()

	boundaryBox := tm.NewBox(((m.RightMost()+1)*BOX_WIDTH)|tm.PCT, (m.BackMost()+1)*BOX_HEIGHT, GO_TERM_FLAGS)
	tm.Print(tm.MoveTo(boundaryBox.String(), (BOX_WIDTH-1)|tm.PCT, 1))

	for _, seat := range m.Seats() {
		box := tm.NewBox(BOX_WIDTH|tm.PCT, BOX_HEIGHT, GO_TERM_FLAGS)
		if seat.PassengerId() != -1 {
			fmt.Fprint(box, seat.PassengerId())
		}
		tm.Print(tm.MoveTo(box.String(), (seat.Column()*BOX_WIDTH)|tm.PCT, seat.Row()*(BOX_HEIGHT)))
	}

	boundaryBox = tm.NewBox(((m.RightMost()+1)*BOX_WIDTH)|tm.PCT, BOX_HEIGHT, GO_TERM_FLAGS)
	tm.Print(tm.MoveTo(boundaryBox.String(), (BOX_WIDTH-1)|tm.PCT, (m.BackMost()+1)*BOX_HEIGHT))

	tm.Flush()
}
