package geometry2d

/**************************************************************************/
/*  geometry_2d.h                                                         */
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
	clipper "github.com/ctessum/go.clipper"
)

// PolyJoinType defines the type of join for polygon edges.
type JoinType int

const (
	JoinTypeSquare JoinType = iota
	JoinTypeRound
	JoinTypeMiter
)

// PolyEndType defines the end type for open paths.
type EndType int

const (
	EndTypePolygon EndType = iota
	EndTypeJoined
	EndTypeButt
	EndTypeSquare
	EndTypeRound
)

func GetClosestPointsBetweenSegments(p1, q1, p2, q2 vector2.Vector2) float64 {
	d1 := q1.Sub(p1) // Direction vector of segment S1.
	d2 := q2.Sub(p2) // Direction vector of segment S2.
	r := p1.Sub(p2)

	a := d1.Dot(d1) // Squared length of segment S1, always nonnegative.
	e := d2.Dot(d2) // Squared length of segment S2, always nonnegative.
	f := d2.Dot(r)
	var s, t float64
	// Check if either or both segments degenerate into points.
	if a <= zerogdscript.CMP_EPSILON && e <= zerogdscript.CMP_EPSILON {
		// Both segments degenerate into points.
		c1 := p1
		c2 := p2
		return math.Sqrt((c1.Sub(c2)).Dot(c1.Sub(c2)))
	}
	if a <= zerogdscript.CMP_EPSILON {
		// First segment degenerates into a point.
		s = 0.0
		t = f / e // s = 0 => t = (b*s + f) / e = f / e
		t = zerogdscript.Clampf(t, 0.0, 1.0)
	} else {
		c := d1.Dot(r)
		if e <= zerogdscript.CMP_EPSILON {
			// Second segment degenerates into a point.
			t = 0.0
			s = zerogdscript.Clampf(-c/a, 0.0, 1.0) // t = 0 => s = (b*t - c) / a = -c / a
		} else {
			// The general nondegenerate case starts here.
			b := d1.Dot(d2)
			denom := a*e - b*b // Always nonnegative.
			// If segments not parallel, compute closest point on L1 to L2 and
			// clamp to segment S1. Else pick arbitrary s (here 0).
			if denom != 0.0 {
				s = zerogdscript.Clampf((b*f-c*e)/denom, 0.0, 1.0)
			} else {
				s = 0.0
			}
			// Compute point on L2 closest to S1(s) using
			// t = Dot((P1 + D1*s) - P2,D2) / Dot(D2,D2) = (b*s + f) / e
			t = (b*s + f) / e

			//If t in [0,1] done. Else clamp t, recompute s for the new value
			// of t using s = Dot((P2 + D2*t) - P1,D1) / Dot(D1,D1)= (t*b - c) / a
			// and clamp s to [0, 1].
			if t < 0.0 {
				t = 0.0
				s = zerogdscript.Clampf(-c/a, 0.0, 1.0)
			} else if t > 1.0 {
				t = 1.0
				s = zerogdscript.Clampf((b-c)/a, 0.0, 1.0)
			}
		}
	}
	c1 := p1.Add(d1.Mulf(s))
	c2 := p2.Add(d2.Mulf(t))
	return math.Sqrt((c1.Sub(c2)).Dot(c1.Sub(c2)))
}

func GetClosestPointToSegment(point vector2.Vector2, segment [2]vector2.Vector2) vector2.Vector2 {
	p := point.Sub(segment[0])
	n := segment[1].Sub(segment[0])
	l2 := n.LengthSquared()
	if l2 < 1e-20 {
		return segment[0] // Both points are the same, just give any.
	}

	d := n.Dot(p) / l2

	if d <= 0.0 {
		return segment[0] // Before first point.
	} else if d >= 1.0 {
		return segment[1] // After first point.
	} else {
		return segment[0].Add(n.Mulf(d)) // Inside.
	}
}

func GetDistanceToSegment(point vector2.Vector2, segment [2]vector2.Vector2) float64 {
	return point.DistanceTo(GetClosestPointToSegment(point, segment))
}

func GetDistanceSquaredToSegment(point vector2.Vector2, segment [2]vector2.Vector2) float64 {
	return point.DistanceSquaredTo(GetClosestPointToSegment(point, segment))
}

func GetClosestPointToSegmentUncapped(point vector2.Vector2, segment [2]vector2.Vector2) vector2.Vector2 {
	p := point.Sub(segment[0])
	n := segment[1].Sub(segment[0])
	l2 := n.LengthSquared()
	if l2 < 1e-20 {
		return segment[0] // Both points are the same, just give any.
	}

	d := n.Dot(p) / l2

	return segment[0].Add(n.Mulf(d)) // Inside.
}

func LineIntersectsLine(from_a, dir_a, from_b, dir_b vector2.Vector2) vector2.Vector2 {
	// See http://paulbourke.net/geometry/pointlineplane/

	denom := dir_b.Y*dir_a.X - dir_b.X*dir_a.Y
	if zerogdscript.IsZeroApprox(denom) { // Parallel?
		return vector2.Zero()
	}

	v := from_a.Sub(from_b)
	t := (dir_b.X*v.Y - dir_b.Y*v.X) / denom
	return from_a.Add(dir_a.Mulf(t))
}

func SegmentIntersectsSegment(from_a, to_a, from_b, to_b vector2.Vector2) vector2.Vector2 {
	B := to_a.Sub(from_a)
	C := from_b.Sub(from_a)
	D := to_b.Sub(from_a)

	ABlen := B.Dot(B)
	if ABlen <= 0 {
		return vector2.Zero()
	}
	Bn := B.Divf(ABlen)
	C = vector2.New(C.X*Bn.X+C.Y*Bn.Y, C.Y*Bn.X-C.X*Bn.Y)
	D = vector2.New(D.X*Bn.X+D.Y*Bn.Y, D.Y*Bn.X-D.X*Bn.Y)

	// Fail if C x B and D x B have the same sign (segments don't intersect).
	if (C.Y < -zerogdscript.CMP_EPSILON && D.Y < -zerogdscript.CMP_EPSILON) || (C.Y > zerogdscript.CMP_EPSILON && D.Y > zerogdscript.CMP_EPSILON) {
		return vector2.Zero()
	}

	// Fail if segments are parallel or colinear.
	// (when A x B == zero, i.e (C - D) x B == zero, i.e C x B == D x B)
	if zerogdscript.IsEqualApprox(C.Y, D.Y) {
		return vector2.Zero()
	}

	ABpos := D.X + (C.X-D.X)*D.Y/(D.Y-C.Y)

	// Fail if segment C-D crosses line A-B outside of segment A-B.
	if (ABpos < 0) || (ABpos > 1) {
		return vector2.Zero()
	}

	// Apply the discovered position to line A-B in the original coordinate system.
	return from_a.Add(B.Mulf(ABpos))
}

func OffsetPolygon(polygon []vector2.Vector2, delta float64, joinType JoinType) [][]vector2.Vector2 {
	return doOffset(polygon, delta, clipper.JoinType(joinType), clipper.EtClosedPolygon)
}

func OffsetPolyline(polygon []vector2.Vector2, delta float64, joinType JoinType, endType EndType) [][]vector2.Vector2 {
	if endType == EndTypePolygon {
		return [][]vector2.Vector2{}
	}
	return doOffset(polygon, delta, clipper.JoinType(joinType), clipper.EndType(endType))
}

// IsPolygonClockwise determines if the given polygon points are in a clockwise order.
func IsPolygonClockwise(polygon []*vector2.Vector2) bool {
	c := len(polygon)
	if c < 3 {
		return false
	}

	sum := 0.0
	for i := 0; i < c; i++ {
		v1 := polygon[i]
		v2 := polygon[(i+1)%c]
		sum += (v2.X - v1.X) * (v2.Y + v1.Y)
	}

	return sum > 0
}

func toFixedPointPrecision(x, y float64) *clipper.IntPoint {
	return clipper.NewIntPointFromFloat(x*100000000, y*100000000)
}

func toFloatingPointPrecision(value *clipper.IntPoint) vector2.Vector2 {
	return vector2.New(float64(value.X), float64(value.Y)).Divf(100000000)
}

func doOffset(polygon []vector2.Vector2, delta float64, jt clipper.JoinType, et clipper.EndType) [][]vector2.Vector2 {
	clip := clipper.NewClipperOffset()
	path := clipper.NewPath()
	for _, pt := range polygon {
		iPt := toFixedPointPrecision(pt.X, pt.Y)
		path = append(path, iPt)
	}
	clip.AddPath(path, jt, et)

	clip.ArcTolerance = 0.0
	clip.MiterLimit = 4.0

	solutions := clip.Execute(delta * 100000000)
	if len(solutions) == 0 {
		return [][]vector2.Vector2{}
	}

	res := make([][]vector2.Vector2, len(solutions))
	for _, solution := range solutions {
		points := make([]vector2.Vector2, len(solution))
		for _, pt := range solution {
			points = append(points, toFloatingPointPrecision(pt))
		}
		res = append(res, points)
	}

	return res
}
