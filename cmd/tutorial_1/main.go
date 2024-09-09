package main

import "fmt"

func main(){
  printPls("Hello U All!!");
  fmt.Println(divisionPls(7,5));

}

func printPls(str string){
  fmt.Println(str);

}

func divisionPls(n1 ,n2 int) (int, int){
	divAns:=n1/n2;
	rimAns:=n1%n2;
 return divAns, rimAns;
}
