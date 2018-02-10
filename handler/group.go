package handler

type AdminGroup struct {
	adminRouter Router
}

type RestrictedGroup struct {
	restrictedHogeRouter Router
	restrictedHugaRouter Router
}
