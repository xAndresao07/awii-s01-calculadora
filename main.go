package main

import (
	"fmt"
	"semana03-taller-relaciones/internal/cafeteria"
)

func main() {
	var repo cafeteria.Repository = cafeteria.NewRepoMemoria()
	repo.GuardarCliente(cafeteria.Cliente{ID: 1, Nombre: "Juan", Carrera: "TI", Saldo: 100.0})
	repo.GuardarCliente(cafeteria.Cliente{ID: 2, Nombre: "Anthony", Carrera: "Turismo", Saldo: 200.0})
	repo.GuardarProducto(cafeteria.Producto{ID: 1, Nombre: "Dona", Precio: 5.0, Stock: 100, Categoria: cafeteria.Categoria{ID: 1, Nombre: "Alimentos"}})
	repo.GuardarProducto(cafeteria.Producto{ID: 2, Nombre: "FuzeTea", Precio: 20.0, Stock: 50, Categoria: cafeteria.Categoria{ID: 2, Nombre: "Bebidas"}})
	repo.GuardarProducto(cafeteria.Producto{ID: 3, Nombre: "Torta", Precio: 15.0, Stock: 75, Categoria: cafeteria.Categoria{ID: 1, Nombre: "Alimentos"}})
	repo.GuardarPedido(cafeteria.Pedido{ID: 1, Cliente: cafeteria.Cliente{ID: 1, Nombre: "Juan", Carrera: "TI", Saldo: 100.0}, Producto: cafeteria.Producto{ID: 1, Nombre: "Dona", Precio: 5.0, Stock: 100, Categoria: cafeteria.Categoria{ID: 1, Nombre: "Alimentos"}}, Cantidad: 2, Total: 10.0})

	//Cliente existente
	c, err := repo.ObtenerCliente(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Cliente encontrado:", c.Nombre)
	}

	//Cliente no existente
	c2, err := repo.ObtenerCliente(99)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Cliente encontrado:", c2.Nombre)
	}

	//Listar productos
	productos := repo.ListarProductos()
	fmt.Println("ID: Nombre: Precio: Stock: Categoria:")
	for _, p := range productos {
		fmt.Printf("%d: %s: %.2f: %d: %s\n", p.ID, p.Nombre, p.Precio, p.Stock, p.Categoria.Nombre)
	}

	//Listar pedidos
	pedidos := repo.ListarPedidos()
	for _, p := range pedidos {
		fmt.Printf("Pedido ID: %d, Cliente: %s, Producto: %s, Cantidad: %d, Total: %.2f\n", p.ID, p.Cliente.Nombre, p.Producto.Nombre, p.Cantidad, p.Total)
	}

	//Listar productos
	productoss := repo.ListarProductos()
	fmt.Println("ID: Nombre: Precio: Stock: Categoria:")
	for _, p := range productoss {
		fmt.Printf("%d: %s: %.2f: %d: %s\n", p.ID, p.Nombre, p.Precio, p.Stock, p.Categoria.Nombre)
	}

}
