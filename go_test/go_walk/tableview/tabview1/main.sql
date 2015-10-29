/*
Navicat SQLite Data Transfer

Source Server         : test-tab
Source Server Version : 30714
Source Host           : :0

Target Server Type    : SQLite
Target Server Version : 30714
File Encoding         : 65001

Date: 2015-10-29 13:11:28
*/

PRAGMA foreign_keys = OFF;

-- ----------------------------
-- Table structure for setting
-- ----------------------------
DROP TABLE IF EXISTS "main"."setting";
CREATE TABLE setting (ip text not null primary key, user text, passwd text);

-- ----------------------------
-- Records of setting
-- ----------------------------

-- ----------------------------
-- Table structure for version
-- ----------------------------
DROP TABLE IF EXISTS "main"."version";
CREATE TABLE "version" (
"index"  int NOT NULL,
"masterver"  text,
"version"  text,
"pack"  text,
"tag"  text,
"tagpath"  text,
"packtime"  TEXT,
PRIMARY KEY ("index" ASC)
);

-- ----------------------------
-- Records of version
-- ----------------------------
INSERT INTO "main"."version" VALUES (1, 2, 2.22, '11,22,33', '2.22.11', 'www.baidu.com', '2015-10-10 12:00:50');
INSERT INTO "main"."version" VALUES (2, 3, 3.11, '88,99', 3.88, 'www.hupu.net', '2015-10-10 12:00:56');
INSERT INTO "main"."version" VALUES (3, 5, 5.5, '22,55', 5.99, 'www.hupu.cn', '2015-10-20 12:00:56');
INSERT INTO "main"."version" VALUES (4, 4, 4, 444, 444, '', '2015-10-25 12:00:56');
INSERT INTO "main"."version" VALUES (66, 8, 5, 55, 55, '', '2015-10-28 12:00:56');

-- ----------------------------
-- Table structure for _version_old_20151027
-- ----------------------------
DROP TABLE IF EXISTS "main"."_version_old_20151027";
CREATE TABLE "_version_old_20151027" ("index" int not null primary key, masterver text, version text, pack text, tag text, tagpath text, packtime datatype);

-- ----------------------------
-- Records of _version_old_20151027
-- ----------------------------
INSERT INTO "main"."_version_old_20151027" VALUES (1, 2, 2.22, '11,22,33', '2.22.11', 'www.baidu.com', '2015-10-10 12:00:50');
INSERT INTO "main"."_version_old_20151027" VALUES (2, 3, 3.11, '88,99', 3.88, 'www.hupu.net', '2015-10-10 12:00:56');
INSERT INTO "main"."_version_old_20151027" VALUES (3, 5, 5.5, '22,55', 5.99, 'www.hupu.cn', '2015-10-20 12:00:56');
