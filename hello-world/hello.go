package main

import "fmt"

const englishPrefix="Hello, "
const spanishPrefix="Hola, "

func Hello(name,language string) string{
	if name==""{
		return englishPrefix + "world";
	}
	if language=="Spanish"{
		return spanishPrefix + name;
	}
	return englishPrefix + name;
}
func main () {
	fmt.Println(Hello("Pratik","Spanish"));
	fmt.Println(Hello("Pratik",""));
	fmt.Println(Hello("",""));
}