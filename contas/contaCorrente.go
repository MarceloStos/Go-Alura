package contas

import "Banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(quantidade float64) (string, float64) {
	if quantidade > 0 && quantidade <= c.Saldo {
		c.Saldo -= quantidade
		return "Saque realizado com sucesso! Saldo atualizado:", c.Saldo
	} else {
		return "Saldo insuficiente! Saldo:", c.Saldo
	}
}

func (c *ContaCorrente) Depositar(quantidade float64) (string, float64) {

	if quantidade > 0 {
		c.Saldo += quantidade
		return "Deposito realizado com sucesso! Saldo atualizado:", c.Saldo
	} else {
		return "Confira o valor do deposito!", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(quantidade float64, contaDestino *ContaCorrente) bool {
	if quantidade > 0 && quantidade < c.Saldo {
		c.Saldo -= quantidade
		contaDestino.Depositar(quantidade)
		return true
	} else {
		return false
	}
}
