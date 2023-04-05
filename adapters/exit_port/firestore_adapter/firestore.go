package firestore_adapter

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"cloud.google.com/go/firestore"
)

type Firestore_adapter struct {
	client    *firestore.Client
	ctx       context.Context
	projectID string
}

func GetFirestoreAdapterImplementation() Firestore_adapter {
	return Firestore_adapter{}
}

func (f *Firestore_adapter) getCountOfDataFromFirestore() (int, error) {

	countName := "Event_Count"

	query, err := f.client.Collection(os.Getenv("Pokemon_Collection_Name")).NewAggregationQuery().WithCount(countName).Get(f.ctx)

	queryCopy := make(map[string]interface{})

	for k, v := range query {
		queryCopy[k] = v
	}

	queryValue := queryCopy[countName]

	value := strings.Split(fmt.Sprint(queryValue), ":")

	dataCount, err := strconv.Atoi(value[1])

	if err != nil {
		return 0, fmt.Errorf("[GetCountOfDataFromFirestore] error: %v", err)

	}

	return dataCount, nil

}
