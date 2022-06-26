package boid

import "math"

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
func (v Vector) DivideValue(value float64) Vector {

	// TODO check for division by 0 error.
	return Vector{
		X: v.X / value,
		Y: v.Y / value,
	}
}

// Limit limits upper and lower boundaries for v.X and v.Y.
func (v Vector) Limit(lower, upper float64) Vector {
	return Vector{
		X: math.Min(math.Max(v.X, lower), upper),
		Y: math.Min(math.Max(v.Y, lower), upper),
	}
}

// Distance calculates the distance between v and the target vector.
func (v Vector) Distance(target Vector) float64 {
	return math.Sqrt(math.Pow(v.X-target.X, 2) + math.Pow(v.Y-target.Y, 2))
}
