pragma solidity  >=0.5.16 <0.7.0;

contract identity {
    address private Owner = msg.sender;

    modifier sec() {
        require(
            msg.sender == Owner,
            "sender not authorized"
        );
        _;
    }

    modifier onlyBy(address account) {
        require(
            msg.sender == account,
            "sender not authorized"
        );
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