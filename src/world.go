package src

// Physical World Class

type World struct {
	Bodies []Body
}

func NewWorld() *World {
	return &World{}
}

// Appending new objects to the world
func (world *World) AddBody(body *Body) {
	world.Bodies = append(world.Bodies, *body)
}

// Generating Friction
func (world *World) GenerateFirction() {
	for i := range world.Bodies {
		world.Bodies[i].ApplyForce(*world.Bodies[i].Vel.Get().Multi(-1).Multi(0.5))
		if world.Bodies[i].Vel.X < 0.3 && world.Bodies[i].Vel.X > -0.3 {
			world.Bodies[i].Vel.X = 0
		}
		if world.Bodies[i].Vel.Y < 0.3 && world.Bodies[i].Vel.Y > -0.3 {
			world.Bodies[i].Vel.Y = 0
		}
	}
}
