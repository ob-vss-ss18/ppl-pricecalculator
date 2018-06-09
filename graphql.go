package main

import (
	"github.com/graphql-go/graphql"
	"encoding/json"
)

/**
	This file handles the complete GraphQL-API.
 */

type Item struct {
	Id        int       `json:"id"`
	PriceNew  float32   `json:"price_new"`
	ItemType  ItemType  `json:"itemType"`
	Condition Condition `json:"condition"`
}

type Condition int

const (
	NEW  Condition = 0
	USED Condition = 1
)

type ItemType int

const (
	SKI   ItemType = 0
	STICK ItemType = 1
)

var (
	priceCalculationSchema graphql.Schema
)

// initialization of the complete GraphQL-API
func initGraphQl() {

	// Each request starts with the root query
	// and walks down the graph
	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{

			//name of the query-function of the GraphQL-API
			"calculateTotalPrice": &graphql.Field{

				//result-type for the response
				Type: graphql.Int, //return total price
				//Type: graphql.NewList(itemOutput), //result-type for an array of items

				//Configure all input types of a request
				Args: graphql.FieldConfigArgument{
					"items": &graphql.ArgumentConfig{
						Type: graphql.NewList(itemInput),
					},
					"family_discount": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},

				//Resolve forwards the input data to the calculation algorithm
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					//get the values from the request
					familyDiscountQuery, familyDiscountIsOK := p.Args["family_discount"].(int)

					//convert input to []Item{}-Object
					//not the best solution but works
					jsonItemArray, _ := json.Marshal(p.Args["items"])
					//fmt.Printf("%s \n", jsonItemArray)
					items := []Item{}
					err := json.Unmarshal(jsonItemArray, &items)

					if familyDiscountIsOK && err == nil {
						//do the calculation it input was correct
						return calculateTotalPrice(familyDiscountQuery, items)
					}
					return nil, nil
				},
			},
		},
	})

	//definition of the schema
	priceCalculationSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
}

var conditionEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "condition",
	Description: "The condition of the item",
	Values: graphql.EnumValueConfigMap{
		"NEW": &graphql.EnumValueConfig{
			Value: NEW,
		},
		"USED": &graphql.EnumValueConfig{
			Value: USED,
		},
	},
})

var itemTypeEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "itemType",
	Description: "The type of the item",
	Values: graphql.EnumValueConfigMap{
		"SKI": &graphql.EnumValueConfig{
			Value: SKI,
		},
		"STICK": &graphql.EnumValueConfig{
			Value: STICK,
		},
	},
})

/*
	Converting GraphQL-Input to an item
*/
var itemInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "itemInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"price_new": &graphql.InputObjectFieldConfig{
				Type: graphql.Float,
			},
			"itemType": &graphql.InputObjectFieldConfig{
				Type: itemTypeEnum,
			},
			"condition": &graphql.InputObjectFieldConfig{
				Type: conditionEnum,
			},
		},
	},
)
