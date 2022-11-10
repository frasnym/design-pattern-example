package main

func main() {
	store1 := Store{
		Name:     "Store X",
		City:     "Jakarta",
		Country:  "Indonesia",
		Category: "Gadget",
	}
	store1.print()

	store2 := store1.clone()
	store2.setName("Store Y")
	store2.print()

	store3 := store1.clone()
	store3.setName("Store Z")
	store3.setCity("Bandung")
	store3.print()
}
