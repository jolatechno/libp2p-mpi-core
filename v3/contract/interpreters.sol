pragma solidity  >=0.5.16 <0.7.0;

import "./utils.sol";
import "./task.sol";

contract repository is identity, random {
    interpreter[] private interpreters;
    uint256[] private sizes;

    function AddInterpreter(string memory IpfsHash, uint256 size) public sec() {
        interpreter inter = new interpreter(IpfsHash);
        interpreters.push(inter);
        sizes.push(size);
    }

    function getInterpreter(uint256 max_size) public view returns(interpreter) {
        uint256 available = 0;
        for(uint256 i = 0; i < sizes.length; i++) {
            if(sizes[i] <= max_size) {
                available++;
            }
        }

        require(available > 0, "no interpreter with small enough size");

        uint256 task_idx = rand(available) + 1; //random number

        for(uint256 i = 0; i < sizes.length; i++) {
            if(sizes[i] <= max_size) {
                task_idx--;
            }

            if(task_idx == 0) {
                return interpreters[i];
            }
        }

        revert("no interpreter with small enough size");
    }
}

contract interpreter is random, identity {
    string private ipfsHash; //interpretter folder
    address[] private owners;
    task[] private list;

    constructor(string memory IpfsHash) public {
        ipfsHash = IpfsHash;
    }

    function advertise(task Task) public {
        list.push(Task);
        owners.push(msg.sender);
    }

    function done(task Task) public onlyBy(address(Task)) {
        for(uint256 i = 0; i < list.length; i++) {
            if(list[i] == Task) {
                delete list[i];
                delete owners[i];
            }
        }
    }

    function getTask() public view returns(task Task) {
        require(list.length > 0, "no task currentrly available");

        uint256 task_idx = rand(list.length); //random number
        return list[task_idx];
    }
}
