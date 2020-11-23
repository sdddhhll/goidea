package customer

type Customer struct {
	ID int64
	Name string
	Gender string
	Age byte
	Tel string
	Mail string
}

func (cs *Customer) AddCustomer(css *[]Customer){
	*css = append(*css,*cs)
}


func (cs *Customer) DeleteCustomer(css *[]Customer,ID int64 ){
	var newCu []Customer = make([]Customer,0)
	for index,v := range *css {
		if v.ID != ID {
			newCu = append(newCu,(*css)[index])
		}
	}

	//这里能使用css = &newCu
	*css = newCu
}

func (cs *Customer) GetMaxCustomer(css *[]Customer) Customer{
	if len(*css) == 0 {
		return Customer{0,"","",0,"",""}
	}
	return (*css)[len(*css) - 1]
}

func (cs *Customer)GetCustomerById(css *[]Customer,ID int64) (*Customer,bool) {
	for index,v := range *css{
		if v.ID == ID {
			return &(*css)[index],true
		}
	}
	var cus Customer
	return &cus,false
}

