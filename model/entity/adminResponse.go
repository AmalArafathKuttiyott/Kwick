package models

type AdminHomePageResponse struct {
	Users    int
	Products int
	Orders   int
	Revenue  int
	Sales    Sales
}
