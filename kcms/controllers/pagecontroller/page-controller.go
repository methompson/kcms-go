package pagecontroller

// PageController handles all page duties
type PageController interface {
	CheckUserForSiteMod()
	GetPageByID(id string)
	GetPageBySlug(slug string)
	AddPage()
	EditPage()
	DeletePage()
}

// BasePageController is a base implementation of the PageController with
// definitions of common functions.
type BasePageController struct{}

// CheckUserForSiteMod checks if a user's credentials are allowed to modify pages
func (inst BasePageController) CheckUserForSiteMod() {}

// GetPageByID gets a page object using an id
func (inst BasePageController) GetPageByID(id string) {}

// GetPageBySlug gets a page object using a slug
func (inst BasePageController) GetPageBySlug(slug string) {}

// AddPage adds a page object to storage
func (inst BasePageController) AddPage() {}

// EditPage edits a page in storage
func (inst BasePageController) EditPage() {}

// DeletePage removes a page from storage
func (inst BasePageController) DeletePage() {}
