package dogfood

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"os"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Swagger(w http.ResponseWriter, r *http.Request) {
	file, e := ioutil.ReadFile("/config/swagger.json")
  if e != nil {
      fmt.Printf("File error: %v\n", e)
      os.Exit(1)
  }
  fmt.Fprintf(w, string(file))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Woof Woof")
}

var routes = Routes{
	Route{
		"Swagger",
		"GET",
		"/v2/swagger",
		Swagger,
	},

	Route{
		"Index",
		"GET",
		"/v2/",
		Index,
	},

	Route{
		"Adddog",
		"POST",
		"/v2/dog",
		Adddog,
	},

	Route{
		"Deletedog",
		"DELETE",
		"/v2/dog/{dogId}",
		Deletedog,
	},

	Route{
		"FinddogsByStatus",
		"GET",
		"/v2/dog/findByStatus",
		FinddogsByStatus,
	},

	Route{
		"FinddogsByTags",
		"GET",
		"/v2/dog/findByTags",
		FinddogsByTags,
	},

	Route{
		"GetdogById",
		"GET",
		"/v2/dog/{dogId}",
		GetdogById,
	},

	Route{
		"Updatedog",
		"PUT",
		"/v2/dog",
		Updatedog,
	},

	Route{
		"UpdatedogWithForm",
		"POST",
		"/v2/dog/{dogId}",
		UpdatedogWithForm,
	},

	Route{
		"UploadFile",
		"POST",
		"/v2/dog/{dogId}/uploadImage",
		UploadFile,
	},

	Route{
		"DeleteOrder",
		"DELETE",
		"/v2/store/order/{orderId}",
		DeleteOrder,
	},

	Route{
		"GetInventory",
		"GET",
		"/v2/store/inventory",
		GetInventory,
	},

	Route{
		"GetOrderById",
		"GET",
		"/v2/store/order/{orderId}",
		GetOrderById,
	},

	Route{
		"PlaceOrder",
		"POST",
		"/v2/store/order",
		PlaceOrder,
	},

	Route{
		"CreateUser",
		"POST",
		"/v2/user",
		CreateUser,
	},

	Route{
		"CreateUsersWithArrayInput",
		"POST",
		"/v2/user/createWithArray",
		CreateUsersWithArrayInput,
	},

	Route{
		"CreateUsersWithListInput",
		"POST",
		"/v2/user/createWithList",
		CreateUsersWithListInput,
	},

	Route{
		"DeleteUser",
		"DELETE",
		"/v2/user/{username}",
		DeleteUser,
	},

	Route{
		"GetUserByName",
		"GET",
		"/v2/user/{username}",
		GetUserByName,
	},

	Route{
		"LoginUser",
		"GET",
		"/v2/user/login",
		LoginUser,
	},

	Route{
		"LogoutUser",
		"GET",
		"/v2/user/logout",
		LogoutUser,
	},

	Route{
		"UpdateUser",
		"PUT",
		"/v2/user/{username}",
		UpdateUser,
	},

}
