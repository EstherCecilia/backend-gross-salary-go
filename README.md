# Calculadora de Salário Líquido 🧮

Uma aplicação simples em Go para calcular o salário líquido com base nos descontos de INSS e Imposto de Renda (IR). Suporta modelos de cálculo padrão e novos, considerando regras específicas para isenção e alíquotas progressivas.

## 🚀 Funcionalidades

- **Cálculo do INSS**: Baseado nas alíquotas oficiais progressivas.
- **Cálculo do IR**: Com deduções e faixas de isenção.
- **Novo Modelo de IR**: 
  - Isenção para salários até R$ 5.000.  
  - Alíquota adicional de 10% para salários acima de R$ 50.000.
- **Cálculo do Salário Líquido**: Inclui opção de cálculo com o novo modelo de IR.

## 📂 Estrutura de Arquivos

- `main.go`: Ponto de entrada da aplicação. Exemplos de uso.
- `utils/salary.go`: Contém a lógica de cálculo do salário líquido, INSS e IR.

## 🛠️ Como Usar

### 1. Clone o repositório

```bash
git clone https://github.com/seu-usuario/nome-repositorio.git
cd nome-repositorio
```

### 2. Compile e execute
Certifique-se de ter o Go instalado:

```bash
go run main.go
```

### 3. Personalize os valores
No arquivo main.go, altere o valor do salário na estrutura Salary:

```go
salary := Salary{Value: 10000}
```

### 📊 Exemplo de Saída
```plaintext
Salário Bruto: R$ 10.000,00
Desconto INSS: R$ 713,10
Desconto IR: R$ 869,36
Salário Líquido (modelo atual): R$ 8.417,54
Salário Líquido (novo modelo): R$ 8.786,90
```
