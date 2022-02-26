
# Crud api with mongodb and Golang 

# Installation Steps
 
# Quick Run Project
     go get github.com/pavanilla/appointy 
     
               or 
   
     git clone https://github.com/pavanilla/appointy.git
     
    cd appointy
    
    go build main.go


## Need to import following packages

	  1.go get go.mongodb.org/mongo-driver/mongo

	  2.go get go.mongodb.org/mongo-driver/mongo/options

	  3. go get go.mongodb.org/mongo-driver/bson
  
      4. go get go.mongodb.org/mongo-driver/bson/primitive 
    
      5. go get github.com/gorilla/mux
      

      
      
## Route end points 
    
	  1.curl -v -H "Accept: application/json" -H "Content-type: application/json" -X GET http://localhost:8080/users/getUser 

	  2.curl -v -H "Accept: application/json" -H "Content-type: application/json" -X GET http://localhost:8080/users?id={}

	  3.curl -v -H "Accept: application/json" -H "Content-type: application/json" -X GET  http://localhost:8080/users?name={}&password={}
	
      4.curl -v -H "Accept: application/json" -H "Content-type: application/json" -X  POST http://localhost:8080/users/user

	  5.curl -i -H "Accept: application/json" -H "Content-type: application/json" -X PUT http://localhost:8080/users?id={}
    
      6.curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST http://localhost:8080/posts/post
    
      7.curl -i -H "Accept: application/json" -H "Content-type: application/json" -X PUT http://localhost:8080/posts/edit?id={}
    
      8.curl -i -H "Accept: application/json" -H "Content-type: application/json" -X GET http://localhost:8080/posts/list
    
      9.curl -i -H "Accept: application/json" -H "Content-type: application/json" -X DELETE http://localhost:8080/posts/delete?id={}
   
  
