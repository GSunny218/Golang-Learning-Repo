package auth;
func extractSession() string {  //Private function, not accessible outside this package
	return "Logged in session for user";
}
func GetSession() string {
	return extractSession();  //Calling the private function from a public function within the same package
}