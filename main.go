 package main

import "fmt"

func main(){
  var day int
  var guest int
  var sum int
  
  fmt.Println("Введите день недели:")
  fmt.Scan(&day)

  fmt.Println("Введите число гостей:")
  fmt.Scan(&guest)

  fmt.Println("Введите сумму чека:")
  fmt.Scan(&sum)
 
  if day == 1 {
     if guest > 5 {
      discount:= sum /100 *10
      bonus := sum / 100 * 10
      sum = sum - discount + bonus 
      fmt.Println("Скидка по понедельникам:",discount,"рублей")
      fmt.Println("Надбавка на обслуживание:", bonus)
      fmt.Println("Сумма к оплате:", sum,"рублей")
    } else if {
      discount:= sum /100 *10
       sum -= discount
       fmt.Println("Скидка по понедельникам:",discount,"рублей")
       fmt.Println("Сумма к оплате:", sum,"рублей")
      
}