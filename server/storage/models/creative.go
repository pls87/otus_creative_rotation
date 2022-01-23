package models

import "time"

type Creative struct {
	ID   ID     `db:"ID"`
	Desc string `db:"description"`
}

type SlotCreative struct {
	SlotID     ID `db:"slot_id"`
	CreativeID ID `db:"creative_id"`
}

type Impression struct {
	ID         ID        `db:"ID"`
	SlotID     ID        `db:"slot_id"`
	CreativeID ID        `db:"creative_id"`
	SegmentID  ID        `db:"segment_id"`
	Time       time.Time `db:"time"`
}

type Conversion struct {
	ID         ID        `db:"ID"`
	SlotID     ID        `db:"slot_id"`
	CreativeID ID        `db:"creative_id"`
	SegmentID  ID        `db:"segment_id"`
	Time       time.Time `db:"time"`
}

type Stats struct {
	SlotID      ID     `db:"slot_id"`
	CreativeID  ID     `db:"creative_id"`
	Impressions uint64 `db:"impressions"`
	Conversions uint64 `db:"conversions"`
}
