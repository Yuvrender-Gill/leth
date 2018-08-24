pragma solidity ^0.4.24;

contract Example {
	address owner;

	mapping(address => uint) balance;

	event Deposit(address sender, uint value);

	constructor() {
		owner = msg.sender;
	}

	function () public payable {
		balance[msg.sender] = msg.value;
		emit Deposit(msg.sender, msg.value);
	}
}