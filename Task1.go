// task 1 :- create the Person class using a struct in Go to represent individuals with attributes like name, age, and methods to introduce themselves, update their age, and check if they are eligible to vote.

package main
import "fmt"

type Person struct{
	Name string
	Age int
}

func(p Person)Introduce(){
	fmt.Printf("Hello... My name is %s and i am %d years old!\n",p.Name,p.Age)
}

func (p *Person)updateAge(newAge int){
	p.Age=newAge
}

func(p Person)canVote()bool{
return p.Age >= 18
}

func main(){
person := Person{Name: "Meghashree", Age:10}
person.Introduce()
person.updateAge(22)
person.Introduce()
person.canVote()
 
if person.canVote(){
	fmt.Println("You are eligible to vote")
}else{
	fmt.Println("You are not eligible to vote")
}
}



