-- MySQL dump 10.13  Distrib 5.1.72, for pc-linux-gnu (i686)
--
-- Host: localhost    Database: hupunac
-- ------------------------------------------------------
-- Server version   5.1.72

DELIMITER ;

SET FOREIGN_KEY_CHECKS=0;

DROP PROCEDURE if exists add_col_homework;
DROP PROCEDURE  if exists add_public;
DELIMITER //
-- 判断表或者列是否存在
CREATE PROCEDURE add_public(IN tabname CHAR(64), IN colname CHAR(64), OUT err INT) BEGIN
    IF EXISTS (SELECT column_name FROM  information_schema.columns where TABLE_NAME = tabname and COLUMN_NAME = colname)
    THEN
      -- SELECT column_name FROM  information_schema.columns where TABLE_NAME = tabname and COLUMN_NAME = colname;
        set err=0;
    ELSE
        -- SELECT column_name FROM  information_schema.columns where TABLE_NAME = tabname and COLUMN_NAME = colname;
        set err=1;
    END IF;


END//

CREATE PROCEDURE add_col_homework() BEGIN 

--  认证流程表新增字段: ifenablelogin ienableotherlogin iloginmethod iloginnum
    call add_public('tauthpolicyapprove', 'ifenablelogin', @err);
    IF (@err = 1)
    THEN
        ALTER TABLE `tauthpolicyapprove` ADD COLUMN `ifenablelogin`  tinyint(1) NULL DEFAULT 0 COMMENT '是否开启登录设置 0:不开启 1：开启' AFTER `iadminaudit`;
    END IF;
    
    call add_public('tauthpolicyapprove', 'ienableotherlogin', @err);
    IF (@err = 1)
    THEN
        ALTER TABLE `tauthpolicyapprove` ADD COLUMN `ienableotherlogin`  tinyint(1) NULL DEFAULT 0 COMMENT '是否禁止其他设备登录 0:禁止 1:允许' AFTER `ifenablelogin`;
    END IF;

    call add_public('tauthpolicyapprove', 'iloginmethod', @err);
    IF (@err = 1)
    THEN
        ALTER TABLE `tauthpolicyapprove` ADD COLUMN `iloginmethod`  tinyint(1) NULL DEFAULT 0 COMMENT '登录方式 0:不限制 1:下线 2:限制数量' AFTER `ienableotherlogin`;
    END IF;

    call add_public('tauthpolicyapprove', 'iloginnum', @err);
    IF (@err = 1)
    THEN
        ALTER TABLE `tauthpolicyapprove` ADD COLUMN `iloginnum`  tinyint(2) NULL DEFAULT 2 COMMENT '登录的用户数量' AFTER `iloginmethod`;
    END IF;
    
--  认证流程系统安检控制新增字段: ifappsecurity ifnetworksecurity

    call add_public('tauthpolicysecuritycheck', 'ifappsecurity', @err);
    IF (@err = 1)
    THEN
        ALTER TABLE `tauthpolicysecuritycheck` ADD COLUMN `ifappsecurity`  tinyint(1) NULL DEFAULT NULL COMMENT '启用助手安检' AFTER `isecuritycheckperiods`;
    END IF;

    call add_public('tauthpolicysecuritycheck', 'ifnetworksecurity', @err);
    IF (@err = 1)
    THEN
        ALTER TABLE `tauthpolicysecuritycheck` ADD COLUMN `ifnetworksecurity`  tinyint(1) NULL DEFAULT NULL COMMENT '启用网络安检' AFTER `ifappsecurity`;
    END IF;

--  Radius设置表tauthradius新增字段: socketTimeout authtype
    call add_public('tauthradius', 'socketTimeout', @err);
    IF (@err = 1)
    THEN
        ALTER TABLE `tauthradius` ADD COLUMN `socketTimeout`  int(11) NULL DEFAULT NULL AFTER `isenable`;
    END IF;

    call add_public('tauthradius', 'authtype', @err);
    IF (@err = 1)
    THEN
        ALTER TABLE `tauthradius` ADD COLUMN `authtype`  varchar(4) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL AFTER `socketTimeout`;
    END IF;

--  Radius设置表tauthradius修改字段: iport isenable
    call add_public('tauthradius', 'iport', @err);
    IF (@err = 0)
    THEN
        ALTER TABLE `tauthradius` MODIFY COLUMN `iport`  smallint(5) UNSIGNED NOT NULL DEFAULT 1812 COMMENT 'radius端口号,0-65535' AFTER `spwd`;
    END IF;

    call add_public('tauthradius', 'isenable', @err);
    IF (@err = 0)
    THEN
        ALTER TABLE `tauthradius` MODIFY COLUMN `isenable`  tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否开启了radius认证,0未开启，1开启' AFTER `iport`;
    END IF;
    
--  新增移动安检表 tcheckinstallsoftwarecontentmobile
    call add_public('tcheckinstallsoftwarecontentmobile', 'installsoftwaremobileid', @err);
    IF (@err = 1)
    THEN
        CREATE TABLE `tcheckinstallsoftwarecontentmobile` (
        `installsoftwaremobileid`  int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id' ,
        `ipolicyauthid`  int(11) NULL DEFAULT NULL COMMENT '主键id' ,
        `softname`  varchar(168) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '安装软件名称' ,
        `softnamedescrible`  varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '安装软件描述' ,
        `smodifyer`  varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '修改人' ,
        `dmodifytime`  datetime NULL DEFAULT NULL COMMENT '修改时间' ,
        PRIMARY KEY (`installsoftwaremobileid`)
        )
        ENGINE=MyISAM
        DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
        CHECKSUM=0
        ROW_FORMAT=Dynamic
        DELAY_KEY_WRITE=0
        ;
    END IF;

--  新增移动安检配置表 tcheckinstallsoftwaremobile
    call add_public('tcheckinstallsoftwaremobile', 'ipolicyauthid', @err);
    IF (@err = 1)
    THEN
        CREATE TABLE `tcheckinstallsoftwaremobile` (
        `ipolicyauthid`  int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id' ,
        `scompanycode`  varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '公司编码' ,
        `idepartmentid`  int(11) NULL DEFAULT NULL COMMENT '部门id' ,
        `iuserid`  int(11) NULL DEFAULT NULL COMMENT '用户id' ,
        `spolicyname`  varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '策略名称' ,
        `isornotcheckp`  int(11) NULL DEFAULT NULL COMMENT '是否配置了安装软件' ,
        `smodifyer`  varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '修改人' ,
        `dmodifytime`  datetime NULL DEFAULT NULL COMMENT '修改时间' ,
        `sstarttime`  varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '响应时间段' ,
        `sendtime`  varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '提示信息' ,
        `dstartdate`  datetime NULL DEFAULT NULL COMMENT '策略开始日期' ,
        `denddate`  datetime NULL DEFAULT NULL COMMENT '策略结束日期' ,
        `iintervaldays`  int(11) NULL DEFAULT NULL COMMENT '执行间隔天数' ,
        `spolicyaction`  int(11) NULL DEFAULT NULL COMMENT '策略动作' ,
        `ilogrecord`  int(1) NULL DEFAULT NULL COMMENT '日志记录：   1.记录           0.不记录' ,
        `iwarnlevelid`  int(11) NULL DEFAULT NULL COMMENT '警报级别id' ,
        `ipriority`  int(11) NULL DEFAULT NULL COMMENT '优先级' ,
        PRIMARY KEY (`ipolicyauthid`)
        )
        ENGINE=MyISAM
        DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
        CHECKSUM=0
        ROW_FORMAT=Dynamic
        DELAY_KEY_WRITE=0
        ;
    END IF;

END//

DELIMITER ;  
CALL add_col_homework();
DROP PROCEDURE  IF EXISTS add_col_homework;
DROP PROCEDURE  IF EXISTS add_public;


-- ----------------------------
-- Table structure for `tpageresource`
-- ----------------------------
DROP TABLE IF EXISTS `tpageresource`;
CREATE TABLE `tpageresource` (
  `ipageresourceid` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `spageresourcekey` varchar(32) NOT NULL COMMENT '资源信息的键',
  `zh` varchar(526) DEFAULT NULL COMMENT '中文',
  `en` varchar(526) DEFAULT NULL COMMENT '英文',
  `sdesc` varchar(128) DEFAULT NULL COMMENT '备注',
  `ft` varchar(526) DEFAULT NULL COMMENT '功能按钮',
  PRIMARY KEY (`ipageresourceid`)
) ENGINE=MyISAM AUTO_INCREMENT=162 DEFAULT CHARSET=utf8 COMMENT='页面资源信息';


-- ----------------------------
-- Table structure for `tpermission`
-- ----------------------------
DROP TABLE IF EXISTS `tpermission`;
CREATE TABLE `tpermission` (
  `ipermissionid` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `spermissionkey` varchar(32) NOT NULL COMMENT '权限在js里的变量名称',
  `spermissioncode` varchar(32) NOT NULL COMMENT '权限编号',
  `sparentpermissioncode` varchar(32) NOT NULL COMMENT '父级权限编号',
  `spermissionscope` varchar(2) DEFAULT '0' COMMENT '权限所属范围：0.管理员，1.审计员',
  `spermissiontype` varchar(2) DEFAULT '0' COMMENT '权限类型（0：模块；1：页面菜单权限；2：按钮权限）',
  `sifenabled` varchar(2) DEFAULT '1' COMMENT '是否启用（0：不启用；1：启用）',
  `zh` varchar(32) DEFAULT NULL COMMENT '中文',
  `en` varchar(32) DEFAULT NULL COMMENT '英文',
  `sdesc` varchar(256) DEFAULT NULL COMMENT '备注',
  `ft` varchar(32) DEFAULT NULL COMMENT '功能',
  `icon` varchar(256) DEFAULT NULL COMMENT '图标',
  PRIMARY KEY (`ipermissionid`)
) ENGINE=MyISAM AUTO_INCREMENT=251 DEFAULT CHARSET=utf8 COMMENT='权限表';


-- ----------------------------
-- Records of tpermission
-- 页面功能菜单表
-- ----------------------------
INSERT INTO `tpermission` VALUES ('1', 'Home_Page_Name', 'A00', '0', '0', '0', '1', 'Home ', 'Home', '首页模块', '首頁', null);
INSERT INTO `tpermission` VALUES ('2', 'Network_Monitoring_Name', 'B00', '0', '0', '0', '1', '网络', 'Network Monitoring', '网络监控', '准入控制', 'nacCore/nacCoreApp/nacApp/images/icons/network.gif');
INSERT INTO `tpermission` VALUES ('4', 'Strategy_Management_Name', 'D00', '0', '0', '0', '1', '用户', 'Strategy Management', '用户管理', '用戶', 'nacCore/nacCoreApp/nacApp/images/icons/user.gif');
INSERT INTO `tpermission` VALUES ('5', 'Statistical_Report_Name', 'E00', '0', '0', '0', '1', '策略', 'Statistical Report', '策略管理', '策略', 'nacCore/nacCoreApp/nacApp/images/icons/strategy.gif');
INSERT INTO `tpermission` VALUES ('6', 'System_Manage_Name', 'F00', '0', '0', '0', '1', '报表', 'System Manage', '报表管理', '報表', 'nacCore/nacCoreApp/nacApp/images/icons/report.gif');
INSERT INTO `tpermission` VALUES ('9', 'System_Manage', 'H00', '0', '0', '0', '1', '系统', 'System_Manage', '系统管理', '系統', 'nacCore/nacCoreApp/nacApp/images/icons/system.gif');
INSERT INTO `tpermission` VALUES ('11', 'Authentication_Manage', 'D0001', 'D00', '0', '1', '1', '用户管理', 'Authentication Manage', '认证用户管理', '认证用户管理', 'nacCore/nacCoreApp/nacApp/images/mainmenu/user_manager.gif');
INSERT INTO `tpermission` VALUES ('13', 'Sys_Manage_Account_Admin', 'H0001', 'H00', '0', '1', '1', '管理员账户', 'Admin Account', '管理员账户管理', '管理員帳戶', 'nacCore/nacCoreApp/nacApp/images/mainmenu/admin_account.gif');
INSERT INTO `tpermission` VALUES ('14', 'Sys_Manage_Account_Employee', 'D0003', 'D00', '0', '1', '1', '部门管理', 'Employee Account', '员工账户管理', '部門管理', 'nacCore/nacCoreApp/nacApp/images/mainmenu/dept_manager.gif');
INSERT INTO `tpermission` VALUES ('17', 'Role_Manage', 'H0002', 'H00', '0', '1', '1', '管理员角色', 'Role_Manage', '管理员角色', '管理员角色', 'nacCore/nacCoreApp/nacApp/images/mainmenu/admin_role.gif');
INSERT INTO `tpermission` VALUES ('57', 'Log_Manager', 'H0006', 'H00', '1', '1', '1', '日志管理', 'Log_Manager', '日志管理', '日誌管理', 'nacCore/nacCoreApp/nacApp/images/mainmenu/op_log_manager.png');
INSERT INTO `tpermission` VALUES ('73', 'Certificate_Strategy', 'E0002', 'E00', '0', '1', '1', '系统安检', 'Certificate Strategy', '系统安检', '系统安检', 'nacCore/nacCoreApp/nacApp/images/mainmenu/system_strategy.gif');
INSERT INTO `tpermission` VALUES ('80', 'Sys_client config', 'H0004', 'H00', '0', '1', '1', '客户端配置', 'client config', '客户端配置', '客戶端配置', 'nacCore/nacCoreApp/nacApp/images/mainmenu/client_config.gif');
INSERT INTO `tpermission` VALUES ('81', 'Device_manage', 'B0004', 'B00', '0', '1', '0', '设备管理', 'device manage', '设备管理', '設備管理', 'nacCore/nacCoreApp/nacApp/images/mainmenu/device_manager.gif');
INSERT INTO `tpermission` VALUES ('82', 'Network_config\r\n', 'E0004', 'E00', '0', '1', '1', '控制器配置', 'control config', '控制器配置', '控制器配置', 'nacCore/nacCoreApp/nacApp/images/mainmenu/control_config.gif');
INSERT INTO `tpermission` VALUES ('84', 'Guest_Manage', 'D0002', 'D00', '0', '1', '1', '来宾管理', 'Guest_Manage', '来宾用户管理', '来宾用户管理', 'nacCore/nacCoreApp/nacApp/images/mainmenu/guest_manager.gif');
INSERT INTO `tpermission` VALUES ('85', 'Station_Manage', 'B0005', 'B00', '0', '1', '1', '工位管理', 'Station_Manage', '员工办公位置定位', '工位管理', 'nacCore/nacCoreApp/nacApp/images/mainmenu/station_manager.gif');
INSERT INTO `tpermission` VALUES ('104', 'Strategy_Base', 'E0007', 'E00', '0', '1', '1', '策略基础', 'Strategy_Base', '策略基础', '策略基礎', 'nacCore/nacCoreApp/nacApp/images/mainmenu/strategy_base.gif');
INSERT INTO `tpermission` VALUES ('109', 'User_Report_Form', 'F0002', 'F00', '0', '1', '0', '用户报表', 'User_Report_Form', '用户报表', '用户报表', 'nacCore/nacCoreApp/nacApp/images/mainmenu/user_report.gif');
INSERT INTO `tpermission` VALUES ('110', 'Device_Report_Form', 'F0001', 'F00', '0', '1', '0', '设备报表', 'Device_Report_Form', '设备报表', '设备报表', 'nacCore/nacCoreApp/nacApp/images/mainmenu/device_report.gif');
INSERT INTO `tpermission` VALUES ('111', 'Guest_Report_Form', 'F0003', 'F00', '0', '1', '0', '来宾报表', 'Guest_Report_Form', '来宾报表', '来宾报表', 'nacCore/nacCoreApp/nacApp/images/mainmenu/guest_report.gif');
INSERT INTO `tpermission` VALUES ('121', 'Devicetype_Manage', 'B0006', 'B00', '0', '1', '1', '设备类型', 'Devicetype Manage', '设备类型', '设备类型', 'nacCore/nacCoreApp/nacApp/images/mainmenu/device_type.gif');
INSERT INTO `tpermission` VALUES ('122', 'Client_Auth_Code', 'H0005', 'H00', '0', '1', '1', '客户端卸载码', 'Client_Auth_Code', '客户端卸载码', '客户端卸载码', 'nacCore/nacCoreApp/nacApp/images/mainmenu/client_uninstall.gif');
INSERT INTO `tpermission` VALUES ('123', 'Auth_Process_Policy', 'E0001', 'E00', '0', '1', '1', '认证流程', 'Auth_Process_Policy', '认证流程策略', '认证流程策略', 'nacCore/nacCoreApp/nacApp/images/mainmenu/auth_process.gif');
INSERT INTO `tpermission` VALUES ('124', 'Bind_Policy', 'E0003', 'E00', '0', '1', '1', '绑定策略', 'Bind Policy', '绑定策略', '绑定策略', 'nacCore/nacCoreApp/nacApp/images/mainmenu/bind_strategy.gif');
INSERT INTO `tpermission` VALUES ('125', 'System_Config', 'H0007', 'H00', '0', '1', '1', '系统配置', 'System Config', '系统配置', '系統配置', 'nacCore/nacCoreApp/nacApp/images/mainmenu/system_config.png');
INSERT INTO `tpermission` VALUES ('126', 'Network_Diagnosis', 'H0008', 'H00', '0', '1', '1', '网络诊断', 'Network_Diagnosis', '网络诊断', '网络诊断', 'nacCore/nacCoreApp/nacApp/images/mainmenu/network.png');
INSERT INTO `tpermission` VALUES ('127', 'Data_Backup', 'H0009', 'H00', '0', '1', '1', '数据库备份', 'Data_Backup', '数据库备份', '数据库备份', 'nacCore/nacCoreApp/nacApp/images/mainmenu/db_backup.png');
INSERT INTO `tpermission` VALUES ('128', 'AD_Manage', 'D0004', 'D00', '0', '1', '1', 'AD域管理', 'AD_Manage', 'AD域管理', 'AD域管理', 'nacCore/nacCoreApp/nacApp/images/mainmenu/ad.png');
INSERT INTO `tpermission` VALUES ('129', 'Guest_Code_Manage', 'H0010', 'H00', '0', '1', '0', '来宾上网码', 'Guest_Code_Manage', '来宾上网码', '来宾上网码', 'nacCore/nacCoreApp/nacApp/images/mainmenu/guestCode.png');
INSERT INTO `tpermission` VALUES ('131', 'Sys_Individualization', 'H0011', 'H00', '0', '1', '1', '个性化', 'Sys_Individualization', '个性化', '个性化', 'nacCore/nacCoreApp/nacApp/images/mainmenu/Individualization.png');
INSERT INTO `tpermission` VALUES ('132', 'Sys_Debug_Log', 'H0012', 'H00', '0', '1', '1', '调试日志', 'Debug Log', '调试日志', '調試日誌', 'nacCore/nacCoreApp/nacApp/images/mainmenu/debug_log.png');
INSERT INTO `tpermission` VALUES ('133', 'Original_Log', 'F0004', 'F00', '0', '1', '1', '原始日志', 'Original Log', '原始日志', '原始日誌', 'nacCore/nacCoreApp/nacApp/images/mainmenu/original_log.png');
INSERT INTO `tpermission` VALUES ('134', 'WarnLog_Export_Email', 'E0006', 'E00', '0', '1', '1', '警报外发', 'WarnLog_Export_Email', '邮件报警', '邮件报警', 'nacCore/nacCoreApp/nacApp/images/mainmenu/warntoemail.png');
INSERT INTO `tpermission` VALUES ('136', 'Alarm_Strategy', 'E0005', 'E00', '0', '1', '1', '警报策略', 'Alarm_Strategy', '警报策略', '警报策略', 'nacCore/nacCoreApp/nacApp/images/mainmenu/warn_policy.png');
INSERT INTO `tpermission` VALUES ('137', 'Warn_Log', 'F0005', 'F00', '0', '1', '1', '警报日志', 'Warn Log', '警报日志', '警報日誌', 'nacCore/nacCoreApp/nacApp/images/mainmenu/warn_log.png');
INSERT INTO `tpermission` VALUES ('138', 'Guest_Classify', 'D0005', 'D00', '0', '1', '1', '来宾分类', 'GuestClassify', '来宾分类', '來賓分類', 'nacCore/nacCoreApp/nacApp/images/mainmenu/warntosms.png');
INSERT INTO `tpermission` VALUES ('139', 'Ip_Config', 'B0001', 'B00', '0', '1', '1', 'IP管理', 'IP Config', 'IP管理', 'IP管理', 'nacCore/nacCoreApp/nacApp/images/mainmenu/ip.png');
INSERT INTO `tpermission` VALUES ('140', 'MAC_Config', 'B0002', 'B00', '0', '1', '1', 'MAC管理', 'MAC_Config', 'MAC管理', 'MAC管理', 'nacCore/nacCoreApp/nacApp/images/mainmenu/mac_manage.jpg');
INSERT INTO `tpermission` VALUES ('141', 'Port_Config', 'B0003', 'B00', '0', '1', '1', 'PORT管理', 'PORT Config', 'PORT管理', 'PORT管理', 'nacCore/nacCoreApp/nacApp/images/mainmenu/port_manage.jpg');
INSERT INTO `tpermission` VALUES ('142', 'Sys_Bind_Cycle_Config', 'H0013', 'H00', '0', '1', '1', '绑定巡检周期', 'Bind_Cycle_Config', '绑定巡检周期', '綁定巡檢週期', 'nacCore/nacCoreApp/nacApp/images/mainmenu/bind_cycle.png');
INSERT INTO `tpermission` VALUES ('242', 'Send_Info_Config', 'H0014', 'H00', '0', '1', '1', '邮件/短信配置', 'Send Info Config', '邮件/短信配置', '邮件/短信配置', 'nacCore/nacCoreApp/nacApp/images/mainmenu/sendinfo_config.png');
INSERT INTO `tpermission` VALUES ('243', 'Bind_Relation_Report', 'F0006', 'F00', '0', '1', '1', '绑定报表', 'Bind Relation Report', '绑定报表', '绑定报表', 'nacCore/nacCoreApp/nacApp/images/mainmenu/bind_relation.png');
INSERT INTO `tpermission` VALUES ('244', 'Active_Device_TopN', 'F0007', 'F00', '0', '1', '1', '活跃设备报表', 'ActiveDeviceTopN', '活跃设备报表', '活跃设备报表', 'nacCore/nacCoreApp/nacApp/images/mainmenu/active_device.png');
INSERT INTO `tpermission` VALUES ('245', 'Active_User_TopN', 'F0008', 'F00', '0', '1', '1', '活跃用户报表', 'ActiveUserTopN', '活跃用户报表', '活跃用户报表', 'nacCore/nacCoreApp/nacApp/images/mainmenu/active_user.png');
INSERT INTO `tpermission` VALUES ('246', 'Sys_Upgrade_log', 'H0015', 'H00', '0', '1', '1', '升级日志', 'System_Upgrade_Log', '升级日志', '升级日志', 'nacCore/nacCoreApp/nacApp/images/mainmenu/active_user.png');
INSERT INTO `tpermission` VALUES ('247', 'Device_Fingerprint', 'B0007', 'B00', '0', '1', '1', '设备指纹', 'Device_Fingerprint', '设备指纹', '设备指纹', 'nacCore/nacCoreApp/nacApp/images/mainmenu/fingerprint.png');
INSERT INTO `tpermission` VALUES ('248', 'Hot_Standby', 'H0016', 'H00', '0', '1', '1', '双机热备', 'Hot_Standby', 'Hot Standby', 'Hot Standby', 'nacCore/nacCoreApp/nacApp/images/mainmenu/hot_standby.png');
INSERT INTO `tpermission` VALUES ('249', 'Radius_Manage', 'D0006', 'D00', '0', '1', '1', 'Radius管理', 'Radius_Manage', 'Radius管理', 'Radius管理', 'nacCore/nacCoreApp/nacApp/images/mainmenu/hupu_radius.png');
INSERT INTO `tpermission` VALUES ('250', 'Mobile_Secruity', 'E0008', 'E00', '0', '1', '1', '移动安检', 'Mobile_Secruity', '移动安检', '移动安检', 'nacCore/nacCoreApp/nacApp/images/mainmenu/mobile_security.png');



-- ----------------------------
-- Records of tpageresource
-- 页面菜单表元素表
-- ----------------------------
INSERT INTO `tpageresource` VALUES ('1', 'frame_current_user', '用户', 'Current User', '主页面TOP板块，当前登录用户', '當前用戶');
INSERT INTO `tpageresource` VALUES ('2', 'frame_login_time', '登录', 'Login Time', '主页面TOP板块，上次登录时间', '登錄時間');
INSERT INTO `tpageresource` VALUES ('3', 'frame_loginout', '注销', 'Login Out', '主页面，注销按钮', '註銷');
INSERT INTO `tpageresource` VALUES ('4', 'frame_exit', '退出', 'Exit', '主页面，退出', '退出');
INSERT INTO `tpageresource` VALUES ('5', 'frame_selectlanguage', '语言', 'Languge', '主页面，语言选择下拉框前的标签', '語言');
INSERT INTO `tpageresource` VALUES ('6', 'dept_manage_deptgrid_title', '部门信息', 'Dept Info', '部门管理页面部门grid的Title信息', '部門信息');
INSERT INTO `tpageresource` VALUES ('7', 'dept_manage_postgrid_title', '岗位信息', 'Post Info', '部门管理页面岗位grid的Title信息', '崗位信息');
INSERT INTO `tpageresource` VALUES ('8', 'dept_manage_postgrid_add', '添加岗位', 'Add Post', '添加岗位按钮', '添加崗位');
INSERT INTO `tpageresource` VALUES ('9', 'dept_manage_postgrid_update', '修改岗位', 'Modify Post', '修改岗位按钮', '修改崗位');
INSERT INTO `tpageresource` VALUES ('10', 'dept_manage_postgrid_delete', '删除岗位', 'Delete Post', '删除岗位', '刪除崗位');
INSERT INTO `tpageresource` VALUES ('11', 'common_refresh', '刷新', 'Refresh', '通用刷新', '刷新');
INSERT INTO `tpageresource` VALUES ('12', 'common_search', '查询', 'Search', '通用查询', '查詢');
INSERT INTO `tpageresource` VALUES ('13', 'common_keyword', '关键字', 'KeyWord', '通用查询关键字', '關鍵字');
INSERT INTO `tpageresource` VALUES ('14', 'common_keyin_please', '请输入关键字', 'Please KeyIn KeyWord', '通用的请输入关键字', '請輸入關鍵字');
INSERT INTO `tpageresource` VALUES ('15', 'dept_manage_deptgrid_add', '添加部门', 'Add Dept', '添加部门按钮', '添加部門');
INSERT INTO `tpageresource` VALUES ('16', 'dept_manage_deptgrid_addchild', '添加子部门', 'Add Child Dept', '添加子部门按钮', '添加子部門');
INSERT INTO `tpageresource` VALUES ('17', 'dept_manage_deptgrid_update', '修改部门', 'Modify Dept', '修改部门按钮', '修改部門');
INSERT INTO `tpageresource` VALUES ('18', 'dept_manage_deptgrid_delete', '删除部门', 'Delete Dept', '删除部门', '刪除部門');
INSERT INTO `tpageresource` VALUES ('19', 'common_expandall', '展开', 'ExpandAll', '展开全部', '展開');
INSERT INTO `tpageresource` VALUES ('20', 'common_collapseall', '收缩', 'CollapseAll', '收缩', '收縮');
INSERT INTO `tpageresource` VALUES ('21', 'dept_manage_column_postname', '岗位名称', 'Post Name', '岗位名称', '崗位名稱');
INSERT INTO `tpageresource` VALUES ('22', 'dept_manage_column_postdesc', '岗位描述', 'Post Desc', '岗位描述', '崗位描述');
INSERT INTO `tpageresource` VALUES ('23', 'dept_manage_column_deptname', '部门名称', 'Dept Name', '部门名称', '部門名稱');
INSERT INTO `tpageresource` VALUES ('24', 'frame_handup', '挂起', 'Hand Up', '主页面，挂起', '掛起');
INSERT INTO `tpageresource` VALUES ('25', 'frame_seek', '搜索', 'Search', '主页面搜索', '搜索');
INSERT INTO `tpageresource` VALUES ('26', 'tree_my_store_up', '我的收藏', 'My Store Up', '树节点', '我的收藏');
INSERT INTO `tpageresource` VALUES ('27', 'frame_advanced_search', '高级', 'Advanced', '主页面，高级搜索', '高級');
INSERT INTO `tpageresource` VALUES ('28', 'home_asm_state', '设备端口连接状态', 'Device Port Connection Status', '首页，标题', '設備端口連接狀態');
INSERT INTO `tpageresource` VALUES ('29', 'home_asm_state_now', '设备端口连接状态：', 'Equipment Port Connection Status:', '首页，信息', '設備端口連接狀態：');
INSERT INTO `tpageresource` VALUES ('30', 'frame_title', '网官网络接入管理', 'Network Terminal Access Management', '主页面，公司性质', '網官網絡接入管理');
INSERT INTO `tpageresource` VALUES ('31', 'vg_switch_list', '交换机管理列表', 'Switch Manage List', 'VG+页面title', '交換機管理列表');
INSERT INTO `tpageresource` VALUES ('32', 'vg_switch_name', '交换机名称', 'Switch Name', 'VG+页面', '交換機名稱');
INSERT INTO `tpageresource` VALUES ('33', 'vg_switch_add', '添加交换机', 'Add Switch', 'VG+页面', '添加交換機');
INSERT INTO `tpageresource` VALUES ('34', 'vg_switch_edit', '编辑交换机', 'Edit Switch', 'VG+页面', '編輯交換機');
INSERT INTO `tpageresource` VALUES ('35', 'vg_switch_delete', '删除交换机', 'Delete Switch', 'VG+页面', '刪除交換機');
INSERT INTO `tpageresource` VALUES ('36', 'vg_switch_ip', 'IP地址', 'IP Address', 'VG+页面', 'IP地址');
INSERT INTO `tpageresource` VALUES ('37', 'vg_switch_factory', '交换机厂商', 'Switch Factory', 'VG+页面', '交換機廠商');
INSERT INTO `tpageresource` VALUES ('38', 'vg_switch_model', '交换机型号', 'Switch Model', 'VG+页面', '交換機型號');
INSERT INTO `tpageresource` VALUES ('39', 'login_name', '登录用户名', 'Login Name', 'VG+页面', '登錄用戶名');
INSERT INTO `tpageresource` VALUES ('40', 'mode_of_operation', '操作方式', 'Mode Of Operation', 'VG+页面', '操作方式');
INSERT INTO `tpageresource` VALUES ('41', 'redirect_url', '重定向URL', 'Redirect URL', '基本参数设置页面', '重定向URL');
INSERT INTO `tpageresource` VALUES ('42', 'redirect_url_address', '重定向URL地址', 'Redirect URL Address', '基本参数设置页面', '重定向URL地址');
INSERT INTO `tpageresource` VALUES ('43', 'access_type', '接入技术类型', 'Access Technology type', '基本参数设置页面', '接入技術類型');
INSERT INTO `tpageresource` VALUES ('44', 'access_choose', '请选择您要启用的接入技术', 'Please Choose Your To Enable Access Technology', '基本参数设置页面', '請選擇您要啟用的接入技術');
INSERT INTO `tpageresource` VALUES ('45', 'access_choose_type', '策略路由', 'Policy Routing', '基本参数设置页面', '策略路由');
INSERT INTO `tpageresource` VALUES ('46', 'access_next_ip', '下一跳IP地址', 'The Next IP Address', '基本参数设置页面', '下一跳IP地址');
INSERT INTO `tpageresource` VALUES ('47', 'access_next_mac', '下一跳MAC地址', 'The Next MAC Address', '基本参数设置页面', '下一跳MAC地址');
INSERT INTO `tpageresource` VALUES ('48', 'access_get_next', '获取下一跳MAC地址', 'Get The Next Hop MAC Address', null, '獲取下一跳MAC地址');
INSERT INTO `tpageresource` VALUES ('49', 'test_cycle', '网络应用客户端检查', 'Network Application Client Inspection', '基本参数设置页面', '網絡應用客戶的檢查');
INSERT INTO `tpageresource` VALUES ('50', 'test_cycle_choose', '是否启用', 'Whether To Enable', '基本参数设置页面', '是否啟用');
INSERT INTO `tpageresource` VALUES ('51', 'enable', '启用', 'Enable', '全局', '啟用');
INSERT INTO `tpageresource` VALUES ('52', 'close', '关闭', 'Close', '全局', '關閉');
INSERT INTO `tpageresource` VALUES ('53', 'test_cycle_time', '周期时长', 'Cycle Time', '基本参数设置页面', '週期時長');
INSERT INTO `tpageresource` VALUES ('54', 'test_cycle_num', '周期次数', 'Cycle Numbers', '基本参数设置页面', '週期次數');
INSERT INTO `tpageresource` VALUES ('55', 'test_cycle_ips', '服务器IP组', 'IP Server', '基本参数设置页面', '服務器IP組');
INSERT INTO `tpageresource` VALUES ('56', 'seconds', '秒', 'Seconds', '全局', '秒');
INSERT INTO `tpageresource` VALUES ('57', 'numbers', '次', 'Numbers', '基本参数设置页面', '次');
INSERT INTO `tpageresource` VALUES ('58', 'nat_settings', 'NAT设置', 'NAT Settings', '基本参数设置页面', 'NAT設置');
INSERT INTO `tpageresource` VALUES ('59', 'nat_settings_choose', '是否启用NAT网段设置', 'Whether To Enable NAT Network Settings', '基本参数设置页面', '是否啟用NAT網段設置');
INSERT INTO `tpageresource` VALUES ('60', 'nat_network_settings', 'NAT网段设置', 'NAT Network Settings', '基本参数设置页面', 'NAT網段設置');
INSERT INTO `tpageresource` VALUES ('61', 'help_text', '帮助说明', ' Help Text', '基本参数设置页面', '幫助說明');
INSERT INTO `tpageresource` VALUES ('69', 'save_config', '确定', 'confirm', '基本参数设置页面', '确定');
INSERT INTO `tpageresource` VALUES ('70', 'reset', '重置', 'reset', '基本参数设置页面', '重置');
INSERT INTO `tpageresource` VALUES ('77', 'prompt', '提示', 'Prompt', '全局', '提示');
INSERT INTO `tpageresource` VALUES ('91', 'ip_illegal', 'IP不合法！', 'IP Illegal!', '全局', 'IP不合法！');
INSERT INTO `tpageresource` VALUES ('92', 'next_hop', '下一跳IP格式不正确!', 'Next Hop IP Format Is Incorrect!', '', '下一跳IP格式不正確！');
INSERT INTO `tpageresource` VALUES ('93', 'save_success', '保存成功!', 'Successfully saved!', '全局', '保存成功！');
INSERT INTO `tpageresource` VALUES ('94', 'save_faile', '保存失败!', 'Save failed!', '全局', '保存失败!');
INSERT INTO `tpageresource` VALUES ('95', 'next_hop_isnull', '下一跳IP为空，无法获取下一跳MAC', 'The next hop IP is empty, unable to get the next hop MAC', '', '下一跳IP為空，無法獲取下一跳MAC');
INSERT INTO `tpageresource` VALUES ('96', 'next_get_faile', '获取MAC地址失败！请核实您填写的IP地址是否正确！', 'Get MAC address failed! Please verify that the IP address you fill is correct!', '', '獲取MAC地址失败！請核實您填寫的IP地址是否正確！');
INSERT INTO `tpageresource` VALUES ('97', 'test_cycle_five', '服务器ip组不能超过5个!', 'server ip group can not be more than 5!', '', '服務器ip組不能超過5個!');
INSERT INTO `tpageresource` VALUES ('98', 'input_data', '输入数据有重复!', 'Input data duplication!', '', '輸入數據有重複!');
INSERT INTO `tpageresource` VALUES ('99', 'ips_isnull', '服务器ip组设置不能为空!', 'server ip group setting can not be empty!', '', '服務器ip組設置不能為空!');
INSERT INTO `tpageresource` VALUES ('100', 'nat_error', 'nat设置的网段不合法!', 'nat network settings are not legal!', '', 'nat設置的網段不合法!');
INSERT INTO `tpageresource` VALUES ('101', 'nat_error1', 'nat设置的ip不能跨网段!', 'ip nat settings can not cross segment!', '', 'nat設置的ip不能跨網段!');
INSERT INTO `tpageresource` VALUES ('102', 'nat_error2', 'nat设置的网段起始ip不能大于结束ip!', 'segment starting ip nat setting can not exceed the end of the ip!', '', 'nat設置的網段起始ip不能大於結束ip!');
INSERT INTO `tpageresource` VALUES ('103', 'nat_error3', 'ip网段有重复ip段!', 'Duplicate ip ip network segment!', '', 'ip网段有重复ip段!');
INSERT INTO `tpageresource` VALUES ('104', 'this_page', ' 当前页', 'This page', '全局', '當前頁');
INSERT INTO `tpageresource` VALUES ('105', 'page', '页', 'page', '全局', '頁');
INSERT INTO `tpageresource` VALUES ('106', 'altogether', '共', 'altogether', '全局', '共');
INSERT INTO `tpageresource` VALUES ('107', 'showing_records', '显示记录', 'Showing Records', '全局', '顯示記錄');
INSERT INTO `tpageresource` VALUES ('108', 'item', '条', 'item', '全局', '條');
INSERT INTO `tpageresource` VALUES ('109', 'no_data', '没有数据显示', 'NO Data', '全局', '沒有顯示數據');
INSERT INTO `tpageresource` VALUES ('110', 'previous', '上一页', 'Previous', '全局', '上一頁');
INSERT INTO `tpageresource` VALUES ('111', 'next', '下一页', 'Next', '全局', '下一頁');
INSERT INTO `tpageresource` VALUES ('112', 'last_page', '最后页', 'Last Page', '全局', '最後頁');
INSERT INTO `tpageresource` VALUES ('113', 'refresh', '刷新', 'Refresh', '全局', '刷新');
INSERT INTO `tpageresource` VALUES ('114', 'first_page', '第一页', 'First Page', '全局', '第一頁');
INSERT INTO `tpageresource` VALUES ('115', 'vlan_title', 'Vlan映射设置', 'Vlan Mapping Settings', 'vlan页面', 'Vlan映射設置');
INSERT INTO `tpageresource` VALUES ('116', 'vlan_serial', '序列号', 'Serial Number', 'vlan页面', '序列號');
INSERT INTO `tpageresource` VALUES ('117', 'vlan_before', '认证前（隔离）Vlan', 'Certification Ago (Isolated)Vlan', 'vlan页面', '認證前（隔離）Vlan');
INSERT INTO `tpageresource` VALUES ('118', 'vlan_after', '认证后（正常）Vlan', 'After Certification (Normal)Vlan', 'vlan页面', '認證后（正常）Vlan');
INSERT INTO `tpageresource` VALUES ('119', 'vlan_remark', '备注信息', 'Remarks', 'vlan页面', '備註信息');
INSERT INTO `tpageresource` VALUES ('120', 'vlan_add', '添加Vlan', 'Add Vlan', 'vlan页面', '添加Vlan');
INSERT INTO `tpageresource` VALUES ('121', 'delete_item', '删除', 'Delete', '全局', '刪除');
INSERT INTO `tpageresource` VALUES ('122', 'vlan_delete', '删除Vlan', 'Delete Vlan', 'vlan页面', '刪除Vlan');
INSERT INTO `tpageresource` VALUES ('123', 'vlan_text1', '认证前Vlan：设备未进行认证或安全检查不合格时所属的Vlan。', 'Before certification Vlan: Vlan when the device is not unqualified belongs authentication or security checks.', 'vlan页面', '認證前Vlan：設備未進行認證或安全檢查不合格時所屬的Vlan。');
INSERT INTO `tpageresource` VALUES ('125', 'vlan_text2', '认证后Vlan：设备安全检查合格后所属的Vlan，即设备正常工作的Vlan。', 'After certification Vlan: Vlan equipment safety checks after passing belongs, namely Vlan devices work properly.', 'vlan页面', '認證后Vlan：設備安全檢查合格后所屬的Vlan，即設備正常工作的Vlan。');
INSERT INTO `tpageresource` VALUES ('129', 'vlan_text3', '1.Vlan映射必须一一对应，不能交叉重复。', '1.Vlan mapping must correspond, not overlapping.', 'vlan页面', '1.Vlan映射必須一一對應，不能交叉重複。');
INSERT INTO `tpageresource` VALUES ('130', 'vlan_text4', '2.添加、删除Vlan后，需要点击<保存配置>才能生效。', '2 add, delete Vlan, need to click the<Save Configuration>to take effect.', 'vlan页面', '2.添加、刪除Vlan后，需要點擊<保存配置>才能生效');
INSERT INTO `tpageresource` VALUES ('134', 'vlan_text5', '3.备注信息不能超过10个字符。', '3. Remarks can not exceed 10 characters.', 'vlan页面', '3.備註信息不能超過10個字符。');
INSERT INTO `tpageresource` VALUES ('135', 'vlan_error1', 'Vlan ID只能为1-4096内的正整数,请重新输入!', 'Vlan ID can only be a positive integer 1-4096 inside, please re-enter!', 'vlan页面', 'Vlan ID只能為1-4096內的正整數,請重新輸入!');
INSERT INTO `tpageresource` VALUES ('136', 'vlan_error2', '认证前的Vlan ID 不能为负数,请重新输入!', 'Vlan ID authentication before can not be negative, please re-enter!', 'vlan页面', '認證前的Vlan ID 不能為負數,請重新輸入!');
INSERT INTO `tpageresource` VALUES ('137', 'vlan_error3', '认证前的Vlan ID不能为0,请重新输入!', 'Vlan ID authentication ago can not be 0, please re-enter!', 'vlan页面', '認證前的Vlan ID不能為0,請重新輸入!');
INSERT INTO `tpageresource` VALUES ('138', 'vlan_error4', '认证前的Vlan ID不能为空,请重新输入!', 'Vlan ID authentication can not be empty before, please re-enter!', 'vlan页面', '認證前的Vlan ID不能為空,請重新輸入!');
INSERT INTO `tpageresource` VALUES ('139', 'vlan_error5', '认证前的Vlan ID只能为1-4096内的正整数,请重新输入!', 'Vlan ID authentication before the only positive integer 1-4096, please re-enter!', 'vlan页面', '認證前的Vlan ID只能為1-4096内的正整數,請重新輸入!');
INSERT INTO `tpageresource` VALUES ('140', 'vlan_error6', '认证后的Vlan ID不能为0,请重新输入!', 'Certification after Vlan ID can not be 0, please re-enter!', 'vlan页面', '認證后的Vlan ID不能為0,請重新輸入!');
INSERT INTO `tpageresource` VALUES ('141', 'vlan_error7', '认证后的Vlan ID不能为空,请重新输入!', 'Certification after Vlan ID can not be empty, please re-enter!', 'vlan页面', '認證后的Vlan ID不能為空,請重新輸入!');
INSERT INTO `tpageresource` VALUES ('142', 'vlan_error8', '输入的备注信息在0到10个字符之间!', 'Remarks entered between 0-10 characters!', 'vlan页面', '輸入的備註信息在0到10个字符之間!');
INSERT INTO `tpageresource` VALUES ('143', 'vlan_error9', '认证后的Vlan ID 不能为负数,请重新输入!', 'Certification after Vlan ID can not be negative, please re-enter!', 'vlan页面', '認證后的Vlan ID 不能為負數,請重新輸入!');
INSERT INTO `tpageresource` VALUES ('144', 'vlan_error10', '认证后的Vlan ID只能为1-4096内的正整数,请重新输入!', 'After Vlan ID authentication can only be a positive integer 1-4096 inside, please re-enter!', 'vlan页面', '認證后的Vlan ID只能為1-4096内的正整數,請重新輸入!');
INSERT INTO `tpageresource` VALUES ('145', 'test_cycle_text', '多个服务器ip设置用英文分号\';\'隔开,如：\'192.168.1.1;192.168.1.2\',ip限制为5个', 'Multiple server ip settings Granville Shield English semicolon \';\' separated, such as: \'192.168.1.1; 192.168.1.2\', ip is limited to five', '基本参数设置页面', '多個服務器ip設置用英文分號\';\'隔開,如：\'192.168.1.1;192.168.1.2,ip限制為5個');
INSERT INTO `tpageresource` VALUES ('146', 'nat_network_text', '多个NAT网段设置用英文分号\';\'隔开,如：\'192.168.1.1-192.168.1.255;192.168.2.1-192.168.2.255\'', 'Multiple NAT network settings with a semicolon \';\' separated, such as: \'192.168.1.1-192.168.1.255; 192.168.2.1-192.168.2.255\'', '基本参数设置页面', '多個NAT網段設置用英文分號\';\'隔開,如：\'192.168.1.1-192.168.1.255;192.168.2.1-192.168.2.255\'');
INSERT INTO `tpageresource` VALUES ('147', 'help_text_one', '1.重定向URL地址：如果没有参数，则必须以\'/\'结尾。例：http://127.0.0.1/。', '1 Redirect URL: If no parameters, you must \'/\' at the end. Example: http://127.0.0.1/.', '基本参数设置页面', '1.重定向URL地址：如果沒有參數，則必須以\'/\'結尾。例：http://127.0.0.1/。');
INSERT INTO `tpageresource` VALUES ('148', 'help_text_two', '1.服务器ip组设置：如：192.168.1.1。只能设置5个ip。', '1 server ip group settings: eg: 192.168.1.1. You can only set 5 ip', '基本参数设置页面', '1.服務器ip組設置：如：192.168.1.1。只能設置5個ip。');
INSERT INTO `tpageresource` VALUES ('149', 'help_text_three', '1.多个NAT网段设置用英文分号\';\'隔开,如：192.168.1.1-192.168.1.255;192.168.2.1-192.168.2.255。', '1 NAT network settings with a semicolon \';\' separated, as :192.168.1.1-192 .168.1.255; 192.168.2.1-192.168.2.255.', '基本参数设置页面', '1.多個NAT網段設置用英文分號\';\'隔開,如：192.168.1.1-192.168.1.255;192.168.2.1-192.168.2.255。');
INSERT INTO `tpageresource` VALUES ('151', 'nat_error4', 'NAT网段设置不能为空!', 'NAT network settings can not be empty!', '基本参数设置页面', 'NAT網段設置不能為空!');
INSERT INTO `tpageresource` VALUES ('152', 'frame_companyname', '单位', 'company name', '全局参数', '公司名稱');
INSERT INTO `tpageresource` VALUES ('153', 'cycle_error1', '周期时长不能为空！', 'Cycle length can not be empty!', '基本参数设置页面', '週期時長不能為空');
INSERT INTO `tpageresource` VALUES ('154', 'cycle_error2', '周期时长不能为0！', 'Length can not be 0 when the cycle!', '基本参数设置页面', '週期時長不能為0！');
INSERT INTO `tpageresource` VALUES ('155', 'cycle_error3', '周期时长不能为负数！', 'Cycle length can not be negative!', '基本参数设置页面', '週期時長不能為負數！');
INSERT INTO `tpageresource` VALUES ('156', 'cycle_error4', '周期时长只能在30到90之间的正整数！', 'Long only positive integer between 30 to 90 when the cycle!', '基本参数设置页面', '週期時長只能在30到90之間的正整數！');
INSERT INTO `tpageresource` VALUES ('157', 'cycle_error5', '周期次数不能为空！', 'Cycle times can not be empty!', '基本参数设置页面', '週期次數不能為空！');
INSERT INTO `tpageresource` VALUES ('158', 'cycle_error6', '周期次数不能为0！', 'Cycle times can not be 0!', '基本参数设置页面', '週期次數不能為0！');
INSERT INTO `tpageresource` VALUES ('159', 'cycle_error7', '周期次数不能为负数！', 'Cycle times can not be negative!', '基本参数设置页面', '週期次數不能為負數！');
INSERT INTO `tpageresource` VALUES ('160', 'cycle_error8', '周期次数只能是1到5之间的正整数！', 'Cycle times can only be a positive integer between 1-5!', '基本参数设置页面', '週期次數只能是1到5之間的正整數！');
INSERT INTO `tpageresource` VALUES ('161', 'qucik_location', '快速定位', 'quick location', '快速定位', '快速定位');


-- ----------------------------
-- Table structure for `tmanufactory`
-- 交换机厂商表
-- ----------------------------
DROP TABLE IF EXISTS `tmanufactory`;
CREATE TABLE `tmanufactory` (
  `imanufactoryid` int(11) NOT NULL AUTO_INCREMENT COMMENT '厂商ID',
  `smanufactoryname` varchar(32) NOT NULL COMMENT '交换机厂商名称',
  `ssimplespell` varchar(32) NOT NULL COMMENT '简写形式',
  `sifsupported` varchar(2) NOT NULL DEFAULT '1' COMMENT '0:不支持;1:支持',
  `swebsite` varchar(255) DEFAULT NULL COMMENT '网址',
  `sdesc` varchar(255) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`imanufactoryid`)
) ENGINE=MyISAM AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tmanufactory
-- ----------------------------
INSERT INTO `tmanufactory` VALUES ('1', '思科', 'CISCO', '1', null, null);
INSERT INTO `tmanufactory` VALUES ('2', 'H3C', 'H3C', '1', null, null);
INSERT INTO `tmanufactory` VALUES ('3', '华为', 'HUAWEI', '1', null, null);
INSERT INTO `tmanufactory` VALUES ('4', '锐捷', 'RJ', '1', null, null);
INSERT INTO `tmanufactory` VALUES ('5', '博达', 'BD', '1', null, null);
INSERT INTO `tmanufactory` VALUES ('6', '中兴', 'ZTE', '1', null, null);
INSERT INTO `tmanufactory` VALUES ('7', '神州数码', 'SHENZHOUSHUMA', '1', null, null);
INSERT INTO `tmanufactory` VALUES ('8', '戴尔', 'DELL', '1', null, null);
INSERT INTO `tmanufactory` VALUES ('9', '惠普', 'HP', '1', null, null);
INSERT INTO `tmanufactory` VALUES ('10', '迈普', 'MP', '1', null, null);
INSERT INTO `tmanufactory` VALUES ('11', 'FORCE10', 'FORCE10', '1', null, null);
INSERT INTO `tmanufactory` VALUES ('12', 'FOUNDRY', 'FOUNDRY', '1', null, null);
INSERT INTO `tmanufactory` VALUES ('13', 'H3C无线控制器', 'H3C_WLC', '1', null, null);
INSERT INTO `tmanufactory` VALUES ('14', '凯创', 'KC', '1', null, null);
INSERT INTO `tmanufactory` VALUES ('15', '阿尔卡特', 'ALCATEL', '1', null, null);


-- ----------------------------
-- Table structure for `thardwareversion`
-- 交换机型号表
-- ----------------------------
DROP TABLE IF EXISTS `thardwareversion`;
CREATE TABLE `thardwareversion` (
  `ihardwareversionid` int(11) NOT NULL AUTO_INCREMENT,
  `smanufactorysimplespell` varchar(16) NOT NULL COMMENT '厂商简拼',
  `shardwareversionnumber` varchar(16) NOT NULL COMMENT '硬件版本号',
  `sifsupported` varchar(2) NOT NULL DEFAULT '1' COMMENT '0:不支持；1:支持',
  `sdesc` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ihardwareversionid`)
) ENGINE=MyISAM AUTO_INCREMENT=77 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of thardwareversion
-- ----------------------------
INSERT INTO `thardwareversion` VALUES ('1', 'CISCO', '1900', '1', null);
INSERT INTO `thardwareversion` VALUES ('2', 'CISCO', '2900', '1', null);
INSERT INTO `thardwareversion` VALUES ('3', 'CISCO', '2948', '1', null);
INSERT INTO `thardwareversion` VALUES ('4', 'CISCO', '2950', '1', null);
INSERT INTO `thardwareversion` VALUES ('5', 'CISCO', '2960', '1', null);
INSERT INTO `thardwareversion` VALUES ('6', 'CISCO', '3500', '1', null);
INSERT INTO `thardwareversion` VALUES ('7', 'CISCO', '3548', '1', null);
INSERT INTO `thardwareversion` VALUES ('8', 'CISCO', '3550', '1', null);
INSERT INTO `thardwareversion` VALUES ('9', 'CISCO', '3560', '1', null);
INSERT INTO `thardwareversion` VALUES ('10', 'CISCO', '3750', '1', null);
INSERT INTO `thardwareversion` VALUES ('11', 'CISCO', '4506', '1', null);
INSERT INTO `thardwareversion` VALUES ('12', 'CISCO', '6505', '1', null);
INSERT INTO `thardwareversion` VALUES ('13', 'CISCO', '其它', '1', null);
INSERT INTO `thardwareversion` VALUES ('14', 'H3C', '3100', '1', null);
INSERT INTO `thardwareversion` VALUES ('15', 'H3C', '3600', '1', null);
INSERT INTO `thardwareversion` VALUES ('16', 'H3C', '5000', '1', null);
INSERT INTO `thardwareversion` VALUES ('17', 'H3C', '5500', '1', null);
INSERT INTO `thardwareversion` VALUES ('18', 'H3C', '5510', '1', null);
INSERT INTO `thardwareversion` VALUES ('19', 'H3C', '7500', '1', null);
INSERT INTO `thardwareversion` VALUES ('20', 'H3C', '12000', '1', null);
INSERT INTO `thardwareversion` VALUES ('21', 'H3C', '12508', '1', null);
INSERT INTO `thardwareversion` VALUES ('22', 'HUAWEI', '2300', '1', null);
INSERT INTO `thardwareversion` VALUES ('23', 'HUAWEI', '2700', '1', null);
INSERT INTO `thardwareversion` VALUES ('24', 'HUAWEI', '3100', '1', null);
INSERT INTO `thardwareversion` VALUES ('25', 'HUAWEI', '3300', '1', null);
INSERT INTO `thardwareversion` VALUES ('26', 'HUAWEI', '3328', '1', null);
INSERT INTO `thardwareversion` VALUES ('27', 'HUAWEI', '5100', '1', null);
INSERT INTO `thardwareversion` VALUES ('28', 'HUAWEI', '5300', '1', null);
INSERT INTO `thardwareversion` VALUES ('29', 'HUAWEI', '5700', '1', null);
INSERT INTO `thardwareversion` VALUES ('30', 'HUAWEI', '7500', '1', null);
INSERT INTO `thardwareversion` VALUES ('31', 'HUAWEI', '8500', '1', null);
INSERT INTO `thardwareversion` VALUES ('32', 'HUAWEI', '9300', '1', null);
INSERT INTO `thardwareversion` VALUES ('33', 'HUAWEI', '9500', '1', null);
INSERT INTO `thardwareversion` VALUES ('34', 'HUAWEI', '其它', '1', null);
INSERT INTO `thardwareversion` VALUES ('35', 'RJ', '2600', '1', null);
INSERT INTO `thardwareversion` VALUES ('36', 'RJ', '2625', '1', null);
INSERT INTO `thardwareversion` VALUES ('37', 'RJ', '2924', '1', null);
INSERT INTO `thardwareversion` VALUES ('38', 'RJ', '其它', '1', null);
INSERT INTO `thardwareversion` VALUES ('39', 'BD', '3224', '1', null);
INSERT INTO `thardwareversion` VALUES ('40', 'BD', '其它', '1', null);
INSERT INTO `thardwareversion` VALUES ('41', 'ZTE', '3252', '1', null);
INSERT INTO `thardwareversion` VALUES ('42', 'ZTE', '2952', '1', null);
INSERT INTO `thardwareversion` VALUES ('43', 'ZTE', '其它', '1', null);
INSERT INTO `thardwareversion` VALUES ('44', 'SHENZHOUSHUMA', '3526', '1', null);
INSERT INTO `thardwareversion` VALUES ('45', 'SHENZHOUSHUMA', '3600', '1', null);
INSERT INTO `thardwareversion` VALUES ('46', 'SHENZHOUSHUMA', '3950', '1', null);
INSERT INTO `thardwareversion` VALUES ('47', 'SHENZHOUSHUMA', '其它', '1', null);
INSERT INTO `thardwareversion` VALUES ('48', 'DELL', '5212', '1', null);
INSERT INTO `thardwareversion` VALUES ('49', 'DELL', '其它', '1', null);
INSERT INTO `thardwareversion` VALUES ('50', 'HP', '2626', '1', null);
INSERT INTO `thardwareversion` VALUES ('51', 'HP', '2650', '1', null);
INSERT INTO `thardwareversion` VALUES ('52', 'HP', '5308', '1', null);
INSERT INTO `thardwareversion` VALUES ('53', 'HP', '其它', '1', null);
INSERT INTO `thardwareversion` VALUES ('54', 'MP', '3100', '1', null);
INSERT INTO `thardwareversion` VALUES ('55', 'MP', '3152', '1', null);
INSERT INTO `thardwareversion` VALUES ('56', 'MP', '其它', '1', null);
INSERT INTO `thardwareversion` VALUES ('57', 'FORCE10', 'S60', '1', null);
INSERT INTO `thardwareversion` VALUES ('58', 'FORCE10', '其它', '1', null);
INSERT INTO `thardwareversion` VALUES ('59', 'FOUNDRY', '2402', '1', null);
INSERT INTO `thardwareversion` VALUES ('60', 'FOUNDRY', '其它', '1', null);
INSERT INTO `thardwareversion` VALUES ('61', 'H3C_WLC', '3010', '1', null);
INSERT INTO `thardwareversion` VALUES ('62', 'H3C_WLC', '其它', '1', null);
INSERT INTO `thardwareversion` VALUES ('63', 'KC', 'C2H124', '1', null);
INSERT INTO `thardwareversion` VALUES ('64', 'KC', '其它', '1', null);
INSERT INTO `thardwareversion` VALUES ('65', 'ALCATEL', '6248', '1', null);
INSERT INTO `thardwareversion` VALUES ('66', 'ALCATEL', '其它', '1', null);
INSERT INTO `thardwareversion` VALUES ('67', 'H3C', 'S2600', '1', null);
INSERT INTO `thardwareversion` VALUES ('68', 'CISCO', '4503', '1', null);
INSERT INTO `thardwareversion` VALUES ('69', 'H3C', '5024', '1', null);
INSERT INTO `thardwareversion` VALUES ('70', 'H3C', '2610', '1', '');
INSERT INTO `thardwareversion` VALUES ('73', 'H3C', '5120', '1', null);
INSERT INTO `thardwareversion` VALUES ('72', 'CISCO', '2918', '1', null);
INSERT INTO `thardwareversion` VALUES ('74', 'HP', '2530', '1', null);
INSERT INTO `thardwareversion` VALUES ('75', 'HP', '2510', '1', null);
INSERT INTO `thardwareversion` VALUES ('76', 'HP', '5412', '1', null);


-- ----------------------------
-- Table structure for `trolepermission`
-- 角色权限表
-- ----------------------------
DROP TABLE IF EXISTS `trolepermission`;
CREATE TABLE `trolepermission` (
  `irolepermissionid` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `iroleid` int(11) DEFAULT NULL COMMENT '角色ID',
  `spermissioncode` varchar(32) DEFAULT NULL COMMENT '权限编号',
  `sdesc` varchar(128) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`irolepermissionid`)
) ENGINE=MyISAM AUTO_INCREMENT=68 DEFAULT CHARSET=utf8 COMMENT='角色权限对应关系表';

-- ----------------------------
-- Records of trolepermission
-- ----------------------------
INSERT INTO `trolepermission` VALUES ('1', '1', 'A00', null);
INSERT INTO `trolepermission` VALUES ('2', '1', 'B00', null);
INSERT INTO `trolepermission` VALUES ('3', '1', 'D00', null);
INSERT INTO `trolepermission` VALUES ('4', '1', 'E00', null);
INSERT INTO `trolepermission` VALUES ('5', '1', 'F00', null);
INSERT INTO `trolepermission` VALUES ('6', '1', 'H00', null);
INSERT INTO `trolepermission` VALUES ('7', '1', 'B0001', null);
INSERT INTO `trolepermission` VALUES ('8', '1', 'B0002', null);
INSERT INTO `trolepermission` VALUES ('9', '1', 'B0003', null);
INSERT INTO `trolepermission` VALUES ('10', '1', 'D0001', null);
INSERT INTO `trolepermission` VALUES ('11', '1', 'D0002', null);
INSERT INTO `trolepermission` VALUES ('12', '1', 'D0003', null);
INSERT INTO `trolepermission` VALUES ('13', '1', 'E0001', null);
INSERT INTO `trolepermission` VALUES ('14', '1', 'E0002', null);
INSERT INTO `trolepermission` VALUES ('15', '1', 'E0003', null);
INSERT INTO `trolepermission` VALUES ('16', '1', 'E0004', null);
INSERT INTO `trolepermission` VALUES ('17', '1', 'E0005', null);
INSERT INTO `trolepermission` VALUES ('18', '1', 'F0001', null);
INSERT INTO `trolepermission` VALUES ('19', '1', 'F0002', null);
INSERT INTO `trolepermission` VALUES ('20', '1', 'F0003', null);
INSERT INTO `trolepermission` VALUES ('21', '1', 'H0001', null);
INSERT INTO `trolepermission` VALUES ('22', '1', 'H0002', null);
INSERT INTO `trolepermission` VALUES ('24', '1', 'H0004', null);
INSERT INTO `trolepermission` VALUES ('25', '1', 'H0005', null);
INSERT INTO `trolepermission` VALUES ('26', '2', 'H00', null);
INSERT INTO `trolepermission` VALUES ('27', '2', 'H0006', null);
INSERT INTO `trolepermission` VALUES ('56', '4', 'B0003', null);
INSERT INTO `trolepermission` VALUES ('55', '4', 'B0002', null);
INSERT INTO `trolepermission` VALUES ('54', '4', 'B0001', null);
INSERT INTO `trolepermission` VALUES ('53', '4', 'B00', null);
INSERT INTO `trolepermission` VALUES ('32', '1', 'H0007', null);
INSERT INTO `trolepermission` VALUES ('33', '1', 'H0008', null);
INSERT INTO `trolepermission` VALUES ('34', '1', 'H0009', null);
INSERT INTO `trolepermission` VALUES ('35', '1', 'D0004', null);
INSERT INTO `trolepermission` VALUES ('36', '1', 'H0010', null);
INSERT INTO `trolepermission` VALUES ('37', '1', 'E0006', null);
INSERT INTO `trolepermission` VALUES ('38', '9', 'E00', null);
INSERT INTO `trolepermission` VALUES ('39', '9', 'E0003', null);
INSERT INTO `trolepermission` VALUES ('40', '9', 'E0001', null);
INSERT INTO `trolepermission` VALUES ('41', '1', 'H0011', null);
INSERT INTO `trolepermission` VALUES ('42', '1', 'H0012', null);
INSERT INTO `trolepermission` VALUES ('43', '1', 'F0004', null);
INSERT INTO `trolepermission` VALUES ('44', '1', 'E0007', null);
INSERT INTO `trolepermission` VALUES ('45', '1', 'E0008', null);
INSERT INTO `trolepermission` VALUES ('46', '1', 'E0009', null);
INSERT INTO `trolepermission` VALUES ('47', '1', 'F0005', null);
INSERT INTO `trolepermission` VALUES ('48', '1', 'D0005', null);
INSERT INTO `trolepermission` VALUES ('49', '1', 'B0004', null);
INSERT INTO `trolepermission` VALUES ('50', '1', 'B0006', null);
INSERT INTO `trolepermission` VALUES ('51', '1', 'B0005', null);
INSERT INTO `trolepermission` VALUES ('52', '1', 'H0013', null);
INSERT INTO `trolepermission` VALUES ('57', '4', 'H00', null);
INSERT INTO `trolepermission` VALUES ('58', '4', 'H0001', null);
INSERT INTO `trolepermission` VALUES ('59', '1', 'H0014', null);
INSERT INTO `trolepermission` VALUES ('60', '1', 'F0006', null);
INSERT INTO `trolepermission` VALUES ('61', '1', 'F0007', null);
INSERT INTO `trolepermission` VALUES ('62', '1', 'F0008', null);
INSERT INTO `trolepermission` VALUES ('63', '1', 'H0015', null);
INSERT INTO `trolepermission` VALUES ('64', '1', 'B0007', null);
INSERT INTO `trolepermission` VALUES ('65', '1', 'H0016', null);
INSERT INTO `trolepermission` VALUES ('66', '1', 'D0006', null);

-- 新增存储过程 p_searchcheckmobileinstallsoftwaretimemax
DELIMITER //

DROP PROCEDURE IF EXISTS `p_searchcheckmobileinstallsoftwaretimemax`;

CREATE DEFINER = `root`@`%` PROCEDURE `p_searchcheckmobileinstallsoftwaretimemax`(scompanycode VARCHAR(8),
iuserid VARCHAR(11),

idepartmentid VARCHAR(50),
nowdate VARCHAR(50))
BEGIN
    DECLARE userid INT;
    DECLARE deptid INT;
    DECLARE countFlag INT(8);
    SET @countFlag = 0;
    DROP TEMPORARY TABLE IF EXISTS tmpTableSoftware;
    CREATE TEMPORARY TABLE IF NOT EXISTS tmpTableSoftware  
         (  
             iden INT(8),
             spolicyaction INT(11),
             dstartdate datetime,
             denddate datetime,
             ilogrecord INT(1),
             sweekday VARCHAR(32),
             stimeperiod VARCHAR(255),
             spromptinfo VARCHAR(2048),
             iwarnlevel INT(11),
             swarnname VARCHAR(20)
         ) ENGINE = MEMORY;  
    WHILE idepartmentid IS NOT NULL AND idepartmentid <> 'null' AND idepartmentid <> '' DO
        IF iuserid IS NOT NULL AND iuserid <> 'null' AND iuserid <> '' THEN 
            SET userid = CAST(iuserid AS DECIMAL(8,0)); 
            
            INSERT INTO tmpTableSoftware SELECT @countFlag:=@countFlag + 1 AS iden,tt.spolicyaction,tt.dstartdate,tt.denddate,tt.ilogrecord,t.sweekday sweekday,
                    t.stimeperiod stimeperiod,tp.spromptinfo spromptinfo,tw.iwarnlevel iwarnlevel,tw.swarnname swarnname
                    FROM tcheckinstallsoftwaremobile tt 
                    LEFT JOIN tperiod t ON t.iperiodid=tt.sstarttime
                    LEFT JOIN tpromptinfo tp ON tp.ipromptinfoid=tt.sendtime
                    LEFT JOIN twarnlevel tw ON tw.iwarnlevelid = tt.iwarnlevelid
             WHERE tt.scompanycode= scompanycode 
                        AND tt.iuserid= userid
                        AND date_format(tt.dstartdate,'%Y-%m-%d') <= date_format(nowdate,'%Y-%m-%d')
                        AND date_format(tt.denddate,'%Y-%m-%d') >=date_format(nowdate,'%Y-%m-%d')
                        AND tt.isornotcheckp = 1 
            ORDER BY tt.ipriority DESC;
            
            
            SET iuserid = NULL; 
        ELSE
            
            SET deptid = CAST(idepartmentid AS DECIMAL(8,0)); 

            INSERT INTO tmpTableSoftware SELECT @countFlag:=@countFlag + 1 AS iden,tt.spolicyaction,tt.dstartdate,tt.denddate,tt.ilogrecord,t.sweekday sweekday,
                    t.stimeperiod stimeperiod,tp.spromptinfo spromptinfo,tw.iwarnlevel iwarnlevel,tw.swarnname swarnname
                    FROM tcheckinstallsoftwaremobile tt 
                    LEFT JOIN tperiod t ON t.iperiodid=tt.sstarttime
                    LEFT JOIN tpromptinfo tp ON tp.ipromptinfoid=tt.sendtime
                    LEFT JOIN twarnlevel tw ON tw.iwarnlevelid = tt.iwarnlevelid
             WHERE tt.scompanycode= scompanycode 
                        AND tt.idepartmentid= deptid
                        AND date_format(tt.dstartdate,'%Y-%m-%d') <= date_format(nowdate,'%Y-%m-%d')
                        AND date_format(tt.denddate,'%Y-%m-%d') >=date_format(nowdate,'%Y-%m-%d')
                        AND tt.isornotcheckp = 1 
            ORDER BY tt.ipriority DESC;

            
            IF deptid = -1 THEN
                SET idepartmentid = NULL; 
            ELSEIF deptid = -3 OR deptid = -2 THEN 
                SET idepartmentid = '-1';
            ELSE
                
                IF deptid > 0 THEN 
                    SELECT iparentdepartmentid INTO idepartmentid FROM tdepartment d WHERE d.idepartmentid = idepartmentid;
                    IF idepartmentid = '-1' THEN
                        SET idepartmentid = '-3';
                    END IF;
                ELSE 
                    SELECT iguestclassifyparentid INTO idepartmentid FROM tguestclassify g WHERE g.iguestclassifyid = ABS(idepartmentid);
                    IF idepartmentid <> '-2' THEN
                        SET idepartmentid = CONCAT('-',idepartmentid);
                    END IF;
                END IF;
            END IF;
        END IF;
    END WHILE;
    
    SELECT t.iden ipriority,t.spolicyaction,t.dstartdate,t.denddate,t.ilogrecord,t.sweekday sweekday,
        t.stimeperiod stimeperiod,t.spromptinfo spromptinfo,t.iwarnlevel iwarnlevel,t.swarnname swarnname 
    FROM tmpTableSoftware t GROUP BY t.stimeperiod,t.sweekday,t.spolicyaction;
    DROP TEMPORARY TABLE tmpTableSoftware; 
END//

DELIMITER ;

-- 删除云用户触发器 tr_tclouduser
DROP TRIGGER IF EXISTS `tr_tclouduser` ;

-- 删除警报外发表
DROP TABLE `twarnlogoutward`;

SET FOREIGN_KEY_CHECKS=1;