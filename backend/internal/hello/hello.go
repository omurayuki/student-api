package hello

// User user type
type User struct {
	ID   int64    `json:"id,omitempty"`
	Name string   `json:"name,omitempty"`
	Addr *Address `json:"addr,omitempty"`
}

// Address address type
type Address struct {
	City   string
	ZIP    int
	LatLng [2]float64
}

var alex = User{
	ID:   0,
	Name: "Yuhei Nakasaka",
	Addr: &Address{
		City: "",
		ZIP:  0,
		LatLng: [2]float64{
			0.0,
			0.0,
		},
	},
}

// Hello writes a welcome string
func Hello() string {
	return "Hello, " + alex.Name
}
