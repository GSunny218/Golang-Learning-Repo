package main

// maps -> hash, object, dictionary
import (
	"fmt"
	"maps"
)

func main() {
	//creating a map
	m := make(map[string]string);
	m["name"] = "golang";
	m["area"] = "backend";
	//get an element
	//IMP: if key does not exist in the map then it return zero value
	fmt.Println(m["name"], m["area"]);
	fmt.Println(m["phone"]);

	m1 := make(map[string]int);
	m1["age"] = 30;
	m1["height"] = 180;
	fmt.Println(m1["age"]);
	fmt.Println(len(m1)); //length of map

	delete(m1, "age"); //delete an element from map
	fmt.Println(m1);
	clear(m1); //clear the map
	fmt.Println(m1);  //Output: map[]

	m2 := map[string]int{"price" : 40, "phone" : 3};
	fmt.Println(m2); // output: map[phone:3 price:40]
	v, ok := m2["phone"];
	fmt.Println(v);
	if(!ok) {
		fmt.Println("all ok");
	} else {
		fmt.Println("not ok");
	}

	//Checking map equal or not
	m3 := map[string]int{"price": 40, "phone": 99}
	m4 := map[string]int{"price": 40, "phone": 99}
	fmt.Println(maps.Equal(m3, m4)); //Output: true,      Only use maps function for map objects.
	m5 := map[string]int{"price": 40, "phone": 98}
	m6 := map[string]int{"price": 40, "phone": 99}
	fmt.Println(maps.Equal(m5, m6)); // Output: false
}
