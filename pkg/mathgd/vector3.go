package mathgd

/**************************************************************************/
/*  vector3.h                                                             */
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

import (
	"math"
)

type Vector3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

func NewVector3(x, y, z float64) Vector3 {
	return Vector3{
		X: x,
		Y: y,
		Z: z,
	}
}

func CopyVector3(vector Vector3) Vector3 {
	return NewVector3(vector.X, vector.Y, vector.Z)
}

func ZeroVector3() Vector3 {
	return NewVector3(0, 0, 0)
}

func OneVector3() Vector3 {
	return NewVector3(1, 1, 1)
}

func CrossVector3(a, b Vector3) Vector3 {
	return a.Cross(b)
}

func DotVector3(a, b Vector3) float64 {
	return a.Dot(b)
}

func (v *Vector3) set(x, y, z float64) {
	v.X = x
	v.Y = y
	v.Z = z
}

func (v Vector3) IsUp() bool {
	return v.IsEqualApprox(NewVector3(0, 0, -1))
}

func (v Vector3) IsDown() bool {
	return v.IsEqualApprox(NewVector3(0, 0, 1))
}

func (v Vector3) Add(with Vector3) Vector3 {
	v.set(v.X+with.X, v.Y+with.Y, v.Z+with.Z)
	return v
}

func (v Vector3) Addf(with float64) Vector3 {
	v.set(v.X+with, v.Y+with, v.Z+with)
	return v
}

func (v Vector3) Sub(with Vector3) Vector3 {
	v.set(v.X-with.X, v.Y-with.Y, v.Z-with.Z)
	return v
}

func (v Vector3) Subf(with float64) Vector3 {
	v.set(v.X-with, v.Y-with, v.Z-with)
	return v
}

func (v Vector3) Mul(with Vector3) Vector3 {
	v.set(v.X*with.X, v.Y*with.Y, v.Z*with.Z)
	return v
}

func (v Vector3) Mulf(with float64) Vector3 {
	v.set(v.X*with, v.Y*with, v.Z*with)
	return v
}

func (v Vector3) Div(with Vector3) Vector3 {
	if with.X == 0 {
		v.X = math.Inf(1)
	} else {
		v.X /= with.X
	}

	if with.Y == 0 {
		v.Y = math.Inf(1)
	} else {
		v.Y /= with.Y
	}

	if with.Z == 0 {
		v.Z = math.Inf(1)
	} else {
		v.Z /= with.Z
	}
	return v
}

func (v Vector3) Divf(with float64) Vector3 {
	if with == 0 {
		v.set(math.Inf(1), math.Inf(1), math.Inf(1))
	} else {
		v.set(v.X/with, v.Y/with, v.Z/with)
	}
	return v
}

func (v Vector3) Cross(with Vector3) Vector3 {
	v.set(
		(v.Y*with.Z)-(v.Z*with.Y),
		(v.Z*with.X)-(v.X*with.Z),
		(v.X*with.Y)-(v.Y*with.X),
	)
	return v
}

func (v Vector3) Dot(with Vector3) float64 {
	return v.X*with.X + v.Y*with.Y + v.Z*with.Z
}

func (v Vector3) Abs() Vector3 {
	v.set(math.Abs(v.X), math.Abs(v.Y), math.Abs(v.Z))
	return v
}

func (v Vector3) Sign() Vector3 {
	v.set(Sign(v.X), Sign(v.Y), Sign(v.Z))
	return v
}

func (v Vector3) Floor() Vector3 {
	v.set(math.Floor(v.X), math.Floor(v.Y), math.Floor(v.Z))
	return v
}

func (v Vector3) Ceil() Vector3 {
	v.set(math.Ceil(v.X), math.Ceil(v.Y), math.Ceil(v.Z))
	return v
}

func (v Vector3) Round() Vector3 {
	v.set(math.Round(v.X), math.Round(v.Y), math.Round(v.Z))
	return v
}

func (v Vector3) Lerp(to Vector3, weight float64) Vector3 {
	v.set(
		Lerp(v.X, to.X, weight),
		Lerp(v.Y, to.Y, weight),
		Lerp(v.Z, to.Z, weight),
	)
	return v
}

func (v Vector3) Slerp(to Vector3, weight float64) Vector3 {
	// This method seems more complicated than it really is, since we write out
	// the internals of some methods for efficiency (mainly, checking length).
	sl2 := v.LengthSquared()
	el2 := to.LengthSquared()
	if sl2 == 0.0 || el2 == 0.0 {
		// Zero length vectors have no angle, so the best we can do is either lerp or throw an error.
		return v.Lerp(to, weight)
	}
	axis := v.Cross(to)
	al2 := axis.LengthSquared()
	if al2 == 0.0 {
		// Colinear vectors have no rotation axis or angle between them, so the best we can do is lerp.
		return v.Lerp(to, weight)
	}
	axis = axis.Divf(math.Sqrt(al2))
	sl := math.Sqrt(sl2)
	rl := Lerp(sl, math.Sqrt(el2), weight)
	angle := v.AngleTo(to)
	return v.Rotated(axis, angle*weight).Mulf(rl / sl)
}

func (v Vector3) CubicInterpolate(b Vector3, pre_a Vector3, post_b Vector3, weight float64) Vector3 {
	v.X = CubicInterpolate(v.X, b.X, pre_a.X, post_b.X, weight)
	v.Y = CubicInterpolate(v.Y, b.Y, pre_a.Y, post_b.Y, weight)
	v.Z = CubicInterpolate(v.Z, b.Z, pre_a.Z, post_b.Z, weight)
	return v
}

func (v Vector3) CubicInterpolateInTime(b, pre_a, post_b Vector3, weight, b_t, pre_a_t, post_b_t float64) Vector3 {
	v.X = CubicInterpolateInTime(v.X, b.X, pre_a.X, post_b.X, weight, b_t, pre_a_t, post_b_t)
	v.Y = CubicInterpolateInTime(v.Y, b.Y, pre_a.Y, post_b.Y, weight, b_t, pre_a_t, post_b_t)
	v.Z = CubicInterpolateInTime(v.Z, b.Z, pre_a.Z, post_b.Z, weight, b_t, pre_a_t, post_b_t)
	return v
}

func (v Vector3) BezierInterpolate(control_1, control_2, end Vector3, t float64) Vector3 {
	v.X = BezierInterpolate(v.X, control_1.X, control_2.X, end.X, t)
	v.Y = BezierInterpolate(v.Y, control_1.Y, control_2.Y, end.Y, t)
	v.Z = BezierInterpolate(v.Z, control_1.Z, control_2.Z, end.Z, t)
	return v
}

func (v Vector3) BezierDerivative(control_1, control_2, end Vector3, t float64) Vector3 {
	v.X = BezierDerivative(v.X, control_1.X, control_2.X, end.X, t)
	v.Y = BezierDerivative(v.Y, control_1.Y, control_2.Y, end.Y, t)
	v.Z = BezierDerivative(v.Z, control_1.Z, control_2.Z, end.Z, t)
	return v
}

func (v Vector3) DistanceTo(to Vector3) float64 {
	return to.Sub(v).Length()
}

func (v Vector3) DistanceSquaredTo(to Vector3) float64 {
	return to.Sub(v).LengthSquared()
}

func (v Vector3) Posmod(mod float64) Vector3 {
	return NewVector3(Fposmod(v.X, mod), Fposmod(v.Y, mod), Fposmod(v.Z, mod))
}

func (v Vector3) Posmodv(modv Vector3) Vector3 {
	return NewVector3(Fposmod(v.X, modv.X), Fposmod(v.Y, modv.Y), Fposmod(v.Z, modv.Z))
}

func (v Vector3) Project(to Vector3) Vector3 {
	return to.Mulf((v.Dot(to) / to.LengthSquared()))
}

func (v Vector3) AngleTo(to Vector3) float64 {
	return math.Atan2(v.Cross(to).Length(), v.Dot(to))
}

func (v Vector3) SignedAngleTo(to, axis Vector3) float64 {
	cross_to := v.Cross(to)
	unsigned_angle := math.Atan2(cross_to.Length(), v.Dot(to))
	sign := cross_to.Dot(axis)
	if sign < 0 {
		return -unsigned_angle
	}
	return unsigned_angle
}

func (v Vector3) DirectionTo(to Vector3) Vector3 {
	ret := NewVector3(to.X-v.X, to.Y-v.Y, to.Z-v.Z)
	ret.Normalize()
	return ret
}

func (v Vector3) Length() float64 {
	x2 := v.X * v.X
	y2 := v.Y * v.Y
	z2 := v.Z * v.Z

	return math.Sqrt(x2 + y2 + z2)
}

func (v Vector3) LengthSquared() float64 {
	x2 := v.X * v.X
	y2 := v.Y * v.Y
	z2 := v.Z * v.Z

	return x2 + y2 + z2
}

func (v *Vector3) Normalize() {
	lengthsq := v.LengthSquared()
	if lengthsq == 0 {
		v.set(0, 0, 0)
	} else {
		length := math.Sqrt(lengthsq)
		v.X /= length
		v.Y /= length
		v.Z /= length
	}
}

func (v Vector3) Normalized() Vector3 {
	v.Normalize()
	return v
}

func (v Vector3) IsNormalized() bool {
	// use length_squared() instead of length() to avoid sqrt(), makes it more stringent.
	return IsEqualApprox(v.LengthSquared(), 1.0)
}

func (v Vector3) IsEqualApprox(b Vector3) bool {
	return IsEqualApprox(v.X, b.X) && IsEqualApprox(v.Y, b.Y) && IsEqualApprox(v.Z, b.Z)
}

func (v Vector3) Inverse() Vector3 {
	v.set(1.0/v.X, 1.0/v.Y, 1.0/v.Z)
	return v
}

// slide returns the component of the vector along the given plane, specified by its normal vector.
func (v Vector3) Slide(normal Vector3) Vector3 {
	if !normal.IsNormalized() {
		return v
	}
	return v.Sub(normal.Mulf(v.Dot(normal)))
}

func (v Vector3) Bounce(normal Vector3) Vector3 {
	return v.Reflect(normal).Mulf(-1.0)
}

func (v Vector3) Reflect(normal Vector3) Vector3 {
	if !normal.IsNormalized() {
		return v
	}
	return normal.Mulf(v.Dot(normal)).Mulf(2.0).Sub(v)
	//return 2.0 * normal * Dot(normal) - v
}

// Rotate the current Vector3 around the provided axis by the specified angle.
func (v *Vector3) Rotate(axis Vector3, angle float64) {
	basis := NewBasisFromAxisAndAngle(axis, angle)
	*v = basis.Xform(*v)
}

// Return a new Vector3 that is the result of rotating the current Vector3 around the provided axis by the specified angle.
func (v Vector3) Rotated(axis Vector3, angle float64) Vector3 {
	rotatedVector := v
	rotatedVector.Rotate(axis, angle)
	return rotatedVector
}

func (v Vector3) ToVector2() Vector2 {
	return NewVector2(v.X, v.Y)
}
