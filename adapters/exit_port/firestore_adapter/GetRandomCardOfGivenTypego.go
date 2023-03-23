package firestore_adapter

import (
	"fmt"
	"math/rand"
	"time"
)

func (f *Firestore_adapter) GetRandomEventCard() error {

	countName := "Event_Count"

	query, err := f.client.Collection("pokemons").NewAggregationQuery().WithCount(countName).Get(f.ctx)

	if err != nil {
		return fmt.Errorf("[GetRandomCardOfType] error:%v", err)
	}

	dataCount, err := f.GetCountOfDataFromFirestore(query, countName)

	if err != nil {
		return fmt.Errorf("[GetRandomCardOfGivenType] func => %v")
	}

	rand.Seed(time.Now().UnixNano())

	chosenCard := rand.Intn(dataCount)

	return nil
}
