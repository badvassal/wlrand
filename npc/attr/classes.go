package attr

const (
	AttrClassIDAthlete int = iota
	AttrClassIDScholar
	AttrClassIDArtisan
	AttrClassIDCount
)

var AttrClasses = []AttrClass{
	AttrClassIDAthlete: AttrClass{
		Name: "Athlete",
		Weights: []float64{
			AttrIdxStrength:  6.0,
			AttrIdxIQ:        1.0,
			AttrIdxLuck:      1.0,
			AttrIdxSpeed:     3.0,
			AttrIdxAgility:   3.0,
			AttrIdxDexterity: 1.0,
			AttrIdxCharisma:  1.0,
		},
	},

	AttrClassIDScholar: AttrClass{
		Name: "Scholar",
		Weights: []float64{
			AttrIdxStrength:  1.0,
			AttrIdxIQ:        3.0,
			AttrIdxLuck:      1.0,
			AttrIdxSpeed:     1.0,
			AttrIdxAgility:   1.0,
			AttrIdxDexterity: 1.0,
			AttrIdxCharisma:  1.0,
		},
	},

	AttrClassIDArtisan: AttrClass{
		Name: "Artisan",
		Weights: []float64{
			AttrIdxStrength:  2.0,
			AttrIdxIQ:        1.0,
			AttrIdxLuck:      1.0,
			AttrIdxSpeed:     1.0,
			AttrIdxAgility:   2.0,
			AttrIdxDexterity: 4.0,
			AttrIdxCharisma:  1.0,
		},
	},
}
