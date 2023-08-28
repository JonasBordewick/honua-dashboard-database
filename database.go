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

// Variable für das Singeltone Pattern
var instance *HonuaDashboardDatabase

// GetInstance Methode für das Singletone Pattern. Diese Methode benötigt Parameter, damit, falls keine Instanz existiert, eine neue erstellt werden kann.
//
// db: ist ein String und gibt an wie die Datenbank heißt, zu der man sich verbinden möchte.
//
// user & password: sind die Login Daten, damit man zu mongodb eine Verbindung herstellen kann.
//
// host & port: sind die IP-Addresse + der Port
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

// AddDashboard fügt zur Datenbank ein (neues) Dashboard hinzu. Falls bereits
// ein Dashboard mit der ID existiert, dann wird das bereits existierende 
// Dashboard gelöscht und erst dann wird das neue hinzugefügt.
// Die Dashboard ID ist die identity des entsprechenden backends.
// Die Methode ist ADD + EDIT gleichzeitig
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

// GetDashboard gibt das Dashboard mit der angegebenen ID zurück
func (hddb *HonuaDashboardDatabase) GetDashboard(id string) (*models.Dashboard, error) {
	filter := bson.M{"_id": id}
	var result *models.Dashboard
	err := hddb.collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteDashboard löscht das Dashboard mit der angegebenen ID
func (hddb *HonuaDashboardDatabase) DeleteDashboard(id string) error {
	filter := bson.M{"_id": id}
	_, err := hddb.collection.DeleteOne(context.Background(), filter)
	return err
}

// Die Private Methode exists_dashboard checkt, ob bereits ein Dashboard unter
// der angegebenen ID existiert. Und gibt true zurück falls diese bereits
// existiert und false wenn es kein Dashboard mit der angegebenen ID gibt.
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
