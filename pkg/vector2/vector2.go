package vector2

/**************************************************************************/
/*  vector2.h                                                             */
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

	zerogdscript "github.com/Anaxarchus/zero-gdscript"
)

type Vector2 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func New(x, y float64) Vector2 {
	return Vector2{X: x, Y: y}
}

func Zero() Vector2 {
	return New(0, 0)
}

func One() Vector2 {
	return New(1, 1)
}

func (v Vector2) Add(b Vector2) Vector2 {
	v.X += b.X
	v.Y += b.Y
	return v
}

func (v Vector2) Sub(b Vector2) Vector2 {
	v.X -= b.X
	v.Y -= b.Y
	return v
}

func (v Vector2) Mul(b Vector2) Vector2 {
	v.X *= b.X
	v.Y *= b.Y
	return v
}

func (v Vector2) Div(b Vector2) Vector2 {
	if b.X == 0 {
		v.X = math.Inf(1)
	} else {
		v.X /= b.X
	}

	if b.Y == 0 {
		v.Y = math.Inf(1)
	} else {
		v.Y /= b.Y
	}
	return v
}

func (v Vector2) Addf(s float64) Vector2 {
	v.X += s
	v.Y += s
	return v
}

func (v Vector2) Subf(s float64) Vector2 {
	v.X -= s
	v.Y -= s
	return v
}

func (v Vector2) Mulf(s float64) Vector2 {
	v.X *= s
	v.Y *= s
	return v
}

func (v Vector2) Divf(s float64) Vector2 {
	if s == 0 {
		v.X = math.Inf(1)
		v.Y = math.Inf(1)
	} else {
		v.X /= s
		v.Y /= s
	}
	return v
}

func (v Vector2) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

func (v Vector2) FromAngle(angle float64) Vector2 {
	v.X = math.Cos(angle)
	v.Y = math.Sin(angle)
	return v
}

func (v Vector2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector2) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v *Vector2) Normalize() {
	l := v.X*v.X + v.Y*v.Y
	if l != 0 {
		l = math.Sqrt(l)
		v.X /= l
		v.Y /= l
	}
}

func (v Vector2) Normalized() Vector2 {
	v.Normalize()
	return v
}

func (v Vector2) IsNormalized() bool {
	// use length_squared() instead of length() to avoid sqrt(), makes it more stringent.
	return zerogdscript.IsEqualApprox(v.LengthSquared(), 1)
}

func (v Vector2) DistanceTo(b Vector2) float64 {
	return math.Sqrt((v.X-b.X)*(v.X-b.X) + (v.Y-b.Y)*(v.Y-b.Y))
}

func (v Vector2) DistanceSquaredTo(b Vector2) float64 {
	return (v.X-b.X)*(v.X-b.X) + (v.Y-b.Y)*(v.Y-b.Y)
}

func (v Vector2) DirectionTo(p_to Vector2) Vector2 {
	v.X = p_to.X - v.X
	v.Y = p_to.Y - v.Y
	v.Normalize()
	return v
}

func (v Vector2) AngleTo(b Vector2) float64 {
	return math.Atan2(v.Cross(b), v.Dot(b))
}

func (v Vector2) AngleToPoint(b Vector2) float64 {
	return b.Sub(v).Angle()
}

func (v Vector2) Dot(b Vector2) float64 {
	return v.X*b.X + v.Y*b.Y
}

func (v Vector2) Cross(b Vector2) float64 {
	return v.X*b.Y - v.Y*b.X
}

func (v Vector2) Sign() Vector2 {
	v.X = zerogdscript.Sign(v.X)
	v.Y = zerogdscript.Sign(v.Y)
	return v
}

func (v Vector2) Floor() Vector2 {
	v.X = math.Floor(v.X)
	v.Y = math.Floor(v.Y)
	return v
}

func (v Vector2) Ceil() Vector2 {
	v.X = math.Ceil(v.X)
	v.Y = math.Ceil(v.Y)
	return v
}

func (v Vector2) Round() Vector2 {
	v.X = math.Round(v.X)
	v.Y = math.Round(v.Y)
	return v
}

func (v Vector2) Rotated(x float64) Vector2 {
	sine := math.Sin(x)
	cosi := math.Cos(x)
	v.X = v.X*cosi - v.Y*sine
	v.Y = v.X*sine + v.Y*cosi
	return v
}

func (v Vector2) Posmod(x float64) Vector2 {
	v.X = zerogdscript.Fposmod(v.X, x)
	v.Y = zerogdscript.Fposmod(v.Y, x)
	return v
}

func (v Vector2) Posmodv(b Vector2) Vector2 {
	v.X = zerogdscript.Fposmod(v.X, b.X)
	v.Y = zerogdscript.Fposmod(v.Y, b.Y)
	return v
}

func (v Vector2) Project(b Vector2) Vector2 {
	return b.Mulf((v.Dot(b) / b.LengthSquared()))
}

func (v Vector2) Clampi(min, max Vector2) Vector2 {
	v.X = zerogdscript.Clampf(v.X, min.X, max.X)
	v.Y = zerogdscript.Clampf(v.Y, min.Y, max.Y)
	return v
}

func (v Vector2) Clampf(min, max float64) Vector2 {
	v.X = zerogdscript.Clampf(v.X, min, max)
	v.Y = zerogdscript.Clampf(v.Y, min, max)
	return v
}

func (v Vector2) Snapped(to Vector2) Vector2 {
	v.X = zerogdscript.Snapped(v.X, to.X)
	v.Y = zerogdscript.Snapped(v.Y, to.Y)
	return v
}

func (v Vector2) Snappedf(to float64) Vector2 {
	v.X = zerogdscript.Snapped(v.X, to)
	v.Y = zerogdscript.Snapped(v.Y, to)
	return v
}

func (v Vector2) LimitLength(maxLength float64) Vector2 {
	l := v.Length()
	res := v
	if l > 0 && maxLength < l {
		res = res.Divf(l)
		res = res.Mulf(maxLength)
	}
	return res
}

func (v Vector2) MoveToward(to Vector2, delta float64) Vector2 {
	vd := to.Sub(v)
	len := vd.Length()
	if len <= delta || len <= zerogdscript.CMP_EPSILON {
		return to
	}
	return vd.Divf(len).Mulf(delta).Add(v)
}

// slide returns the component of the vector along the given plane, specified by its normal vector.
func (v Vector2) Slide(normal Vector2) Vector2 {
	if !normal.IsNormalized() {
		panic("normal:Vector2 must be normalized before function:Vector2.Slide")
	}
	return v.Sub(normal.Mulf(v.Dot(normal)))
}

func (v Vector2) Bound(b Vector2) Vector2 {
	return v.Reflect(b).Mulf(-1)
}

func (v Vector2) Reflect(normal Vector2) Vector2 {
	if !normal.IsNormalized() {
		panic("normal:Vector2 must be normalized before function:Vector2.Slide")
	}
	//return 2.0f * p_normal * dot(p_normal) - *this;
	return normal.Mulf(2.0).Mulf(v.Dot(normal)).Sub(v)
}

func (v Vector2) IsEqual(b Vector2) bool {
	return v.X == b.X && v.Y == b.Y
}

func (v Vector2) IsEqualApprox(b Vector2) bool {
	return zerogdscript.IsEqualApprox(v.X, b.X) && zerogdscript.IsEqualApprox(v.Y, b.Y)
}

func (v Vector2) IsZeroApprox() bool {
	return zerogdscript.IsZeroApprox(v.X) && zerogdscript.IsZeroApprox(v.Y)
}

func (v Vector2) IsFinite() bool {
	return !math.IsInf(v.X, 1) && !math.IsInf(v.Y, 1)
}
