package pagecontroller

import "fmt"

// MySQLPageController is an implementation of PageController with MySQL
// implementations for retrieving data
type MySQLPageController struct {
	BasePageController
}

func (inst MySQLPageController) GetPageByID(id string) {
	fmt.Println("Inside MySQLPageController GetPageByID")
}

func (inst MySQLPageController) GetPageBySlug(slug string) {
	fmt.Println("Inside MySQLPageController GetPageBySlug")
}

func (inst MySQLPageController) AddPage() {
	fmt.Println("Inside MySQLPageController AddPage")
}

func (inst MySQLPageController) EditPage() {
	fmt.Println("Inside MySQLPageController EditPage")
}

func (inst MySQLPageController) DeletePage() {
	fmt.Println("Inside MySQLPageController DeletePage")
}
