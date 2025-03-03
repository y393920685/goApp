/*
 Navicat Premium Dump SQL

 Source Server         : 10.10.10.250_3306
 Source Server Type    : MySQL
 Source Server Version : 90001 (9.0.1)
 Source Host           : 10.10.10.250:3306
 Source Schema         : go_jxc

 Target Server Type    : MySQL
 Target Server Version : 90001 (9.0.1)
 File Encoding         : 65001

 Date: 24/02/2025 07:22:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for bank
-- ----------------------------
DROP TABLE IF EXISTS `bank`;
CREATE TABLE `bank`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '银行id',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '银行名称',
  `holder` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '开户人',
  `account` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '银行账号',
  `address` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '开户银行地址',
  `status` enum('0','1') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '银行状态 0:禁用 1:启用',
  `active` enum('0','1') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '银行默认 0:否 1:是',
  `opening_balance` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '期初余额',
  `balance` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '余额',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_bank_name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 30 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '银行表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for bank_record
-- ----------------------------
DROP TABLE IF EXISTS `bank_record`;
CREATE TABLE `bank_record`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '银行记录id',
  `bank_id` int UNSIGNED NOT NULL COMMENT '银行id',
  `type` enum('CG','XS') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '记录类型 CG:采购入库 XS:销售出库',
  `type_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '记录类型名称',
  `bill_no` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '单据号',
  `amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '交易金额',
  `balance` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '交易后余额',
  `user_id` int UNSIGNED NOT NULL COMMENT '创建人id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_bank_record_bill_no`(`bill_no` ASC) USING BTREE,
  INDEX `idx_bank_record_bank_id`(`bank_id` ASC) USING BTREE,
  INDEX `idx_bank_record_user_id`(`user_id` ASC) USING BTREE,
  CONSTRAINT `fk_bank_record_bank_id` FOREIGN KEY (`bank_id`) REFERENCES `bank` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_bank_record_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '银行记录表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for brand
-- ----------------------------
DROP TABLE IF EXISTS `brand`;
CREATE TABLE `brand`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '品牌id',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '品牌名称',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_brand_name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '品牌表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for buy
-- ----------------------------
DROP TABLE IF EXISTS `buy`;
CREATE TABLE `buy`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '采购id',
  `buy_no` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '采购单号',
  `buy_date` datetime NOT NULL COMMENT '采购日期',
  `buy_type` enum('0','1') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '采购类型 0:普通采购 1:其他入库',
  `buy_type_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '采购类型名称',
  `customer_id` int UNSIGNED NOT NULL COMMENT '供货商id',
  `warehouse_id` int UNSIGNED NOT NULL COMMENT '仓库id',
  `bank_id` int UNSIGNED NOT NULL COMMENT '银行id',
  `receivable_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '应付金额',
  `received_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '实付金额',
  `arrears_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '欠款金额',
  `create_user_id` int UNSIGNED NOT NULL COMMENT '创建人id',
  `update_user_id` int UNSIGNED NOT NULL COMMENT '修改人id',
  `audit_user_id` int UNSIGNED NULL DEFAULT NULL COMMENT '审核人id',
  `audit_status` enum('0','1') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '采购状态 0:未审核 1:已审核',
  `remark` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `audit_at` datetime NULL DEFAULT NULL COMMENT '审核时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_buy_buy_no`(`buy_no` ASC) USING BTREE,
  INDEX `idx_buy_customer_id`(`customer_id` ASC) USING BTREE,
  INDEX `idx_buy_warehouse_id`(`warehouse_id` ASC) USING BTREE,
  INDEX `idx_buy_bank_id`(`bank_id` ASC) USING BTREE,
  INDEX `idx_buy_create_user_id`(`create_user_id` ASC) USING BTREE,
  INDEX `idx_buy_update_user_id`(`update_user_id` ASC) USING BTREE,
  INDEX `idx_buy_audit_user_id`(`audit_user_id` ASC) USING BTREE,
  CONSTRAINT `fk_buy_audit_user_id` FOREIGN KEY (`audit_user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_buy_bank_id` FOREIGN KEY (`bank_id`) REFERENCES `bank` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_buy_create_user_id` FOREIGN KEY (`create_user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_buy_customer_id` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_buy_update_user_id` FOREIGN KEY (`update_user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_buy_warehouse_id` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouse` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '采购表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for buy_detail
-- ----------------------------
DROP TABLE IF EXISTS `buy_detail`;
CREATE TABLE `buy_detail`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '采购明细id',
  `buy_no` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '采购单号',
  `product_id` int UNSIGNED NOT NULL COMMENT '产品id',
  `product_num` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '采购数量',
  `product_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '采购单价',
  `product_total` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '采购总价',
  `remark` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_buy_detail_buy_no`(`buy_no` ASC) USING BTREE,
  INDEX `idx_buy_detail_product_id`(`product_id` ASC) USING BTREE,
  CONSTRAINT `fk_buy_detail_product_id` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '采购明细表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '产品类别id',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '产品类别名称',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_category_name`(`name` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '产品类别表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for color
-- ----------------------------
DROP TABLE IF EXISTS `color`;
CREATE TABLE `color`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '颜色id',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '颜色名称',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_color_name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '颜色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for customer
-- ----------------------------
DROP TABLE IF EXISTS `customer`;
CREATE TABLE `customer`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '客户id',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '客户名称',
  `pinyin` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '客户名称拼音',
  `contact` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户联系人',
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户联系电话',
  `company_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户公司名称',
  `address` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户地址',
  `tax_number` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户纳税号',
  `bank_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户开户行',
  `bank_account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户银行账号',
  `bank_holder` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '银行开户人',
  `opening_balance` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '客户期初余额',
  `status` enum('0','1') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '客户状态 0:禁用 1:启用',
  `active` enum('0','1') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '客户默认 0:否 1:是',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_customer_name`(`name` ASC) USING BTREE,
  INDEX `idx_customer_pinyin`(`pinyin` ASC) USING BTREE,
  INDEX `idx_customer_contact`(`contact` ASC) USING BTREE,
  INDEX `idx_customer_phone`(`phone` ASC) USING BTREE,
  INDEX `idx_customer_company_name`(`company_name` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '客户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '产品id',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '产品名称',
  `pinyin` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '拼音码',
  `category_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '产品类别id',
  `unitd` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '单位',
  `color` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '颜色',
  `brand` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '品牌',
  `spec` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '规格',
  `code` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '条码',
  `buy_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '销售价格',
  `sale_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '进货价格',
  `vip_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT 'VIP会员价',
  `low_sale_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '最低售价',
  `hight_buy_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '最高进价',
  `stock_upper` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '库存上限',
  `stock_lower` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '库存下限',
  `status` enum('0','1') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '产品状态 0:禁用 1:启用',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_product_name`(`name` ASC) USING BTREE,
  INDEX `idx_product_pinyin`(`pinyin` ASC) USING BTREE,
  INDEX `idx_product_category_id`(`category_id` ASC) USING BTREE,
  CONSTRAINT `fk_product_category_id` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '产品表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色id',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `value` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色值',
  `status` enum('0','1') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '角色状态 0:禁用 1:启用',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_role_name`(`name` ASC) USING BTREE,
  UNIQUE INDEX `uni_role_value`(`value` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sale
-- ----------------------------
DROP TABLE IF EXISTS `sale`;
CREATE TABLE `sale`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '销售id',
  `sale_no` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '销售单号',
  `sale_date` datetime NOT NULL COMMENT '销售日期',
  `sale_type` enum('0','1') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '销售类型 0:普通销售 1:其他出库',
  `sale_type_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '销售类型名称',
  `customer_id` int UNSIGNED NOT NULL COMMENT '客户id',
  `warehouse_id` int UNSIGNED NOT NULL COMMENT '仓库id',
  `bank_id` int UNSIGNED NOT NULL COMMENT '银行id',
  `receivable_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '应收金额',
  `received_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '实收金额',
  `sale_profit` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '销售利润',
  `arrears_amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '欠款金额',
  `create_user_id` int UNSIGNED NOT NULL COMMENT '创建人id',
  `update_user_id` int UNSIGNED NOT NULL COMMENT '修改人id',
  `audit_user_id` int UNSIGNED NULL DEFAULT NULL COMMENT '审核人id',
  `audit_status` enum('0','1') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '销售状态 0:未审核 1:已审核',
  `remark` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `audit_at` datetime NULL DEFAULT NULL COMMENT '审核时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_sale_sale_no`(`sale_no` ASC) USING BTREE,
  INDEX `idx_sale_customer_id`(`customer_id` ASC) USING BTREE,
  INDEX `idx_sale_warehouse_id`(`warehouse_id` ASC) USING BTREE,
  INDEX `idx_sale_bank_id`(`bank_id` ASC) USING BTREE,
  INDEX `idx_sale_create_user_id`(`create_user_id` ASC) USING BTREE,
  INDEX `idx_sale_update_user_id`(`update_user_id` ASC) USING BTREE,
  INDEX `idx_sale_audit_user_id`(`audit_user_id` ASC) USING BTREE,
  CONSTRAINT `fk_sale_audit_user_id` FOREIGN KEY (`audit_user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_sale_bank_id` FOREIGN KEY (`bank_id`) REFERENCES `bank` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_sale_create_user_id` FOREIGN KEY (`create_user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_sale_customer_id` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_sale_update_user_id` FOREIGN KEY (`update_user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_sale_warehouse_id` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouse` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '销售表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sale_detail
-- ----------------------------
DROP TABLE IF EXISTS `sale_detail`;
CREATE TABLE `sale_detail`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '销售明细id',
  `sale_no` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '销售单号',
  `product_id` int UNSIGNED NOT NULL COMMENT '产品id',
  `product_num` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '销售数量',
  `product_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '销售单价',
  `product_total` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '销售总价',
  `profit` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '销售利润',
  `remark` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sale_detail_sale_no`(`sale_no` ASC) USING BTREE,
  INDEX `idx_sale_detail_product_id`(`product_id` ASC) USING BTREE,
  CONSTRAINT `fk_sale_detail_product_id` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '销售明细表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for stock
-- ----------------------------
DROP TABLE IF EXISTS `stock`;
CREATE TABLE `stock`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '库存id',
  `product_id` int UNSIGNED NOT NULL COMMENT '产品id',
  `warehouse_id` int UNSIGNED NOT NULL COMMENT '仓库id',
  `stock_num` int NOT NULL DEFAULT 0 COMMENT '库存数量',
  `stock_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '库存单价',
  `stock_total` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '库存总价',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_stock_product_id_warehouse_id`(`product_id` ASC, `warehouse_id` ASC) USING BTREE,
  INDEX `idx_stock_product_id`(`product_id` ASC) USING BTREE,
  INDEX `idx_stock_warehouse_id`(`warehouse_id` ASC) USING BTREE,
  CONSTRAINT `fk_stock_product_id` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `fk_stock_warehouse_id` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouse` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '库存表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for supplier
-- ----------------------------
DROP TABLE IF EXISTS `supplier`;
CREATE TABLE `supplier`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '供货商id',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '供货商名称',
  `pinyin` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '供货商名称拼音',
  `contact` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '供货商联系人',
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '供货商联系电话',
  `company_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '供货商公司名称',
  `address` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '供货商地址',
  `tax_number` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '供货商纳税号',
  `bank_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '供货商开户行',
  `bank_account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '供货商银行账号',
  `bank_holder` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '银行开户人',
  `opening_balance` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '供货商期初余额',
  `status` enum('0','1') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '供货商状态 0:禁用 1:启用',
  `active` enum('0','1') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '供货商默认 0:否 1:是',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_supplier_name`(`name` ASC) USING BTREE,
  INDEX `idx_supplier_pinyin`(`pinyin` ASC) USING BTREE,
  INDEX `idx_supplier_contact`(`contact` ASC) USING BTREE,
  INDEX `idx_supplier_phone`(`phone` ASC) USING BTREE,
  INDEX `idx_supplier_company_name`(`company_name` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '供货商表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for unit
-- ----------------------------
DROP TABLE IF EXISTS `unit`;
CREATE TABLE `unit`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '单位id',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '单位名称',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_unit_name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '单位表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名称',
  `password` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户密码',
  `real_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '真实姓名',
  `age` tinyint UNSIGNED NULL DEFAULT 18 COMMENT '年龄',
  `gender` enum('0','1','2') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '性别 0:未知 1:男 2:女',
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '联系电话',
  `address` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址',
  `status` enum('0','1') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '用户状态 0:禁用 1:启用',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_user_username`(`username` ASC) USING BTREE,
  INDEX `idx_user_real_name`(`real_name` ASC) USING BTREE,
  INDEX `idx_user_phone`(`phone` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户角色id',
  `user_id` int UNSIGNED NOT NULL COMMENT '用户id',
  `role_id` int UNSIGNED NOT NULL COMMENT '角色id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_user_role_user_id_role_id`(`user_id` ASC, `role_id` ASC) USING BTREE,
  INDEX `idx_user_role_user_id`(`user_id` ASC) USING BTREE,
  INDEX `idx_user_role_role_id`(`role_id` ASC) USING BTREE,
  CONSTRAINT `fk_user_role_role_id` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `fk_user_role_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for warehouse
-- ----------------------------
DROP TABLE IF EXISTS `warehouse`;
CREATE TABLE `warehouse`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '仓库id',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '仓库名称',
  `contact` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '仓库联系人',
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '仓库联系电话',
  `status` enum('0','1') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '仓库状态 0:禁用 1:启用',
  `active` enum('0','1') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '仓库默认 0:否 1:是',
  `address` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '仓库地址',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_warehouse_name`(`name` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '仓库表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
