package testdata

type HasUnexportedField struct {
	//lint:ignore U1000 used only for testing
	a string
	A string
}
