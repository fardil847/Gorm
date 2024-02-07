package main

import (
	"Gorm/database"
	"Gorm/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()

	// getUserById(1)
	// createUser("lalalala@gmail.com")
	// updateUserById(1, "johnjohn@gmail.com")
	// getAllUsers()
	// createProduct(1, "YLO", "YYYY")
	// getUserWithProducts()
	// deleteProductById(1)

}

func createUser(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data:", err)
		return
	}
	fmt.Println("New User Data: ", User)
}

func getUserById(id uint) {
	db := database.GetDB()

	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found ")
			return
		}
		print("Error finding user:", err)
	}
	fmt.Printf("User Data: %v \n", user)
}

func getAllUsers() {
	db := database.GetDB()

	var users []models.User
	err := db.Find(&users).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		fmt.Println("Error finding users: ", err)
		return
	}
	fmt.Printf("User Data : %v\n", users)
}

func updateUserById(id uint, email string) {
	db := database.GetDB()

	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error

	if err != nil {
		fmt.Println("Error updatiing user data", err)
		return

	}
	fmt.Println("Update user's email %v \n", user.Email)
}

func createProduct(userId uint, brand string, name string) {
	db := database.GetDB()

	Product := models.Product{
		UserID: userId,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error creating product data: ", err.Error())
		return
	}

	fmt.Println("New Product Data: ", Product)
}

func getUserWithProducts() {
	db := database.GetDB()
	users := models.User{}
	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user datas with products:", err.Error())
		return
	}

	fmt.Println("User Datas With Products")
	fmt.Printf("%+v", users)
}

func deleteProductById(id uint) {
	db := database.GetDB()
	product := models.Product{}
	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error deleting product:", err.Error())
		return
	}
	fmt.Printf("Product with id %d has been successfully deleted", id)
}
