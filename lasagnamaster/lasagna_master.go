package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, n int)int{
    if n == 0{
        n = 2
    }
    totalPreprationTime := len(layers) * n
    return totalPreprationTime
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int , float64){
    noodles := 0
    sauce := 0.0
    for _, items := range layers{
		if items == "noodles" {
			noodles +=50

		}
        if items == "sauce"{
			sauce +=0.2
		}     
        
    }
    return noodles, sauce
    
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}



// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe(quantities []float64, n int) []float64 {
    sliceD := make([]float64, len(quantities))
    for i, a := range quantities {
        sliceD[i] = a / 2 * float64(n)
    }
    return sliceD
}