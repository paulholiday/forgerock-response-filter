# forgerock-response-filter

Simple script that takes a file directory consisting of JSON files in the structure of a ForgeRock IDM REST API user query response and prints out all the
usernames in the response.

If the query response contains more than just usernames then the Data struct will need the extra fields adding to it