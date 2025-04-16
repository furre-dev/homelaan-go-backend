package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func InitFirebaseApp() (*firebase.App, error) {
	opt := option.WithCredentialsFile("C:/Users/Furre/Documents/GitHub/homelaan-gql/secrets/serviceKey.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}
