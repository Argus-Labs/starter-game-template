package types

import ecs "github.com/argus-labs/world-engine/cardinal/ecs"

type IArchetype struct {
	Label      string
	Components []ecs.IComponentType
}

type IConstant struct {
	Label string
	Value any
}
