package firestore_adapter

import (
	"context"
	"fmt"
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

func (i *Firestore_adapter) GetCountOfDataFromFirestore(originalMap map[string]interface{}, countName string) (int, error) {

	mapCopy := make(map[string]interface{})

	for k, v := range originalMap {
		mapCopy[k] = v
	}

	queryValue := mapCopy[countName]

	value := strings.Split(fmt.Sprint(queryValue), ":")

	dataCount, err := strconv.Atoi(value[1])

	if err != nil {
		return 0, fmt.Errorf("[GetCountOfDataFromFirestore] error: %v", err)

	}

	return dataCount, nil

}
