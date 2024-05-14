package mathgd

import "math"

/**************************************************************************/
/*  math_funcs.h, math_defs.h                                             */
/**************************************************************************/
/*                         This file is part of:                          */
/*                             GODOT ENGINE                               */
/*                        https://godotengine.org                         */
/*                                                                        */
/*                        Ported to Go on 5/2024 from					  */
/*                    Godot Engine v4.2.1.stable.official                 */
/*                                                                        */
/**************************************************************************/
/* Copyright (c) 2014-present Godot Engine contributors (see AUTHORS.md). */
/* Copyright (c) 2007-2014 Juan Linietsky, Ariel Manzur.                  */
/*                                                                        */
/* Permission is hereby granted, free of charge, to any person obtaining  */
/* a copy of this software and associated documentation files (the        */
/* "Software"), to deal in the Software without restriction, including    */
/* without limitation the rights to use, copy, modify, merge, publish,    */
/* distribute, sublicense, and/or sell copies of the Software, and to     */
/* permit persons to whom the Software is furnished to do so, subject to  */
/* the following conditions:                                              */
/*                                                                        */
/* The above copyright notice and this permission notice shall be         */
/* included in all copies or substantial portions of the Software.        */
/*                                                                        */
/* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,        */
/* EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF     */
/* MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. */
/* IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY   */
/* CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,   */
/* TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE      */
/* SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.                 */
/**************************************************************************/

// CMP_EPSILON represents the tolerance value used for floating-point comparison.
const CMP_EPSILON = 0.00001

// CMP_EPSILON2 represents the square of CMP_EPSILON.
const CMP_EPSILON2 = CMP_EPSILON * CMP_EPSILON

// CMP_NORMALIZE_TOLERANCE represents the tolerance value used for normalizing vectors.
const CMP_NORMALIZE_TOLERANCE = 0.000001

// CMP_POINT_IN_PLANE_EPSILON represents the tolerance value used for checking if a point lies on a plane.
const CMP_POINT_IN_PLANE_EPSILON = 0.00001

// TAU represents the mathematical constant Tau (2 * Pi).
const TAU = 6.2831853071795864769252867666

// PI represents the mathematical constant Pi.
const PI = 3.1415926535897932384626433833

type EulerOrder int

const (
	EulerOrderXYZ EulerOrder = iota
	EulerOrderXZY
	EulerOrderYXZ
	EulerOrderYZX
	EulerOrderZXY
	EulerOrderZYX
)

// IsZeroApprox checks if a floating-point number is approximately zero within a certain tolerance.
func IsZeroApprox(x float64) bool {
	return math.Abs(x) < CMP_EPSILON
}

// IsEqualApprox checks if two floating-point numbers are approximately equal within a certain tolerance.
func IsEqualApprox(x, y float64) bool {
	return IsZeroApprox(x - y)
}

// Sign returns the sign of a floating-point number.
// It returns 1 if x is positive, -1 if x is negative, and 0 if x is zero.
func Sign(x float64) float64 {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

// Clamp clamps a value within a specified range.
// If val is less than min, it returns min.
// If val is greater than max, it returns max.
// Otherwise, it returns val.
func Clampi(val, min, max int) int {
	if val < min {
		return min
	} else if val > max {
		return max
	} else {
		return val
	}
}

// Clamp clamps a value within a specified range.
// If val is less than min, it returns min.
// If val is greater than max, it returns max.
// Otherwise, it returns val.
func Clampf(val, min, max float64) float64 {
	if val < min {
		return min
	} else if val > max {
		return max
	} else {
		return val
	}
}

// Snapped returns the nearest value to 'from' that is a multiple of 'to'.
// If 'to' is zero, it returns 0.
func Snapped(from, to float64) float64 {
	if to == 0 {
		return 0
	}
	return math.Round(from/to) * to
}

// Fposmod returns the positive floating-point modulus of x modulo y.
// If the result of the modulo operation is negative, it wraps around to ensure a positive result.
func Fposmod(x, y float64) float64 {
	result := math.Mod(x, y)
	if result < 0 {
		result += y
	}
	return result
}

// DegToRad converts degrees to radians.
// It multiplies the input value 'deg' by the ratio of Pi to 180.
func DegToRad(deg float64) float64 {
	return deg * (PI / 180.0)
}

// RadToDeg converts radians to degrees.
// It multiplies the input value 'rad' by the ratio of 180 to Pi.
func RadToDeg(rad float64) float64 {
	return rad * (180.0 / PI)
}

// Lerp performs linear interpolation between two values.
// It returns a value between 'p_from' and 'p_to' based on the interpolation weight 'p_weight'.
func Lerp(p_from, p_to, p_weight float64) float64 {
	return p_from + (p_to-p_from)*p_weight
}

// CubicInterpolate performs cubic interpolation between two values.
// It interpolates between 'p_from' and 'p_to' using 'p_pre' and 'p_post' as control points,
// with 'p_weight' determining the position between 'p_from' and 'p_to'.
func CubicInterpolate(p_from, p_to, p_pre, p_post, p_weight float64) float64 {
	return 0.5 *
		((p_from * 2.0) +
			(-p_pre+p_to)*p_weight +
			(2.0*p_pre-5.0*p_from+4.0*p_to-p_post)*(p_weight*p_weight) +
			(-p_pre+3.0*p_from-3.0*p_to+p_post)*(p_weight*p_weight*p_weight))
}

// CubicInterpolateAngle performs cubic interpolation between two angles represented in radians.
// It ensures smooth interpolation by handling angle wrapping around the unit circle.
func CubicInterpolateAngle(p_from, p_to, p_pre, p_post, p_weight float64) float64 {
	from_rot := math.Mod(p_from, TAU)

	pre_diff := math.Mod(p_pre-from_rot, TAU)
	pre_rot := from_rot + math.Mod(2.0*pre_diff, TAU) - pre_diff

	to_diff := math.Mod(p_to-from_rot, TAU)
	to_rot := from_rot + math.Mod(2.0*to_diff, TAU) - to_diff

	post_diff := math.Mod(p_post-to_rot, TAU)
	post_rot := to_rot + math.Mod(2.0*post_diff, TAU) - post_diff

	return CubicInterpolate(from_rot, to_rot, pre_rot, post_rot, p_weight)
}

// CubicInterpolateInTime performs cubic interpolation in time domain using the Barry-Goldman method.
// It interpolates between 'p_from' and 'p_to' using 'p_pre' and 'p_post' as control points,
// with 'p_weight' determining the position between 'p_to' and 'p_from' in time.
// 'p_to_t', 'p_pre_t', and 'p_post_t' represent the time values corresponding to 'p_to', 'p_pre', and 'p_post' respectively.
func CubicInterpolateInTime(p_from, p_to, p_pre, p_post, p_weight, p_to_t, p_pre_t, p_post_t float64) float64 {
	/* Barry-Goldman method */
	t := Lerp(0.0, p_to_t, p_weight)
	a1 := p_pre
	if p_pre_t != 0 {
		a1 = Lerp(p_pre, p_from, (t-p_pre_t)/-p_pre_t)
	}

	a2 := p_from
	if p_to_t == 0 {
		a2 = Lerp(p_from, p_to, 0.5)
	} else {
		a2 = Lerp(p_from, p_to, t/p_to_t)
	}

	a3 := p_post
	f := p_post_t - p_to_t
	if f != 0 {
		a3 = Lerp(p_to, p_post, (t-p_to_t)/f)
	}

	b1 := a1
	f = p_to_t - p_pre_t
	if f != 0 {
		b1 = Lerp(a1, a2, (t-p_pre_t)/f)
	}

	b2 := a3
	if p_post_t != 0 {
		b2 = Lerp(a2, a3, t/p_post_t)
	}

	if p_to_t == 0 {
		return Lerp(b1, b2, 0.5)
	}
	return Lerp(b1, b2, t/p_to_t)
}

// CubicInterpolateAngleInTime performs cubic interpolation in time domain between two angles represented in radians.
// It ensures smooth interpolation by handling angle wrapping around the unit circle.
// The interpolation is done using the Barry-Goldman method.
func CubicInterpolateAngleInTime(p_from, p_to, p_pre, p_post, p_weight, p_to_t, p_pre_t, p_post_t float64) float64 {
	from_rot := math.Mod(p_from, TAU)

	pre_diff := math.Mod(p_pre-from_rot, TAU)
	pre_rot := from_rot + math.Mod(2.0*pre_diff, TAU) - pre_diff

	to_diff := math.Mod(p_to-from_rot, TAU)
	to_rot := from_rot + math.Mod(2.0*to_diff, TAU) - to_diff

	post_diff := math.Mod(p_post-to_rot, TAU)
	post_rot := to_rot + math.Mod(2.0*post_diff, TAU) - post_diff

	return CubicInterpolateInTime(from_rot, to_rot, pre_rot, post_rot, p_weight, p_to_t, p_pre_t, p_post_t)
}

// BezierInterpolate interpolates between two points using a cubic Bezier curve.
// It returns the interpolated value at position 'p_t' between 'p_start' and 'p_end' with control points 'p_control_1' and 'p_control_2'.
func BezierInterpolate(p_start, p_control_1, p_control_2, p_end, p_t float64) float64 {
	/* Formula from Wikipedia article on Bezier curves. */
	omt := (1.0 - p_t)
	omt2 := omt * omt
	omt3 := omt2 * omt
	t2 := p_t * p_t
	t3 := t2 * p_t

	return p_start*omt3 + p_control_1*omt2*p_t*3.0 + p_control_2*omt*t2*3.0 + p_end*t3
}

// BezierDerivative calculates the derivative of a cubic Bezier curve at a given position.
// It returns the derivative value at position 'p_t' between 'p_start' and 'p_end' with control points 'p_control_1' and 'p_control_2'.
func BezierDerivative(p_start, p_control_1, p_control_2, p_end, p_t float64) float64 {
	/* Formula from Wikipedia article on Bezier curves. */
	omt := (1.0 - p_t)
	omt2 := omt * omt
	t2 := p_t * p_t

	d := (p_control_1-p_start)*3.0*omt2 + (p_control_2-p_control_1)*6.0*omt*p_t + (p_end-p_control_2)*3.0*t2
	return d
}

// AngleDifference calculates the difference between two angles in radians.
// It returns the difference between 'p_from' and 'p_to' taking into account angle wrapping around the unit circle.
func AngleDifference(p_from, p_to float64) float64 {
	difference := math.Mod(p_to-p_from, TAU)
	return math.Mod(2.0*difference, TAU) - difference
}

// LerpAngle performs linear interpolation between two angles represented in radians.
// It returns the interpolated angle at position 'p_weight' between 'p_from' and 'p_to'.
func LerpAngle(p_from, p_to, p_weight float64) float64 {
	return p_from + AngleDifference(p_from, p_to)*p_weight
}

// InverseLerp calculates the interpolation parameter ('t') between two values 'p_from' and 'p_to' based on a given value 'p_value'.
// It returns the interpolation parameter that corresponds to 'p_value' relative to the range between 'p_from' and 'p_to'.
func InverseLerp(p_from, p_to, p_value float64) float64 {
	return (p_value - p_from) / (p_to - p_from)
}

// Remap remaps a value from one range to another.
// It linearly interpolates the value 'p_value' from the range defined by 'p_istart' and 'p_istop'
// to the range defined by 'p_ostart' and 'p_ostop'.
func Remap(p_value, p_istart, p_istop, p_ostart, p_ostop float64) float64 {
	return Lerp(p_ostart, p_ostop, InverseLerp(p_istart, p_istop, p_value))
}

// Smoothstep interpolates smoothly between two values based on a third value.
// It returns a value between 'p_from' and 'p_to' based on 'p_s', using Hermite interpolation.
func Smoothstep(p_from, p_to, p_s float64) float64 {
	if IsEqualApprox(p_from, p_to) {
		return p_from
	}
	s := Clampf((p_s-p_from)/(p_to-p_from), 0.0, 1.0)
	return s * s * (3.0 - 2.0*s)
}

// MoveToward moves a value towards another value by a given delta amount.
// It returns the value moved from 'p_from' towards 'p_to' by 'p_delta' amount.
func MoveToward(p_from, p_to, p_delta float64) float64 {
	if math.Abs(p_to-p_from) <= p_delta {
		return p_to
	}
	return p_from + Sign(p_to-p_from)*p_delta
}

// RotateToward rotates a value towards another value by a given delta amount.
// It returns the value rotated from 'p_from' towards 'p_to' by 'p_delta' amount.
func RotateToward(p_from, p_to, p_delta float64) float64 {
	difference := AngleDifference(p_from, p_to)
	abs_difference := math.Abs(difference)
	// When `p_delta < 0` move no further than to PI radians away from `p_to` (as PI is the max possible angle distance).
	r := p_from + Clampf(p_delta, abs_difference-PI, abs_difference)
	if difference >= 0.0 {
		return r * 1.0
	}
	return -1.0
}

// LinearToDb converts a linear value to decibels.
// It returns the logarithmic conversion of 'p_linear' value to decibels.
func LinearToDb(p_linear float64) float64 {
	return math.Log(p_linear) * 8.6858896380650365530225783783321
}

// DbToLinear converts a decibel value to linear scale.
// It returns the exponential conversion of 'p_db' value from decibels to linear scale.
func DbToLinear(p_db float64) float64 {
	return math.Exp(p_db * 0.11512925464970228420089957273422)
}

// Wrapi wraps an integer value within a specified range.
// It returns the wrapped value of 'value' within the range defined by 'min' and 'max'.
func Wrapi(value, min, max int) int {
	r := max - min
	if r == 0 {
		return min
	}
	return min + ((((value - min) % r) + r) % r)
}

// Wrapf wraps a float64 value within a specified range.
// It returns the wrapped value of 'value' within the range defined by 'min' and 'max'.
func Wrapf(value, min, max float64) float64 {
	rng := max - min
	if IsZeroApprox(rng) {
		return min
	}
	result := value - (rng * math.Floor((value-min)/rng))
	if IsEqualApprox(result, max) {
		return min
	}
	return result
}

// Fract returns the fractional part of a floating-point number.
// It returns the fractional part of 'value' after subtracting the integer part.
func Fract(value float64) float64 {
	return value - math.Floor(value)
}

// Pingpong calculates the ping-pong value within a specified length.
// It returns the ping-pong value of 'value' within the range defined by 'length'.
func Pingpong(value, length float64) float64 {
	if length == 0.0 {
		return length
	}
	return math.Abs(Fract((value-length)/(length*2.0))*length*2.0 - length)
}

// SnapScalar snaps a value to the nearest multiple of a step size.
// It returns the snapped value of 'p_target' relative to 'p_offset' and 'p_step'.
func SnapScalar(p_offset, p_step, p_target float64) float64 {
	if p_step == 0 {
		return Snapped(p_target-p_offset, p_step) + p_offset
	}
	return p_target
}

// SnapScalarSeparation snaps a value to the nearest multiple of a step size with a specified separation.
// It returns the snapped value of 'p_target' relative to 'p_offset', 'p_step', and 'p_separation'.
func SnapScalarSeparation(p_offset, p_step, p_target, p_separation float64) float64 {
	if p_step != 0 {
		a := Snapped(p_target-p_offset, p_step+p_separation) + p_offset
		b := a
		if p_target >= 0 {
			b -= p_separation
		} else {
			b += p_step
		}

		if math.Abs(p_target-a) < math.Abs(p_target-b) {
			return a
		}
		return b
	}
	return p_target
}
