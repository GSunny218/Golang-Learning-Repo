package main
import "fmt";
//This are function signatures
//               |           |   
func concat(s1 string, s2 string) string {
	return s1 + s2;
}
func main() {
	fmt.Println(concat("Sunny ","Hi"))
}