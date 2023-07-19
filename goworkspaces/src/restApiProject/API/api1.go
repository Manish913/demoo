package main

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
)

// This example shows how to use methods as RouteFunctions for WebServices.
// The ProductResource has a Register() method that creates and initializes
// a WebService to expose its methods as REST operations.
// The WebService is added to the restful.DefaultContainer.
// A ProductResource is typically created using some data access object.
//
// GET http://localhost:8080/products/1
// POST http://localhost:8080/products
// <Product><Id>1</Id><Title>The First</Title></Product>

type Product struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type ProductResource struct {
	// typically reference a DAO (data-access-object)
	usr map[string]Product
}

func (p ProductResource) getOne(req *restful.Request, resp *restful.Response) {
	//id := req.PathParameter("id")
	//log.Println("getting product with id:" + id)
	//resp.WriteEntity(Product{Id: id, Title: "test"})
	list := []Product{}
	for _, each := range p.usr {
		list = append(list, each)
	}
	resp.WriteAsJson("Getting")
	r := resp.WriteEntity(list)
	log.Println(r)
}

func (p ProductResource) postOne(request *restful.Request, response *restful.Response) {
	us := Product{Id: request.PathParameter("user-id")}
	err := request.ReadEntity(&us)
	if err == nil {
		p.usr[us.Id] = us
		response.WriteHeaderAndEntity(http.StatusCreated, us)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func (p ProductResource) Register() {
	ws := new(restful.WebService)
	ws.Path("/products")
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/{id}").To(p.getOne).
		Doc("get the product by its id").
		Param(ws.PathParameter("id", "identifier of the product").DataType("string")))

	ws.Route(ws.POST("/").To(p.postOne).
		Doc("update or create a product").
		Param(ws.BodyParameter("Product", "a Product (Json)").DataType("main.Product")))

	restful.Add(ws)
}

func main() {
	ProductResource{}.Register()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
