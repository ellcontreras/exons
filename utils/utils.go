package utils

// CheckErr verify if exists an error in the var err
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

