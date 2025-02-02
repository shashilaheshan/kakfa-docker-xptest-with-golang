#Sample kafka consumer and producer test with single node , with 2 partitions

#You can run the docker-compose.yaml file inside the project root using docker-compose up -d

#This will install kafka , zookeeper, kafka-ui in ur docker env . 

#Open any browser and go to [localhost:8080](http://localhost:8080/ui/) to check your kafka cluster and create any topic name u want , here i used heshan-docker-xp-debug topic name which i have used inside my go app (consumer and producer) with two partitions attached.

#Go inside project root and run the go dependancy init using go mod tidy command . This will install required package to run the go application.
#Then run the consumer by running go run main.go command inside the project root terminal.

#Then move to producer/ folder to run the producer.go . Type go run producer.go to run it . It will produce sample test messages in the given topic.

#If all good consumer will consume the message correctly. Cheers

