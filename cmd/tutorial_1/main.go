package main

import (
	//	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

func main(){
  
//  ErroHandel();
    

// ArrayT();

// Strings();

// Structs();  // About Structs and interfaces
 
Pointers(); // pointers in go

}


func ErrorHandel(){
printPls("Hello U All!!");
	var result, remainder,err = divisionPls(6,0); 
 if(err!=nil){
   fmt.Println(err.Error());	
 }
  fmt.Printf("The result of the integer division is %v with remainder %v", result,remainder)
}

func printPls(str string){
  fmt.Println(str);
}

func divisionPls(n1 ,n2 int) (int, int, error){
        var err error
	if n2 == 0{
	err = errors.New("Cannot Divide by Zero");
	return 0,0, err;
	}
	divAns:=n1/n2;
	rimAns:=n1%n2;
 return divAns, rimAns, err;
 }

//Arrays


func ArrayT(){
   
  //var arr [4]int=[4]int{1,2,3,4};
	arr:= [3]int{11,22,33}

  fmt.Println(arr[1]);

  
  var sliceArr []int = []int {1,22,33};
  
   sliceArr= append(sliceArr, 2);
   fmt.Printf("The length of the array is %v \n", len(sliceArr));

   fmt.Printf("The cpacity of the array is %v", cap(sliceArr));
   
   
   var sliceArr1 []int = []int{55,66,77,88,99};
   sliceArr = append(sliceArr,sliceArr1...);

   fmt.Println(sliceArr)
   fmt.Printf("the size of the array is %v", len(sliceArr));
   fmt.Printf("the capacity of the array is %v \n", cap(sliceArr));
   


    var myMap map[string]int= make(map[string]int);
     fmt.Println(myMap);

    var myMap2  = map[string]int{"Addam":23, "Sarah":-44};
	fmt.Println(myMap2["Addam"]);

	if _,asd:=myMap2["Json"]; asd {
      fmt.Println(myMap2["Json"]);
     }else{
      fmt.Println("No name is Present");
     }
     // delete(<mapName>, <key Name>)
     delete(myMap2, "Addam");

    //Loops in Go 
    
     for name, age := range myMap2{
	 fmt.Printf("Name: %v, Age: %v \n", name, age);
	}
     
     for i, v := range sliceArr1{
         fmt.Printf("Index: %v, Value: %v \n", i , v);

	}
    
      //No while loop in Go but for is there to save the  day 

     var i int =0;
	for i<10{
		fmt.Println(i);
		i+=1;
	}
  

     n:=10000;
     var mySlice11 []int = []int{}; 
     var mySliceCap []int = make([]int , 0, n);

   fmt.Printf("The time taken to excute wihtout preallocation: %v \n" , timeLoop(mySlice11,n));
   fmt.Printf("The time taken to excute with preallocation: %v \n", timeLoop(mySliceCap,n));
}


func timeLoop(mySlice []int , n int)time.Duration{
  var t0 = time.Now();
        var i int=0;
	for i<n{
		mySlice = append(mySlice,i);
		i++;
	}
  
  return time.Since(t0);

}



func Strings(){
 
  var myString = "résumé"  
   fmt.Println(len(myString)) //it shows the bits it use to store résumé as é uses 2 bits
   fmt.Println(utf8.RuneCountInString(myString))  //this method is used to find the len

   var indexed = myString[1];
   fmt.Printf("%v %T \n", indexed, indexed); //This shows different value as it shows only 1 bit's val
   
   for i,v := range myString{
            		
          fmt.Println(i, v); // here range is doing extra work and jumped from 1 to 3 as é takes 2 bits to store also cuz of this val of string is also correct as it was not correct above;
	}
  
   //To skip all this mombo jombo we can do Strings declareation using an array of runes which is an alias for int32.

	var myString1 = []rune("résumé");  
   fmt.Println(len(myString1)) //it now shows the legth correctly its still showing the same thing number of bytes it takes to store but each cell in rune is of 32 bits or 4 bytes no UTF8 uses more than 4 bytes.

   var indexedd = myString[1];
   fmt.Printf("%v %T \n", indexedd, indexedd); //This is also showing the correct output now essentially cuz of runes
   
   for i,v := range myString{
            		
          fmt.Println(i, v); // here range is also now jumping the correct jumps that is one step at a time
	}

  // String Buildres

 // While it does supports concationation but always at time of concatination it creates a new string with new location all the time pretty inefficent right insted we can use String Builders now by importing strings package.


	var strSlice = []string{"g", "i" ,"t"," ", "i","s"," ","a","w","e","s","o","m","e"}  
        
	var strBuilder strings.Builder;

	for i := range strSlice{
		strBuilder.WriteString(strSlice[i]);
	}

	var catStr = strBuilder.String();
	fmt.Printf("\n %v", catStr)
     
        
}

type engine interface{
    milesLeft()   uint
}

type electricEngine struct {
	mpkwh uint
	kwh   int
}

func (e electricEngine)milesLeft()uint{
	return e.mpkwh * uint(e.kwh);
}
 

type gasEngine struct{
	mpg uint
	gallons int
	ownerName owner
}

//We can also give methods to a struct
func (gE gasEngine) milesLeft()uint{
	return gE.mpg * uint(gE.gallons);
}

type owner struct{
	name string
}

func canMakeIt(e engine, miles int){
   
	if miles<=int(e.milesLeft()){
		fmt.Println("You can make it there");
	}else{
		fmt.Println("You cannot make it fuel up fast!!");
	}

}


func Structs(){
	var myEngine gasEngine = gasEngine{gallons: 23, mpg: 45,ownerName:  owner{"Akash"}}
	 fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerName.name);

	myEngine = gasEngine{23, 45, owner{"Abhay"}};
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerName.name)
       
	//Anonymous Structs
	// these structs are declared and used at that time itself and cannot be reused

	var myEngine2 = struct{
		name string
		age int
	}{"Akash", 23};

	fmt.Println(myEngine2.name, myEngine2.age);
         

	fmt.Println(myEngine.milesLeft());

	canMakeIt(myEngine, 11);
	
	var myEEngine electricEngine = electricEngine{4, 23};
	canMakeIt(myEEngine, 98);
      
}


func Pointers(){
  
	//var p *int32 =;

	//fmt.Println(p);

}
