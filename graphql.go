package main

import (
	"github.com/graphql-go/graphql"
	"encoding/json"
)

/**
	This file handles the complete GraphQL-API.
 */

type Item struct {
	Family_discount float64 `json:"family_discount"`
	Discout_perc float64 `json:"discout_perc"`
	Additional_stuff float64 `json:"additional_stuff"`

	Id        int       `json:"id"`
	ItemType  ItemType  `json:"itemType"`

	PriceNew  float64   `json:"price_new"`
	Condition Condition `json:"condition"`
	Amortisation_factor float64 `json:"amortisation_factor"`

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
				Type: graphql.Float, //return total price
				//Type: graphql.NewList(itemOutput), //result-type for an array of items

				//Configure all input types of a request
				Args: graphql.FieldConfigArgument{
					"items": &graphql.ArgumentConfig{
						Type: graphql.NewList(itemInput),
					},
				},

				//Resolve forwards the input data to the calculation algorithm
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					//convert input to []Item{}-Object
					//not the best solution but works
					jsonItemArray, _ := json.Marshal(p.Args["items"])
					//fmt.Printf("%s \n", jsonItemArray)
					items := []Item{}
					err := json.Unmarshal(jsonItemArray, &items)

					if err == nil {
						//do the calculation it input was correct
						return calculateTotalPrice(items)
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
			"family_discount": &graphql.InputObjectFieldConfig{
				Type: graphql.Float,
			},
			"dicount_perc": &graphql.InputObjectFieldConfig{
				Type: graphql.Float,
			},
			"additional_stuff": &graphql.InputObjectFieldConfig{
				Type: graphql.Float,
			},
			"id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"itemType": &graphql.InputObjectFieldConfig{
				Type: itemTypeEnum,
			},
			"price_new": &graphql.InputObjectFieldConfig{
				Type: graphql.Float,
			},
			"condition": &graphql.InputObjectFieldConfig{
				Type: conditionEnum,
			},
			"amortisation_factor": &graphql.InputObjectFieldConfig{
				Type: graphql.Float,
			},
		},
	},
)