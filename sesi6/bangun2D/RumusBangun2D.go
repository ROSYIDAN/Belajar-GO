package bangun2D

type DataLuasPersegi struct {
	Sisi int
	Luas float64
}

// type struct menjadi method yang dipanggil pada fungsi LuasPersegi
// return menjadi var dataLuas yang ter assign dengan fungsi struct
func LuasPersegi(sisi int) DataLuasPersegi {
	var luas = sisi * sisi
	var dataLuas = DataLuasPersegi{
		Sisi: sisi,
		Luas: float64(luas),
	}
	return dataLuas
}
