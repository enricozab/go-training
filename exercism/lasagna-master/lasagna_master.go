// Package lasagna prepares and cooks your brilliant lasagna
package lasagna

// TODO: define the 'PreparationTime()' function
// PreparationTime returns number of minutes need to prepare a lasagna based on the layers and the average preparation time per layer provided
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		return len(layers) * 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
// Quantities returns the amount of noodles and sauce need based on the how many times they will occur in the layers provided
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64

	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
// AddSecretIngredient updates the myList by adding a secret ingreditiend from the friendList provided
func AddSecretIngredient(friendsList []string, myList []string) {
	myList = append(myList[:len(myList)-1], friendsList[len(friendsList)-1:]...)
}

// TODO: define the 'ScaleRecipe()' function
// ScaleRecipe returns a slice containing the scaled up/down of the quantities based on the portions provided
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var newQuantities []float64

	for _, quantity := range quantities {
		if portions == 0 {
			newQuantities = append(newQuantities, 0)
		} else {
			newQuantities = append(newQuantities, quantity*(float64(portions)/2))
		}
	}

	return newQuantities
}
