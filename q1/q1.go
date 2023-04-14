package q1

import "fmt"

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	var soma float64
	var z float64
	var desconto float64
	z = 0

	for i := 0; i < len(purchaseHistory); i++ {
		soma = soma + purchaseHistory[i]
		z++
	}
	media := soma / z

	if currentPurchase <= 0 {
		return 0, fmt.Errorf("Valor da compra atual é invalido")
	}

	if media > 1000 {
		desconto = 20 / 100 * currentPurchase
		return desconto, nil
	} else if len(purchaseHistory) == 0 {
		desconto = 10 / 100 * currentPurchase
		return desconto, nil
	}

	return 0, nil
}
