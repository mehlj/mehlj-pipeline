package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
  "bytes"
)

func TestDefaultHandler(t *testing.T) {
  req, err := http.NewRequest("GET", "/", nil)

  // fail test if query returns error
  if err != nil {
    t.Fatal(err)
  }

  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(DefaultHandler)
  handler.ServeHTTP(rr, req)

  // ensure status code is OK
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: got %v want %v",
      status, http.StatusOK)
  }

  // ensure response body is what we expect
  expected := "Hello world\n"
  if rr.Body.String() != expected {
    t.Errorf("handler returned unexpected body: got %v want %v",
      rr.Body.String(), expected)
  }
}

func TestQueryAllProducts(t *testing.T) {
  req, err := http.NewRequest("GET", "/products", nil)

  // fail test if query returns error
  if err != nil {
    t.Fatal(err)
  }

  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(QueryAllProducts)
  handler.ServeHTTP(rr, req)

  // ensure status code is OK
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: got %v want %v",
      status, http.StatusOK)
  }

  // ensure response body is what we expect
  expected := `[{"Name":"apple","quantity":54},{"Name":"pear","quantity":12}]` + "\n"

  if rr.Body.String() != expected {
    t.Errorf("handler returned unexpected body: got %v want %v",
      rr.Body.String(), expected)
  }
}

func TestCreateProduct(t *testing.T) {
  var newProduct = []byte(`{"Name":"testobject","quantity":3}`)

  req, err := http.NewRequest("POST", "/product", bytes.NewBuffer(newProduct))

  // fail test if query returns error
  if err != nil {
    t.Fatal(err)
  }

  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(CreateProduct)
  handler.ServeHTTP(rr, req)

  // ensure status code is OK
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: got %v want %v",
      status, http.StatusOK)
  }

  // ensure response body is what we expect
  expected := `{"Name":"testobject","quantity":3}` + "\n"

  if rr.Body.String() != expected {
    t.Errorf("handler returned unexpected body: got %v want %v",
      rr.Body.String(), expected)
  }

  // ensure that the product is actually added
  req, err = http.NewRequest("GET", "/products", nil)
  rr = httptest.NewRecorder()
  handler = http.HandlerFunc(QueryAllProducts)
  handler.ServeHTTP(rr, req)
  expected = `[{"Name":"apple","quantity":54},{"Name":"pear","quantity":12},{"Name":"testobject","quantity":3}]` + "\n"

  if rr.Body.String() != expected {
    t.Errorf("handler did not successfully add object. got %v want %v",
      rr.Body.String(), expected)
  }
}

func TestUpdateProduct(t *testing.T) {
  var newProductDetails = []byte(`{"Name":"testobject","quantity":6}`)

  req, err := http.NewRequest("PUT", "/product", bytes.NewBuffer(newProductDetails))

  // fail test if query returns error
  if err != nil {
    t.Fatal(err)
  }

  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(UpdateProduct)
  handler.ServeHTTP(rr, req)

  // ensure status code is OK
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: got %v want %v",
      status, http.StatusOK)
  }

  // ensure response body is what we expect
  expected := `{"Name":"testobject","quantity":6}` + "\n"

  if rr.Body.String() != expected {
    t.Errorf("handler returned unexpected body: got %v want %v",
      rr.Body.String(), expected)
  }

  // ensure that the product is actually updated
  req, err = http.NewRequest("GET", "/products", nil)
  rr = httptest.NewRecorder()
  handler = http.HandlerFunc(QueryAllProducts)
  handler.ServeHTTP(rr, req)
  expected = `[{"Name":"apple","quantity":54},{"Name":"pear","quantity":12},{"Name":"testobject","quantity":6}]` + "\n"

  if rr.Body.String() != expected {
    t.Errorf("handler did not successfully add object. got %v want %v",
      rr.Body.String(), expected)
  }
}

func TestQueryProduct(t *testing.T) {
  req, err := http.NewRequest("GET", "/product?name=testobject", nil)

  // fail test if query returns error
  if err != nil {
    t.Fatal(err)
  }

  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(QueryProduct)
  handler.ServeHTTP(rr, req)

  // ensure status code is OK
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: got %v want %v",
      status, http.StatusOK)
  }

  // ensure response body is what we expect
  expected := `{"Name":"testobject","quantity":6}` + "\n"

  if rr.Body.String() != expected {
    t.Errorf("handler returned unexpected body: got %v want %v",
      rr.Body.String(), expected)
  }
}


func TestDeleteProduct(t *testing.T) {
  var productToBeDeleted = []byte(`{"Name":"testobject","quantity":3}`)

  req, err := http.NewRequest("DELETE", "/product", bytes.NewBuffer(productToBeDeleted))

  // fail test if query returns error
  if err != nil {
    t.Fatal(err)
  }

  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(DeleteProduct)
  handler.ServeHTTP(rr, req)

  // ensure status code is OK
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: got %v want %v",
      status, http.StatusOK)
  }

  // ensure response body is as we expect
  expected := `{"Name":"testobject","quantity":3}` + "\n"

  if rr.Body.String() != expected {
    t.Errorf("handler did not successfully delete object. got %v want %v",
      rr.Body.String(), expected)
  }

  // ensure that the product is actually deleted
  req, err = http.NewRequest("GET", "/products", nil)
  rr = httptest.NewRecorder()
  handler = http.HandlerFunc(QueryAllProducts)
  handler.ServeHTTP(rr, req)
  expected = `[{"Name":"apple","quantity":54},{"Name":"pear","quantity":12}]` + "\n"

  if rr.Body.String() != expected {
    t.Errorf("handler did not successfully delete object. got %v want %v",
      rr.Body.String(), expected)
  }
}
