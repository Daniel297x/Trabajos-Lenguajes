package main

import "fmt"

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	var existe bool = false
	for i := 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			(*l)[i].cantidad = (*l)[i].cantidad + cantidad
			existe = true
			if (*l)[i].precio != precio {
				(*l)[i].precio = precio
			}
		}
	}
	if !existe {
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}
	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente
	//listo**
}

// Hacer una función adicional que reciba una cantidad potencialmente infinita de productos previamente creados y los agregue a la lista

func (l *listaProductos) buscarProducto(nombre string) (*producto, int) { //el retorno es el índice del producto encontrado y -1 si no existe
	var err = -1
	var p *producto = nil
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			err = 0
			p = &((*l)[i])
		}
	}
	return p, err
}

//modifique la función para que esta vez solo retorne el producto como tal y que retorne otro valor adicional "err" conteniendo un 0 si no hay error y números mayores si existe algún
//tipo de error como por ejemplo que el producto no exista. Una vez implementada la nueva función, cambie los usos de la anterior función en funciones posteriores, realizando los cambios
//que el uso de la nueva función ameriten

func (l *listaProductos) venderProducto(nombre string, cant int) {
	var prod, err = l.buscarProducto(nombre)

	if err != -1 {
		if (*prod).cantidad == 0 {
			for i := 0; i < len(*l); i++ {
				if (*l)[i].nombre == (*prod).nombre {
					fmt.Println("no hay existencias disponibles, se ha eliminado el producto: " + nombre)
					*l = append((*l)[:i], (*l)[i+1:]...)

				}

			}
		} else if (*prod).cantidad-cant < 0 {
			var vendido = &((*prod).cantidad)
			for i := 0; i < len(*l); i++ {
				if (*l)[i].nombre == (*prod).nombre {
					fmt.Println("se han vendido solo", *vendido, " y se ha eliminado el producto: "+nombre)
					*l = append((*l)[:i], (*l)[i+1:]...)

				}

			}

		} else {
			(*prod).cantidad -= cant
			if (*prod).cantidad == 0 {
				for i := 0; i < len(*l); i++ {
					if (*l)[i].nombre == (*prod).nombre {

						*l = append((*l)[:i], (*l)[i+1:]...)
						fmt.Println("se han vendido todas las existencias, se ha eliminado el producto: " + nombre)
					}
				}
			}

		}

		//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista" y notificar dicha acción

	}

}

// haga una función para a partir del nombre del producto, se pueda modificar el precio del mismo Pero utilizando la función buscarProducto MODIFICADA PREVIAMENTE
func (l listaProductos) modificarPrecio(nombre string, precio int) {
	var prod, err = l.buscarProducto(nombre)
	if err != -1 {
		(*prod).precio = precio
	}

}

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)
	lProductos.agregarProducto("café", 13, 4300)
	lProductos.modificarPrecio("café", 2522)
}

func (l *listaProductos) listarProductosMinimos() listaProductos {
	// debe retornar una nueva lista con productos con existencia mínima
	newSlice := make(listaProductos, 0)

	for _, p := range *l {
		if p.cantidad <= existenciaMinima {
			newSlice = append(newSlice, p)
		}
	}
	return newSlice
}

func main() {
	llenarDatos()
	fmt.Println(lProductos)
	//venda productos

	lProductos.venderProducto("arroz", 10)
	fmt.Println(lProductos)
	lProductos.venderProducto("arroz", 4)
	fmt.Println(lProductos)
	lProductos.venderProducto("arroz", 1)
	fmt.Println(lProductos)
	lProductos.venderProducto("frijoles", 2)
	fmt.Println(lProductos)

	pminimos := lProductos.listarProductosMinimos()
	fmt.Println(pminimos)

}
