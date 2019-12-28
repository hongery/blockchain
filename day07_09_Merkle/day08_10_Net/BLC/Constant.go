package BLC

const DBNAME = "blockchain_%s.db"  //数据库名
const BLOCKTABLENAME = "blocks" //表名



const PROTOCOL  = "tcp"
const COMMANDLENGTH  = 12
const NODE_VERSION  = 1


//12个字节 + 结构体序列化的字节数组


// 命令
const COMMAND_VERSION  = "version"
const COMMAND_ADDR  = "addr"
const COMMAND_BLOCK  = "block"
const COMMAND_INV  = "inv"
const COMMAND_GETBLOCKS  = "getblocks"
const COMMAND_GETDATA  = "getdata"
const COMMAND_TX  = "tx"

// 类型
const BLOCK_TYPE  = "block"
const TX_TYPE  = "tx"