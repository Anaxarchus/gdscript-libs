package mathgd

import (
	"errors"
	"math"
)

/**************************************************************************/
/*  basis.h                                                               */
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

// A 3×3 matrix used for representing 3D rotation and scale.
// Usually used as an orthogonal basis for a Transform3D.

// Contains 3 vector fields X, Y and Z as its columns,
// which are typically interpreted as the local basis vectors of a transformation.
// For such use, it is composed of a scaling and a rotation matrix, in that order (M = R.S).

// Basis can also be accessed as an array of 3D vectors.
// These vectors are usually orthogonal to each other,
// but are not necessarily normalized (due to scaling).

type Basis struct {
	Rows [3][3]float64
}

func NewBasis() Basis {
	return Basis{
		Rows: [3][3]float64{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
	}
}

func NewBasisFromAxisAndAngle(axis Vector3, angle float64) Basis {
	basis := NewBasis()
	basis.SetAxisAngle(axis, angle)
	return basis
}

func (b *Basis) Set(pXX, pXY, pXZ, pYX, pYY, pYZ, pZX, pZY, pZZ float64) {
	b.Rows[0] = [3]float64{pXX, pXY, pXZ}
	b.Rows[1] = [3]float64{pYX, pYY, pYZ}
	b.Rows[2] = [3]float64{pZX, pZY, pZZ}
}

// SetColumns sets the columns of the basis matrix.
func (b *Basis) SetColumns(x, y, z Vector3) {
	b.SetColumn(0, x)
	b.SetColumn(1, y)
	b.SetColumn(2, z)
}

// GetColumn returns the specified column of the basis matrix.
func (b Basis) GetColumn(index int) Vector3 {
	// Get actual basis axis column (we store transposed as Rows for performance).
	return Vector3{X: b.Rows[0][index], Y: b.Rows[1][index], Z: b.Rows[2][index]}
}

// SetColumn sets the specified column of the basis matrix.
func (b *Basis) SetColumn(index int, value Vector3) {
	// Set actual basis axis column (we store transposed as Rows for performance).
	b.Rows[0][index] = value.X
	b.Rows[1][index] = value.Y
	b.Rows[2][index] = value.Z
}

// GetMainDiagonal returns the main diagonal of the basis matrix.
func (b Basis) GetMainDiagonal() Vector3 {
	return Vector3{X: b.Rows[0][0], Y: b.Rows[1][1], Z: b.Rows[2][2]}
}

// TransposeXform returns the result of transposing and multiplying the provided basis matrix with this basis matrix.
func (b Basis) TransposeXform(m Basis) Basis {
	return Basis{
		Rows: [3][3]float64{
			{
				b.Rows[0][0]*m.Rows[0][0] + b.Rows[1][0]*m.Rows[1][0] + b.Rows[2][0]*m.Rows[2][0],
				b.Rows[0][0]*m.Rows[0][1] + b.Rows[1][0]*m.Rows[1][1] + b.Rows[2][0]*m.Rows[2][1],
				b.Rows[0][0]*m.Rows[0][2] + b.Rows[1][0]*m.Rows[1][2] + b.Rows[2][0]*m.Rows[2][2],
			},
			{
				b.Rows[0][1]*m.Rows[0][0] + b.Rows[1][1]*m.Rows[1][0] + b.Rows[2][1]*m.Rows[2][0],
				b.Rows[0][1]*m.Rows[0][1] + b.Rows[1][1]*m.Rows[1][1] + b.Rows[2][1]*m.Rows[2][1],
				b.Rows[0][1]*m.Rows[0][2] + b.Rows[1][1]*m.Rows[1][2] + b.Rows[2][1]*m.Rows[2][2],
			},
			{
				b.Rows[0][2]*m.Rows[0][0] + b.Rows[1][2]*m.Rows[1][0] + b.Rows[2][2]*m.Rows[2][0],
				b.Rows[0][2]*m.Rows[0][1] + b.Rows[1][2]*m.Rows[1][1] + b.Rows[2][2]*m.Rows[2][1],
				b.Rows[0][2]*m.Rows[0][2] + b.Rows[1][2]*m.Rows[1][2] + b.Rows[2][2]*m.Rows[2][2],
			},
		},
	}
}

// Set the basis matrix to represent a rotation around the given axis by the specified angle.
func (b *Basis) SetAxisAngle(axis Vector3, angle float64) {
	// Ensure axis is normalized (optional check)
	// if !axis.IsNormalized() {
	//     // Optionally handle error
	//     return
	// }

	// Compute squared components of the axis
	axisSq := Vector3{X: axis.X * axis.X, Y: axis.Y * axis.Y, Z: axis.Z * axis.Z}

	// Compute cosine and sine of the angle
	cosine := math.Cos(angle)
	sine := math.Sin(angle)

	// Compute intermediate values
	t := 1 - cosine
	xyzt := axis.X * axis.Y * t
	zyxs := axis.Z * sine

	// Set elements of the basis matrix
	b.Rows[0][0] = axisSq.X + cosine*(1.0-axisSq.X)
	b.Rows[1][1] = axisSq.Y + cosine*(1.0-axisSq.Y)
	b.Rows[2][2] = axisSq.Z + cosine*(1.0-axisSq.Z)

	b.Rows[0][1] = xyzt - zyxs
	b.Rows[1][0] = xyzt + zyxs

	xyzt = axis.X * axis.Z * t
	zyxs = axis.Y * sine
	b.Rows[0][2] = xyzt + zyxs
	b.Rows[2][0] = xyzt - zyxs

	xyzt = axis.Y * axis.Z * t
	zyxs = axis.X * sine
	b.Rows[1][2] = xyzt - zyxs
	b.Rows[2][1] = xyzt + zyxs
}

func (b Basis) Xform(pVector Vector3) Vector3 {
	return NewVector3(
		b.rowToVector3(0).Dot(pVector),
		b.rowToVector3(1).Dot(pVector),
		b.rowToVector3(2).Dot(pVector),
	)
}

func (b *Basis) rowToVector3(rowIdx int) Vector3 {
	row := b.Rows[rowIdx]
	vec := NewVector3(row[0], row[1], row[2])
	return vec
}

func (b *Basis) Determinant() float64 {
	return b.Rows[0][0]*(b.Rows[1][1]*b.Rows[2][2]-b.Rows[2][1]*b.Rows[1][2]) -
		b.Rows[1][0]*(b.Rows[0][1]*b.Rows[2][2]-b.Rows[2][1]*b.Rows[0][2]) +
		b.Rows[2][0]*(b.Rows[0][1]*b.Rows[1][2]-b.Rows[1][1]*b.Rows[0][2])
}

// cofac calculates the cofactor of a 3x3 matrix.
func cofac(rows [3][3]float64, row1, col1, row2, col2 int) float64 {
	return rows[row1][col1]*rows[row2][col2] - rows[row1][col2]*rows[row2][col1]
}

// Invert inverts the Basis matrix.
func (b *Basis) Invert() error {
	co := [3]float64{
		cofac(b.Rows, 1, 1, 2, 2),
		cofac(b.Rows, 1, 2, 2, 0),
		cofac(b.Rows, 1, 0, 2, 1),
	}

	det := b.Rows[0][0]*co[0] + b.Rows[0][1]*co[1] + b.Rows[0][2]*co[2]

	// Check for zero determinant
	if det == 0 {
		return errors.New("matrix is not invertible, determinant is zero")
	}

	s := 1.0 / det

	// Set the new values of the matrix
	b.Rows[0] = [3]float64{co[0] * s, cofac(b.Rows, 0, 2, 2, 1) * s, cofac(b.Rows, 0, 1, 1, 2) * s}
	b.Rows[1] = [3]float64{co[1] * s, cofac(b.Rows, 0, 0, 2, 2) * s, cofac(b.Rows, 0, 2, 1, 0) * s}
	b.Rows[2] = [3]float64{co[2] * s, cofac(b.Rows, 0, 1, 2, 0) * s, cofac(b.Rows, 0, 0, 1, 1) * s}

	return nil
}
