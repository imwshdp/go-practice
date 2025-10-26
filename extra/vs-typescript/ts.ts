const customTypesBehavior = () => {
	type Celsius = number;
	type Fahrenheit = number;

	const celsius: Celsius = 1;
	const fahrenheit: Fahrenheit = 1;

	console.log('is celsius equal to fahrenheit', celsius === fahrenheit);
};

const arrayComparison = () => {
	const array1: Array<number> = [1, 2, 3];
	const array2: Array<number> = [1, 2, 3];

	console.log('are number arrays equal', array1 === array2);
};

const structsArrayComparison = () => {
	type TestStruct = {
		str: string;
		num: number;
	};

	let object1: TestStruct = { str: '1', num: 1 },
		object2: TestStruct = { str: '2', num: 2 },
		object3: TestStruct = { str: '3', num: 3 };

	const array1: TestStruct[] = [object1, object2, object3];

	// even if uncomment - there is no "true" result because of different addresses of arrays
	// object3 = { str: '4', num: 4 };

	const array2: TestStruct[] = [object1, object2, object3];

	console.log('are object arrays equal', array1 === array2);
};

function main() {
	// 1. Custom types comparison
	// In JavaScript it is possible to compare custom types when they have the same foundation
	customTypesBehavior();

	// 2. Arrays comparison
	// In JavaScript arrays are compared by reference to a memory location
	arrayComparison();

	// 3. Object arrays comparison
	// Same logic, arrays are compared by reference to a memory location
	structsArrayComparison();
}

main();
