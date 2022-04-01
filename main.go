package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/hyson007/goschool/assignment1/category"
	"github.com/hyson007/goschool/assignment1/shopitem"
)

var Categories = category.CategorySlice
var shopItems []shopitem.ShopItem

func getCategoryIndex(name string) int {
	for idx, item := range Categories {
		if item == name {
			return idx
		}
	}
	log.Printf("unable to locate %v from category, return -1", name)
	return -1
}

func itemInCategories(item string) bool {
	for _, c := range Categories {
		if item == c {
			return true
		}
	}
	return false
}

func modifyItem() {
	var oldName, newName, cat string
	var quantity int
	var cost float64
	fmt.Println("Modify Item")
	fmt.Println("Which item would you wish to modify")
	fmt.Scanln(&oldName)

	fmt.Println("Enter new name, enter for no change")
	fmt.Scanln(&newName)
	fmt.Println("Enter new Category, enter for no change")
	fmt.Scanln(&cat)
	fmt.Println("Enter new Quantity, enter for no change")
	fmt.Scanln(&quantity)
	fmt.Println("Enter new Unit cost, enter for no change")
	fmt.Scanln(&cost)

	for idx, shopitem := range shopItems {
		if shopitem.Name == oldName {
			if cat == "" || cat == Categories[shopitem.Category] {
				fmt.Println("No changes to item category made")
				cat = Categories[shopitem.Category]
			}

			if newName == "" || oldName == newName {
				fmt.Println("No changes to item name made")
			}

			if quantity == 0 || quantity == shopitem.Quantity {
				fmt.Println("No changes to item quantity made")
				quantity = shopitem.Quantity
			}

			if cost == 0 || cost == shopitem.Cost {
				fmt.Println("No changes to item cost made")
				cost = shopitem.Cost
			}

			if cost > 0 && quantity > 0 {
				if newName != "" {
					// meaning the item name is changed
					newItem := shopitem.ModifyItem(getCategoryIndex(cat), quantity, idx, cost, newName)
					shopItems[idx] = newItem

				} else {
					// same old name
					newItem := shopitem.ModifyItem(getCategoryIndex(cat), quantity, idx, cost, oldName)
					shopItems[idx] = newItem
				}
				fmt.Println("change made!")
			}

		}

	}

}

// in the map implementation, it's not possible to add two items with same name
// we carry on this restriction, even if it's possible to do with slice of struct
func addItem() {
	var name, cat string
	var quantity int
	var cost float64
	preLen := len(shopItems)
	fmt.Println("What's the Name of your item?")
	fmt.Scanln(&name)
	fmt.Println("What category does it belong to?")
	fmt.Scanln(&cat)
	fmt.Println("How many units are there?")
	fmt.Scanln(&quantity)
	fmt.Println("How much does it cost per unit?")
	fmt.Scanln(&cost)

	if !itemInCategories(cat) {
		fmt.Printf("%s is not yet in the category list, please try again...\n", cat)
		return
	}

	for _, shopitem := range shopItems {
		if shopitem.Name == name {
			fmt.Printf("Item %s already exist, ignoring...\n", name)
			return
		}
	}
	shopItems = append(shopItems, shopitem.NewShopItem(
		getCategoryIndex(cat),
		quantity,
		cost,
		name,
	))

	fmt.Println("item added!")
	fmt.Printf("current number of total item(key): %d, it was %d \n",
		len(shopItems), preLen)
}

func generateShopReport() {
	var choice int
	var reportText = `
	Generate Report
	1. Total Cost of each category.
	2. List of item by category.
	3. Main Menu
	`
	fmt.Println(reportText)
	fmt.Scanln(&choice)
	switch {
	case choice == 1:
		var hold = make(map[int]float64)
		for _, shopItem := range shopItems {
			hold[shopItem.Category] += shopItem.Cost * float64(shopItem.Quantity)
		}
		fmt.Println("Total cost by Category")
		for k, v := range hold {
			fmt.Printf("%s cost : %.1f\n", Categories[k], v)

		}

	case choice == 2:
		//sort the slice of struct first then print
		sort.Slice(shopItems, func(i, j int) bool {
			return shopItems[i].Category < shopItems[j].Category
		})

		for _, shopitem := range shopItems {
			fmt.Print(shopitem.PrintItem())
		}

	default:
		return
	}
}

func deleteItem() {
	var name string
	fmt.Println("Enter item name to delete")
	fmt.Scanln(&name)
	for idx, shopitem := range shopItems {
		if name == shopitem.Name {
			shopItems = append(shopItems[:idx], shopItems[idx+1:]...)
		}
	}
}

func addCategory() {
	fmt.Println("Add New Category Name")
	var cat string
	fmt.Scanln(&cat)
	if cat == "" {
		fmt.Println("No Input Found!")
		return
	}
	for idx, category := range Categories {
		if cat == category {
			fmt.Printf("Category: %s already exist at index %d\n", cat, idx)
			return
		}
	}
	Categories = append(Categories, cat)
	fmt.Printf("New category: %s added at index %d\n", cat, len(Categories)+1)
}

func modifyCategory() {
	var curCat, newCat string
	fmt.Println("please provide the current and new name for category to be modified, with space in between")
	fmt.Scanf("%s %s", &curCat, &newCat)
	if newCat == "" || curCat == "" {
		fmt.Println("invalid category")
		return
	}
	for idx, cat := range Categories {
		if curCat == cat {
			Categories[idx] = newCat

			// there is no need to update the rest of items as the category
			// index remains the same
			fmt.Println("category modified!, new category as below:")
			fmt.Println(Categories)
			return
		}
	}
	fmt.Println("the input current category is not in list!")
}

func deleteCategory() []shopitem.ShopItem {
	if len(shopItems) == 0 {
		fmt.Println("Nothing left!")
		return nil
	}

	var curCat string
	var NewShopItems []shopitem.ShopItem
	fmt.Println("please provide the category to be delete")
	fmt.Scanln(&curCat)
	if curCat == "" {
		fmt.Println("invalid category")
		return nil
	}
	for idx, cat := range Categories {

		if cat == curCat {

			// delete the other items within this category
			for _, shopitem := range shopItems {
				if Categories[shopitem.Category] != curCat {

					//update all other items categories as index has been shuffled.
					if shopitem.Category > idx {
						shopitem.UpdateCategory()
					}
					NewShopItems = append(NewShopItems, shopitem)
				}
			}

			// delete category from slice
			Categories = append(Categories[:idx], Categories[idx+1:]...)
			fmt.Println("category and items in that category has been deleted!")
			return NewShopItems

		}

	}
	fmt.Println("invalid category, not found in list")
	return nil
}

func saveToJson() {
	b, err := json.Marshal(shopItems)
	if err != nil {
		log.Panic("error during save", err)
	}
	err = os.WriteFile("db.json", b, 0644)
	if err != nil {
		log.Panic("error during save", err)
	}
	fmt.Println("file saved!")
}

func loadFromJson() {
	b, err := ioutil.ReadFile("db.json")
	if err != nil {
		log.Panic("error during load", err)
	}

	err = json.Unmarshal(b, &shopItems)

	fmt.Println("data loaded from local storage!")
}

func main() {

	mainText := `
	Shopping List Application
	===============================
	1. View entire shopping list.
	2. Generate Shipping List Report.
	3. Add Items.
	4. Modify Items.
	5. Delete Item.
	6. Print Current Data.
	7. Add New Category Name.
	8. Modify Category
	9. Delete Category (All items in that category will be removed!)
	10. Save/Store to localstorage (default: db.json in same directory)
	11. Load from localstorage (default db.json in same directory)
	Select your choice:
	`
	var choice int
	//The following test data shall be preloaded during runtime

	//Category Item Quantity Unit Cost
	//Household Fork
	//Household Plates
	//Household Cups
	//Food Bread
	//Food Cake
	//Drinks Coke
	//Drinks Sprite

	//For advanced options.
	//Category Item
	//Snacks Chips 10
	//Stationary Pencil 5

	//loading test data during run time
	shopItems = append(shopItems, shopitem.NewShopItem(
		getCategoryIndex("Household"), 4, 3, "fork"))

	shopItems = append(shopItems, shopitem.NewShopItem(
		getCategoryIndex("Drinks"), 5, 2, "Coke"))

	shopItems = append(shopItems, shopitem.NewShopItem(
		getCategoryIndex("Household"), 4, 3, "Plates"))

	shopItems = append(shopItems, shopitem.NewShopItem(
		getCategoryIndex("Household"), 5, 3, "Cups"))

	shopItems = append(shopItems, shopitem.NewShopItem(
		getCategoryIndex("Food"), 3, 1, "Cake"))

	shopItems = append(shopItems, shopitem.NewShopItem(
		getCategoryIndex("Food"), 2, 2, "Bread"))

	shopItems = append(shopItems, shopitem.NewShopItem(
		getCategoryIndex("Drinks"), 5, 2, "Sprite"))

	shopItems = append(shopItems, shopitem.NewShopItem(
		getCategoryIndex("Snacks"), 10, 3, "Chips"))

	shopItems = append(shopItems, shopitem.NewShopItem(
		getCategoryIndex("Stationary"), 5, 1, "Pencil"))

	for {
		fmt.Println(mainText)
		fmt.Scanln(&choice)
		switch {
		case choice == 1:
			if len(shopItems) == 0 {
				fmt.Println("Nothing in the shopping list")
			}

			for _, shopitem := range shopItems {
				fmt.Print(shopitem.PrintItem())
			}

		case choice == 2:
			generateShopReport()
		case choice == 3:
			addItem()
		case choice == 4:
			modifyItem()
		case choice == 5:
			deleteItem()
		case choice == 6:
			for _, shopitem := range shopItems {
				fmt.Print(shopitem)
			}

		case choice == 7:
			addCategory()
		case choice == 8:
			modifyCategory()
		case choice == 9:
			shopItems = deleteCategory()

		case choice == 10:
			saveToJson()
		case choice == 11:
			loadFromJson()
		default:
			continue
		}
	}

}
