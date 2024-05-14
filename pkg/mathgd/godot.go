package mathgd

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
