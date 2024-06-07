# resq

Simple CLI REST Client.

## Example

resq -r request.yaml -o response.json

## Request Definition File 

```yaml
version: "1.0"
name: "Name of your request"
url: "https://example.com/{param1}/{param2}" # URL of the request
method: GET                                  # Method

path:                                        # Path parameters
  param1: value1
  param2: value2

query:                                       # Query parameters
  page: 1
  search: "example"

headers:                                     # Request headers
  Content-Type: application/json
  Accept: application/json

auth:                                        # Auth helper    
  type: basic                                # Auth method api_key, basic
  
  username: your_username                    # User name for basic auth
  password: your_password                    # Password for basic auth
  
  key: your_api_key                          # Key for api_key auth
  header_name: your_header_name              # Header to include key in api_key auth

body: |-                                     # Body of the request
  {
    "key": "value"
  }

```