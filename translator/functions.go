package translator

func Soul(name string) string {
	vp := translate("soul")
	return vp.Sub(name).GetString("vsub")
}

func Kirin(name string) string {
	vp := translate("evol")
	return vp.GetString(name)
}

func Shikigami(name string, method string) string {
	vp := translate("shikigami")
	return vp.Sub(name).GetString(method)
}

func ErrCode(name string) string {
	vp := translate("err_code")
	return vp.GetString(name)
}
