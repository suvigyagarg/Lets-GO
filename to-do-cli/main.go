package main


func main() {
	todos := Todos{}
	todos.add("Buy Milk")
	todos.add("Shop Grocery")
	todos.toggle(0)
	todos.print()
	
}