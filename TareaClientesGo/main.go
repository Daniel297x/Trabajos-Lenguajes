package main

import (
	"fmt"
	"sort"
	"strings"
)

type infoCliente struct {
	nombre string
	correo string
	edad   int32
}

type clientes []infoCliente

var listaClientes clientes

func (c *clientes) agregarCliente(nombre string, correo string, edad int32) {

	*c = append(*c, infoCliente{nombre, correo, edad})
}

// Función genérica Map
func MapG[T, U any](list []T, f func(T) U) []U {
	mapped := make([]U, len(list))
	for i, v := range list {
		mapped[i] = f(v)
	}
	return mapped
}

// Función genérica Filter
func Filter[T any](list []T, f func(T) bool) []T {
	var filtered []T
	for _, v := range list {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func reduce(list []int) int {
	var result int = 0
	for _, v := range list {
		result += v
	}
	return result
}

func obtenerApellido(cliente *infoCliente) string {
	apellido := strings.ToLower((strings.Fields(cliente.nombre))[1])
	return apellido
}
func obtenerApellidoCorrero(cliente *infoCliente) string {
	apellido := strings.Split(cliente.correo, "@")
	return (apellido[0])[1:]
}
func verificarApellidoCorreo(cliente infoCliente) bool {
	apellidoCorreo := false
	if obtenerApellido(&cliente) == obtenerApellidoCorrero(&cliente) {
		apellidoCorreo = true
	}
	return apellidoCorreo

}

func listaClientesApellidoEnCorreo(c *clientes) clientes {
	return Filter(*c, verificarApellidoCorreo)

}
func verificarDominio(cliente infoCliente) bool {
	dominioCr := false
	correo := strings.Split(cliente.correo, ".")
	if correo[len(correo)-1] == "cr" {
		dominioCr = true
	}
	return dominioCr
}

func cantidadCorreosCR(c *clientes) int {
	correosDom := Filter(*c, verificarDominio)
	correosDom2 := MapG(correosDom, func(ca infoCliente) int {
		return 1
	})
	total := reduce(correosDom2)
	return total
}

func verificarNoApellidoCorreo(cliente infoCliente) bool {
	apellidoCorreo := false
	if obtenerApellido(&cliente) != obtenerApellidoCorrero(&cliente) {
		apellidoCorreo = true
	}
	return apellidoCorreo

}
func cambiarCorreo(cliente infoCliente) infoCliente {
	restoCorreo := (strings.Split(cliente.correo, "@"))[1]
	cliente.correo = (strings.ToLower(string(cliente.nombre[0]))) + obtenerApellido(&cliente) + "" + "@" + restoCorreo
	return cliente

}

func clientesSugerenciasCorreos(c *clientes) clientes {
	listaNoApellidos := Filter(*c, verificarNoApellidoCorreo)
	ListaCambiados := MapG(listaNoApellidos, cambiarCorreo)

	return Filter(append(*c, ListaCambiados...), verificarApellidoCorreo)
}
func correos(cliente infoCliente) string {
	correoC := cliente.correo
	return correoC
}
func ordenarClientes() {

}
func correosOrdenadosAlfabeticos(c *clientes) clientes {
	var nuevaListaOrdena clientes
	listaCorreos := MapG(*c, correos)
	sort.Strings(listaCorreos)
	for _, v := range listaCorreos {
		for _, cl := range *c {
			if cl.correo == v {
				nuevaListaOrdena = append(nuevaListaOrdena, cl)
			}
		}
	}
	return nuevaListaOrdena
}

func agregarDatos() {
	listaClientes.agregarCliente("Oscar Viquez", "oviquez@tec.ac.cr", 44)
	listaClientes.agregarCliente("Pedro Perez", "elsegundo@gmail.com", 30)
	listaClientes.agregarCliente("Maria Lopez", "mlopez@hotmail.com", 18)
	listaClientes.agregarCliente("Juan Rodriguez", "jrodriguez@gmail.com", 25)
	listaClientes.agregarCliente("Luisa Gonzalez", "muyseguro@tec.ac.cr", 67)
	listaClientes.agregarCliente("Marco Rojas", "loquesea@hotmail.com", 47)
	listaClientes.agregarCliente("Marta Saborio", "msaborio@gmail.com", 33)
	listaClientes.agregarCliente("Camila Segura", "csegura@ice.co.cr", 19)
	listaClientes.agregarCliente("Fernando Rojas", "frojas@estado.gov", 27)
	listaClientes.agregarCliente("Rosa Ramirez", "risuenna@gmail.com", 50)

}
func main() {
	agregarDatos()
	fmt.Println("lista original--")
	fmt.Println(listaClientes)
	fmt.Println("\nlista de clientes con el apellido en el correo----")
	fmt.Println(listaClientesApellidoEnCorreo(&listaClientes))

	fmt.Println("\ncantidad de correos con el dominio .cr")
	fmt.Println(cantidadCorreosCR(&listaClientes))
	fmt.Println("\nlista con el correo cambiado a correo con apellido + clientes originales")
	fmt.Println(clientesSugerenciasCorreos(&listaClientes))
	fmt.Println("\nlista de clientes ordenada por orden alfabetico segun el correo")
	fmt.Println(correosOrdenadosAlfabeticos(&listaClientes))

}
