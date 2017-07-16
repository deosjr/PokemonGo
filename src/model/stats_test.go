package model

import (
	"testing"
)

func TestUpdateStages(t *testing.T) {
	for i, tt := range []struct {
		stats     Stats
		changes   Stats
		newStats  Stats
		effective Stats
		maxed     [6]bool
	}{
		{
			stats:     Stats{0, 0, 0, 0, 0, 0},
			changes:   Stats{0, 0, 0, 0, 0, 0},
			newStats:  Stats{0, 0, 0, 0, 0, 0},
			effective: Stats{0, 0, 0, 0, 0, 0},
			maxed:     [6]bool{false, false, false, false, false, false},
		},
		{
			stats:     Stats{5, 0, 6, 0, 0, 0},
			changes:   Stats{1, 0, 0, 0, 0, 0},
			newStats:  Stats{6, 0, 6, 0, 0, 0},
			effective: Stats{1, 0, 0, 0, 0, 0},
			maxed:     [6]bool{false, false, false, false, false, false},
		},
		{
			stats:     Stats{0, -4, 0, 0, -6, 0},
			changes:   Stats{0, -2, 0, 0, 0, 0},
			newStats:  Stats{0, -6, 0, 0, -6, 0},
			effective: Stats{0, -2, 0, 0, 0, 0},
			maxed:     [6]bool{false, false, false, false, false, false},
		},
		{
			stats:     Stats{6, 0, 0, 0, -6, 0},
			changes:   Stats{3, 0, 0, 0, -4, 0},
			newStats:  Stats{6, 0, 0, 0, -6, 0},
			effective: Stats{0, 0, 0, 0, 0, 0},
			maxed:     [6]bool{true, false, false, false, true, false},
		},
	} {
		ns, e, m := tt.stats.updateStages(tt.changes)
		if ns != tt.newStats || e != tt.effective || m != tt.maxed {
			t.Errorf("%d): got %v %v %v want %v %v %v", i, ns, e, m, tt.newStats, tt.effective, tt.maxed)
		}
	}
}
