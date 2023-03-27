package requests

type ProcessType struct {
	BulkProcess       bool `json:"BulkProcess"`
	IndividualProcess bool `json:"IndividualProcess"`
	PluralitySpec     bool `json:"PluralitySpec"`
	RangeSpec         bool `json:"RangeSpec"`
}
