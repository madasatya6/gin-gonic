package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	//"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"firebase.google.com/go/db"
)

//@reference https://medium.com/google-cloud/firebase-developing-serverless-functions-in-go-963cb011265d
var Client *db.Client
var CredentialFile = "coba-firebase.json" //file ini terletak di sejajar main.go
var DatabaseURL = "https://coba-1234-default-rtdb.firebaseio.com/"

func FirebaseConns(ctx context.Context) (*db.Client, *firebase.App, error) {
	
	//create connection 
	var (
		app *firebase.App
		err error 
	)

	opt := option.WithCredentialsFile(CredentialFile)
	conf := &firebase.Config{
		DatabaseURL: DatabaseURL, //https://your-project.firebaseio.com
	}

	app, err = firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Println(err.Error())
		return nil, nil, err
	}

	initDb, err := app.Database(ctx)
	if err != nil {
		log.Println(err.Error())
		return nil, nil, err
	}

	return initDb, app, nil
}

