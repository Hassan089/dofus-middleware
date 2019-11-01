package main


type Fighter struct {
	Id string
	Name string
	Level int
	TeamId int
	CellId int
	Life int
}

type Fight struct {
	Fighters []Fighter
}

var CurrentFight *Fight = nil

func getFighter(fight Fight, fighterId string) *Fighter {
	for i, c := range fight.Fighters {
		if c.Id == fighterId || c.Name == fighterId {
			return &fight.Fighters[i]
		}
	}
	return nil
}

func moveFigterToCellId(fight Fight, fighterId int, cellId int) {
	getFighter(fight, string(fighterId)).CellId = cellId
}

func updateFighter(fight Fight, fighter Fighter) {
	f := getFighter(fight, fighter.Id)
	if f != nil {
		f.CellId = fighter.CellId
	} else {
		fight.Fighters = append(CurrentFight.Fighters, fighter)
	}
}