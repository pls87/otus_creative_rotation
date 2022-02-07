package queries

import "fmt"

const LocationRelation = "slot_creative"

type LocationQueries struct{}

func (l *LocationQueries) GetFor(primary, secondary string) string {
	return fmt.Sprintf(`SELECT s.* FROM "%s" sc
		INNER JOIN "%s" s ON sc.%s_id=s."ID"
		WHERE sc.%s_id = $1`, LocationRelation, secondary, secondary, primary)
}

func (l *LocationQueries) Create() string {
	return fmt.Sprintf(`INSERT INTO "%s" (creative_id, slot_id) VALUES ($1, $2)`, LocationRelation)
}

func (l *LocationQueries) Exists() string {
	return fmt.Sprintf(`SELECT * FROM "%s" WHERE creative_id = $1 AND slot_id = $2`, LocationRelation)
}

func (l *LocationQueries) Delete() string {
	return fmt.Sprintf(`DELETE FROM "%s" WHERE creative_id = $1 AND slot_id=$2`, LocationRelation)
}

func (l *LocationQueries) All() string {
	return fmt.Sprintf(`SELECT sc.slot_id, sc.creative_id, s.description as slot_desc, 
	cr.description as creative_desc 
		FROM "%s" sc 
		INNER JOIN "%s" s ON sc.slot_id=s."ID"
		INNER JOIN "%s" cr ON sc.creative_id=cr."ID"`, LocationRelation, SlotRelation, CreativeRelation)
}

var Location = LocationQueries{}
