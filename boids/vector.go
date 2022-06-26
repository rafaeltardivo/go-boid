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

// AddValue adds the specified value to v.X and v.Y.
func (v Vector) AddValue(value float64) Vector {
	return Vector{
		X: v.X + value,
		Y: v.Y + value,
	}
}

// SubtractValue subtracts the specified value from v.X and v.Y.
func (v Vector) SubtractValue(value float64) Vector {
	return Vector{
		X: v.X - value,
		Y: v.Y - value,
	}
}

// MultiplyValue multiplies the specified value by v.X and v.Y.
func (v Vector) MultiplyValue(value float64) Vector {
	return Vector{
		X: v.X * value,
		Y: v.Y * value,
	}
}

// DivideValue divides the specified value by v.X and v.Y.
// TODO check for division by 0 error.
func (v Vector) DivideValue(value float64) Vector {
	return Vector{
		X: v.X / value,
		Y: v.Y / value,
	}
}
