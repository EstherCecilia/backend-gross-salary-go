package main

import (
	"fmt"
	"gross-salary/utils"
)

func main() {
	salary := utils.Salary{Value: 10437.00}

	fmt.Println("Cálculo do Salário Líquido Padrão")
	fmt.Printf("Salário Bruto: R$ %.2f\n", salary.Value)
	fmt.Printf("Desconto INSS: R$ %.2f\n", salary.CalculateINSS())
	fmt.Printf("Desconto IR: R$ %.2f\n", salary.CalculateIR())
	fmt.Printf("Salário Líquido: R$ %.2f\n", salary.CalculateSalarioLiquido(false))

	fmt.Println()

	fmt.Println("Cálculo do Salário Líquido Novo")
	fmt.Printf("Salário Bruto: R$ %.2f\n", salary.Value)
	fmt.Printf("Desconto INSS: R$ %.2f\n", salary.CalculateINSS())
	fmt.Printf("Desconto IR: R$ %.2f\n", salary.CalculateNewIR())
	fmt.Printf("Salário Líquido: R$ %.2f\n", salary.CalculateSalarioLiquido(true))

}
