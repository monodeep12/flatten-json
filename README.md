# flattenjson

## What is flattenjson?
flattenjson is a go library to flatten nested JSON structure. 

## How does it work?
```go
package main

import (
    "fmt"

    "github.com/monodeep12/flattenjson"
)
func main() {
    input := []byte(`{
  "_id": "57145f82f7579b1b172afdef",
  "brand_info": {
      "media_plan_name": "Media Plan 3",
      "brand_name": "Dabur",
      "product_category": "Shoes",
      "product_sub_category": "Formal Shoes",
      "gender": ["Male"],
      "age": ["35+","15-35"],
      "audience_type": ["Urban"],
      "geography": ["BLR"]
  }
}`)
  // passing `flattenSlice` flag as `true` will result in flattening the inner slices as well 
	out1, err := flattenjson.JSONByte(input, ".", true)
	// passing `flattenSlice` flag as `false` will result in not flattening the inner slices
	out2, err := flattenjson.JSONByte(input, ".", false)
	if err != nil {
		// handle error
	}
	fmt.Println(string(out1))
	fmt.Println(string(out2))
}
```

Output1:
```javascript
{"_id":"57145f82f7579b1b172afdef","brand_info.age.0":"35+","brand_info.age.1":"15-35","brand_info.audience_type.0":"Urban","brand_info.brand_name":"Dabur","brand_info.gender.0":"Male","brand_info.geography.0":"BLR","brand_info.media_plan_name":"Media Plan 3","brand_info.product_category":"Shoes","brand_info.product_sub_category":"Formal Shoes"}
```

Output2:
```javascript
{"_id":"57145f82f7579b1b172afdef","brand_info.age":["35+","15-35"],"brand_info.audience_type":["Urban"],"brand_info.brand_name":"Dabur","brand_info.gender":["Male"],"brand_info.geography":["BLR"],"brand_info.media_plan_name":"Media Plan 3","brand_info.product_category":"Shoes","brand_info.product_sub_category":"Formal Shoes"}
```
