# gRPC_User_Service
A Golang gRPC User Service

## How to use?

### Copy this repo
```bash
git clone https://github.com/nabin3/userInfo
```

### Building docker image
```bash
cd userInfo/
docker build . -t userinfoimage
```

### Spawn a docker instance of newly created image
```bash
docker run --name userinfoservice -d -p 3000:3000 userinfoimage
```
Now you can access different endpoint with **curl** or **thunder** on http://localhost:3000

### Different endpoints:

#### /adduser 
This endpoint is for creating a new user

    * Method: POST
    * Request Content-Type: application/json
    * Response Content-Type: application/json
    
    Example Request:
    {
        "fname": "Steve",
        "city": "NewYork",
        "Phone": "1234567890",
        "height": 6.0,
        "is_married": true
    } 
        
    Example Response:
    {
        "user_id": "01ef1f17-3f52-68e5-9fcd-48ba4ee522a4"
    }  

#### /getuser
This endpoint is for retrieveing a existing user by user_id

    * Method: GET
    * Request Content-Type: application/json
    * Response Content-Type: application/json
    
    Example Request:
    {
        "user_id": "01ef1f17-3f52-68e5-9fcd-48ba4ee522a4"
    }    

    Example Response:
    {
        "id": "01ef1f17-3f52-68e5-9fcd-48ba4ee522a4",
        "fname": "Steve",
        "city": "NewYork",
        "phone": "1234567890",
        "height": 6,
        "is_married": true
    }
 
#### /get_multiple_users
This endpoint is for retrieveing multiple users by a list of user_ids

    * Method: GET
    * Request Content-Type: application/json
    * Response Content-Type: application/json
    
    Example Request:
    {
        "user_id_list": [
                            "01ef1f17-3f52-68e5-9fcd-48ba4ee522a4",
                            "01ef1f17-c82e-6386-9fcd-48ba4ee522a4"
                        ]
    } 

    Example Response:
    [
        {
            "id": "01ef1f17-3f52-68e5-9fcd-48ba4ee522a4",
            "fname": "Steve",
            "city": "NewYork",
            "phone": "1234567890",
            "height": 6,
            "is_married": true
        },
        {
            "id": "01ef1f17-c82e-6386-9fcd-48ba4ee522a4",
            "fname": "Nabin",
            "city": "Kolkata",
            "phone": "1255567898",
            "height": 5.6,
            "is_married": false
        }
    ]
 
 #### /search_users
This endpoint is for searching multiple users by a search criteria name and data related to that mentioned criteria

    * Method: GET
    * Request Content-Type: application/json
    * Response Content-Type: application/json
    
    Example Request with search_criteria **fname** : 
    {
        "search_criteria": "fname",
        "value": "Nabin"
    }
       
    Example Response with search_criteria **fname** : 
    [
        {
            "id": "01ef1f17-c82e-6386-9fcd-48ba4ee522a4",
            "fname": "Nabin",
            "city": "Kolkata",
            "phone": "1255567898",
            "height": 5.6,
            "is_married": false
        }
    ]    

    Example Request with search_criteria **is_married**:
    {
        "search_criteria": "is_married",
        "value": true
    }

    Example Request with search_criteria **is_married**:
    [
        {
            "id": "01ef1f17-3f52-68e5-9fcd-48ba4ee522a4",
            "fname": "Steve",
            "city": "NewYork",
            "phone": "1234567890",
            "height": 6,
            "is_married": true
        },
        {
            "id": "01ef1f19-611c-628a-9fcd-48ba4ee522a4",
            "fname": "Sam",
            "city": "Kolkata",
            "phone": "12559999",
            "height": 5.6,
            "is_married": true
        }
    ]  

    Example Request with search criteria **city**:
    {
        "search_criteria": "city",
        "value": "Kolkata"
    }
        
    Example Response with search criteria **city**:
    [
        {
            "id": "01ef1f17-c82e-6386-9fcd-48ba4ee522a4",
            "fname": "Nabin",
            "city": "Kolkata",
            "phone": "1255567898",
            "height": 5.6,
            "is_married": false
        },
        {
            "id": "01ef1f19-611c-628a-9fcd-48ba4ee522a4",
            "fname": "Sam",
            "city": "Kolkata",
            "phone": "12559999",
            "height": 5.6,
            "is_married": true
        }
    ]

    Example Request with search_criteria **phone** :
    {
        "search_criteria": "phone",
        "value": "12559999"
    }

    Example Request with search_criteria **phone** :
    [
        {
            "id": "01ef1f19-611c-628a-9fcd-48ba4ee522a4",
            "fname": "Sam",
            "city": "Kolkata",
            "phone": "12559999",
            "height": 5.6,
            "is_married": true
        }
    ]