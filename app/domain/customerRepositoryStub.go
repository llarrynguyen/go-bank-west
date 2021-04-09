package domain


type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer,error) {
	return s.customers,nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer {
		{"1001","Larry","Las Vegas","33333","1988-01-01","1"},
		{"1002","Larry","Las Vegas","33333","1988-01-01","1"},
	}
	return CustomerRepositoryStub{customers: customers}
}

