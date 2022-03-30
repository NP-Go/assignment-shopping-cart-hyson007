package main

  
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var CategorySlice = []string{"Household", "Food", "Drinks", "Snacks", "Stationary"}

type shopItemSt struct {
	Category int
	Quantity int
	Cost     float64
}

type shopItemHelper struct {
	Category int     `json:"category"`
	Quantity int     `json:"quantity"`
	Cost     float64 `json:"cost"`
	Name     string  `json:"name"`
}

var shopItemHelperSlice = []shopItemHelper{}

type shopItemMapT map[string]shopItemSt

var shopItemMap = make(shopItemMapT)

func (s shopItemMapT) String() string {
	output := "Print Current Data.\n"
	for item, value := range shopItemMap {
		output += fmt.Sprintf("%s - %v\n", item, value)
	}
	return output
}

func (s shopItemMapT) printItem() {
	for k, v := range s {
		fmt.Printf("Catgory: %s - Item: %s Quantity: %d Unit Cost: %.1f\n",
			CategorySlice[v.Category], k, v.Quantity, v.Cost)
	}
}

func (s shopItemMapT) generateShopReportTotal() {
	var hold = make(map[string]float64)
	for _, v := range shopItemMap {
		hold[CategorySlice[v.Category]] += v.Cost * float64(v.Quantity)
	}
	fmt.Println("Total cost by Category")
	for k, v := range hold {
		fmt.Printf("%s cost : %.1f\n", k, v)

	}
}

func (s shopItemMapT) generateShopReportByCat() {
	for catIdx, _ := range CategorySlice {
		for k, v := range shopItemMap {
			if v.Category == catIdx {
				fmt.Printf("Catgory: %s - Item: %s Quantity: %d Unit Cost: %.1f\n",
					CategorySlice[v.Category], k, v.Quantity, v.Cost)
			}
		}
	}
}

func (s shopItemMapT) deleteItem() {
	var name string
	fmt.Println("Enter item name to delete")
	fmt.Scanln(&name)

	if _, ok := shopItemMap[name]; ok {
		delete(shopItemMap, name)
		fmt.Printf("item %s delete\n", name)
	} else {
		fmt.Println("Item not found, nothing to delete")
		return
	}
}

func (s shopItemMapT) modifyItem() {
	var oldName, newName, cat string
	var quantity int
	var cost float64
	fmt.Println("Modify Item")
	fmt.Println("Which item would you wish to modify")
	fmt.Scanln(&oldName)
	if value, ok := shopItemMap[oldName]; ok {
		fmt.Printf("Current item name is %s - Category is %s - Quantity is %d - Unit cost %.1f\n",
			oldName, CategorySlice[value.Category], value.Quantity, value.Cost)

		fmt.Println("Enter new name, enter for no change")
		fmt.Scanln(&newName)
		fmt.Println("Enter new Category, enter for no change")
		fmt.Scanln(&cat)
		fmt.Println("Enter new Quantity, enter for no change")
		fmt.Scanln(&quantity)
		fmt.Println("Enter new Unit cost, enter for no change")
		fmt.Scanln(&cost)

		if cat == "" || cat == CategorySlice[value.Category] {
			fmt.Println("No changes to item category made")
		}

		if newName == "" || oldName == newName {
			fmt.Println("No changes to item name made")
		}

		if quantity == 0 || quantity == value.Quantity {
			fmt.Println("No changes to item quantity made")
		}

		if cost == 0 || cost == value.Cost {
			fmt.Println("No changes to item cost made")
		}

		if cost > 0 && quantity > 0 {
			if newName != "" {
				shopItemMap[newName] = shopItemSt{
					Category: value.Category,
					Quantity: quantity,
					Cost:     cost,
				}
			} else {
				shopItemMap[oldName] = shopItemSt{
					Category: value.Category,
					Quantity: quantity,
					Cost:     cost,
				}
			}
			fmt.Println("change made!")
			value := shopItemMap[newName]
			fmt.Printf("new item name is %s - Category is %s - Quantity is %d - Unit cost %.1f\n",
				newName, CategorySlice[value.Category], value.Quantity, value.Cost)
		}

	} else {
		fmt.Println("no such item!")
		return
	}
}

func (s shopItemMapT) addItem() {
	var name, cat string
	var quantity int
	var cost float64
	preLen := len(shopItemMap)
	fmt.Println("What's the Name of your item?")
	fmt.Scanln(&name)
	fmt.Println("What category does it belong to?")
	fmt.Scanln(&cat)
	fmt.Println("How many units are there?")
	fmt.Scanln(&quantity)
	fmt.Println("How much does it cost per unit?")
	fmt.Scanln(&cost)
	if _, found := shopItemMap[name]; found {
		fmt.Printf("Item %s already exist in map, ignoring...\n", name)
		return
	}

	shopItemMap[name] = shopItemSt{
		getCategoryIndex(cat),
		quantity,
		cost,
	}
	fmt.Println("item added!")
	fmt.Printf("current number of total item(key) in map: %d, it was %d \n",
		len(shopItemMap), preLen)
}

func getCategoryIndex(name string) int {
	for idx, item := range CategorySlice {
		if item == name {
			return idx
		}
	}
	log.Printf("unable to locate %v from category, return -1", name)
	return -1
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
		shopItemMap.generateShopReportTotal()
	case choice == 2:
		shopItemMap.generateShopReportByCat()
	default:
		return
	}
}

// func deleteItem() {
// 	var name string
// 	fmt.Println("Enter item name to delete")
// 	fmt.Scanln(&name)
// 	if _, ok := shopItemMap[name]; ok {
// 		delete(shopItemMap, name)
// 		fmt.Printf("item %s delete\n", name)
// 	} else {
// 		fmt.Println("Item not found, nothing to delete")
// 		return
// 	}
// }

func addCategory() {
	fmt.Println("Add New Category Name")
	var cat string
	fmt.Scanln(&cat)
	if cat == "" {
		fmt.Println("No Input Found!")
		return
	}
	for idx, category := range CategorySlice {
		if cat == category {
			fmt.Printf("Category: %s already exist at index %d\n", cat, idx)
			return
		}
	}
	CategorySlice = append(CategorySlice, cat)
	fmt.Printf("New category: %s added at index %d\n", cat, len(CategorySlice)+1)
}

func modifyCategory() {
	var curCat, newCat string
	fmt.Println("please provide the current and new name for category to be modified")
	fmt.Scanf("%s %s", &curCat, &newCat)
	if newCat == "" || curCat == "" {
		fmt.Println("invalid category")
		return
	}
	for idx, cat := range CategorySlice {
		if curCat == cat {
			CategorySlice[idx] = newCat

			// there is no need to update the rest of items as the category
			// index remains the same
			fmt.Println("category modified!, new category as below:")
			fmt.Println(CategorySlice)
			return
		}
	}
	fmt.Println("the input current category is not in list!")
}

func deleteCategory() {
	var curCat string
	fmt.Println("please provide the category to be delete")
	fmt.Scanln(&curCat)
	if curCat == "" {
		fmt.Println("invalid category")
		return
	}
	for idx, cat := range CategorySlice {
		if cat == curCat {
			// delete category from slice
			CategorySlice = append(CategorySlice[:idx], CategorySlice[idx+1:]...)
			// delete the other items within this category
			for k, v := range shopItemMap {
				if v.Category == idx {
					delete(shopItemMap, k)
				}
			}
			fmt.Println("category and items in that category has been deleted!")

		}
	}
	fmt.Println("the input current category is not in list!")
}

func saveToJson() {
	//populate the helper struct before saving
	for k, v := range shopItemMap {

		shopItemHelperSlice = append(shopItemHelperSlice,
			shopItemHelper{
				v.Category,
				v.Quantity,
				v.Cost,
				k,
			})
	}
	// fmt.Println(shopItemHelperSlice)
	b, err := json.Marshal(shopItemHelperSlice)
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

	err = json.Unmarshal(b, &shopItemHelperSlice)

	//repopulate the map from slice
	shopItemMap = make(map[string]shopItemSt)

	for _, item := range shopItemHelperSlice {
		shopItemMap[item.Name] = shopItemSt{
			Category: item.Category,
			Quantity: item.Quantity,
			Cost:     item.Cost,
		}
	}

	if err != nil {
		log.Panic("error during load", err)
	}

	fmt.Println("data loaded from local storage!")
	fmt.Println(shopItemMap)
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

	shopItemMap["Fork"] = shopItemSt{
		Category: getCategoryIndex("Household"),
		Quantity: 4,
		Cost:     3,
	}

	shopItemMap["Plates"] = shopItemSt{
		getCategoryIndex("Household"),
		4,
		3,
	}

	shopItemMap["Cups"] = shopItemSt{
		getCategoryIndex("Household"),
		5,
		3,
	}

	shopItemMap["Bread"] = shopItemSt{
		getCategoryIndex("Food"),
		2,
		2,
	}

	shopItemMap["Cake"] = shopItemSt{
		getCategoryIndex("Food"),
		3,
		1,
	}

	shopItemMap["Coke"] = shopItemSt{
		getCategoryIndex("Drinks"),
		5,
		2,
	}

	shopItemMap["Sprite"] = shopItemSt{
		getCategoryIndex("Drinks"),
		5,
		2,
	}

	shopItemMap["Chips"] = shopItemSt{
		getCategoryIndex("Snacks"),
		10,
		3,
	}

	shopItemMap["Pencil"] = shopItemSt{
		getCategoryIndex("Stationary"),
		5,
		1,
	}

	for {
		fmt.Println(mainText)
		fmt.Scanln(&choice)
		switch {
		case choice == 1:
			shopItemMap.printItem()
		case choice == 2:
			generateShopReport()
		case choice == 3:
			shopItemMap.addItem()
		case choice == 4:
			shopItemMap.modifyItem()
		case choice == 5:
			shopItemMap.deleteItem()
		case choice == 6:
			fmt.Println(shopItemMap)
		case choice == 7:
			addCategory()
		case choice == 8:
			modifyCategory()
		case choice == 9:
			deleteCategory()
		case choice == 10:
			saveToJson()
		case choice == 11:
			loadFromJson()
		default:
			continue
		}
	}

