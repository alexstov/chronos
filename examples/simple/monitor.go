// Package examples provides chronos customization examples
package main

// API operation ID.
type API int

const (
	// APIUnknown unknown API.
	APIUnknown API = iota
	// ChronosSampleSimple example API.
	ChronosSampleSimple
	// ChronosSampleFull example API.
	ChronosSampleFull
)

func (a API) String() string {
	return [...]string{
		"Unknown",
		"ChronosSampleSimple",
		"ChronosSampleFull",
	}[a]
}

// AreaID operation ID.
type AreaID int

const (
	// OpsTotal total operation time.
	OpsTotal AreaID = iota
	// OpsUnknown unknown operation.
	OpsUnknown
	// OpsDatabase .
	OpsDatabase
	// OpsMiddleware .
	OpsMiddleware
	// OpsIO .
	OpsIO
)

func (o AreaID) String() string {
	return [...]string{"Total", "Unknown", "OpsDatabase", "OpsMiddleware", "OpsIO"}[o]
}
