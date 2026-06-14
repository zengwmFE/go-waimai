-- =============================================
-- 苍穹外卖 (Sky Take-Out) Database Schema
-- =============================================

CREATE DATABASE IF NOT EXISTS sky_take_out DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE sky_take_out;

-- ----------------------------
-- 1. 员工表 (employee)
-- ----------------------------
DROP TABLE IF EXISTS `employee`;
CREATE TABLE `employee` (
  `id`          BIGINT       NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name`        VARCHAR(32)  NOT NULL COMMENT '姓名',
  `username`    VARCHAR(32)  NOT NULL COMMENT '用户名',
  `password`    VARCHAR(64)  NOT NULL COMMENT '密码',
  `phone`       VARCHAR(11)  NOT NULL COMMENT '手机号',
  `sex`         VARCHAR(2)   NOT NULL COMMENT '性别',
  `id_number`   VARCHAR(18)  NOT NULL COMMENT '身份证号',
  `status`      INT          NOT NULL DEFAULT 1 COMMENT '状态 1:启用 0:禁用',
  `create_time` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_user` BIGINT       NOT NULL COMMENT '创建人ID',
  `update_user` BIGINT       NOT NULL COMMENT '修改人ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='员工表';

-- ----------------------------
-- 2. 分类表 (category)
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id`          BIGINT       NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name`        VARCHAR(32)  NOT NULL COMMENT '分类名称',
  `type`        INT          NOT NULL COMMENT '类型 1:菜品分类 2:套餐分类',
  `sort`        INT          NOT NULL DEFAULT 0 COMMENT '排序',
  `status`      INT          NOT NULL DEFAULT 1 COMMENT '状态 1:启用 0:禁用',
  `create_time` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_user` BIGINT       NOT NULL COMMENT '创建人ID',
  `update_user` BIGINT       NOT NULL COMMENT '修改人ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='分类表';

-- ----------------------------
-- 3. 菜品表 (dish)
-- ----------------------------
DROP TABLE IF EXISTS `dish`;
CREATE TABLE `dish` (
  `id`          BIGINT         NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name`        VARCHAR(32)    NOT NULL COMMENT '菜品名称',
  `category_id` BIGINT         NOT NULL COMMENT '分类ID',
  `price`       DECIMAL(10,2)  NOT NULL COMMENT '价格',
  `image`       VARCHAR(255)   NOT NULL COMMENT '图片',
  `description` VARCHAR(255)   DEFAULT NULL COMMENT '描述',
  `status`      INT            NOT NULL DEFAULT 1 COMMENT '状态 1:在售 0:停售',
  `create_time` DATETIME       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_user` BIGINT         NOT NULL COMMENT '创建人ID',
  `update_user` BIGINT         NOT NULL COMMENT '修改人ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜品表';

-- ----------------------------
-- 4. 菜品口味表 (dish_flavor)
-- ----------------------------
DROP TABLE IF EXISTS `dish_flavor`;
CREATE TABLE `dish_flavor` (
  `id`      BIGINT       NOT NULL AUTO_INCREMENT COMMENT '主键',
  `dish_id` BIGINT       NOT NULL COMMENT '菜品ID',
  `name`    VARCHAR(32)  NOT NULL COMMENT '口味名称',
  `value`   VARCHAR(255) NOT NULL COMMENT '口味值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜品口味表';

-- ----------------------------
-- 5. 套餐表 (setmeal)
-- ----------------------------
DROP TABLE IF EXISTS `setmeal`;
CREATE TABLE `setmeal` (
  `id`          BIGINT         NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name`        VARCHAR(32)    NOT NULL COMMENT '套餐名称',
  `category_id` BIGINT         NOT NULL COMMENT '分类ID',
  `price`       DECIMAL(10,2)  NOT NULL COMMENT '价格',
  `image`       VARCHAR(255)   NOT NULL COMMENT '图片',
  `description` VARCHAR(255)   DEFAULT NULL COMMENT '描述',
  `status`      INT            NOT NULL DEFAULT 1 COMMENT '状态 1:在售 0:停售',
  `create_time` DATETIME       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_user` BIGINT         NOT NULL COMMENT '创建人ID',
  `update_user` BIGINT         NOT NULL COMMENT '修改人ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='套餐表';

-- ----------------------------
-- 6. 套餐菜品关系表 (setmeal_dish)
-- ----------------------------
DROP TABLE IF EXISTS `setmeal_dish`;
CREATE TABLE `setmeal_dish` (
  `id`         BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键',
  `setmeal_id` BIGINT NOT NULL COMMENT '套餐ID',
  `dish_id`    BIGINT NOT NULL COMMENT '菜品ID',
  `copies`     INT    NOT NULL DEFAULT 1 COMMENT '份数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='套餐菜品关系表';

-- ----------------------------
-- 7. 用户表 (user)
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id`          BIGINT       NOT NULL AUTO_INCREMENT COMMENT '主键',
  `openid`      VARCHAR(45)  NOT NULL COMMENT '微信openid',
  `name`        VARCHAR(32)  DEFAULT NULL COMMENT '姓名',
  `phone`       VARCHAR(11)  DEFAULT NULL COMMENT '手机号',
  `sex`         VARCHAR(2)   DEFAULT NULL COMMENT '性别',
  `id_number`   VARCHAR(18)  DEFAULT NULL COMMENT '身份证号',
  `avatar`      VARCHAR(500) DEFAULT NULL COMMENT '头像',
  `create_time` DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_openid` (`openid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- ----------------------------
-- 8. 地址表 (address_book)
-- ----------------------------
DROP TABLE IF EXISTS `address_book`;
CREATE TABLE `address_book` (
  `id`            BIGINT       NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id`       BIGINT       NOT NULL COMMENT '用户ID',
  `consignee`     VARCHAR(32)  NOT NULL COMMENT '收货人',
  `sex`           VARCHAR(2)   NOT NULL COMMENT '性别',
  `phone`         VARCHAR(11)  NOT NULL COMMENT '手机号',
  `province_code` VARCHAR(12)  DEFAULT NULL COMMENT '省份编码',
  `province_name` VARCHAR(32)  DEFAULT NULL COMMENT '省份名称',
  `city_code`     VARCHAR(12)  DEFAULT NULL COMMENT '城市编码',
  `city_name`     VARCHAR(32)  DEFAULT NULL COMMENT '城市名称',
  `district_code` VARCHAR(12)  DEFAULT NULL COMMENT '区县编码',
  `district_name` VARCHAR(32)  DEFAULT NULL COMMENT '区县名称',
  `detail`        VARCHAR(200) DEFAULT NULL COMMENT '详细地址',
  `label`         VARCHAR(100) DEFAULT NULL COMMENT '标签',
  `is_default`    TINYINT      NOT NULL DEFAULT 0 COMMENT '是否默认 1:是 0:否',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='地址表';

-- ----------------------------
-- 9. 购物车表 (shopping_cart)
-- ----------------------------
DROP TABLE IF EXISTS `shopping_cart`;
CREATE TABLE `shopping_cart` (
  `id`          BIGINT         NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name`        VARCHAR(32)    DEFAULT NULL COMMENT '商品名称',
  `image`       VARCHAR(255)   DEFAULT NULL COMMENT '图片',
  `user_id`     BIGINT         NOT NULL COMMENT '用户ID',
  `dish_id`     BIGINT         DEFAULT NULL COMMENT '菜品ID',
  `setmeal_id`  BIGINT         DEFAULT NULL COMMENT '套餐ID',
  `dish_flavor` VARCHAR(50)    DEFAULT NULL COMMENT '口味',
  `number`      INT            NOT NULL DEFAULT 1 COMMENT '数量',
  `amount`      DECIMAL(10,2)  NOT NULL COMMENT '单价',
  `create_time` DATETIME       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='购物车表';

-- ----------------------------
-- 10. 订单表 (orders)
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `id`                      BIGINT         NOT NULL AUTO_INCREMENT COMMENT '主键',
  `number`                  VARCHAR(50)    NOT NULL COMMENT '订单号',
  `status`                  INT            NOT NULL COMMENT '订单状态 1:待付款 2:待接单 3:已接单 4:派送中 5:已完成 6:已取消',
  `user_id`                 BIGINT         NOT NULL COMMENT '用户ID',
  `address_book_id`         BIGINT         NOT NULL COMMENT '地址ID',
  `order_time`              DATETIME       NOT NULL COMMENT '下单时间',
  `checkout_time`           DATETIME       DEFAULT NULL COMMENT '结账时间',
  `pay_method`              INT            DEFAULT NULL COMMENT '支付方式 1:微信 2:支付宝',
  `pay_status`              TINYINT        NOT NULL DEFAULT 0 COMMENT '支付状态 0:未支付 1:已支付 2:退款',
  `amount`                  DECIMAL(10,2)  NOT NULL COMMENT '实收金额',
  `remark`                  VARCHAR(100)   DEFAULT NULL COMMENT '备注',
  `phone`                   VARCHAR(11)    DEFAULT NULL COMMENT '手机号',
  `address`                 VARCHAR(255)   DEFAULT NULL COMMENT '地址',
  `user_name`               VARCHAR(32)    DEFAULT NULL COMMENT '用户名',
  `consignee`               VARCHAR(32)    DEFAULT NULL COMMENT '收货人',
  `cancel_reason`           VARCHAR(255)   DEFAULT NULL COMMENT '取消原因',
  `rejection_reason`        VARCHAR(255)   DEFAULT NULL COMMENT '拒单原因',
  `cancel_time`             DATETIME       DEFAULT NULL COMMENT '取消时间',
  `estimated_delivery_time` DATETIME       DEFAULT NULL COMMENT '预计送达时间',
  `delivery_status`         TINYINT        NOT NULL DEFAULT 1 COMMENT '配送状态 1:立即送出 0:定时',
  `delivery_time`           DATETIME       DEFAULT NULL COMMENT '送达时间',
  `pack_amount`             INT            DEFAULT 0 COMMENT '打包费',
  `tableware_number`        INT            DEFAULT 0 COMMENT '餐具数量',
  `tableware_status`        TINYINT        NOT NULL DEFAULT 1 COMMENT '餐具状态 1:按数量 0:无需餐具',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_number` (`number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';

-- ----------------------------
-- 11. 订单明细表 (order_detail)
-- ----------------------------
DROP TABLE IF EXISTS `order_detail`;
CREATE TABLE `order_detail` (
  `id`          BIGINT         NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name`        VARCHAR(32)    DEFAULT NULL COMMENT '商品名称',
  `image`       VARCHAR(255)   DEFAULT NULL COMMENT '图片',
  `order_id`    BIGINT         NOT NULL COMMENT '订单ID',
  `dish_id`     BIGINT         DEFAULT NULL COMMENT '菜品ID',
  `setmeal_id`  BIGINT         DEFAULT NULL COMMENT '套餐ID',
  `dish_flavor` VARCHAR(50)    DEFAULT NULL COMMENT '口味',
  `number`      INT            NOT NULL DEFAULT 1 COMMENT '数量',
  `amount`      DECIMAL(10,2)  NOT NULL COMMENT '单价',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单明细表';

-- =============================================
-- 骑手表
-- =============================================
CREATE TABLE IF NOT EXISTS rider (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    phone VARCHAR(11) NOT NULL UNIQUE,
    status INT NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO rider (name, phone) VALUES
('骑手张三', '13800000001'),
('骑手李四', '13800000002'),
('骑手王五', '13800000003'),
('骑手赵六', '13800000004'),
('骑手陈七', '13800000005');
