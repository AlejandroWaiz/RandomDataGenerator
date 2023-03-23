package firestore_adapter

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func (f *Firestore_adapter) GetRandomEventCard() error {

	countName := "Event_Count"

	query, err := f.client.Collection(os.Getenv("Pokemon_Collection_Name")).NewAggregationQuery().WithCount(countName).Get(f.ctx)

	if err != nil {
		return fmt.Errorf("[GetRandomCardOfType] error:%v", err)
	}

	dataCount, err := f.GetCountOfDataFromFirestore(query, countName)

	if err != nil {
		return fmt.Errorf("[GetRandomCardOfGivenType] func => %v")
	}

	rand.Seed(time.Now().UnixNano())

	chosenCard := rand.Intn(dataCount)

	log.Println(chosenCard)

	//TODO: Agregar TODAS las cartas con un id unico con el siguiente metodo.
	//result, err := f.client.Collection(os.Getenv("Pokemon_Collection_Name")).NewDoc().Set()

	return nil
}
