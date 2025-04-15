// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract Begging {

    // 用于记录每个捐赠者的捐赠金额
    mapping(address => uint256) public donations;

    // 合约所有者
    address public owner;

    // Run only once
    constructor() {
        //msg.sender sender which is the person calling the function
        owner = msg.sender;
    }

    // 允许用户向合约发送以太币，并记录捐赠信息  payable修饰 意味着用户可以接受ETH
    function donate() public payable {
        donations[msg.sender] += msg.value;
    }

    // 允许合约所有者提取所有资金
    function withdraw() public {
        //msg.sender ===> 方法调用者信息
        require(msg.sender == owner, "Only contract owner can withdraw");
        //address(this) corresponds to the address of the current smart contract (only owner can withdraw eht)
        payable(owner).transfer(address(this).balance);
    }

    // 允许查询某个地址的捐赠金额
    function getDonation(address donor) public view returns (uint256) {
        return donations[donor];
    }
}


//代码解释：
//
//1.** SPDX - License - Identifier**：这是指定合约的开源许可证，这里使用的是MIT许可证。
//
//2.pragma solidity ^0.8.10：指定合约所使用的Solidity编译器版本，这里要求编译器版本在0.8.10及以上，但低于0.9.0。
//
//3.contract Begging：定义了名为“Begging”的合约。
//
//4.mapping(address => uint256) public donations：创建了一个映射（mapping），键是捐赠者的地址（address类型），值是捐赠的金额（uint256类型），并且这个映射是公开的，可以通过合约外部访问。
//
//5.address public owner：定义了合约的所有者地址，并且是公开的。
//
//6.constructor：合约的构造函数，在合约部署时执行，将部署合约的地址设置为合约所有者。
//
//7.function donate() public payable：这是捐赠函数，payable修饰符表示该函数可以接收以太币。每次调用该函数时，将发送者的捐赠金额累加到其对应的记录中。
//
//8.function withdraw() public：提取资金的函数，只有合约所有者可以调用。require语句用于检查调用者是否为合约所有者，然后将合约中的所有以太币转账给合约所有者。
//
//9.function getDonation(address donor) public view returns (uint256)：查询函数，接收一个地址作为参数，返回该地址对应的捐赠金额。view修饰符表示该函数不会修改区块链状态，只是读取状态。
