# Release Branch:

https://ppl-pricecalculator.herokuapp.com/

[![Build (Master) Status](https://travis-ci.org/ob-vss-ss18/ppl-pricecalculator.svg?branch=release)](https://travis-ci.org/ob-vss-ss18/ppl-pricecalculator)


# Development Branch:

https://ppl-pricecalculator-dev.herokuapp.com/

[![Build (Master) Status](https://travis-ci.org/ob-vss-ss18/ppl-pricecalculator.svg?branch=master)](https://travis-ci.org/ob-vss-ss18/ppl-pricecalculator)

# GraphQL

### GrapQL-API

```graphql
type RootQuery {
    calculateTotalPrice(items: [Item]): Float
}

input Item {
  family_discount: Int!
  discout_perc: Float!
  additional_stuff: Float!

  # you need id and itemType or
  # price_new, condition and amortisation_factor
  id: Int
  itemType: ItemType

  price_new: Float
  condition: Condition
  amortisation_factor: Float
}

enum ItemType {
  SKI
  STICK
}

enum Condition {
  NEW
  USED
}
```

### Beispiel GraphQL-Abfrage

* Programm zum testen: GraphiQL (https://electronjs.org/apps/graphiql)
* GraphQL Endpoint: http://localhost:8000/graphql
* Method: GET

**Request:**
```
{
  calculateTotalPrice(items: [
    {family_discount: 0.0, dicount_perc: 0.0, additional_stuff: 50.0, id: 22, price_new: 500, itemType: SKI, condition: USED, amortisation_factor: 0.5},
    {family_discount: 0.0, dicount_perc: 0.0, additional_stuff: 50.0, id: 22, price_new: 500, itemType: SKI, condition: USED, amortisation_factor: 0.5}
  ])
}
```

**Response:**
```
{
  "data": {
    "calculateTotalPrice": 398
  }
}
```