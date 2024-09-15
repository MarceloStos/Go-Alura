package contas

import "Banco/clientes"

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (c *ContaPoupanca) Sacar(quantidade float64) (string, float64) {
	if quantidade > 0 && quantidade <= c.saldo {
		c.saldo -= quantidade
		return "Saque realizado com sucesso! Saldo atualizado:", c.saldo
	} else {
		return "Saldo insuficiente! Saldo:", c.saldo
	}
}

func (c *ContaPoupanca) Depositar(quantidade float64) (string, float64) {

	if quantidade > 0 {
		c.saldo += quantidade
		return "Deposito realizado com sucesso! Saldo atualizado:", c.saldo
	} else {
		return "Confira o valor do deposito!", c.saldo
	}
}

func (c *ContaPoupanca) Transferir(quantidade float64, contaDestino *ContaPoupanca) bool {
	if quantidade > 0 && quantidade < c.saldo {
		c.saldo -= quantidade
		contaDestino.Depositar(quantidade)
		return true
	} else {
		return false
	}
}
