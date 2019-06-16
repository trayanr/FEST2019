package models

type Points struct {
	ID     int
	Value  float64
	UserId int
	Level  int
}

var types = map[int]float64{
	7:  1,
	8:  2.5,
	1:  1.8,
	82: 3,
}

var levels = map[int]float64{
	0: 0,
	1: 100,
	2: 300,
	3: 900,
	4: 2700,
	5: 8100,
}

func (p *Points) Calculate(activityType int, activityDuration uint) {
	var multiplier = types[activityType]
	act := activityDuration / 100000
	activityDuration = uint(act)
	var points float64
	if activityDuration != 0 {
		points = float64(activityDuration) * multiplier
	} else {
		points = multiplier
	}

	p.Value += points

	//Смята левала
	for i := 1; i <= len(levels); i++ {
		if p.Value < levels[i] && p.Value > levels[i-1] {
			p.Level = i
		}
	}

}
