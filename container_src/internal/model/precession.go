// Package model provides precession calculations for coordinate transformation.
//
// The IAU constellation boundaries are officially defined at epoch B1875.0, but modern
// star catalogs use J2000.0 coordinates. To correctly identify which constellation
// contains a given star, we must precess the J2000.0 coordinates back to B1875.0
// before checking against the boundary data.
//
// This file implements the standard precession formula using Lieske (1979) precession
// angles. The math involves rotating the coordinate system through three Euler angles
// (zeta_A, z_A, theta_A) that account for the slow wobble of Earth's axis over time.
//
// Reference: Lieske, J.H. et al. (1977), "Expressions for the Precession Quantities
// Based upon the IAU (1976) System of Astronomical Constants", Astron. Astrophys. 58, 1-16.
package model

import "math"

func precessJ2000ToB1875(ra, dec float64) (float64, float64) {
	raRad := ra * math.Pi / 180.0
	decRad := dec * math.Pi / 180.0

	const T = -1.249670833

	T2 := T * T
	T3 := T2 * T

	zetaA := 2306.2181*T + 1.39656*T2 - 0.000139*T3
	zA := 2306.2181*T + 1.09468*T2 + 0.018203*T3
	thetaA := 2004.3109*T - 0.85330*T2 - 0.000217*T3

	zetaARad := zetaA * math.Pi / (180.0 * 3600.0)
	zARad := zA * math.Pi / (180.0 * 3600.0)
	thetaARad := thetaA * math.Pi / (180.0 * 3600.0)

	cosRA := math.Cos(raRad)
	sinRA := math.Sin(raRad)
	cosDec := math.Cos(decRad)
	sinDec := math.Sin(decRad)

	x0 := cosDec * cosRA
	y0 := cosDec * sinRA
	z0 := sinDec

	cosZetaA := math.Cos(zetaARad)
	sinZetaA := math.Sin(zetaARad)
	cosZA := math.Cos(zARad)
	sinZA := math.Sin(zARad)
	cosThetaA := math.Cos(thetaARad)
	sinThetaA := math.Sin(thetaARad)

	x1 := cosZetaA*x0 - sinZetaA*y0
	y1 := sinZetaA*x0 + cosZetaA*y0
	z1 := z0

	x2 := cosThetaA*x1 + sinThetaA*z1
	y2 := y1
	z2 := -sinThetaA*x1 + cosThetaA*z1

	x3 := cosZA*x2 - sinZA*y2
	y3 := sinZA*x2 + cosZA*y2
	z3 := z2

	newDec := math.Asin(z3)
	newRA := math.Atan2(y3, x3)

	newRADeg := newRA * 180.0 / math.Pi
	newDecDeg := newDec * 180.0 / math.Pi

	for newRADeg < 0 {
		newRADeg += 360.0
	}
	for newRADeg >= 360.0 {
		newRADeg -= 360.0
	}

	return newRADeg, newDecDeg
}
