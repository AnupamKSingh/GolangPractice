package main
import (
	"fmt"
)

type Customer struct {
	Name string
	Age int
	Addr []Address
}

type Address struct {
	City string
	Country string
}


var cust = Customer {
		Name : "Anupam",
		Age:24,
		Addr : []Address {
				Address {
					City :"Bangalore",
					Country :"India",
					},
				Address {
					City :"Mangalore",
					Country : "India",
					},
		},
	}

func main () {
	fmt.Println (cust)
	cust.Name = "Deepoo"
	fmt.Println (cust)
	cust.Addr[1].Country = "Bharat"
	fmt.Println (cust)
	newCust := Customer {
		Name : "Deepoo",
		Age:23,
		Addr: []Address {
					Address {
						City: "Bengaluru",
						Country :"India",
					},
					Address {
						City:"Poanta Sahib",
						Country :"India",
					},
			},
	}
	fmt.Println (newCust)
			
}
