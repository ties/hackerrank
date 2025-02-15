package main
/**
It's New Year's Day and everyone's in line for the Wonderland rollercoaster
ride! There are a number of people queued up, and each person wears a sticker
indicating their initial position in the queue. Initial positions increment by 
from  at the front of the line to  at the back.

Any person in the queue can bribe the person directly in front of them to swap
positions. If two people swap positions, they still wear the same sticker
denoting their original places in line. One person can bribe at most two others.
For example, if $n=8$ and $Person 5$ bribes $Person 4$, the queue will look like
this: $[1,2,3,5,4,6,7,8]$.
*/
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
    /**
    Naïeve approach: Number of bribes needed is total number of positions
    people shifted in front of their original position.

    Does not work: 
    [1,2,3] => [3,2,1]
    [1,2,3] => [3,1,2] => [3,2,1]

    Three bribes while only one number moved by two.

    Alternative:
    |moved forward| + |moved backward| => does not work either

    Alternative two:
      * Check for > 2 moves.
      * Perform bubble sort, every swap is a bribe.
    */
    // Check for > 2 moves.
    for key, value := range q {
        moved := int(value) - (key + 1)
        if moved > 2 {
            fmt.Println("Too chaotic")
            return
        }
    }

    var bribes = 0
    var swapped = true

    // Bubble sort
    // Potential improvement: store the bounds of the first and last inversion
    // instead of iterating over whole list.
    for swapped {
        swapped = false
        // Last i elements are already sorted in place
        for j := 0; j < len(q) - 1; j++ {
            if q[j] > q[j+1] {
                bribes += 1
                tmp := q[j]
                q[j] = q[j+1]
                q[j+1] = tmp
                // Indicate this iteration swapped items
                swapped = true
            }
        }
    }

    fmt.Printf("%d\n", bribes)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        n := int32(nTemp)

        qTemp := strings.Split(readLine(reader), " ")

        var q []int32

        for i := 0; i < int(n); i++ {
            qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
            checkError(err)
            qItem := int32(qItemTemp)
            q = append(q, qItem)
        }

        minimumBribes(q)
    }
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
