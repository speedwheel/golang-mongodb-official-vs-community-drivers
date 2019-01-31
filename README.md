# golang-mongodb-official-vs-community-drivers
![mongodb golang official drive vs community driver](https://i.imgur.com/QTRUv16.jpg)

```
$ go get -u github.com/codesenberg/bombardier

$ go run main.go

$ bombardier -n 10000 -c 100 localhost:8080/api/store/cars/5c5227e2fe5d37009859ac9f
```

mongodb document:
```
{ 
    "_id" : ObjectId("5c5227e2fe5d37009859ac9f"), 
    "test_id" : NumberInt(347996663), 
    "vehicle_id" : NumberInt(146941831), 
    "test_date" : ISODate("2017-04-10T00:00:00.000+0000"), 
    "test_class_id" : NumberInt(4), 
    "test_type" : "NT", 
    "test_result" : "P", 
    "test_mileage" : "42356", 
    "postcode_area" : "HP", 
    "make" : "VAUXHALL", 
    "model" : "INSIGNIA", 
    "colour" : "GREY", 
    "fuel_type" : "DI", 
    "cylinder_capacity" : NumberInt(1956), 
    "first_use_date" : ISODate("2014-06-09T00:00:00.000+0000")
}
```
