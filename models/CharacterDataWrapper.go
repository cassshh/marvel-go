package models

// CharacterDataWrapper struct
type CharacterDataWrapper struct {
	Code            int
	Status          string
	Copyright       string
	AttributionText string
	AttributionHTML string
	Data            struct{}
	Etag            string
}
