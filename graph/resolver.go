package graph

import (
	"com.methompson/kcms-go/kcms"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the bridge between the application and the database instance
type Resolver struct {
	KCMS *kcms.KCMS
}
