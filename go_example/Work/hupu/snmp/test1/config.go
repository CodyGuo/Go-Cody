package main

const (
	ABILITY_DEVALL_INFO    = 26
	DH_SERIALNO_LEN        = 48
	DH_MAX_SNMP_COMMON_LEN = 64 // snmp 读写数据长度
	DH_DEV_SNMP_CFG        = 0X005f
	DH_DEV_TIMECFG         = 0x0008
)

type EM_LOGIN_SPAC_CAP_TYPE int

const (
	EM_LOGIN_SPEC_CAP_TCP            EM_LOGIN_SPAC_CAP_TYPE = iota // TCP登陆, 默认方式
	EM_LOGIN_SPEC_CAP_ANY                                          // 无条件登陆
	EM_LOGIN_SPEC_CAP_SERVER_CONN                                  // 主动注册的登入
	EM_LOGIN_SPEC_CAP_MULTICAST                                    // 组播登陆
	EM_LOGIN_SPEC_CAP_UDP                                          // UDP方式下的登入
	EM_LOGIN_SPEC_CAP_MAIN_CONN_ONLY                               // 只建主连接下的登入
	EM_LOGIN_SPEC_CAP_SSL                                          // SSL加密方式登陆

	EM_LOGIN_SPEC_CAP_INTELLIGENT_BOX EM_LOGIN_SPAC_CAP_TYPE = iota + 9 // 登录智能盒远程设备
	EM_LOGIN_SPEC_CAP_NO_CONFIG                                         // 登录设备后不做取配置操作
	EM_LOGIN_SPEC_CAP_U_LOGIN                                           // 用U盾设备的登入
	EM_LOGIN_SPEC_CAP_LDAP                                              // LDAP方式登录
	EM_LOGIN_SPEC_CAP_AD                                                // AD（ActiveDirectory）登录方式
	EM_LOGIN_SPEC_CAP_RADIUS                                            // Radius 登录方式
	EM_LOGIN_SPEC_CAP_SOCKET_5                                          // Socks5登陆方式
	EM_LOGIN_SPEC_CAP_CLOUD                                             // 云登陆方式
	EM_LOGIN_SPEC_CAP_AUTH_TWICE                                        // 二次鉴权登陆方式
	EM_LOGIN_SPEC_CAP_TS                                                // TS码流客户端登陆方式
	EM_LOGIN_SPEC_CAP_P2P                                               // 为P2P登陆方式
	EM_LOGIN_SPEC_CAP_MOBILE                                            // 手机客户端登陆
	EM_LOGIN_SPEC_CAP_INVALID                                           // 无效的登陆方式
)

const (
	EN_FTP                         = iota // FTP 按位,1：传送录像文件 2：传送抓图文件
	EN_SMTP                               // SMTP 按位,1：报警传送文本邮件 2：报警传送图片 3:支持健康邮件功能
	EN_NTP                                // NTP  按位：1：调整系统时间
	EN_AUTO_MAINTAIN                      // 自动维护 按位：1：重启 2：关闭 3:删除文件
	EN_VIDEO_COVER                        // 区域遮挡 按位：1：多区域遮挡
	EN_AUTO_REGISTER                      // 主动注册 按位：1：注册后sdk主动登陆
	EN_DHCP                               // DHCP 按位：1：DHCP
	EN_UPNP                               // UPNP 按位：1：UPNP
	EN_COMM_SNIFFER                       // 串口抓包 按位：1:CommATM
	EN_NET_SNIFFER                        // 网络抓包 按位： 1：NetSniffer
	EN_BURN                               // 刻录功能 按位：1：查询刻录状态
	EN_VIDEO_MATRIX                       // 视频矩阵 按位：1：是否支持视频矩阵 2:是否支持SPOT视频矩阵 3:是否支持HDMI视频矩阵
	EN_AUDIO_DETECT                       // 音频检测 按位：1：是否支持音频检测
	EN_STORAGE_STATION                    // 存储位置 按位：1：Ftp服务器(Ips) 2：SMB 3：NFS 4：ISCSI 16：DISK 17：U盘
	EN_IPSSEARCH                          // IPS存储查询 按位：1：IPS存储查询
	EN_SNAP                               // 抓图  按位：1：分辨率2：帧率3：抓图方式4：抓图文件格式5：图画质量
	EN_DEFAULTNIC                         // 支持默认网卡查询 按位 1：支持
	EN_SHOWQUALITY                        // CBR模式下显示画质配置项 按位 1:支持
	EN_CONFIG_IMEXPORT                    // 配置导入导出功能能力 按位 1:支持
	EN_LOG                                // 是否支持分页方式的日志查询 按位 1：支持
	EN_SCHEDULE                           // 录像设置的一些能力 按位 1:冗余 2:预录 3:录像时间段
	EN_NETWORK_TYPE                       // 网络类型按位表示 1:以态网 2:无线局域 3:CDMA/GPRS 4:CDMA/GPRS多网卡配置
	EN_MARK_IMPORTANTRECORD               // 标识重要录像 按位:1：设置重要录像标识
	EN_ACFCONTROL                         // 活动帧率控制 按位：1：支持活动帧率控制, 2:支持定时报警类型活动帧率控制(不支持动检),该能力与ACF能力互斥
	EN_MULTIASSIOPTION                    // 多路辅码流 按位：1：支持三路辅码流, 2:支持辅码流编码压缩格式独立设置
	EN_DAVINCIMODULE                      // 组件化模块 按位：1,时间表分开处理 2:标准I帧间隔设置
	EN_GPS                                // GPS功能 按位：1：Gps定位功能
	EN_MULTIETHERNET                      // 支持多网卡查询 按位 1：支持
	EN_LOGIN_ATTRIBUTE                    // Login属性 按位：1：支持Login属性查询
	EN_RECORD_GENERAL                     // 录像相关 按位：1,普通录像；2：报警录像；3：动态检测录像；4：本地存储；5：远程存储；6：冗余存储；7：本地紧急存储；8：支持区分主辅码流的远程存储
	EN_JSON_CONFIG                        // Json格式配置:按位：1支持Json格式, 2: 使用F6的NAS配置, 3: 使用F6的Raid配置, 4：使用F6的MotionDetect配置, 5：完整支持三代配置(V3),通过F6命令访问
	EN_HIDE_FUNCTION                      // 屏蔽功能：按位：1,屏蔽PTZ云台功能, 2: 屏蔽3G的保活时段功能
	EN_DISK_DAMAGE                        // 硬盘坏道信息支持能力: 按位：1,硬盘坏道信息查询支持能力
	EN_PLAYBACK_SPEED_CTRL                // 支持回放网传速度控制:按位:1,支持回放加速
	EN_HOLIDAYSCHEDULE                    // 支持假期时间段配置:按位:1,支持假期时间段配置
	EN_FETCH_MONEY_TIMEOUT                // ATM取钱超时
	EN_BACKUP_VIDEO_FORMAT                // 备份支持的格式,按位：1:DAV, 2:ASF
	EN_QUERY_DISK_TYPE                    // 支持硬盘类型查询
	EN_CONFIG_DISPLAY_OUTPUT              // 支持设备显示输出（VGA等）配置,按位: 1:画面分割轮巡配置
	EN_SUBBITRATE_RECORD_CTRL             // 支持扩展码流录像控制设置, 按位：1-辅码流录像控制设置
	EN_IPV6                               // 支持IPV6设置, 按位：1-IPV6配置
	EN_SNMP                               // SNMP（简单网络管理协议）
	EN_QUERY_URL                          // 支持获取设备URL地址, 按位: 1-查询配置URL地址
	EN_ISCSI                              // ISCSI（Internet小型计算机系统接口配置）
	EN_RAID                               // 支持Raid功能
	EN_HARDDISK_INFO                      // 支持磁盘信息F5查询
	EN_PICINPIC                           // 支持画中画功能 按位:1,画中画设置; 2,画中画预览、录像存储、查询、下载;3,支持画中画编码配置,同时支持画中画通道查询
	EN_PLAYBACK_SPEED_CTRL_SUPPORT        // 同 EN_PLAYBACK_SPEED_CTRL ,只为了兼容协议
	EN_LF_XDEV                            // 支持24、32、64路LF-X系列,标注这类设备特殊的编码能力计算方式
	EN_DSP_ENCODE_CAP                     // F5 DSP编码能力查询
	EN_MULTICAST                          // 组播能力查询
	EM_NET_LIMIT                          // 网络限制能力查询,按位,1-网络发送码流大小限
	EM_COM422                             // 串口422
	EM_PROTOCAL_FRAMEWORK                 // 是否支持三代协议框架（需要实现listMethod(),listService()）,通过F6命令访问
	EM_WRITE_DISK_OSD                     // 刻录OSD叠加, 按位, 1-刻录OSD叠加配置
	EM_DYNAMIC_MULTI_CONNECT              // 动态多连接, 按位, 1-请求视频数据应答
	EM_CLOUDSERVICE                       // 云服务,按位,1-支持私有云服务
	EM_RECORD_INFO                        // 录像信息上报, 按位, 1-录像信息主动上报, 2-支持录像帧数查询
	EN_DYNAMIC_REG                        // 主动注册能力,按位,1-支持动态主动注册
	EM_MULTI_PLAYBACK                     // 多通道预览回放,按为,1-支持多通道预览回放
	EN_ENCODE_CHN                         // 编码通道, 按位, 1-支持纯音频通道, 2-监视支持音视频分开获取
	EN_SEARCH_RECORD                      // 录像查询, 按位, 1-支持异步查询录像, 2-支持三代协议查询录像
	EN_UPDATE_MD5                         // 支持升级文件传输完成后做MD5验证,1-支持MD5验证
)

var (
	names = []string{
		EN_FTP:                         "FTP",
		EN_SMTP:                        "SMTP",
		EN_NTP:                         "NTP",
		EN_AUTO_MAINTAIN:               "自动维护",
		EN_VIDEO_COVER:                 "区域遮挡",
		EN_AUTO_REGISTER:               "主动注册",
		EN_DHCP:                        "DHCP",
		EN_UPNP:                        "UPNP",
		EN_COMM_SNIFFER:                "串口抓包",
		EN_NET_SNIFFER:                 "网络抓包",
		EN_BURN:                        "刻录功能 ",
		EN_VIDEO_MATRIX:                "视频矩阵",
		EN_AUDIO_DETECT:                "音频检测",
		EN_STORAGE_STATION:             "存储位置",
		EN_IPSSEARCH:                   "IPS存储查询",
		EN_SNAP:                        "抓图",
		EN_DEFAULTNIC:                  "支持默认网卡查询",
		EN_SHOWQUALITY:                 "CBR模式下显示画质配置项",
		EN_CONFIG_IMEXPORT:             "配置导入导出功能能力",
		EN_LOG:                         "是否支持分页方式的日志查询",
		EN_SCHEDULE:                    "录像设置的一些能力",
		EN_NETWORK_TYPE:                "网络类型按位表示",
		EN_MARK_IMPORTANTRECORD:        "标识重要录像",
		EN_ACFCONTROL:                  "活动帧率控制",
		EN_MULTIASSIOPTION:             "多路辅码流",
		EN_DAVINCIMODULE:               "组件化模块",
		EN_GPS:                         "GPS功能",
		EN_MULTIETHERNET:               "支持多网卡查询",
		EN_LOGIN_ATTRIBUTE:             "Login属性",
		EN_RECORD_GENERAL:              "录像相关",
		EN_JSON_CONFIG:                 "Json格式配置",
		EN_HIDE_FUNCTION:               "屏蔽功能",
		EN_DISK_DAMAGE:                 "硬盘坏道信息支持能力",
		EN_PLAYBACK_SPEED_CTRL:         "支持回放网传速度控制",
		EN_HOLIDAYSCHEDULE:             "支持假期时间段配置",
		EN_FETCH_MONEY_TIMEOUT:         "ATM取钱超时",
		EN_BACKUP_VIDEO_FORMAT:         "备份支持的格式",
		EN_QUERY_DISK_TYPE:             "支持硬盘类型查询",
		EN_CONFIG_DISPLAY_OUTPUT:       "支持设备显示输出（VGA等）配置",
		EN_SUBBITRATE_RECORD_CTRL:      "支持扩展码流录像控制设置",
		EN_IPV6:                        "支持IPV6设置",
		EN_SNMP:                        "SNMP（简单网络管理协议）",
		EN_QUERY_URL:                   "支持获取设备URL地址",
		EN_ISCSI:                       "ISCSI（Internet小型计算机系统接口配置）",
		EN_RAID:                        "支持Raid功能",
		EN_HARDDISK_INFO:               "支持磁盘信息F5查询",
		EN_PICINPIC:                    "支持画中画功能",
		EN_PLAYBACK_SPEED_CTRL_SUPPORT: "同 EN_PLAYBACK_SPEED_CTRL",
		EN_LF_XDEV:                     "支持24、32、64路LF-X系列",
		EN_DSP_ENCODE_CAP:              "F5 DSP编码能力查询",
		EN_MULTICAST:                   "组播能力查询",
		EM_NET_LIMIT:                   "网络限制能力查询",
		EM_COM422:                      "串口422",
		EM_PROTOCAL_FRAMEWORK:          "是否支持三代协议框架",
		EM_WRITE_DISK_OSD:              "刻录OSD叠加",
		EM_DYNAMIC_MULTI_CONNECT:       "动态多连接",
		EM_CLOUDSERVICE:                "云服务",
		EM_RECORD_INFO:                 "录像信息上报",
		EN_DYNAMIC_REG:                 "主动注册能力",
		EM_MULTI_PLAYBACK:              "多通道预览回放",
		EN_ENCODE_CHN:                  "编码通道",
		EN_SEARCH_RECORD:               "录像查询",
		EN_UPDATE_MD5:                  "支持升级文件传输完成后做MD5验证,1-支持MD5验证",
	}
)

type NET_PARAM struct {
	nWaittime            int32   // 等待超时时间(毫秒为单位),为0默认5000ms
	nConnectTime         int32   // 连接超时时间(毫秒为单位),为0默认1500ms
	nConnectTryNum       int32   // 连接尝试次数,为0默认1次
	nSubConnectSpaceTime int32   // 子连接之间的等待时间(毫秒为单位),为0默认10ms
	nGetDevInfoTime      int32   // 获取设备信息超时时间,为0默认1000ms
	nConnectBufSize      int32   // 每个连接接收数据缓冲大小(字节为单位),为0默认250*1024
	nGetConnInfoTime     int32   // 获取子连接信息超时时间(毫秒为单位),为0默认1000ms
	nSearchRecordTime    int32   // 按时间查询录像文件的超时时间(毫秒为单位),为0默认为3000ms
	nsubDisconnetTime    int32   // 检测子链接断线等待时间(毫秒为单位),为0默认为60000ms
	byNetType            byte    // 网络类型, 0-LAN, 1-WAN
	byPlaybackBufSize    byte    // 回放数据接收缓冲大小（M为单位）,为0默认为4M
	bDetectDisconnTime   byte    // 心跳检测断线时间(单位为秒),为0默认为60s,最小时间为2s
	bKeepLifeInterval    byte    // 心跳包发送间隔(单位为秒),为0默认为10s,最小间隔为2s
	nPicBufSize          int32   // 实时图片接收缓冲大小（字节为单位）,为0默认为2*1024*1024
	bReserved            [4]byte // 保留字段字段
}

// 设备信息扩展
type NET_DEVICEINFO_Ex struct {
	sSerialNumber    [DH_SERIALNO_LEN]byte // 序列号
	nAlarmInPortNum  int32                 // DVR报警输入个数
	nAlarmOutPortNum int32                 // DVR报警输出个数
	nDiskNum         int32                 // DVR硬盘个数
	nDVRType         int32                 // DVR类型,见枚举NET_DEVICE_TYPE
	nChanNum         int32                 // DVR通道个数
	byLimitLoginTime byte                  // 在线超时时间,为0表示不限制登陆,非0表示限制的分钟数
	byLeftLogTimes   byte                  // 当登陆失败原因为密码错误时,通过此参数通知用户,剩余登陆次数,为0时表示此参数无效
	bReserved        [2]byte               // 保留字节,字节对齐
	nLockLeftTime    int32                 // 当登陆失败,用户解锁剩余时间（秒数）, -1表示设备未设置该参数
	Reserved         [24]byte              // 保留
}

type DHDEV_NET_SNMP_CFG struct {
	bEnable       int                          // SNMP使能
	iSNMPPort     int32                        // SNMP端口
	szReadCommon  [DH_MAX_SNMP_COMMON_LEN]byte // 读共同体
	szWriteCommon [DH_MAX_SNMP_COMMON_LEN]byte // 写共同体
	szTrapServer  [64]byte                     // trap地址
	iTrapPort     int32                        // trap端口
	bSNMPV1       byte                         // 设备是否开启支持版本1格式,0不使能；1使能
	bSNMPV2       byte                         // 设备是否开启支持版本2格式,0不使能；1使能
	bSNMPV3       byte                         // 设备是否开启支持版本3格式,0不使能；1使能
	szReserve     [125]byte
}

type NET_TIME struct {
	dwYear   uint32 // 年
	dwMonth  uint32 // 月
	dwDay    uint32 // 日
	dwHour   uint32 // 时
	dwMinute uint32 // 分
	dwSecond uint32 // 秒
}
