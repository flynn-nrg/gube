package gube

func lerp(v0 float64, v1 float64, t float64) float64 {
	return (1-t)*v0 + t*v1
}
