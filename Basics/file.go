package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
		f, err := os.Open("example.txt"); //Opens files
		if(err != nil) {
			//log the error
			panic(err);
		}
		fileInfo, err := f.Stat();
		if(err != nil) {
			panic(err);
		}
		fmt.Println("File Name: ",fileInfo.Name()); //Returns file name
		fmt.Println("File or Folder: ",fileInfo.IsDir()); //Returns if file type is directory or not
		fmt.Println("File Size: ",fileInfo.Size()); //Returns file size
		fmt.Println("File Permission: ",fileInfo.Mode()); // Returns File mode
		fmt.Println("File modified at ",fileInfo.ModTime()); //Returns last modified time of file

	//Read file
	f, err := os.Open("example.txt");
	if err != nil {
		panic(err);
	}
	defer f.Close(); // Closes file
	buf := make([]byte, 12); //Read file in bytes
	d, err := f.Read(buf);
	if err != nil {
		panic(err);
	}
	for i := 0; i < len(buf); i++ {
		println("data", d, string(buf[i])); // Read file's each letter
	}

	//Read file contents
	data, err := os.ReadFile("example.txt"); //Read all file contents. Not efficient for large files
	if err != nil {
		panic(err);
	}
	fmt.Println(string(data));

	//read folder
	dir, err := os.Open("."); //Opens the current directory
	if err != nil {
		panic(err);
	}
	defer dir.Close();
	fileInfo, err := dir.ReadDir(-1); //Returns all directory names
	//fileInfo, err := dir.ReadDir(3); //Returns names of directory
	for _, fi := range fileInfo {
		fmt.Println(fi.Name(), fi.IsDir());
	}

	//Create a file
	f, err := os.Create("example2.txt") //Creates file
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("Hi, Golang") //Writes in file
	f.WriteString("Nice language");// Continues writing in file.
	bytes := []byte("Hello Golang"); //Overrides existing content of file.
	f.Write(bytes); //Write in bytes(like same as english)

	//Read and write to another file (streaming fashion)
	sourceFile, err := os.Open("example.txt");
	if err != nil {
		panic(err);
	}
	defer sourceFile.Close();
	destFile, err := os.Create("example.txt");
	if err != nil {
		panic(err);
	}
	defer destFile.Close();
	reader := bufio.NewReader(sourceFile); //Returns new buffered reader
	writer := bufio.NewWriter(destFile); //Returns new buffered writer
	for {
		b, err := reader.ReadByte(); //Read bytes and returns a single bytes.
		if err != nil {
			if err.Error() != "EOF" { //Checks until end of file(EOF)
				panic(err);
			}
			break;
		}
		e := writer.WriteByte(b); //Writes a single bytes
		if e != nil {
			panic(e)
		}
	}
	writer.Flush();
	fmt.Println("Written to new file successfully!");

	Delete a file
	err := os.Remove("example2.txt"); //deletes file
	if err != nil {
		panic(err);
	}
	fmt.Println("File deleted successfully!");
}
