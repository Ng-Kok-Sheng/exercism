package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timeToPrepare int) int {
	if timeToPrepare == 0 {
		return len(layers) * 2
	}

	return len(layers) * timeToPrepare
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0

	for layer := 0; layer < len(layers); layer++ {
		if layers[layer] == "noodles" {
			noodles += 50
		}

		if layers[layer] == "sauce" {
			sauce += 0.2
		}
	}

	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantity []float64, portion int) []float64 {
	scaledQuantity := []float64{}
	scaleMultiplier := 0.5
	for i := 0; i < len(quantity); i++ {
		scaledQuantity = append(scaledQuantity, quantity[i]*float64(portion)*scaleMultiplier)
	}

	return scaledQuantity
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
