package models

type Dashboard struct {
	ID      string `bson:"_id, omitempty"`
	Widgets []*Widget `bson:"widgets, omitempty"`
}

type WidgetType int32

const (
	ENTITY WidgetType = iota
	DEVICE
	WEATHER
	GROUP
	DEFAULT
)

type Widget struct {
	WidgetType        WidgetType `bson:"type"`
	Title             string     `bson:"title, omitempty"`
	Unit              string     `bson:"unit, omitempty"`
	EntityID          int32      `bson:"entity_id, omitempty"`
	SecondaryEntityID int32      `bson:"secondary_entity_id, omitempty"`
	Subtitle          string     `bson:"subtitle, omitempty"`
	SwitchRules       bool       `bson:"switch_rules, omitempty"`
	Cards             []*Widget  `bson:"cards, omitempty"`
}