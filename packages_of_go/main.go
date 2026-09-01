package main
//To  make go.mod file, run the command "go mod init github.com/sunny/coder" in the terminal
import (
	"fmt"
	"github.com/sunny/coder/user"  // Import the user package
	"github.com/sunny/coder/auth" // Import the auth package
	"github.com/fatih/color"  // Import the color package. It's a third-party package, so we need to install it first using the command "go get github.com/fatih/color" in the terminal
);
func main() {
	auth.LoginWithCredentials("sunny", "123");
	session := auth.GetSession();
	fmt.Println("Session: ",session);

	user := user.User{
		Email: "user@email.com",
		Name: "Sunny",
	}
	//fmt.Println(user.Email, user.Name);
	color.Red(user.Email); //Print the email in red color using the color package
	color.Green(user.Name); //Print the name in green color using the color package	
}