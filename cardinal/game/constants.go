package game

import "github.com/argus-labs/starter-game-template/types"

// This is where we are declaring various game constants.

type IWorldConstants struct {
	SeedWord    string
	PlayerCount int
}

type IFooConstants struct {
	Foo string
}

var (
	// If you want the constant to be queryable through `query_constant`,
	// make sure to add the constant to the list of exposed constants
	ExposedConstants = []types.IConstant{{
		Label: "world",
		Value: WorldConstants,
	}}

	// WorldConstants is a public constant that can be queried through `query_constant`
	// because it is in the list of ExposedConstants
	WorldConstants = IWorldConstants{
		SeedWord:    "SeedWord",
		PlayerCount: 0,
	}

	// FooConstant is a private constant that cannot be queried through `query_constant`
	// because it is not in the list of ExposedConstants
	FooConstants = IFooConstants{
		Foo: "Bar",
	}
)
