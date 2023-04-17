// @Title
// @Description
// @Author
// @Update

package builtins

import (
	"bufio"
	"fmt"
	"os"
)

/*

Input / output

 - Input

*/

/*
Args

  - message string

  - The message to be displayed prior to recieving the input.

    e.g. "Enter Name: ", or ">> "

Return

  - input string
  - error
*/
func Input(message string) (string, error) {

	var res string

	fmt.Print(message)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	res = scanner.Text()

	return res, scanner.Err()

}
