package file

import (
	"log"
	"os"
	"time"
)

type HeavyApiLoadRepository struct {
}

func NewHeavyApiLoadRepository() *HeavyApiLoadRepository {
	return &HeavyApiLoadRepository{}
}

func (r *HeavyApiLoadRepository) Test() (string, error) {
	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("Началась обработка запроса !!!")
	time.Sleep(60 * time.Second)
	log.Println("Закончилась обработка запроса !!!")
	return "Запрос окончен", nil
}
