package utils

import "log"

// CheckErr verify if exists an error in the var err
func CheckErr(err error, msg string) {
	if err != nil {
		log.Println(err, msg)
	}
}

