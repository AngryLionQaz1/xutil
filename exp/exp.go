package exp

import (
	"log"
	"os"
)

/**
错误处理
*/

func Exp(err error) {

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}
