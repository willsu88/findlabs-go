## Goals

1. Implement multiple sorting functions, including sort by name, address, transaction...
    - Should be able to read from some Json and parse the data and sort as needed
    - this requires me to generate some fake json data.
    - handle bad data
2. Write testing code that tests the sorting functions.
3. 


HTTP Request from Client to know what sorting mechanism to use;

Query Data from Table; 

TODO:
1. Sorting function dynamic API with query params 
2. Filtering function dynamic API
    - filter by name
    - filter by transaction
    - filter by date

3. Refactor SQL query and contracts relationship. Query first then sort or query every time?


Improvements:
- Use gin instead of mux as a web framework
- support sorting for address, transactions. support prefix-find function for address transactions