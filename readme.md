# resq

Simple CLI REST Client.

## Example

resq -f request.yaml -o response.json

## Format 

```yaml
version: "1.0"
name: "Name of your request"
url: "https://example.com/{param1}/{param2}" # URL of the request
method: GET                                  # GET, POST, PUT, DELETE, PATCH

path:                                        # Include path parameters here
  param1: value1
  param2: value2

query:                                       # Include query parameters here
  page: 1
  search: "example"

headers:                                     # Include headers here
  Content-Type: application/json
  Accept: application/json

body: '{"key":"value"}'                      # Body of the request

auth:                                        # Include authentication here       
  type: basic                                # basic, api_key
  
  # Use with type: basic
  username: your_username
  password: your_password
  
  # Use with type: api_key
  key: your_api_key
  header_name: your_header_name

```