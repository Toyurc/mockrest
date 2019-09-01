## MockRest

#### Installation (linux)

* Clone the Repo
```
git clone http://github.com/deolu-asenuga/mockrest
```
* Install it
```
cd mockest/ && sudo make install
```


#### Usage

* to serve a json file (test.json) in the same directoy 
```
mockrest serve test.json
```
 navigate to localhost:3000/mockrest to see the result
* to serve a json file (test.json) from an endpoint "/end" on port 4000
```
mockrest serve test.json -e /end -p 4000
````