pragma solidity ^0.4.24;

contract Example {
	address owner;

	constructor() {
		owner = msg.sender;
	}
}