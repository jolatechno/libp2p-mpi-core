pragma solidity  >=0.5.16 <0.7.0;
pragma experimental ABIEncoderV2;

contract interpreter {
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
        if(list.length == 0) {
            return (Task, false);
        }

        uint256 task_idk = 0;

        //random selection will be implemented later

        return (list[task_idk], true);
    }
}

contract task {
    address private owner;
    mapping(bytes => bytes[]) private stack;
    mapping(bytes => address[]) private senders;

    interpreter private ipfsObject;

    uint128 private safetyProportionTreshold; //the proportion is actually (safetyProportionTreshold - 1)/safetyProportionTreshold
    uint128 private safetyLengthTreshold;

    uint256[] private kernel_size;
    uint256 private kernel_length;
    uint256[] private assigned;

    constructor(interpreter IpfsObject, uint256[] memory Kernel_size,
    bytes[] memory stackAddresses, bytes[] memory stackValue,
    uint128 SafetyProportionTreshold, uint128 SafetyLengthTreshold) public {
        safetyProportionTreshold = SafetyProportionTreshold;
        safetyLengthTreshold = SafetyLengthTreshold;
        kernel_size = Kernel_size;
        owner = msg.sender;
        ipfsObject = IpfsObject;

        uint256 total_size = 1;
        for(uint256 i = 0; i < Kernel_size.length; i++) {
            total_size *= Kernel_size[i];
        }
        kernel_length = total_size;

        assigned = new uint256[](kernel_length);

        if(stackAddresses.length == stackValue.length) {
            for(uint256 i = 0; i < stackAddresses.length; i++) {
                stack[stackAddresses[i]].push(stackValue[i]);
                senders[stackAddresses[i]].push(msg.sender);
            }
        }
    }

    function advertise() public {
        if(msg.sender != owner) {
            return;
        }
        ipfsObject.advertise(this);
    }

    function done() public {
        if(msg.sender != owner) {
            return;
        }
        ipfsObject.done(this);
    }

    function getCommand() public view returns (uint256[] memory kernel_idxs, string memory error) {
        uint256[] memory size = kernel_size;
        uint256 kernel_idx = 0;
        kernel_idxs = new uint256[](kernel_size.length);

        //random selection will be implemented later

        for(uint256 i = 0; i < size.length; i++) {
            kernel_idxs[i] = kernel_idx % size[i];
            kernel_idx = kernel_idx / size[i];
        }

        return (kernel_idxs, "not yet implemented");
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
        if(msg.sender != owner) {
            address[] memory addresses = senders[key];
            for(uint256 i = 0; i < addresses.length; i++) { //to prevent peers from validating themselfs
                if(msg.sender == addresses[i]) {
                    return;
                }
            }
        }

        senders[key].push(msg.sender);
        stack[key].push(value);
    }

    function read(bytes memory key) public view returns (bytes memory value, bool) {
        bytes[] memory values = stack[key];
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
                    if(proportions[j] * safetyProportionTreshold >=  len * (safetyProportionTreshold - 1)) { //<=> proportions[j]/len >= (safetyProportionTreshold - 1)/safetyProportionTreshold
                        return (values[i], true);
                    }
                }
            }
        }

        return (value, false);
    }
}