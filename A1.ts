/**
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-12-04
 * @fileoverview Prompts the sum of integers that are in range and out of range for a given range.
 */


// declare variables
let numberNumber: number;
let inRange: number = 0;
let outRange: number = 0;


// asking for range
const startString: string = prompt("Enter an integer for the start of the range: ") || ("Invalid");
const startNumber: number = parseInt(startString);

const endString: string = prompt("Enter an integer for the start of the range: ") || ("Invalid");
const endNumber: number = parseInt(endString);

// do while loop
do {

  const numberString: string = prompt("Enter an integer or zero(0) to end: ") || ("Invalid") // asks for integer until 0 is pressed.
  numberNumber = parseInt(numberString);

  if (numberNumber >= startNumber && numberNumber <= endNumber) { // adds to total sum
    inRange += numberNumber;
  } else {
    outRange += numberNumber;
  }

} while (numberNumber !== 0); // print out sums
  console.log(`The sum of the integers inside the range is ${inRange}`);
  console.log(`The sum of the integers outside the range is ${outRange}`);

  console.log("\nDone.");
