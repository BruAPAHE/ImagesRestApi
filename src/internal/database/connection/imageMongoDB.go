package connection

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

var ImagesCollection *mongo.Collection
var ImagesConnection *mongo.Database

func init() {

	db := database.NewMongoDB(database.NewConfig())

	client, err := db.OpenConnect(GeoLocationName)
	if err != nil {
		logger.Logger.Fatalw("Cannot open connection to db",
			"name", GeoMagneticName,
			"error", err,
		)
	}

	if err = client.Connect(context.TODO()); err != nil {
		logger.Logger.Fatalw("Cannot open connection to db",
			"name", GeoLocationName,
			"error", err,
		)
	}

	// Check the connection
	if err = client.Ping(context.TODO(), nil); err != nil {
		logger.Logger.Fatalw("Cannot ping mongo database",
			"name", GeoLocationName,
			"error", err,
		)
	}

	connect := client.Database(db.Config[GeoLocationName].DB)

	ImagesCollection = connect.Collection(NameCollection)
	GeoDbConnection = connect
}
