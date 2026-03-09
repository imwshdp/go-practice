export type TestStruct = {
	str: string;
	num: number;
};

export const formatArr = (array: TestStruct[]) =>
	JSON.stringify(array.map((o) => [o.str, o.num]));

export const formatObj = (obj: TestStruct) =>
	JSON.stringify([obj.str, obj.num]);

const customTypesBehavior = () => {
	type Celsius = number;
	type Fahrenheit = number;

	const celsius: Celsius = 1;
	const fahrenheit: Fahrenheit = 1;

	console.log(
		`1) is celsius equal to fahrenheit? Answer is: ${celsius === fahrenheit}\n`,
	);
};

const arrayComparison = () => {
	const array1: Array<number> = [1, 2, 3];
	const array2: Array<number> = [1, 2, 3];

	console.log(
		`2) is array1 (${array1}) equal to array2 (${array2})? Answer is: ${array1 === array2}\n`,
	);
};

const structsArrayComparison = () => {
	const object1: TestStruct = { str: "1", num: 1 },
		object2: TestStruct = { str: "2", num: 2 },
		object3: TestStruct = { str: "3", num: 3 };

	const array1: TestStruct[] = [object1, object2, object3];

	// Even if object3 is changed, the result will still be false because arrays are compared by reference, not by value
	// object3 = { str: "4", num: 4 };

	const array2: TestStruct[] = [object1, object2, object3];

	console.log(
		`3) is array1 (${formatArr(array1)}) equal to array2 (${formatArr(array2)})? Answer is: ${array1 === array2}\n`,
	);
};

const forOfLoop = () => {
	// The for-of loop creates a new variable that is a copy of each array element
	// Modifying this variable does not change the original array
	console.log("4.1) Changing array value using loop variable in for-of");
	const myArray = [1, 2, 3, 4, 5];
	for (let element of myArray) {
		element *= 2;
		console.log(element);
	}
	console.log("result:", myArray);

	// With an array of objects, it works differently because objects are passed by reference
	console.log(
		"\n4.2) Changing array object values using loop variable in for-of",
	);
	const myObjectArray = [
		{ str: "1", num: 1 },
		{ str: "2", num: 2 },
		{ str: "3", num: 3 },
	];
	for (const element of myObjectArray) {
		element.num += 1;
		console.log(element);
	}
	console.log("result:", formatArr(myObjectArray));

	// The for-of loop does not re-evaluate the iterable on each iteration
	// This example would create an endless loop if uncommented
	console.log("\n4.3) Adding new elements to collection with for-of");
	for (const element of myArray) {
		// myArray.push(element * 10);
	}
	console.log("result:", myArray);

	// JavaScript passes objects by reference, so the loop always accesses the current value
	console.log("\n4.4) When for-of value variable is recreating");
	const reduced: Record<number, TestStruct> = {};
	for (const [index, value] of myObjectArray.entries()) {
		reduced[index] = value;
	}
	for (const [key, value] of Object.entries(reduced)) {
		console.log(`${key} - ${formatObj(value)}`);
	}

	// The reduced object stores references to objects from the original array
	// That's why changes to the array are reflected in the reduced object
	console.log("\n4.5) Changing values in the original array");
	myObjectArray[0].num = 777;
	myObjectArray[0].str = "777";
	for (const [key, value] of Object.entries(reduced)) {
		console.log(`${key} - ${formatObj(value)}`);
	}
};

function main() {
	// 1. Custom types comparison
	// In JavaScript, custom types based on the same primitive type can be compared directly
	customTypesBehavior();

	// 2. Arrays comparison
	// In JavaScript, arrays are compared by reference to their memory location
	arrayComparison();

	// 3. Object arrays comparison
	// Same logic: arrays are compared by reference, not by value
	structsArrayComparison();

	// 4. For-of loop nuances
	forOfLoop();
}

main();
