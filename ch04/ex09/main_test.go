package main

import "testing"

func TestWordfreqTestWordfreq(t *testing.T) {
	tests := []struct {
		input string
		want  WordCounts
	}{
		{
			input: "data/input1.txt",
			want: WordCounts{
				"go":   2,
				"bike": 1,
			},
		},
		{
			input: "data/input2.txt",
			want: WordCounts{
				"mid":          1,
				"here":         1,
				"in":           7,
				"source":       3,
				"They":         1,
				"stars":        1,
				"never":        1,
				"before":       2,
				"times":        1,
				"tonight":      2,
				"break":        1,
				"For":          4,
				"see":          1,
				"Day":          1,
				"morning":      2,
				"darkness":     1,
				"all":          2,
				"reality":      1,
				"after":        1,
				"towards":      2,
				"day":          5,
				"hands":        2,
				"Whoa":         1,
				"with":         1,
				"every":        1,
				"pain":         3,
				"their":        3,
				"my":           2,
				"So":           4,
				"thousand":     3,
				"right":        1,
				"Falls":        1,
				"wastelands":   1,
				"lifetime":     3,
				"cracks":       1,
				"descend":      1,
				"you":          1,
				"possibly":     1,
				"eternal":      1,
				"fire":         4,
				"far":          5,
				"Our":          1,
				"raise":        1,
				"alone":        1,
				"this":         2,
				"understand":   1,
				"are":          1,
				"Fightin":      1,
				"become":       1,
				"lost":         4,
				"your":         1,
				"man":          1,
				"time":         3,
				"And":          4,
				"moonlight":    1,
				"re":           5,
				"wilderness":   1,
				"sun":          1,
				"banished":     1,
				"has":          1,
				"On":           3,
				"ride":         1,
				"evil":         1,
				"stand":        1,
				"will":         1,
				"thunderstorm": 1,
				"red":          1,
				"winter":       1,
				"is":           3,
				"souls":        2,
				"days":         3,
				"from":         1,
				"curse":        1,
				"if":           1,
				"falling":      1,
				"free":         3,
				"Deep":         1,
				"own":          1,
				"go":           1,
				"flames":       4,
				"skeletors":    1,
				"heart":        1,
				"gone":         4,
				"heavens":      1,
				"Running":      1,
				"can":          2,
				"carries":      1,
				"inside":       2,
				"hearts":       1,
				"foreign":      1,
				"and":          9,
				"Who":          1,
				"unto":         1,
				"lies":         1,
				"death":        1,
				"s":            4,
				"on":           9,
				"lightning":    1,
				"There":        1,
				"fought":       1,
				"light":        6,
				"hell":         2,
				"now":          3,
				"lands":        1,
				"As":           1,
				"ll":           3,
				"fight":        1,
				"down":         2,
				"fightin":      1,
				"hard":         2,
				"above":        1,
				"burning":      1,
				"back":         1,
				"to":           1,
				"we":           12,
				"know":         1,
				"beyond":       4,
				"destiny":      1,
				"desperation":  1,
				"mind":         1,
				"once":         1,
				"endlessly":    1,
				"find":         1,
				"must":         1,
				"I":            2,
				"When":         1,
				"for":          4,
				"Bodies":       1,
				"freedom":      1,
				"shores":       1,
				"Far":          1,
				"sky":          1,
				"evermore":     1,
				"feel":         4,
				"wait":         3,
				"land":         1,
				"a":            10,
				"of":           9,
				"wings":        2,
				"quest":        1,
				"so":           5,
				"through":      1,
				"that":         1,
				"misery":       1,
				"wasted":       4,
				"blackest":     2,
				"domain":       1,
				"Now":          1,
				"To":           2,
				"need":         1,
				"seal":         1,
				"The":          2,
				"laughter":     1,
				"away":         3,
				"steel":        1,
				"within":       1,
				"blood":        1,
				"cold":         1,
				"We":           8,
				"sound":        1,
				"Lost":         1,
				"reign":        1,
				"watch":        1,
				"roaming":      1,
				"around":       1,
				"world":        1,
				"flame":        1,
				"tied":         1,
				"dream":        1,
				"In":           4,
				"the":          40,
				"tough":        1,
				"All":          1,
				"dreams":       1,
				"Through":      4,
				"flying":       1,
				"carry":        3,
				"again":        1,
				"our":          3,
				"dawning":      1,
			},
		},
	}

	for _, tt := range tests {
		got := wordfreq(tt.input)
		if len(got) != len(tt.want) {
			t.Errorf("len(got) = %d, want = %d", len(got), len(tt.want))
		}
		for k, v := range got {
			if v != tt.want[k] {
				t.Errorf("got[%s] = %d, want = %d", k, v, tt.want[k])
			}
		}
	}
}
