package game

type Constant struct {
	Label string
	Value any
}

type WorldConstants struct {
	SeedWord    string
	PlayerCount int
}

var (
	// ExposedConstants If you want the constant to be queryable through `query_constant`,
	// make sure to add the constant to the list of exposed constants
	ExposedConstants = []Constant{
		{
			Label: "world",
			Value: WorldConstants{
				SeedWord:    "SeedWord1",
				PlayerCount: 0,
			},
		},
	}
)
