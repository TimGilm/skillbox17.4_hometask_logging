// skillbox17_4_hometask
// example of log in file: {"level":"info","msg":"i = 0 a =  0","time":"2022-06-20T03:16:28+04:00"}
package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var a int

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("ошибка открытия файла: %v\n", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("ошибка закрытия файла: %v\n", err)
		}
	}(file)

	for i := 0; i <= 10; i++ {
		a += i * i
		log.SetOutput(file)
		log.Infof("i = %v a = %v", i, a)
	}
}
