package main

import (
  "fmt"
  "log"
  "github.com/streadway/amqp"
)
func main(){
  go client()
  go server()
  var a int
  fmt.Scanln(&a)
}

func client(){
  conn, ch, q := getQueue()
  defer conn.Close()
  defer ch.Close()
  msgs, err := ch.Consume(
    q.Name, //queue,
    "",     //consumer,
    true,   //autoAck,
    false,  //exclusive,
    false,  //noLocal,
    false,  //noWait,
    nil)   //  args)
failOnError(err, "Failed to register a consumer")
for msg := range msgs{
  log.Printf("Recieved Messages with message body: %s",msg.Body)
}

}

func server(){
  conn,ch,q := getQueue()
  defer conn.Close()
  defer ch.Close()

  msg := amqp.Publishing{
    ContentType : "text/plain",
    Body : []byte ("Hello rabbit MQ"),
    }
for{
      ch.Publish("", q.Name, false, false, msg)
    }
}

func getQueue()(*amqp.Connection, *amqp.Channel, *amqp.Queue){
  // Connect
  conn,err := amqp.Dial("amqp://guest@localhost:5672")
  failOnError(err, "Failed to Coonect to RabbitMQ")
  // Open Channel
  ch,err := conn.Channel()
  failOnError(err, "Failed to Open a Channel")
  //Declare a queue
   q, err := ch.QueueDeclare("hello",
     false,          //durable,
     false,           //autoDelete,
     false,           //exclusive,
     false,           //noWait,
     nil)            //   args)
     failOnError(err, "Failed to decalre a queue")
     return conn,ch,&q
}



func failOnError(err error, msg string){
  if err != nil{
    log.Fatalf("%s:%s", msg, err)
    panic(fmt.Sprintf("%s:%s",msg,err))
  }
}
