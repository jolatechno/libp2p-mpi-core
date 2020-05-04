pragma solidity  >=0.5.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./utils.sol";

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

    function getTask() public view returns(task Task, bool) {
        if(list.length == 0) {
            return (Task, false);
        }

        uint256 task_idx = rand(list.length); //random number
        return (list[task_idx], true);
    }
}

contract stack is identity {
    mapping(bytes => bytes[]) private mem;
    mapping(bytes => address[]) private senders;

    address[] autherized;

    uint128 private safetyProportionTreshold; //the proportion is actually (safetyProportionTreshold - 1)/safetyProportionTreshold
    uint128 private safetyLengthTreshold;

    constructor(uint128 SafetyProportionTreshold, uint128 SafetyLengthTreshold) public {
        safetyProportionTreshold = SafetyProportionTreshold;
        safetyLengthTreshold = SafetyLengthTreshold;
        autherized.push(msg.sender);
    }

    function newTask(interpreter IpfsObject, uint256[] memory Kernel_size) public sec() {
        task Task = new task(this, IpfsObject, Kernel_size);
        Task.transfer(msg.sender);
        autherized.push(address(Task));
    }

    function push(address sender, bytes memory key, bytes memory value) public oneOf(autherized) {
        if(sender != owner()) {
            address[] memory addresses = senders[key];
            for(uint256 i = 0; i < addresses.length; i++) { //to prevent peers from validating themselfs
                if(sender == addresses[i]) {
                    return;
                }
            }
        }

        senders[key].push(sender);
        mem[key].push(value);
    }

    function read(bytes memory key) public view returns (bytes memory value, bool) {
        bytes[] memory values = mem[key];
        address[] memory addresses = senders[key];
        uint256 len = values.length;

        if(len < safetyLengthTreshold) {
            return (value, false);
        }

        uint256[] memory proportions = new uint256[](len);

        for(uint256 i = 0; i < len; i++) {
            if(addresses[i] == owner()) {
                return (values[i], true);
            }

            for(uint256 j = 0;  j <= i; j++) {
                if(keccak256(values[i]) == keccak256(values[j])) {
                    proportions[j]++;
                    if(proportions[j] * safetyProportionTreshold >= len * (safetyProportionTreshold - 1)) { //<=> proportions[j]/len >= (safetyProportionTreshold - 1)/safetyProportionTreshold
                        return (values[i], true);
                    }
                }
            }
        }

        return (value, false);
    }
}

contract task is random, identity {
    stack mem;
    interpreter private ipfsObject;

    uint256[] private kernel_size;
    uint256 private kernel_length;
    uint256[] private assigned;

    //event Push(address  sender, bytes  key, bytes  value);

    constructor(stack Stack, interpreter IpfsObject, uint256[] memory Kernel_size) public {
        kernel_size = Kernel_size;
        ipfsObject = IpfsObject;
        mem = Stack;

        uint256 total_size = 1;
        for(uint256 i = 0; i < Kernel_size.length; i++) {
            total_size *= Kernel_size[i];
        }
        kernel_length = total_size;

        assigned = new uint256[](kernel_length);

        ipfsObject.advertise(this);
    }

    function done() public sec() {
        ipfsObject.done(this);
    }

    function getCommand() public view returns (uint256[] memory kernel_idxs, string memory error) {
        kernel_idxs = new uint256[](kernel_size.length);
        for(uint256 i = 0; i < kernel_size.length; i++) {
            kernel_idxs[i] = rand(kernel_size[i]); //random number
        }

        return (kernel_idxs, "");
    }

    function acceptCommand(uint256[] memory kernel_idxs) public {
        uint256 kernel_idx = 0;
        uint256 acc = 1;

        for(uint256 i = 0; i < kernel_idxs.length; i++) {
            kernel_idx += acc*kernel_idxs[i];
            acc *= kernel_size[i];
        }
        assigned[kernel_idx]++; //keeping track of which kernel idx as been assigned
    }

    function push(bytes memory key, bytes memory value) public {
        mem.push(msg.sender, key, value);
    }

    function read(bytes memory key) public view returns (bytes memory value, bool) {
        return mem.read(key);
    }
}