package main

import "fmt"


func PreparationTime(layers []string, n int)int{
    if n == 0{
        n = 2
    }
    totalPreprationTime := len(layers) * n
    return totalPreprationTime
}
func AddSecretIngredient(friendsList , myList []string){
    for _, items := range myList{
        //for _, item := range myList{
            if items == "?"{
                items = friendsList[len(myList)-1]
            }
       // }
    }
}
func main() {
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	total := PreparationTime(layers, 0)
	fmt.Println(total)


	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
myList := []string{"noodles", "meat", "sauce", "mozzarella","?"}

AddSecretIngredient(friendsList, myList)
}