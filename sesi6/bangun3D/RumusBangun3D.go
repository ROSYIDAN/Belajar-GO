package bangun3D

import "math"

type VolumTabung struct {
	Phi    float64
	R      int
	T      int
	Volume float64
}

// (ini input)(ini output){statement}
func VolTabung(r int, t int) VolumTabung {
	vol := math.Pi * math.Pow(float64(r), 2) * float64(t)
	volumeStruct := VolumTabung{
		Phi:    math.Pi,
		R:      r,
		T:      t,
		Volume: vol,
	}
	return volumeStruct
}
