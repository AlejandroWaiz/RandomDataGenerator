package exit_port

import firestore_adapter "github.com/RandomDataGetter/adapters/exit_port/firestore_adapter"

type ExitPort struct {
}

type ExitPort_Interface interface {
	GetRandomEventCard(collection, condition, operator, neededValue string) error
}

func GetExitPortImplementation() ExitPort_Interface {

	f := firestore_adapter.GetFirestoreAdapterImplementation()

	return &f

}
