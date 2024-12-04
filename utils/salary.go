package utils

// Definindo a estrutura Salario
type Salary struct {
	Value float64
}

// Método para calcular o desconto do INSS
func (salary Salary) CalculateINSS() float64 {
	switch {
	case salary.Value <= 1212.00:
		return salary.Value * 0.075 // 7.5%
	case salary.Value <= 2427.35:
		return salary.Value * 0.09 // 9%
	case salary.Value <= 3641.03:
		return salary.Value * 0.12 // 12%
	default:
		return salary.Value * 0.14 // 14%
	}
}

// Método para calcular o desconto do imposto de renda
func (salary Salary) CalculateIR() float64 {
	base := salary.Value - salary.CalculateINSS() // Subtraindo o INSS da base
	switch {
	case base <= 1903.98:
		return 0 // Isento
	case base <= 2826.65:
		return base*0.075 - 142.80 // 7.5% menos dedução
	case base <= 3751.05:
		return base*0.15 - 354.80 // 15% menos dedução
	case base <= 4664.68:
		return base*0.225 - 636.13 // 22.5% menos dedução
	default:
		return base*0.275 - 869.36 // 27.5% menos dedução
	}
}

// Método para calcular o desconto do novo imposto de renda
func (salary Salary) CalculateNewIR() float64 {
	base := salary.Value - salary.CalculateINSS() - 5000
	switch {
	case base <= 0:
		return 0 // Isento
	case base <= 2826.65:
		return base * 0.075
	case base <= 3751.05:
		return base * 0.15
	case base <= 4664.68:
		return base * 0.225
	case base > 50000.00:
		return base*0.275 + base*0.1
	default:
		return base*0.275 - 869.36
	}
}

// Método para calcular o salário líquido
func (salary Salary) CalculateSalarioLiquido(newModel bool) float64 {
	if newModel {
		return salary.Value - salary.CalculateINSS() - salary.CalculateNewIR()
	}
	return salary.Value - salary.CalculateINSS() - salary.CalculateIR()
}
