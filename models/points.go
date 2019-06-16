package models

type Points struct {
	Value float64
	Level int
}

var types = map[int]float64{
	7:  1,
	8:  2.5,
	1:  1.8,
	82: 3,
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

	p.Value = points
}

//TODO::da pazim posledoto dobavqne na to4ki

func GetLastSession(authCode string, lastTimeMS int) {
	// conf := GetConfig()

}
