package main

import (
	"fmt"

//	"math/rand"
)

type entity interface {
	size() float32
}
type Magic struct {
	school string
	level  int8
}

type Person struct {
	firstName string
	lastName  string
}
type Android struct {
	Person
	Armor
}
type Armor struct {
	hardness  int
	thickness int
}

type Game struct {
	Name string
}

//methods

func (m *Person) size() {
	fmt.Print("pretty big")
}
func (m *Magic) cast_spell() {
	fmt.Print("bam!\n")
}

func (m *Armor) dmg_durabiliy() {
	fmt.Print("durability reduced!")
}
func (m *Person) talk() string {
	return "hi"
}

func (p *Person) SetFirstName(newName string) {
	p.firstName = newName
}
func (p *Person) FirstName() string {
	return p.firstName
}
func (m *Game) new_bool() bool {
	return false
}
func (m *Game) new_person() string {
	m.Name = "killer zombies"
	var input string
	fmt.Scanln(&input)
	return "awesome"
}
