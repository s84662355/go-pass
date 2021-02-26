package tcp

import "time"

const (
    ReadTimeout = time.Minute * 5 //当服务器1分钟内没有收到任何数据，断开客户端连接
)

/*

//数据包结构
const (
    TypeLen       = 2                         // 消息类型字节数组长度
    LenLen        = 2                         // 消息长度字节数组长度
    SeqLen        = 4                         // 消息seq字节数组长度
    ContentMaxLen = 1024 * 16                 // 消息体最大长度
    HeadLen       = TypeLen + LenLen + SeqLen // 消息头部字节数组长度（消息类型字节数组长度+消息长度字节数组长度）
    BufLen        = ContentMaxLen + HeadLen   // 缓冲buffer字节数组长度
)
*/
const (
    TcpProtocolLen = 2 //uint16
    ContentMaxLen  = 1024 * 16
    BufLen         = ContentMaxLen + TcpProtocolLen
)

// 数据包类型
const (
    CodeSignIn = 1 // 设备登录

    CodeEventSend = 2 // 客户端事件发送

    CodeHeartbeat = 3 // 心跳

    CodeMessageSend = 4 // 客户端消息发送

    CodeMessagePush = 5 // 服务端消息推送

    CodeEventPush = 6 // 服务端事件推送

    CodeLogout = 7 // 客户端退出登录：用户主动要求退出，或用户相同平台设备登录状态顶替
)
