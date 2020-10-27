/*
 SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

//SmartContract is the data structure which represents this contract and on which  various contract lifecycle functions are attached
type SmartContract struct {
}

type Seller struct {
	ObjectType string `json:"Type"`
	SellerID   string `json:"sellerID"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type Buyer struct {
	ObjectType string `json:"Type"`
	BuyerID    string `json:"buyerID"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type Marketer struct {
	ObjectType string `json:"Type"`
	MarketerID string `json:"marketerID"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type Product struct {
	ObjectType  string `json:"Type"`
	ProductID   string `json:"productID"`
	SellerID    string `json:"sellerID"`
	BuyerID     string `json:"buyerID"`
	ProductName string `json:"productName"`
	Price       string `json:"price"`
	Description string `json:"description"`
	SoldDate    string `json:"soldDate"`
	Status      string `json:"status"`
}

type MarketedProduct struct {
	ObjectType string `json:"Type"`
	MID        string `json:"mID"`
	MarketerID string `json:"marketerID"`
	ProductID  string `json:"productID"`
	SellerID   string `json:"sellerID"`
	SoldDate   string `json:"soldDate"`
	Status     string `json:"status"`
}

func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {

	fmt.Println("Init Firing!")
	return shim.Success(nil)
}

func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("Chaincode Invoke Is Running " + function)
	if function == "addSeller" {
		return t.addSeller(stub, args)
	}
	if function == "querySeller" {
		return t.querySeller(stub, args)
	}
	if function == "queryAllSellers" {
		return t.queryAllSellers(stub, args)
	}
	if function == "querySellerByName" {
		return t.querySellerByName(stub, args)
	}
	if function == "querySellerByID" {
		return t.querySellerByID(stub, args)
	}
	if function == "addBuyer" {
		return t.addBuyer(stub, args)
	}
	if function == "queryBuyer" {
		return t.queryBuyer(stub, args)
	}
	if function == "queryBuyerByID" {
		return t.queryBuyerByID(stub, args)
	}
	if function == "addMarketer" {
		return t.addMarketer(stub, args)
	}
	if function == "queryMarketer" {
		return t.queryMarketer(stub, args)
	}
	if function == "queryMarketerByID" {
		return t.queryMarketerByID(stub, args)
	}
	if function == "addProduct" {
		return t.addProduct(stub, args)
	}
	if function == "queryProduct" {
		return t.queryProduct(stub, args)
	}
	if function == "queryProducts" {
		return t.queryProducts(stub, args)
	}
	if function == "queryProductByID" {
		return t.queryProductByID(stub, args)
	}
	if function == "queryProductbySellerID" {
		return t.queryProductbySellerID(stub, args)
	}
	if function == "queryProductbyBuyerID" {
		return t.queryProductbyBuyerID(stub, args)
	}
	if function == "queryProductbyProductName" {
		return t.queryProductbyProductName(stub, args)
	}
	if function == "queryProductbyStatus" {
		return t.queryProductbyStatus(stub, args)
	}
	if function == "queryProductbyDate" {
		return t.queryProductbyDate(stub, args)
	}
	if function == "addMarketedProduct" {
		return t.addMarketedProduct(stub, args)
	}
	if function == "queryMarketedProduct" {
		return t.queryMarketedProduct(stub, args)
	}
	if function == "queryMarketedProductByID" {
		return t.queryMarketedProductByID(stub, args)
	}
	if function == "queryMarketedProductBySellerID" {
		return t.queryMarketedProductBySellerID(stub, args)
	}
	if function == "queryMarketedProductByMarketerID" {
		return t.queryMarketedProductByMarketerID(stub, args)
	}
	if function == "queryMarketedProductByProductID" {
		return t.queryMarketedProductByProductID(stub, args)
	}
	if function == "queryMarketedProductByDate" {
		return t.queryMarketedProductByDate(stub, args)
	}
	if function == "queryMarketedProductByStatus" {
		return t.queryMarketedProductByStatus(stub, args)
	}
	if function == "updateProduct" {
		return t.updateProduct(stub, args)
	}
	if function == "updateMarketedProduct" {
		return t.updateMarketedProduct(stub, args)
	}

	fmt.Println("Invoke did not find specified function " + function)
	return shim.Error("Invoke did not find specified function " + function)
}

func (t *SmartContract) addSeller(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect Number of Aruments. Expecting 3")
	}

	fmt.Println("Adding new seller")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}

	sellerID := args[0]
	username := args[1]
	password := args[2]

	// ======Check if seller Already exists

	sellerAsBytes, err := stub.GetState(sellerID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if sellerAsBytes != nil {
		return shim.Error("The Inserted seller ID already Exists: " + sellerID)
	}

	// ======Check if seller Already exists

	sellerAsName, err := stub.GetState(username)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if sellerAsName != nil {
		return shim.Error("The Inserted seller username already Exists: " + username)
	}

	// ===== Create seller Object and Marshal to JSON

	objectType := "seller"
	seller := &Seller{objectType, sellerID, username, password}
	sellerJSONasBytes, err := json.Marshal(seller)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save seller to State

	err = stub.PutState(sellerID, sellerJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved seller")
	return shim.Success(nil)
}

func (t *SmartContract) querySeller(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	username := args[0]
	password := args[1]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"seller\",\"username\":\"%s\",\"password\":\"%s\"}}", username, password)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) querySellerByName(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	username := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"seller\",\"username\":\"%s\"}}", username)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) querySellerByID(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	sellerID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"seller\",\"sellerID\":\"%s\"}}", sellerID)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryAllSellers(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"seller\"}}")

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addBuyer(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect Number of Aruments. Expecting 3")
	}

	fmt.Println("Adding new Buyer")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}

	buyerID := args[0]
	username := args[1]
	password := args[2]

	// ======Check if Buyer Already exists

	buyerAsBytes, err := stub.GetState(buyerID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if buyerAsBytes != nil {
		return shim.Error("The Inserted buyer ID already Exists: " + buyerID)
	}

	// ======Check if seller Already exists

	buyerAsName, err := stub.GetState(username)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if buyerAsName != nil {
		return shim.Error("The Inserted buyer username already Exists: " + username)
	}

	// ===== Create Buyer Object and Marshal to JSON

	objectType := "buyer"
	buyer := &Buyer{objectType, buyerID, username, password}
	buyerJSONasBytes, err := json.Marshal(buyer)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save Buyer to State

	err = stub.PutState(buyerID, buyerJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved Buyer")
	return shim.Success(nil)
}

func (t *SmartContract) queryBuyer(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	username := args[0]
	password := args[1]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"buyer\",\"username\":\"%s\",\"password\":\"%s\"}}", username, password)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryBuyerByID(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	buyerID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"buyer\",\"buyerID\":\"%s\"}}", buyerID)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addMarketer(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect Number of Aruments. Expecting 3")
	}

	fmt.Println("Adding new Marketer")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}

	marketerID := args[0]
	username := args[1]
	password := args[2]

	// ======Check if Marketer Already exists

	marketerAsBytes, err := stub.GetState(marketerID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if marketerAsBytes != nil {
		return shim.Error("The Inserted Marketer ID already Exists: " + marketerID)
	}

	// ======Check if seller Already exists

	marketerAsName, err := stub.GetState(username)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if marketerAsName != nil {
		return shim.Error("The Inserted Marketer username already Exists: " + username)
	}

	// ===== Create Buyer Object and Marshal to JSON

	objectType := "marketer"
	marketer := &Marketer{objectType, marketerID, username, password}
	marketerJSONasBytes, err := json.Marshal(marketer)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save marketer to State

	err = stub.PutState(marketerID, marketerJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved Marketer")
	return shim.Success(nil)
}

func (t *SmartContract) queryMarketer(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	username := args[0]
	password := args[1]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"marketer\",\"username\":\"%s\",\"password\":\"%s\"}}", username, password)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryMarketerByID(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	marketerID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"marketer\",\"marketerID\":\"%s\"}}", marketerID)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addProduct(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 8 {
		return shim.Error("Incorrect Number of Aruments. Expecting 8")
	}

	fmt.Println("Adding new Product")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th Argument Must be a Non-Empty String")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th Argument Must be a Non-Empty String")
	}
	if len(args[5]) <= 0 {
		return shim.Error("6th Argument Must be a Non-Empty String")
	}
	if len(args[6]) <= 0 {
		return shim.Error("7th Argument Must be a Non-Empty String")
	}
	if len(args[7]) <= 0 {
		return shim.Error("8th Argument Must be a Non-Empty String")
	}

	productID := args[0]
	sellerID := args[1]
	buyerID := args[2]
	productName := args[3]
	price := args[4]
	description := args[5]
	soldDate := args[6]
	status := args[7]

	// ======Check if product Already exists

	productAsBytes, err := stub.GetState(productID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if productAsBytes != nil {
		return shim.Error("The Inserted productID ID already Exists: " + productID)
	}

	// ===== Create product Object and Marshal to JSON

	objectType := "product"
	product := &Product{objectType, productID, sellerID, buyerID, productName, price, description, soldDate, status}
	productJSONasBytes, err := json.Marshal(product)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save product to State

	err = stub.PutState(productID, productJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved product")
	return shim.Success(nil)
}

func (t *SmartContract) queryProductByID(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	productID := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"product\",\"productID\":\"%s\"}}", productID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryProduct(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	status := args[0]
	productID := args[1]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"product\",\"status\":\"%s\",\"productID\":\"%s\"}}", status, productID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryProducts(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	sellerID := args[0]
	buyerID := args[1]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"product\",\"sellerID\":\"%s\",\"lenderID\":\"%s\"}}", sellerID, buyerID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryProductbySellerID(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	sellerID := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"product\",\"sellerID\":\"%s\"}}", sellerID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryProductbyBuyerID(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	sellerID := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"product\",\"sellerID\":\"%s\"}}", sellerID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryProductbyProductName(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	productName := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"product\",\"productName\":\"%s\"}}", productName)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryProductbyStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	status := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"product\",\"status\":\"%s\"}}", status)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryProductbyDate(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	soldDate := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"product\",\"soldDate\":\"%s\"}}", soldDate)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addMarketedProduct(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 6 {
		return shim.Error("Incorrect Number of Aruments. Expecting 7")
	}

	fmt.Println("Adding new MarketedProduct")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th Argument Must be a Non-Empty String")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th Argument Must be a Non-Empty String")
	}
	if len(args[5]) <= 0 {
		return shim.Error("6th Argument Must be a Non-Empty String")
	}

	mID := args[0]
	marketerID := args[1]
	productID := args[2]
	sellerID := args[3]
	soldDate := args[4]
	status := args[5]

	// ======Check if MarketedProduct Already exists

	marketedAsBytes, err := stub.GetState(mID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if marketedAsBytes != nil {
		return shim.Error("The Inserted marketed ID already Exists: " + mID)
	}

	// ===== Create MarketedProduct Object and Marshal to JSON

	objectType := "marketed"
	marketed := &MarketedProduct{objectType, mID, marketerID, productID, sellerID, soldDate, status}
	marketedJSONasBytes, err := json.Marshal(marketed)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save MarketedProduct to State

	err = stub.PutState(mID, marketedJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved MarketedProduct")
	return shim.Success(nil)
}

func (t *SmartContract) queryMarketedProduct(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	status := args[0]
	marketerID := args[1]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"marketed\",\"status\":\"%s\",\"marketerID\":\"%s\"}}", status, marketerID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryMarketedProductByID(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	mID := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"marketed\",\"mID\":\"%s\"}}", mID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryMarketedProductBySellerID(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	sellerID := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"marketed\",\"sellerID\":\"%s\"}}", sellerID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryMarketedProductByMarketerID(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	marketerID := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"marketed\",\"marketerID\":\"%s\"}}", marketerID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryMarketedProductByProductID(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	productID := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"marketed\",\"productID\":\"%s\"}}", productID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryMarketedProductByDate(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	soldDate := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"marketed\",\"soldDate\":\"%s\"}}", soldDate)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryMarketedProductByStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	status := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"marketed\",\"status\":\"%s\"}}", status)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) updateProduct(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	productID := args[0]
	newStatus := args[1]
	soldDate := args[2]
	buyerID := args[3]
	fmt.Println("- start  ", productID, newStatus, soldDate, buyerID)

	productAsBytes, err := stub.GetState(productID)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if productAsBytes == nil {
		return shim.Error("product does not exist")
	}

	productToUpdate := Product{}
	err = json.Unmarshal(productAsBytes, &productToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	productToUpdate.Status = newStatus //change the status
	productToUpdate.SoldDate = soldDate
	productToUpdate.BuyerID = buyerID

	productJSONasBytes, _ := json.Marshal(productToUpdate)
	err = stub.PutState(productID, productJSONasBytes) //rewrite
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateMarketedProduct(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	mID := args[0]
	newStatus := args[1]
	soldDate := args[2]
	fmt.Println("- start  ", mID, newStatus, soldDate)

	marketedAsBytes, err := stub.GetState(mID)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if marketedAsBytes == nil {
		return shim.Error("marketed does not exist")
	}

	marketedToUpdate := MarketedProduct{}
	err = json.Unmarshal(marketedAsBytes, &marketedToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	marketedToUpdate.Status = newStatus //change the status
	marketedToUpdate.SoldDate = soldDate

	marketedJSONasBytes, _ := json.Marshal(marketedToUpdate)
	err = stub.PutState(mID, marketedJSONasBytes) //rewrite
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

//Main Function starts up the Chaincode
func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Smart Contract could not be run. Error Occured: %s", err)
	} else {
		fmt.Println("Smart Contract successfully Initiated")
	}
}
