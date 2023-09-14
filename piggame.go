package main

import (
	"fmt"
	"math/rand"
)
func menu() string {
	fmt.Println("Enter 'r' to roll again, 'h' to hold.")
	choice:= "-"
	fmt.Scan(&choice)
	return choice
}
func roll()int32{
	roll_value := rand.Intn(6)+1
	return int32(roll_value)
}
func main(){
	fmt.Println("Welcome to the game of Pig!\n")
	var total_score int32 =0;
	// var hold_value int32 = 0;
	var turn_score int32 = 0
	var turn_no int64 = 1
	for total_score<20{
		choice := menu()
		switch choice {
		case "r":
			turn_value := int32(roll())
			fmt.Println("your turn value is : ",turn_value)
			if turn_value==1 {
				turn_score = 0
				fmt.Println("Turn over. No score")
				turn_no++
			    fmt.Println("Turn no ",turn_no)
				fmt.Println("Welcome to the game of Pig!")
			} else{
				turn_score += turn_value
			}
			fmt.Printf("Your turn score is %d and your total score is %d\nIf you hold, you will have %d points.\n", turn_score, total_score,(turn_score+total_score))
		case "h":
			total_score += int32(turn_score)
			fmt.Println("Current score : ", total_score)
		    fmt.Printf("Your turn score is %d and your total score is %d\n",turn_score,total_score)
			turn_score = 0
			if (total_score<20){
			turn_no++
			fmt.Println("Turn no ",turn_no)
			}
		}
	}
	fmt.Println("You win !!!")
}
