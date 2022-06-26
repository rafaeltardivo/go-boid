package boids

// Vector (two dimentional definition).
type Vector struct {
	X float64
	Y float64
}

// Add adds v (X, Y) to operand (X, Y).
func (v Vector) Add(operand Vector) Vector {
	return Vector{
		X: v.X + operand.X,
		Y: v.Y + operand.Y,
	}
}
