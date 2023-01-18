package utils

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}
