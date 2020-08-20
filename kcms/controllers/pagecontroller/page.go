package pagecontroller

// Page is a data representation of a Page row in the database
type Page struct {
	id          string
	name        string
	slug        string
	enabled     bool
	content     string
	meta        string
	dateUpdated int64
	dateAdded   int64
}
