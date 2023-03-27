package requests

type PlannedOrderComponentKey struct {
	PlannedOrder        []int `json:"PlannedOrder"`
	PlannedOrderItem    []int `json:"PlannedOrderItem"`
	IsMarkedForDeletion bool  `json:"IsMarkedForDeletion"`
}
