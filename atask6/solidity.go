package atask6

func main() {

}

//solidity学习

//1.变量 熟悉

//2.函数 说明

/*

setNum : 函数名称
setNum(xx): 参数类型
memory:
public : 可见性修饰符
view : 状态可变性修饰符
returns (int) 返回类型
{}: 函数体


function setNum(int memory a, int memory b) public view returns (int){
	return a + b ;
}
*/
//EVM 交易执行流程
//1. 用户发起交易
//|
//↓
//2. 创建新的 EVM 实例
//|
//↓
//3. 加载合约字节码
//|
//↓
//4. 分配 Stack 空间（1024 slots）
//|
//↓
//5. 执行合约代码
//|
//↓
//6. 更新区块链状态（如果需要）
//|
//↓
//7. 销毁 EVM 实例

//1.用户发起交易
//用户签名交易
//设定 gas limit 和 gas price
//指定目标合约地址和调用数据
//2.创建新的 EVM 实例
//为每笔交易创建独立的 EVM 环境
//初始化执行上下文
//准备内存和存储空间
//3.加载合约字节码
//从区块链状态中读取合约字节码
//将字节码加载到 EVM 中
//准备执行环境
//4.分配 Stack 空间
//分配 1024 个 slots
//每个 slot 256 位
//用于存储临时计算结果
//5.执行合约代码
//逐条执行操作码
//进行状态检查和 gas 计算
//处理函数调用和返回值
//6.更新区块链状态
//写入存储变更
//更新账户余额
//触发事件日志
//7.销毁 EVM 实例
//清理内存
//释放资源
//返回执行结果
//注意事项
//
//整个过程是原子性的：要么全部成功，要么全部失败
//Gas 限制贯穿整个执行过程
//状态变更只在交易成功执行后才会提交

//在 Solidity 中，直接存储在栈（Stack）中的基本类型包括：  https://github.com/Base1-Go/solidity/blob/main/1-solidity-base/1.2%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B/03%E5%87%BD%E6%95%B0%E4%BF%AE%E9%A5%B0%E7%AC%A6.md
//1. 整型（Integer）
//2. 布尔型（Boolean）
//3. 地址（Address）
//4. 固定大小字节数组（Fixed-size Bytes）
//5. 枚举（Enum）

//EVM 存储结构
//
//1. Stack (栈)
//+----------------+
//| 基本类型的值    | <- 函数内的局部变量
//| 引用的地址      | <- 指向其他存储位置
//+----------------+
//
//
//2. Memory (内存)
//+----------------+
//| 临时数据       | <- 函数执行期间的临时数据
//| 函数参数       | <- memory 类型的参数
//| 返回数据       | <- 函数返回值
//+----------------+
//
//
//3. Storage (存储)
//+----------------+
//| 状态变量       | <- 合约的永久存储数据
//| 映射数据       | <- mapping 数据
//| 数组数据       | <- storage 数组
//+----------------+
//
//EVM Stack (固定大小的栈空间)
//+------------------+
//|     空闲空间     | <- 1024 个槽位（slots）
//|        ⬇        |    每个槽位 32 字节（256 位）
//+------------------+
//|    当前使用空间   | <- 随函数执行压入/弹出
//+------------------+
//
//Memory： 存在于 EVM 执行环境中 临时性的，交易执行完就清除 线性寻址（0x00, 0x20, 0x40...）
//
//Storage： 存在于区块链状态中 永久性的，写入区块 使用 slot 和 keccak256 哈希定位

//基本类型 & 引用类型
//contract TypeComparison {
//// 基本类型：直接存储在栈中
//uint256 public stackVar = 123;
//
//// 引用类型：存储引用在栈中，数据在其他位置
//string public stringVar = "hello";  // 数据在存储中
//uint256[] public arrayVar;          // 数据在存储中
//}

//栈中存储
//contract StackUsage {
//function calculate() public pure returns (uint256) {
//// 这些变量都在栈中
//uint256 a = 1;
//uint256 b = 2;
//uint256 c = a + b;
//
//return c;
//} // 函数结束时栈变量自动清除
//}

//全局变量
//1.1.3 全局变量
//因为区块链的特性，solidity 语言有一类变量叫做全局变量，它们是全局工作区中存在的特殊变量，提供有关区块链和交易属性的信息。下面列举几个常用的全局变量：
//
//blockhash(uint blockNumber) returns (bytes32) 给定区块的哈希值 – 只适用于 256 最近区块, 不包含当前区块
//
//block.coinbase (address payable) 当前区块矿工的地址
//
//block.difficulty (uint) 当前区块的难度
//
//block.gaslimit (uint) 当前区块的 gaslimit
//
//block.number (uint) 当前区块的 number
//
//block.timestamp (uint) 当前区块的时间戳，为 unix 纪元以来的秒
//
//gasleft() returns (uint256) 剩余 gas
//
//msg.data (bytes calldata) 完成 calldata
//
//msg.sender (address payable) 消息发送者 (当前 caller)
//
//msg.sig (bytes4) calldata 的前四个字节 (function identifier)
//
//msg.value (uint) 当前消息的 wei 值
//
//now (uint) 当前块的时间戳
//
//tx.gasprice (uint) 交易的 gas 价格
//
//tx.origin (address payable) 交易的发送方
