package component

type HealthComponent struct {
	HP int
}

func (HealthComponent) Name() string {
	return "Health"
}
