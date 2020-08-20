package blogpostcontroller

// BlogPost is a data representation of a BlogPost row in the database
type BlogPost struct {
	id          string
	name        string
	slug        string
	draft       bool
	public      bool
	content     string
	meta        string
	dateUpdated int64
	dateAdded   int64
}
