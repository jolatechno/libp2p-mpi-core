pragma solidity  >=0.5.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./utils.sol";
import "./interpreters.sol";

contract stack is identity {
    address private authorized;

    mapping(bytes => bytes[]) private mem;
    mapping(bytes => address[]) private senders;

    constructor(address Authorized) public {
        authorized = Authorized;
    }

    function push(address sender, bytes memory key, bytes memory value) public {
        if(sender != authorized) {
            address[] memory addresses = senders[key];
            for(uint256 i = 0; i < addresses.length; i++) { //to prevent peers from validating themselfs
                if(sender == addresses[i]) {
                    revert("can't validate yourself");
                }
            }
        }

        senders[key].push(sender);
        mem[key].push(value);
    }

    function read(bytes memory key) public view returns (bytes memory value) {
        bytes[] memory values = mem[key];
        address[] memory addresses = senders[key];
        uint256 len = values.length;

        require(len > 0, "value not according");

        uint256[] memory proportions = new uint256[](len);

        for(uint256 i = 0; i < len; i++) {
            if(addresses[i] == authorized) {
                return values[i];
            }

            for(uint256 j = 0;  j <= i; j++) {
                if(keccak256(values[i]) == keccak256(values[j])) {
                    proportions[j]++;
                    if(proportions[j] > len / 2) { //* safetyProportionTreshold >= len * (safetyProportionTreshold - 1)) { //<=> proportions[j]/len >= (safetyProportionTreshold - 1)/safetyProportionTreshold
                        return values[i];
                    }
                }
            }
        }

        revert("value not according");
    }
}

contract task is random, identity, stake {
    uint256 bid;
    uint256 reward;

    uint256 verify_prob;
    uint256 verify_pool;

    interpreter ipfs_object;
    stack Stack;

    uint256[][] private task_ids;

    //event Push(address  sender, bytes  key, bytes  value);

    constructor(uint256 Bid, uint256 Reward,
    interpreter Ipfs_object, uint256[][] memory Task_ids,
    uint256 Verify_prob, uint256 Verify_pool) public {
        require(Verify_prob * 2 <= Verify_pool, "the verify probability is two high in comparaison to the verify pool size");

        bid = Bid;
        reward = Reward;

        task_ids = Task_ids;
        verify_prob = Verify_prob;
        verify_pool = Verify_pool;

        ipfs_object = Ipfs_object;
        ipfs_object.advertise(this);

        Stack = new stack(msg.sender);
        transfer(address(this));
    }

    function done() public onlyBy(address(this)) {
        ipfs_object.done(this);
        complete(reward);
    }

    function getCommand() public view returns (uint256 task_id) {
        require(task_ids.length > 0, "No task to return");

        uint256 task_id_id = rand(task_ids[0].length);
        return task_ids[0][task_id_id];
    }

    function finish(uint256 task_id, bytes[] memory keys, bytes[] memory values) public payable cost(bid) {
        require(task_ids.length > 0, "already done");
        require(keys.length == values.length, "keys and values length don't match");

        bool pushed = false;
        for(uint256 i = 0; i < task_ids.length; i++) {
            if(task_ids[0][i] == task_id) {
                delete task_ids[i];
                pushed = true;
                break;
            }
        }

        require(pushed, "no corresponding task");

        if(rand(verify_prob) == 0) {
            for(uint256 i = 0; i < verify_pool; i++) {
                task_ids[0].push(task_id);
            }
        }

        for(uint256 i = 0; i < keys.length; i++) {
            Stack.push(msg.sender, keys[i], values[i]);
        }

        if(task_ids[0].length == 0) {
            delete task_ids[0];
        }
        if(task_ids.length == 0) {
            done();
        }
    }
}