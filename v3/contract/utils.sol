pragma solidity  >=0.5.16 <0.7.0;

contract identity {
    address private Owner = msg.sender;

    modifier sec() {
        require(msg.sender == Owner, "sender not authorized");
        _;
    }

    modifier onlyBy(address account) {
        require(msg.sender == account, "sender not authorized");
        _;
    }

    modifier oneOf(address[] storage autherized) {
        bool ok = false;
        for(uint256 i = 0; i < autherized.length; i++) {
            if(autherized[i] == msg.sender) {
                ok = true;
                break;
            }
        }
        require(ok, "sender not authorized");
        _;
    }

    function owner() public view returns(address) {
        return Owner;
    }

    function transfer(address new_owner) public sec() {
        Owner = new_owner;
    }
}

contract stake is identity {
    address[] private payer;
    uint256[] private balance;

    modifier cost(uint256 value) {
        require(msg.value == value, "value is not right");
        payer.push(msg.sender);
        balance.push(msg.value);
        _;
    }

    modifier costAtLeast(uint256 value) {
        require(msg.value >= value, "value is not enough");
        payer.push(msg.sender);
        balance.push(msg.value);
        _;
    }

    function complete(uint256 reward) public payable {
        for(uint256 i = payer.length - 1; i >= 0; i--) {
            if(payer[i] == msg.sender) {
                msg.sender.transfer(balance[i] + reward);
                delete_idx(i);
            }
        }
    }

    function refund() public payable {
        for(uint256 i = payer.length - 1; i >= 0; i--) {
            if(payer[i] == msg.sender) {
                msg.sender.transfer(balance[i]);
                delete_idx(i);
            }
        }
    }

    function delete_user(address user) public sec() {
        for(uint256 i = payer.length - 1; i >= 0; i--) {
            if(payer[i] == user) {
                delete_idx(i);
            }
        }
    }

    function delete_idx(uint256 i) public sec() {
        delete payer[i];
        delete balance[i];
    }
}

contract random {
    function rand(uint256 range) internal view returns(uint256) {
        uint256 seed = uint256(keccak256(abi.encodePacked(
            block.number + uint256(keccak256(abi.encodePacked(msg.sender))) / block.number
        )));

        return seed % range;
    }
}