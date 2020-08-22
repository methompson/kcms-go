package pagecontroller

import (
	"fmt"

	mysqlcontroller "com.methompson/go-test/kcms/db/mysql"
)

// MySQLPageController is an implementation of PageController with MySQL
// implementations for retrieving data
type MySQLPageController struct {
	BasePageController
	DbInstance mysqlcontroller.MySQLCMS
}

// GetPageByID gets a page object using an id
func (inst MySQLPageController) GetPageByID(id string) {
	fmt.Println("Inside MySQLPageController GetPageByID")
}

// GetPageBySlug gets a page object using a slug
func (inst MySQLPageController) GetPageBySlug(slug string) {
	fmt.Println("Inside MySQLPageController GetPageBySlug")
}

// AddPage adds a page object to storage
func (inst MySQLPageController) AddPage() {
	fmt.Println("Inside MySQLPageController AddPage")
}

// EditPage edits a page in storage
func (inst MySQLPageController) EditPage() {
	fmt.Println("Inside MySQLPageController EditPage")
}

// DeletePage removes a page from storage
func (inst MySQLPageController) DeletePage() {
	fmt.Println("Inside MySQLPageController DeletePage")
}
