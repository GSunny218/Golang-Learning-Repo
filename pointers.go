package main;
import ("fmt");
func changeNum(num *int) {
	*num = 5;
	fmt.Println("In changeNum", num);
}
func main() {
	// num := 5;
	// fmt.Println("Memory address", &num); // Prints address

	num := 1;
	changeNum(&num);
	fmt.Println("After changeNum in main: ",num);
}