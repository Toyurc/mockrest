## MockRest

#### Installation (unix)

- Clone the Repo

```
git clone http://github.com/deolu-asenuga/mockrest
```

- Install it

```
cd mockest/ && sudo make install
```

#### Usage

##### Single Endpoint
- to serve a json file (test.json) in the same directoy

```
mockrest serve test.json
```

navigate to localhost:3000/mockrest to see the result

- to serve a json file (test.json) from an endpoint "/end" on port 4000

```
mockrest serve test.json -e /end -p 4000
```
- Navigate to localhost:4000/end  to see the ressponse

##### Multiple Endpoints

- Create a config file ( default is "mockrest.json")
- Populate with configurations eg
```json
{
  "global":{
    "port":"3000"
  },
  "endpoints":[
    {"url":"/" , "payload":{"hello":"world"}}
  ]
}
```
- run mockrest in the same directory
```
mockrest
```

- Navigate to localhost:3000/ to see the result

