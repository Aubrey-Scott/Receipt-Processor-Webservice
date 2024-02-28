# Receipt-Processor-Webservice

-----------------INSTRUCTIONS-------------------

1.) Have the most up to date version of Go installed (which can be found here: https://go.dev/dl/)



2.) Have the most up to date version of Gorilla Mux (download instructions can be found at the top of the page by the installation section here: https://pkg.go.dev/github.com/gorilla/mux#section-readme)



3.) Keep all files in a folder together



4.)Use command "go run main.go pointCalc.go requests.go" to run the program



5.)Send POST requests with a JSON formatted like the examples in the examples folder to "receipts/processor" get a JSON with an id for your stored receipt "score"



6.)Send GET requests formatted as "receipts/{id}/points" with the id being the string from the id JSON created during the POST request, this will return a JSON with your score as an int



7.) Use port 8081 when testing or using the program



8.) Example/Test Jsons are in the Examples folder
