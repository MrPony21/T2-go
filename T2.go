package main

import "fmt"

type Nodo struct {
	nombre    string
	apellido  string
	edad      int
	direccion string
	carrera   string
	curso     string
	estado    int
	siguiente *Nodo
}

type ListaEnlazada struct {
	cabeza *Nodo
}

func (l *ListaEnlazada) Agregar(nombre string, apellido string, edad int, direccion string, carrera string, curso string, estado int) {
	nuevoNodo := &Nodo{nombre: nombre, apellido: apellido, edad: edad, direccion: direccion, carrera: carrera, curso: curso, estado: estado}
	if l.cabeza == nil {
		l.cabeza = nuevoNodo
		return
	}
	temp := l.cabeza
	for temp.siguiente != nil {
		temp = temp.siguiente
	}
	temp.siguiente = nuevoNodo
}

func (l *ListaEnlazada) Editar(nombre string, edad int, direccion string, estado int) {
	temp := l.cabeza
	for temp != nil {
		if temp.nombre == nombre {
			temp.edad = edad
			temp.direccion = direccion
			temp.estado = estado
		}
		temp = temp.siguiente
	}
}

func (l *ListaEnlazada) Recorrer() {
	temp := l.cabeza
	for temp != nil {
		fmt.Println(temp.nombre, temp.apellido, temp.edad, temp.direccion, temp.carrera, temp.curso, temp.estado)
		temp = temp.siguiente
	}
}

func main() {
	lista := &ListaEnlazada{}
	lista.Agregar("Marco", "Solis", 20, "San carlos", "sistemas", "edd", 1)
	lista.Agregar("Antonio", "Gonzalez", 17, "Uni", "Ingenieria", "compi 1", 0)
	lista.Recorrer()

	lista.Editar("Antonio", 58, "Nueva direccion", 1)
	lista.Recorrer()
}
