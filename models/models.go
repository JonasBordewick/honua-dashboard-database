package models

type Dashboard struct {
	ID      string    `bson:"_id, omitempty"`
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
	WidgetType        WidgetType `bson:"type" json:"type"`
	Title             string     `bson:"title,omitempty" json:"title,omitempty"`
	Unit              string     `bson:"unit,omitempty" json:"unit,omitempty"`
	EntityID          int32      `bson:"entity_id,omitempty" json:"entity_id,omitempty"`
	SecondaryEntityID int32      `bson:"secondary_entity_id,omitempty" json:"secondary_entity_id,omitempty"`
	Subtitle          string     `bson:"subtitle,omitempty" json:"subtitle,omitempty"`
	SwitchRules       bool       `bson:"switch_rules,omitempty" json:"switch_rules,omitempty"`
	Expandable        bool       `bson:"expandable,omitempty" json:"expandable,omitempty"`
	Cards             []*Widget  `bson:"cards,omitempty" json:"cards,omitempty"`
}
