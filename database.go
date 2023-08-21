package honuadashboarddatabase

import (
	"context"
	"fmt"
	"log"

	"github.com/JonasBordewick/honua-dashboard-database/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HonuaDashboardDatabase struct {
	collection *mongo.Collection
}

var instance *HonuaDashboardDatabase

func GetHonuaDashboardDatabaseInstance(db, user, password, host string, port int) *HonuaDashboardDatabase {
	if instance == nil {
		// Verbindungsoptionen
		clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", user, password, host, port))

		// Verbindung herstellen
		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		// Verbindung überprüfen
		err = client.Ping(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Erfolgreich zur MongoDB verbunden!")

		// Datenbank und Sammlung auswählen
		database := client.Database(db)
		collection := database.Collection("dashboards")

		instance = &HonuaDashboardDatabase{
			collection: collection,
		}
	}

	return instance
}

func (hddb *HonuaDashboardDatabase) AddDashboard(dashboard *models.Dashboard) error {
	exist, err := hddb.exists_dashboard(dashboard.ID)
	if err != nil {
		return err
	}

	if exist {
		err = hddb.DeleteDashboard(dashboard.ID)
		if err != nil {
			return err
		}
	}
	_, err = hddb.collection.InsertOne(context.Background(), dashboard)
	return err
}

func (hddb *HonuaDashboardDatabase) GetDashboard(id string) (*models.Dashboard, error) {
	filter := bson.M{"_id": id}
	var result *models.Dashboard
	err := hddb.collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (hddb *HonuaDashboardDatabase) DeleteDashboard(id string) error {
	filter := bson.M{"_id": id}
	_, err := hddb.collection.DeleteOne(context.Background(), filter)
	return err
}

func (hddb *HonuaDashboardDatabase) exists_dashboard(id string) (bool, error) {
	filter := bson.M{"_id": id}
	var result *models.Dashboard
	err := hddb.collection.FindOne(context.Background(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
