package game

type IConstant struct {
	Label string
	Value any
}

type IWorldConstants struct {
	SeedWord    string
	PlayerCount int
}

var (
	// ExposedConstants If you want the constant to be queryable through `query_constant`,
	// make sure to add the constant to the list of exposed constants
	ExposedConstants = []IConstant{
		{
			Label: "world",
			Value: WorldConstants,
		},
	}

	WorldConstants = IWorldConstants{
		SeedWord:    "SeedWord1",
		PlayerCount: 0,
	}
)
