package firestore_adapter

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func (f *Firestore_Adapter_Implementation) GetRandomEventCard(collection, condition, operator, neededValue string) error {

	ctx := context.Background()

	err := f.client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {

		docs := f.client.Collection(collection).Where(condition, operator, neededValue).Documents(f.ctx)

		for {
			d, err := docs.Next()

			if err != nil {
				if err == iterator.Done {
					break
				}
				//handle error
			}

			log.Println(d)
			//do something with the document
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
