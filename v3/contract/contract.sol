pragma solidity  >=0.5.16 <0.7.0;
pragma experimental ABIEncoderV2;

contract random {
    function rand(uint256 range) internal view returns(uint256) {
        uint256 seed = uint256(keccak256(abi.encodePacked(
            block.timestamp + block.number +
            uint256(keccak256(abi.encodePacked(block.coinbase))) / block.timestamp +
            uint256(keccak256(abi.encodePacked(msg.sender))) / block.timestamp
        )));

        return seed % range;
    }
}

contract interpreter is random {
    string private ipfsHash; //interpretter folder
    uint256 oppen;

    address[] private owner;
    task[] private list;
    bool[] private Done;

    constructor(string memory IpfsHash) public {
        ipfsHash = IpfsHash;
    }

    function advertise(task Task) public {
        Done.push(false);
        list.push(Task);
        owner.push(msg.sender);
        oppen++;
    }

    function done(task Task) public {
        for(uint256 i = 0; i < list.length; i++) {
            if(list[i] == Task) {
                if(owner[i] == msg.sender && !Done[i]) {
                    Done[i] = true;
                    oppen--;
                }
                return;
            }
        }
    }

    function getTask() public view returns(task Task, bool) {
        if(oppen == 0) {
            return (Task, false);
        }

        uint256 task_idx = rand(oppen); //random number

        for(uint256 i = 0; i < list.length; i++) {
            if(task_idx == 0 && !Done[i]) {
                return (list[i], true);
            }

            if(!Done[i]) {
                task_idx--;
            }
        }

        return (Task, false);
    }
}

contract stack {
    address private owner;
    mapping(bytes => bytes[]) private mem;
    mapping(bytes => address[]) private senders;

    uint128 private safetyProportionTreshold; //the proportion is actually (safetyProportionTreshold - 1)/safetyProportionTreshold
    uint128 private safetyLengthTreshold;

    constructor(uint128 SafetyProportionTreshold, uint128 SafetyLengthTreshold) public {
        owner = msg.sender;
        safetyProportionTreshold = SafetyProportionTreshold;
        safetyLengthTreshold = SafetyLengthTreshold;
    }

    function push(address sender, bytes memory key, bytes memory value) public {
        require(msg.sender == owner, "not owner !");

        if(sender != owner) {
            address[] memory addresses = senders[key];
            for(uint256 i = 0; i < addresses.length; i++) { //to prevent peers from validating themselfs
                if(msg.sender == addresses[i]) {
                    return;
                }
            }
        }

        senders[key].push(msg.sender);
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
            if(addresses[i] == owner) {
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

contract task is random {
    address private owner;

    stack mem;
    interpreter private ipfsObject;

    uint256[] private kernel_size;
    uint256 private kernel_length;
    uint256[] private assigned;

    event Push(address  sender, bytes  key, bytes  value);

    constructor(stack Stack, interpreter IpfsObject, uint256[] memory Kernel_size) public {
        kernel_size = Kernel_size;
        owner = msg.sender;
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

    function done() public {
        require(msg.sender == owner, "not owner !");

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
        emit Push(msg.sender, key, value);
    }

    function read(bytes memory key) public view returns (bytes memory value, bool) {
        return mem.read(key);
    }
}