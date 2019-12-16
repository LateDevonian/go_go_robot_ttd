package main

import (
	"bufio"
	"fmt"

	//"io/ioutil"
	"log"
	"os"
)

// PLACE X,Y,F
// MOVE
// LEFT
// RIGHT
// REPORT

func main() {
	//get instructions from file
	file, err := os.Open("testdata/instructions.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		
		lineText := scanner.Text()
		DoAThingWhereICHECKTHIS(lineText, &robot, tableTop);
		}
	} else {
		fmt.Printf("missing file name \n")
}

	}
}
