# Release Branch:

https://ppl-pricecalculator.herokuapp.com/

[![Build (Master) Status](https://travis-ci.org/ob-vss-ss18/ppl-pricecalculator.svg?branch=release)](https://travis-ci.org/ob-vss-ss18/ppl-pricecalculator)


# Development Branch:

https://ppl-pricecalculator-dev.herokuapp.com/

[![Build (Master) Status](https://travis-ci.org/ob-vss-ss18/ppl-pricecalculator.svg?branch=master)](https://travis-ci.org/ob-vss-ss18/ppl-pricecalculator)

# GraphQL

### GrapQL-API

```
type RootQuery {
    calculateTotalPrice(items: [Item], family_discount: Int): Int
}

input Item {
  id: Int
  price_new: Float
  itemType: ItemType
  condition: Condition
}

enum Condition {
  NEW
  USED
}

enum ItemType {
  SKI
  STICK
}
```

### Beispiel GraphQL-Abfrage

* Programm zum testen: GraphiQL (https://electronjs.org/apps/graphiql)
* GraphQL Endpoint: http://localhost:8000/graphql
* Method: GET

**Request:**
```
{
  calculateTotalPrice(
      items:
          [
            {id: 22, price_new: 324.40, itemType: SKI, condition: USED},
            {id: 24, price_new: 134.43, itemType: STICK, condition: NEW}
          ],
      family_discount: 10
  )
}
```

**Response:**
```
{
  "data": {
    "calculateTotalPrice": 42
  }
}
```