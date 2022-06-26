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

// Subtract subtracts v (X, Y) from operand (X, Y).
func (v Vector) Subtract(operand Vector) Vector {
	return Vector{
		X: v.X - operand.X,
		Y: v.Y - operand.Y,
	}
}

// Multiply multiplies v (X, Y) by operand (X, Y).
func (v Vector) Multiply(operand Vector) Vector {
	return Vector{
		X: v.X * operand.X,
		Y: v.Y * operand.Y,
	}
}
