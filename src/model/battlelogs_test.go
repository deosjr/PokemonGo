package model

import (
	"testing"
)

func TestLogText(t *testing.T) {
	for i, tt := range []struct {
		text string
		want []battleLog
	}{
		{
			text: "Critical hit!",
			want: []battleLog{textLog{Text: "Critical hit!"}},
		},
	} {
		log := NewLogger()
		log.f(tt.text)
		got := log.Logs()[1]
		if len(got) != len(tt.want) {
			t.Errorf("%d) different length output: got %d want %d", i, len(got), len(tt.want))
			continue
		}
		for i, g := range got {
			if g != tt.want[i] {
				t.Errorf("%d) got %q want %q", i, got, tt.want)
				continue
			}
		}
	}
}

func TestDamageLog(t *testing.T) {
	for i, tt := range []struct {
		name   string
		index  int
		damage int
		t      float64
		crit   float64
		want   []battleLog
	}{
		{
			t:      1.0,
			crit:   1.0,
			index:  1,
			damage: 42,
			want:   []battleLog{damageLog{Index: 1, Damage: 42}},
		},
		{
			name: "Pikachu",
			t:    0.0,
			crit: 1.0,
			want: []battleLog{
				damageLog{Index: 0, Damage: 0},
				textLog{Text: "It doesn't affect Pikachu."},
			},
		},
		{
			t:      1.5,
			crit:   2.0,
			damage: 100,
			want: []battleLog{
				damageLog{Index: 0, Damage: 100},
				textLog{Text: "Critical hit!"},
				textLog{Text: "It's super effective!"},
			},
		},
	} {
		log := NewLogger()
		log.damageWithMessages(tt.name, tt.index, tt.damage, tt.t, tt.crit)
		got := log.Logs()[1]
		if len(got) != len(tt.want) {
			t.Errorf("%d) different length output: got %d want %d", i, len(got), len(tt.want))
			continue
		}
		for i, g := range got {
			if g != tt.want[i] {
				t.Errorf("%d) got %q want %q", i, got, tt.want)
				continue
			}
		}
	}
}
