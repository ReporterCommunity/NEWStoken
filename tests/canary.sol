pragma solidity ^0.4.11;

contract Canary {

    function timeStamp() constant returns (uint256) {
        return now;
    }
}