// package main

// import (

//     "cars/config"
//     "cars/routes"
// )

// func main() {
//     config.InitDB()
//     r := routes.SetupRouter()

//	    // Start the server
//	    r.Run(":8080")
//	}
package main

import (
	"cars/config"
	"cars/routes"
	"fmt"
)

func main() {
	// Initialize database connection
	config.InitDB()

	// Set up router
	r := routes.SetupRouter()

	// Start the server
	fmt.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
