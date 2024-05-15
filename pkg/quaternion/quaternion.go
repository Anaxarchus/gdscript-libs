package quaternion

import (
	"math"

	zerogdscript "github.com/Anaxarchus/zero-gdscript"
	"github.com/Anaxarchus/zero-gdscript/pkg/vector3"
)

/**************************************************************************/
/*  quaternion.h                                                          */
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

// Quaternions can be parametrized using both an axis-angle pair or Euler angles.
// Due to their compactness and the way they are stored in memory,
// certain operations (obtaining axis-angle and performing SLERP, in particular)
// are more efficient and robust against floating-point errors.

// Note: Quaternions need to be normalized before being used for rotation.

// A unit quaternion used for representing 3D rotations.
type Quaternion struct {
	X float64
	Y float64
	Z float64
	W float64
}

// Constructs a quaternion defined by the given values.
func New(p_x, p_y, p_z, p_w float64) Quaternion {
	return Quaternion{
		X: p_x,
		Y: p_y,
		Z: p_z,
		W: p_w,
	}
}

// Constructs a default-initialized quaternion with all components set to 0.
func ZERO() Quaternion {
	return New(0, 0, 0, 0)
}

// The identity quaternion, representing no rotation. Equivalent to an identity Basis matrix.
// If a vector is transformed by an identity quaternion, it will not change.
func IDENTITY() Quaternion {
	return New(0, 0, 0, 1)
}

// TODO: Port Basis class from Godot
// Constructs a quaternion from the given Basis.
//func Basis() Quaternion {
//	return New(0, 0, 0, 0)
//}

// Constructs a quaternion that will rotate around the given axis by the specified angle. The axis must be a normalized vector.
func Rotated(axisNormal vector3.Vector3, angle float64) Quaternion {
	if !axisNormal.IsNormalized() {
		return IDENTITY()
	}
	return New(axisNormal.X, axisNormal.Y, axisNormal.Z, angle)
}

// Constructs a Quaternion as a copy of the given Quaternion.
func From(quaternion *Quaternion) Quaternion {
	return New(quaternion.X, quaternion.Y, quaternion.Z, quaternion.W)
}

// Constructs a quaternion representing the shortest arc between two points on the surface of a sphere with a radius of 1.0.
func Between(p_v0, p_v1 vector3.Vector3) Quaternion { // Shortest arc.
	c := p_v0.Cross(p_v1)
	d := p_v0.Dot(p_v1)

	if d < -1.0+zerogdscript.CMP_EPSILON {
		return New(0, 1, 0, 0)
	} else {
		s := math.Sqrt((1.0 + d) * 2.0)
		rs := 1.0 / s
		return New(c.X*rs, c.Y*rs, c.Z*rs, s*0.5)
	}
}
