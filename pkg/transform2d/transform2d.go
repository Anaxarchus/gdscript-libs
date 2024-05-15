package transform2d

/**************************************************************************/
/*  transform_2d.h                                                        */
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
	"github.com/Anaxarchus/zero-gdscript/pkg/vector2"
)

// Transform2D represents a 2D transformation matrix.
type Transform2D struct {
	Columns [3]vector2.Vector2 // A 3x2 matrix, using Vector2 for each column
}

// NewTransform2D creates a new Transform2D given a rotation and a translation vector.
func NewTransform2D(rot float64, pos vector2.Vector2) Transform2D {
	cr := math.Cos(rot)
	sr := math.Sin(rot)

	return Transform2D{
		Columns: [3]vector2.Vector2{
			vector2.New(cr, -sr),
			vector2.New(sr, cr),
			pos,
		},
	}
}

func Transform2DFromCells(xx, xy, yx, yy, ox, oy float64) Transform2D {
	return Transform2D{
		Columns: [3]vector2.Vector2{
			vector2.New(xx, xy),
			vector2.New(yx, yy),
			vector2.New(ox, oy),
		},
	}
}

func Transform2DFromColumns(x, y, origin vector2.Vector2) Transform2D {
	return Transform2D{
		Columns: [3]vector2.Vector2{
			x,
			y,
			origin,
		},
	}
}

func (t *Transform2D) GetRotation() float64 {
	return math.Atan2(t.Columns[0].Y, t.Columns[0].X)
}

func (t *Transform2D) SetRotation(p_rot float64) {
	scale := t.GetScale()
	cr := math.Cos(p_rot)
	sr := math.Sin(p_rot)
	t.Columns[0].X = cr
	t.Columns[0].Y = sr
	t.Columns[1].X = -sr
	t.Columns[1].Y = cr
	t.SetScale(scale)
}

func (t *Transform2D) GetScale() vector2.Vector2 {
	detSign := zerogdscript.Sign(t.determinant())
	return vector2.New(t.Columns[0].Length(), detSign*t.Columns[1].Length())
}

func (t *Transform2D) SetScale(p_scale vector2.Vector2) {
	t.Columns[0].Normalize()
	t.Columns[1].Normalize()
	t.Columns[0] = t.Columns[0].Mulf(p_scale.X)
	t.Columns[1] = t.Columns[1].Mulf(p_scale.Y)
}

func (t Transform2D) Translated(p_offset vector2.Vector2) Transform2D {
	// Equivalent to left multiplication
	return Transform2DFromColumns(t.Columns[0], t.Columns[1], t.Columns[2].Add(p_offset))
}

// ToLocal converts a point from global space to local space.
func (t Transform2D) ToLocal(point vector2.Vector2) vector2.Vector2 {
	return t.AffineInverse().Xform(point)
}

// ToGlobal converts a point from local space to global space.
func (t Transform2D) ToGlobal(point vector2.Vector2) vector2.Vector2 {
	return t.Xform(point)
}

// Inverse returns the inverse of the current transformation if it's a pure rotation.
func (t Transform2D) Inverse() Transform2D {
	// This assumes the matrix is a rotation matrix (no scaling).
	if t.determinant() == 0 {
		return Transform2D{}
	}
	return Transform2D{
		Columns: [3]vector2.Vector2{
			vector2.New(t.Columns[0].X, t.Columns[1].X),
			vector2.New(t.Columns[0].Y, t.Columns[1].Y),
			vector2.New(-t.tdotx(t.Columns[2]), -t.tdoty(t.Columns[2])),
		},
	}
}

// AffineInverse computes the matrix inverse handling potential scalings.
func (t Transform2D) AffineInverse() Transform2D {
	det := t.determinant()
	if det == 0 {
		return Transform2D{}
	}
	idet := 1.0 / det

	return Transform2D{
		Columns: [3]vector2.Vector2{
			vector2.New(t.Columns[1].Y*idet, -t.Columns[0].Y*idet),
			vector2.New(-t.Columns[1].X*idet, t.Columns[0].X*idet),
			vector2.New(-t.tdotx(t.Columns[2])*idet, -t.tdoty(t.Columns[2])*idet),
		},
	}
}

// Xform applies the transformation to a vector.
func (t Transform2D) Xform(vec vector2.Vector2) vector2.Vector2 {
	return vector2.New(t.tdotx(vec), t.tdoty(vec)).Add(t.Columns[2])
}

// tdotx calculates the dot product with the x-axis of the transformation.
func (t Transform2D) tdotx(v vector2.Vector2) float64 {
	return t.Columns[0].X*v.X + t.Columns[1].X*v.Y
}

// tdoty calculates the dot product with the y-axis of the transformation.
func (t Transform2D) tdoty(v vector2.Vector2) float64 {
	return t.Columns[0].Y*v.X + t.Columns[1].Y*v.Y
}

// determinant calculates the determinant of the transformation.
func (t Transform2D) determinant() float64 {
	return t.Columns[0].X*t.Columns[1].Y - t.Columns[1].X*t.Columns[0].Y
}
