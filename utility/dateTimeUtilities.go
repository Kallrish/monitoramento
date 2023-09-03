package utility

import "time"

func RetornaData() string {
	return time.Now().Format("02/01/2006")
}

func RetornaHora() string {
	return time.Now().Format("15:04:05")
}
