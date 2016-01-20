# go-gemfire-rest
=================

This library enables your golang applications to use GemFire as a datastore. (GemFire is a distributed key-value store. A short tutorial can be found at http://goo.gl/rF93fn). This library is an implementation of the python rest api which can be found here https://github.com/gemfire/py-gemfire-rest

## Installation
---------------

Using go get installation is simple
```
    $ go get github.com/mross1080/gemfire-golang
```

## Quick Start
--------------

1. Start the GemFire REST service by [following the instructions](http://gemfire.docs.pivotal.io/latest/userguide/index.html#gemfire_rest/setup_config.html)
2. Create a Region on the server (Region is a distributed ConcurrentMap in which GemFire stores the data). 
```
    gfsh>create region --name=orders --type=PARTITION
```
3.  Setup the Api Conncetion
```golang
    import "github.com/mross1080/gemfire-golang"
    conn := gemfireGolang.Api{"http://127.0.0.1", "8080"}    
    customers := gemfireGolang.Region{conn,"customers"}
```
4. Create a JSON Object to populate the Region with
```golang
    user  := struct{Name string
                    Age int
                    Id string}{
                    "Marty McFly",
                    22,
                    "110"}
    u, err := json.Marshal(user)
    if err != nil {
        fmt.Println(err)
    }
```
5. Insert the new JSON object into the Region
```golang
    customers.Put(user.Id, u)
```

where the order struct has an "id" instance variable. The library does not handle converting the struct to/from json. 