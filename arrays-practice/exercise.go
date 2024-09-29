package main

import "fmt"

func main() {

	// Time to practice what you learned!

	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.

	my_hobbies := [3]string{"Programming", "Cloud Computing", "Bugs"}
	fmt.Println("My Hobbies: ", my_hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list

	fmt.Println("First Hobby: ", my_hobbies[0])
	favoriteHobbies := my_hobbies[1:]
	fmt.Println("Favorite Hobbies: ", favoriteHobbies)

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)

	tech_hobbies := my_hobbies[:2]
	fmt.Println("Tech Hobbies: ", tech_hobbies)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.

	tech_hobbies = tech_hobbies[1:3]
	fmt.Println("Main Hobbies: ", tech_hobbies)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)

	courseGoals := []string{"Learn Go", "Create CLI's"}
	fmt.Println("Course Goals: ", courseGoals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array

	courseGoals[1] = "Create DevOps Tools"
	newCourseGoals := append(courseGoals, "Create CLI's")
	fmt.Println("New Course Goals: ", newCourseGoals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.

	type Product struct {
		id    int64
		title string
		price float64
	}

	productList := []Product{
		{
			id:    100,
			title: "Cleaning Liquid",
			price: 200.00,
		},
		{
			id:    101,
			title: "Detergent",
			price: 249.99,
		},
	}
	fmt.Println("Product List: ", productList)

	updatedProduct := append(productList, Product{
		id:    102,
		title: "Cleaning Kit",
		price: 499.99,
	})
	fmt.Println("Updated Products: ", updatedProduct)
}
