package main

import (
	"Banco/clientes"
	"Banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) (string, float64)
}

func main() {
	clientePrimario := clientes.Titular{"Marcelo", "123.456.789-10", "Desenvolvedor"}
	contaPrimaria := contas.ContaCorrente{clientePrimario, 567, 1234, 1000.0}

	clienteSecundario := clientes.Titular{"Maria", "987.654.321-01", "Desenvolvedora"}
	contaSecundaria := contas.ContaCorrente{clienteSecundario, 321, 5678, 2000.0}

	clienteTerciario := clientes.Titular{"Gicelia", "987.654.321-00", "Engenheira"}
	contaTerciaria := contas.ContaCorrente{clienteTerciario, 001, 4567, 3000.0}

	PagarBoleto(&contaTerciaria, 30)

	// contaPrimaria := contas.ContaCorrente{Titular: clientes.Titular{
	// 	Nome:      "Marcelo",
	// 	CPF:       "123.456.789-10",
	// 	Profissao: "Dev das Galaxias"},
	// 	NumeroAgencia: 567,
	// 	NumeroConta:   1234,
	// 	Saldo:         1000.0}

	// contaSecundaria := contas.ContaCorrente{Titular: clientes.Titular{
	// 	Nome:      "Maria",
	// 	CPF:       "987.654.321.00",
	// 	Profissao: "Nao Sei"},
	// 	NumeroAgencia: 321,
	// 	NumeroConta:   5678,
	// 	Saldo:         2000.0}
	fmt.Println(contaPrimaria)
	fmt.Println(contaSecundaria)
	fmt.Println(contaTerciaria)

	// status := contaPrimaria.Transferir(200, &contaSecundaria)

	// fmt.Println(status)
}
