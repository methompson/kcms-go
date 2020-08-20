package pagecontroller

// PageController handles all page duties
type PageController interface {
	GetPageByID(id string)
	GetPageBySlug(slug string)
	AddPage()
	EditPage()
	DeletePage()
	CheckUserForSiteMod()
}

// BasePageController is a base implementation of the PageController with
// definitions of common functions.
type BasePageController struct{}

func (inst BasePageController) GetPageByID(id string) {}

func (inst BasePageController) GetPageBySlug(slug string) {}

func (inst BasePageController) AddPage() {}

func (inst BasePageController) EditPage() {}

func (inst BasePageController) DeletePage() {}

func (inst BasePageController) CheckUserForSiteMod() {}
