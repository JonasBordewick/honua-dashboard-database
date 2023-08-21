package honuadashboarddatabase

import (
	"testing"

	"github.com/JonasBordewick/honua-dashboard-database/models"
)

var test_instance = GetHonuaDashboardDatabaseInstance("honua", "root", "example", "192.168.0.138", 27017)

var test_dashboard = &models.Dashboard{
	ID: "bordewickbgd",
	Widgets: []*models.Widget{
		{
			WidgetType: models.DEVICE,
			EntityID:   4,
			Title:      "Titel",
		},
	},
}

func TestAddDashboard(t *testing.T) {
	err := test_instance.AddDashboard(test_dashboard)
	if err != nil {
		t.Errorf("FAILED: got error %s", err.Error())
	}
}

func TestGetDashboard(t *testing.T) {
	dashboard, err := test_instance.GetDashboard(test_dashboard.ID)
	if err != nil {
		t.Errorf("FAILED: got error %s", err.Error())
	}
	if len(dashboard.Widgets) != len(test_dashboard.Widgets) {
		t.Errorf("FAILED: got %d as length expected %d", len(dashboard.Widgets), len(test_dashboard.Widgets))
	}
}

func TestDeleteDashboard(t *testing.T) {
	err := test_instance.DeleteDashboard(test_dashboard.ID)
	if err != nil {
		t.Errorf("FAILED: got error %s", err.Error())
	}
}
