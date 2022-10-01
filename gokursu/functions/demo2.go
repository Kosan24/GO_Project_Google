package functions

func QuatreOperations(nombre1 int, nombre2 int) (int, int, int, float32) {
	addition := nombre1 + nombre2
	soustraction := nombre1 - nombre2
	multiplication := nombre1 * nombre2
	division := float32(nombre1 / nombre2)

	return addition, soustraction, multiplication, division

}
