pragma solidity ^0.5.16;

contract interpreter { //will be implemented later
    string[] public cmds; //hand written program

    function execute(task t) public returns (string memory error) { //will be implemented later
        t.push("dummy", "dummy");
        return error;
    }
}

contract task {
    mapping(bytes => bytes[]) public stack;
    interpreter public prgm;
    int public safetyTreshold;

    function push(bytes memory key, bytes memory value) public {
        stack[key].push(value);
    }

    function read(bytes memory key) public view returns (bytes memory value, bool) {
        bytes[] memory values = stack[key];
        uint256 len = values.length;
        
        if(len == 0) {
            return (value, false);
        }

        value = values[0];
        for(uint256 i = 1; i < len; i++) {
            if(keccak256(values[i]) != keccak256(value)) { //should check if the occurence of values[i] rather then check if all value match
                return (value, false);
            }
        }

        return (value, true);
    }
}