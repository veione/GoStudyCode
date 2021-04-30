package main

import (
	"command"
	"common"
	"component/function"
	"component/idgenerator"
	"component/leiting/logbus"
	"component/logger"
	"configtable"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/globalsign/mgo"

	"github.com/globalsign/mgo/bson"
)

// 订单类型
const (
	OrderTypePri uint32 = 1 // 个人订单
	OrderTypePub uint32 = 2 // 多人订单
)

// 订单状态
const (
	OrderWaiting    uint32 = 1 // 订单等待中
	OrderProcessing uint32 = 2 // 订单进行中
	OrderEndedup    uint32 = 3 // 订单已完成
)

// PrivateOrder 私人订单
type PrivateOrder struct {
	orderSystem *OrderSystem // 订单系统
	orderID     uint64       // 订单id
	typeID      uint32       // 配表ID
	status      uint32       // 订单状态
	waitEndTime int64        // 等待结束时间
	orderType   uint32       // 订单类型
}

func newPrivateOrder(os *OrderSystem, orderCfg *configtable.OrderTable, orderID uint64) *PrivateOrder {
	order := &PrivateOrder{
		orderSystem: os,
		orderID:     orderID,
		typeID:      orderCfg.ID,
		status:      OrderWaiting,
		waitEndTime: time.Now().Unix() + 60,
	}
	order.loadDataFromCfgTable(orderCfg)
	return order
}

func newPrivateOrderFromDB(os *OrderSystem, orderInfo *common.PrivateOrder) *PrivateOrder {
	order := &PrivateOrder{
		orderSystem: os,
		orderID:     orderInfo.OrderID,
		typeID:      orderInfo.TypeID,
	}
	orderCfg := configtable.OrderTableMgr_GetMe().Get(orderInfo.TypeID)
	if orderCfg != nil {
		order.loadDataFromCfgTable(orderCfg)
	}
	order.status = orderInfo.Status
	order.waitEndTime = orderInfo.WaitEndTime
	// 已经结束了
	nowTime := time.Now().Unix()
	if order.status == OrderWaiting && nowTime > order.waitEndTime {
		order.waitEndTime = nowTime
		order.status = OrderProcessing
	}
	return order
}

func (priOrder *PrivateOrder) copyFrom(other *PrivateOrder) {
	priOrder.typeID = other.typeID
	priOrder.status = other.status
	priOrder.waitEndTime = other.waitEndTime
	priOrder.orderSystem = other.orderSystem
}

func (priOrder *PrivateOrder) loadDataFromCfgTable(orderCfg *configtable.OrderTable) {
	if orderCfg == nil {
		logger.Errorf("[订单] 配置错误 订单ID:%v", orderCfg.ID)
		return
	}
	priOrder.waitEndTime = time.Now().Unix() + int64(orderCfg.WaitTime*60)
}

func (priOrder *PrivateOrder) appendDataToDB(orderInfo *common.PrivateOrder) {
	orderInfo.OrderID = priOrder.orderID
	orderInfo.Status = priOrder.status
	orderInfo.TypeID = priOrder.typeID
	orderInfo.WaitEndTime = priOrder.waitEndTime
}

func (priOrder *PrivateOrder) appendDataToMsg(msg *command.PrivateOrderInfo) {
	msg.OrderID = priOrder.orderID
	msg.TypeID = priOrder.typeID
	msg.Status = priOrder.status
	msg.WaitEndTime = priOrder.waitEndTime
}

// OrderSystem 订单系统
type OrderSystem struct {
	owner             *Player                  // 拥有者
	collName          string                   // 集合名称
	genTime           int64                    // 订单生成时间
	priAccNum         uint32                   // 个人订单剩余剩余加速次数
	pubAccNum         uint32                   // 多人订单剩余加速次数
	buyPriAccNum      uint32                   // 购买个人订单加速次数
	buyPubAccNum      uint32                   // 购买个人订单加速次数
	deliveryAccNum    uint32                   // 派送加速剩余次数
	buyAccDeliveryNum uint32                   // 购买加速派送次数
	deliveryEndTime   int64                    // 派送等待结束时间
	isDeliveryEnd     bool                     // 订单派送完成
	orderCount        uint32                   //  今日完成订单的数量
	priOrders         map[uint64]*PrivateOrder // 个人订单
	totalOrderNum     uint32                   // 订单格子数
}

func newOrderSystem(owner *Player) *OrderSystem {
	os := &OrderSystem{
		owner:             owner,
		collName:          "OrderSystem",
		genTime:           time.Now().Unix(),
		priAccNum:         5,
		pubAccNum:         3,
		buyPriAccNum:      0,
		buyPubAccNum:      0,
		deliveryAccNum:    2,
		buyAccDeliveryNum: 0,
		deliveryEndTime:   time.Now().Unix(),
		isDeliveryEnd:     false,
		orderCount:        0,
		priOrders:         make(map[uint64]*PrivateOrder),
	}
	os.loadFromDB()
	// 不是同一天了 直接重置订单数据 现在依赖超拉 不能在此重置
	//if !function.InSameDay(time.Now().Unix(), os.genTime) {
	//	os.resetOrder()
	//}
	return os
}

// Clear 清除数据
func (os *OrderSystem) Clear() {
	os.priOrders = nil
	os.owner = nil
}

func (os *OrderSystem) appendDataToDB(osData *common.OrderSystem) {
	osData.PlayerID = os.owner.playerID
	osData.GenTime = os.genTime
	osData.PriAccNum = os.priAccNum
	osData.PubAccNum = os.pubAccNum
	osData.BuyPriAccNum = os.buyPriAccNum
	osData.BuyPubAccNum = os.buyPubAccNum
	osData.DeliveryAccNum = os.deliveryAccNum
	osData.BuyAccDeliveryNum = os.buyAccDeliveryNum
	osData.OrderCount = os.orderCount
	osData.PrivateOrders = make([]*common.PrivateOrder, 0, len(os.priOrders))
	for _, priOrder := range os.priOrders {
		orderInfo := &common.PrivateOrder{}
		priOrder.appendDataToDB(orderInfo)
		osData.PrivateOrders = append(osData.PrivateOrders, orderInfo)
	}
}

func (os *OrderSystem) saveToDB() {
	mdata := bson.M{}
	mdata["gt"] = os.genTime
	mdata["pran"] = os.priAccNum
	mdata["puan"] = os.pubAccNum
	mdata["bPrian"] = os.buyPriAccNum
	mdata["bPuban"] = os.buyPriAccNum
	mdata["dan"] = os.deliveryAccNum
	mdata["badn"] = os.buyAccDeliveryNum
	mdata["oc"] = os.orderCount
	mdata["det"] = os.deliveryEndTime

	privateOrders := make([]*common.PrivateOrder, 0, len(os.priOrders))
	for _, priOrder := range os.priOrders {
		orderInfo := &common.PrivateOrder{}
		priOrder.appendDataToDB(orderInfo)
		privateOrders = append(privateOrders, orderInfo)
	}

	mdata["prio"] = privateOrders
	mdata["tNum"] = os.totalOrderNum

	condition := bson.M{common.PRIMARYKEY: os.owner.playerID}
	update := bson.M{"$set": mdata}
	logicMongo.Update(common.OrderSystemC, condition, update)
}

func (os *OrderSystem) loadFromDB() {
	osData := &common.OrderSystem{}
	query := bson.M{common.PRIMARYKEY: os.owner.playerID}
	err := logicMongo.FindOne(osData.C(), query, nil, osData)
	// 集合中数据不存在插入数据
	if err == mgo.ErrNotFound {
		// 初始化数据
		os.priAccNum = 5
		priAccNumCfg := configtable.SystemTableMgr_GetMe().Get(118)
		if priAccNumCfg != nil {
			os.priAccNum = priAccNumCfg.ValueA
		}
		os.deliveryAccNum = 2
		sysCfg := configtable.SystemTableMgr_GetMe().Get(129)
		if sysCfg != nil {
			os.deliveryAccNum = sysCfg.ValueA
		}
		os.deliveryEndTime = time.Now().Unix()
		os.totalOrderNum = 1
		os.refreshOrder(OrderTypePri, true)
		os.appendDataToDB(osData)
		err = logicMongo.Insert(osData.C(), osData)
		if err != nil {
			logger.Error(err)
		}
	} else {
		os.genTime = osData.GenTime
		os.priAccNum = osData.PriAccNum
		os.pubAccNum = osData.PubAccNum
		os.buyPriAccNum = osData.BuyPriAccNum
		os.buyPubAccNum = osData.BuyPubAccNum
		os.deliveryAccNum = osData.DeliveryAccNum
		os.buyAccDeliveryNum = osData.BuyAccDeliveryNum
		os.orderCount = osData.OrderCount
		os.deliveryEndTime = osData.DeliveryEndTime
		os.totalOrderNum = osData.TotalOrderNum
		for _, priOrder := range osData.PrivateOrders {
			order := newPrivateOrderFromDB(os, priOrder)
			os.priOrders[order.orderID] = order
		}
		if os.totalOrderNum == 0 {
			os.totalOrderNum = 1
		}
	}
}

func (os *OrderSystem) updateWaitTime() {
	// 订单派送时间刷新
	nowTime := time.Now().Unix()
	if !os.isDeliveryEnd && nowTime > os.deliveryEndTime {
		os.isDeliveryEnd = true
		msgSend := &command.SCOrderDelieryFinish{}
		os.owner.sendMsg(command.Command_SCOrderDelieryFinish, msgSend)
		endTime := time.Unix(os.deliveryEndTime, 0)
		logger.Infof("[订单] 玩家:%v 订单派送结束 当前时间:%v 结束时间%v", os.owner.playerID, time.Now(), endTime)
	}

	for _, priOrder := range os.priOrders {
		if priOrder.status != OrderWaiting {
			continue
		}
		// 等待时间结束时需要发消息通知客户端
		if priOrder.status == OrderWaiting && nowTime > priOrder.waitEndTime {
			priOrder.waitEndTime = nowTime
			priOrder.status = OrderProcessing
			msgSend := &command.SCUpdatePrivateOrder{Order: &command.PrivateOrderInfo{}}
			priOrder.appendDataToMsg(msgSend.Order)
			os.owner.sendMsg(command.Command_SCUpdatePrivateOrder, msgSend)
		}
	}
}

// 派送加速
func (os *OrderSystem) accelerateDeliery() {
	msgSend := &command.SCAccelerateDeliery{Result: 1}
	sysCfg := configtable.SystemTableMgr_GetMe().Get(128)
	if sysCfg == nil {
		logger.Error("[订单] 配置错误 未配置加速派送消耗银币参数")
		return
	}

	nowTime := time.Now().Unix()
	// 有免费的次数先用免费的
	if os.deliveryAccNum >= 1 {
		os.deliveryAccNum--
		msgSend.LeftNum = os.deliveryAccNum
	} else {
		// 向上取整
		var deliveryWaitTime uint32
		if nowTime < os.deliveryEndTime {
			deliveryWaitTime = uint32(os.deliveryEndTime - nowTime)
		}
		needCoin := math.Ceil(float64(deliveryWaitTime) / float64(sysCfg.ValueA))

		if os.owner.user.getMoney(command.CoinType(sysCfg.ValueB)) < int(needCoin) {
			os.owner.sendErrMsg(uint16(command.Command_CSAccelerateOrder), command.ErrorCode(sysCfg.ValueB+100))
			return
		}
		os.owner.user.delMoney(command.CoinType(sysCfg.ValueB), int(needCoin), command.ExpStat_ACCDELIERY)
		os.owner.user.syncMony(true)
		os.buyAccDeliveryNum++
	}

	if os.deliveryEndTime > nowTime {
		os.deliveryEndTime = nowTime
	}
	os.owner.sendMsg(command.Command_SCAccelerateDeliery, msgSend)

	// // 保存数据
	// condition := bson.M{common.PRIMARYKEY: os.owner.playerID}
	// logicMongo.UpdateFields(os.collName, condition,
	// 	"dan", "badn", "det", os.deliveryAccNum, os.buyAccDeliveryNum, os.deliveryEndTime)
}

func (os *OrderSystem) unlockOrderTotalNum(index uint32) (uint32, bool) {
	sysCfg := configtable.SystemTableMgr_GetMe().Get(150)

	if sysCfg == nil {
		return os.totalOrderNum, false
	}

	if index != os.totalOrderNum+1 {
		return os.totalOrderNum, false
	}

	unlockLvls := function.SplitStringTo2dUint32Slice(sysCfg.ValueC, "|", ";")

	if int(index) > len(unlockLvls) {
		return os.totalOrderNum, false
	}

	needLvl := unlockLvls[index-1]

	smcLvl := os.owner.user.getLevel(command.SMCJobDef(needLvl[0]))

	if smcLvl < needLvl[1] {
		return os.totalOrderNum, false
	}

	os.totalOrderNum++

	// condition := bson.M{common.PRIMARYKEY: os.owner.playerID}

	// logicMongo.UpdateFields(common.OrderSystemC, condition, "tNum", os.totalOrderNum)

	return os.totalOrderNum, true
}

func (os *OrderSystem) resetOrder() {
	os.genTime = time.Now().Unix()
	os.priAccNum = 5
	priAccNumCfg := configtable.SystemTableMgr_GetMe().Get(118)
	if priAccNumCfg != nil {
		os.priAccNum = priAccNumCfg.ValueA
	}
	os.vipUpgrade(os.owner.petmgr.GetSuperLamuLevel(), false)
	os.pubAccNum = 3
	pubAccNumCfg := configtable.SystemTableMgr_GetMe().Get(119)
	if pubAccNumCfg != nil {
		os.pubAccNum = pubAccNumCfg.ValueA
	}
	os.deliveryAccNum = 2
	sysCfg := configtable.SystemTableMgr_GetMe().Get(129)
	if sysCfg != nil {
		os.deliveryAccNum = sysCfg.ValueA
	}
	os.deliveryEndTime = time.Now().Unix()
	os.orderCount = 0
	os.syncOrderSystemData()
	// condition := bson.M{common.PRIMARYKEY: os.owner.playerID}
	// logicMongo.UpdateFields(os.collName, condition,
	// 	"gt", "pran", "dan", "det", "oc",
	// 	os.genTime, os.priAccNum, os.deliveryAccNum, os.deliveryEndTime, os.orderCount)
}

func (os *OrderSystem) existPriOrder(typeID uint32) bool {
	for _, order := range os.priOrders {
		if order.typeID == typeID {
			return true
		}
	}
	return false
}

// 刷新单人订单
func (os *OrderSystem) refreshOrder(orderType uint32, refreshOld bool) {

	var (
		orders      []*configtable.OrderTable // 满足条件的订单
		probList    []uint32                  // 订单概率
		totalWeight uint32                    // 总权重值
		randValue   int                       // 随机值
		ordersNum   int                       // 当前以获得订单席位
		//pushOrders  []*common.PrivateOrder    // 添加到数据库的订单数据
	)

	if refreshOld { // 是否刷新已有订单
		if orderType == OrderTypePri {
			os.priOrders = make(map[uint64]*PrivateOrder)
		}
	} else {
		if orderType == OrderTypePri {
			ordersNum = len(os.priOrders)
		}
	}
	//pushOrders = make([]*common.PrivateOrder, 0, ordersNum)

	configtable.OrderTableMgr_GetMe().Datas.Range(func(key, value interface{}) bool {
		orderData, _ := value.(*configtable.OrderTable)
		var exist bool
		// 过滤已有的订单
		if orderType == OrderTypePri {
			exist = os.existPriOrder(orderData.ID)
		}
		smcLvl := os.owner.user.getLevel(command.SMCJobDef(orderData.LvType))
		if !exist && orderData.OrderType == orderType &&
			orderData.UnlockLV <= smcLvl && smcLvl <= orderData.MaxLV &&
			orderData.Weight > 0 {
			orders = append(orders, orderData)
		}
		return true
	})

	for i := 0; i < int(os.totalOrderNum)-ordersNum; i++ {

		totalWeight = 0
		// 订单库数量不足
		if len(orders) <= 0 {
			break
		}

		// 根据权重值计算概率
		probList = make([]uint32, 0, len(orders))
		probList = append(probList, orders[0].Weight)
		totalWeight += orders[0].Weight

		for i := 1; i < len(orders); i++ {
			probList = append(probList, probList[i-1]+orders[i].Weight)
			totalWeight += orders[i].Weight
		}

		randValue = rand.Intn(int(totalWeight))
		for idx, prob := range probList {
			if randValue <= int(prob) {
				orderID := idgenerator.FetchID()
				if orderType == OrderTypePri {
					priOrder := newPrivateOrder(os, orders[idx], orderID)
					priOrder.status = OrderProcessing
					os.priOrders[priOrder.orderID] = priOrder
					// 新增订单的时候
					if !refreshOld {
						orderInfo := &common.PrivateOrder{}
						priOrder.appendDataToDB(orderInfo)
						//pushOrders = append(pushOrders, orderInfo)
					}
				}
				orders = append(orders[:idx], orders[idx+1:]...)
				probList = append(probList[:idx], probList[idx+1:]...)
				break
			}
		}
	}

	// // 有新增订单
	// if len(pushOrders) > 0 {
	// 	condition := bson.M{common.PRIMARYKEY: os.owner.playerID}
	// 	update := bson.M{"$push": bson.M{"prio": bson.M{"$each": pushOrders}}, "$set": bson.M{"tNum": os.totalOrderNum}}
	// 	logicMongo.Update(os.collName, condition, update)
	// }
}

func (os *OrderSystem) refreshPriOrderByPlayer(orderID uint64, refreshNow bool) {
	var (
		orders      []*configtable.OrderTable // 满足条件的订单
		probList    []uint32                  // 订单概率
		totalWeight uint32                    // 总权重值
		randValue   int                       // 随机值
		priOrder    *PrivateOrder             // 刷新的订单
		ok          bool                      //
	)

	priOrder, ok = os.priOrders[orderID]
	if !ok {
		logger.Debugf("[订单] 请求刷新的订单不存在 订单id:%v", orderID)
		return
	}

	if priOrder.status != OrderProcessing {
		logger.Debugf("[订单] 请求刷新的订单状态错误 订单状态:%v", priOrder.status)
		return
	}

	logbus.LoggerTA.OrderRefresh(os.owner.playerStrID, orderID, priOrder.typeID, os.owner.getLogbusInfo())

	orders = make([]*configtable.OrderTable, 0, 128)
	configtable.OrderTableMgr_GetMe().Datas.Range(func(key, value interface{}) bool {
		orderData, _ := value.(*configtable.OrderTable)
		// 过滤已有的订单
		exist := os.existPriOrder(orderData.ID)
		smcLvl := os.owner.user.getLevel(command.SMCJobDef(orderData.LvType))
		// 目前设定的条件为花园等级
		if !exist && orderData.OrderType == OrderTypePri &&
			orderData.UnlockLV <= smcLvl && smcLvl <= orderData.MaxLV &&
			orderData.Weight > 0 {
			orders = append(orders, orderData)
		}
		return true
	})

	totalWeight = 0

	// 订单库数量不足
	if len(orders) <= 0 {
		if refreshNow { // 立即刷新,但是订单不足时,返回原订单
			priOrder.waitEndTime = time.Now().Unix()
			priOrder.status = OrderProcessing
			// os.savePrivateOrder(priOrder)
			os.sendPriOrderUpdateMsg(priOrder, os.owner)
		} else {
			os.owner.sendErrMsg(uint16(command.Command_CSRefreshOrder), command.ErrorCode_OrderShortage)
		}
		return
	}

	// 根据权重值计算概率
	probList = make([]uint32, 0, len(orders))
	probList = append(probList, orders[0].Weight)
	totalWeight += orders[0].Weight

	for i := 1; i < len(orders); i++ {
		probList = append(probList, probList[i-1]+orders[i].Weight)
		totalWeight += orders[i].Weight
	}

	randValue = rand.Intn(int(totalWeight))
	for idx, prob := range probList {
		if randValue <= int(prob) {
			newPriOrder := newPrivateOrder(os, orders[idx], priOrder.orderID)
			priOrder.copyFrom(newPriOrder)
			// 立即刷新,不进入等待期
			if refreshNow {
				priOrder.waitEndTime = time.Now().Unix()
				priOrder.status = OrderProcessing
			}
			break
		}
	}

	// os.savePrivateOrder(priOrder)
	os.sendPriOrderUpdateMsg(priOrder, os.owner)
}

// 保存订单数据
func (os *OrderSystem) savePrivateOrder(priOrder *PrivateOrder) {
	// condition := bson.M{common.PRIMARYKEY: os.owner.playerID, "prio.oid": priOrder.orderID}
	// logicMongo.UpdateFields("OrderSystem", condition,
	// 	"prio.$.tid", "prio.$.sta", "prio.$.wet", priOrder.typeID, priOrder.status, priOrder.waitEndTime)
}

func (os *OrderSystem) acceleratePriOrder(orderID uint64, player *Player) {
	msgSend := &command.SCAccelerateOrder{
		OrderType: OrderTypePri,
		OrderID:   orderID,
		Result:    1,
	}
	priOrder, ok := os.priOrders[orderID]
	if !ok {
		logger.Debugf("[订单] 请求加速的订单不存在 订单id:%v", orderID)
		return
	}

	if priOrder.status != OrderWaiting {
		logger.Debugf("[订单] 请求加速的订单状态错误 订单状态:%v", priOrder.status)
		return
	}
	// moneyType := 0
	costMoney := 0
	// condition := bson.M{common.PRIMARYKEY: os.owner.playerID}
	nowTime := time.Now().Unix()
	if os.priAccNum >= 1 {
		os.priAccNum--
		// logicMongo.UpdateFields(os.collName, condition, "pran", os.priAccNum)
	} else {
		sysCfg := configtable.SystemTableMgr_GetMe().Get(128)
		if sysCfg == nil {
			logger.Error("[订单] 配置错误 未配置加速派送消耗银币参数")
			return
		}
		// 向上取整
		var waitTime int64
		if nowTime < priOrder.waitEndTime {
			waitTime = priOrder.waitEndTime - nowTime
		}
		sub := configtable.DivinationBuffVal(os.owner.divi.Buffid, 6, os.owner.divi.Buffend)
		if conf := configtable.SuperLarmTableMgr_GetMe().Get(os.owner.petmgr.GetSuperLamuLevel()); conf != nil {
			sub += conf.AttrNum6
		}
		needCoin := math.Ceil(float64(waitTime) / float64(sysCfg.ValueA))
		needCoin = needCoin * (100.0 - float64(sub)) / 100.0
		costMoney = int(needCoin)
		// moneyType = int(sysCfg.ValueB)
		if player.user.getMoney(command.CoinType(sysCfg.ValueB)) >= int(needCoin) {
			player.user.delMoney(command.CoinType(sysCfg.ValueB), int(needCoin), command.ExpStat_ACCORDER)
			player.user.syncMony(true)
		} else {
			player.sendErrMsg(uint16(command.Command_CSAccelerateOrder), command.ErrorCode(100+sysCfg.ValueB))
			return
		}
		os.buyPriAccNum++
		// logicMongo.UpdateFields(os.collName, condition, "bPrian", os.buyPriAccNum)
	}
	msgSend.AccelerateNum = os.priAccNum
	player.sendMsg(command.Command_SCAccelerateOrder, msgSend)
	priOrder.waitEndTime = nowTime
	priOrder.status = OrderProcessing
	// os.savePrivateOrder(priOrder)
	os.sendPriOrderUpdateMsg(priOrder, player)
	logbus.LoggerTA.OrderAccelerate(player.playerStrID, priOrder.typeID, costMoney, "金豆", player.getLogbusInfo())
}

func (os *OrderSystem) paymentPriOrder(orderID uint64, player *Player) {
	msgSend := &command.SCPaymentOrder{
		OrderType:  1,
		OrderID:    orderID,
		OrderCount: os.orderCount,
		Result:     1,
	}
	priOrder, ok := os.priOrders[orderID]
	if !ok {
		logger.Debugf("[订单] 请求支付的订单不存在 订单id:%v", orderID)
		return
	}

	if priOrder.status != OrderProcessing {
		logger.Debugf("[订单] 请求支付的订单状态错误 订单状态:%v", priOrder.status)
		return
	}

	orderCfg := configtable.OrderTableMgr_GetMe().Get(priOrder.typeID)
	if orderCfg == nil {
		logger.Debugf("[订单] 请求支付的订单配置不存在 订单typeid:%v", priOrder.typeID)
		return
	}

	// 等待订单派送完成,派送完成后才能完成下一个订单
	nowTime := time.Now().Unix()
	if nowTime < os.deliveryEndTime {
		deliveryWaitTime := uint32(os.deliveryEndTime - nowTime)
		player.sendErrMsg(uint16(command.Command_CSPaymentOrder), command.ErrorCode_OrderDeliveryBusy)
		logger.Errorf("[订单] 玩家:%v订单错误 订单派送时间错误:%v", player.playerID, deliveryWaitTime)
		return
	}

	smc := os.owner.user.getSmc(command.SMCJobDef_FARMER)
	if smc == nil {
		return
	}

	id := function.CalculateSMCLevelID(command.SMCJobDef(smc.Job), command.SMCTitleDef(smc.Title), uint32(smc.LV))
	smcCfg := configtable.SMCTableMgr_GetMe().Get(id)
	if smcCfg == nil {
		return
	}
	// vip 加成
	vipadd := configtable.DivinationBuffVal(os.owner.divi.Buffid, 4, os.owner.divi.Buffend)
	if conf := configtable.SuperLarmTableMgr_GetMe().Get(os.owner.petmgr.GetSuperLamuLevel()); conf != nil {
		vipadd += conf.AttrNum4
	}
	// 订单次数已达上限,不能再完成订单了
	if os.orderCount >= smcCfg.DailyOrderNum+vipadd {
		msgSend.Result = 3
		player.sendMsg(command.Command_SCPaymentOrder, msgSend)
		return
	}

	needItems := function.SplitStringTo2dUint32Slice(orderCfg.OrderContent, "|", ",")
	del := make([]common.ContainerInfo, 0, len(needItems))
	// 检测背包中的物品是否满足支付订单需求
	for _, item := range needItems {
		if len(item) < 2 {
			continue
		}
		item[0] = player.checkSpecialID(item[0])
		_, num, _ := player.getItemNum(item[0])
		if num < int32(item[1]) {
			msgSend.Result = 2
			player.sendMsg(command.Command_SCPaymentOrder, msgSend)
			return
		}
		del = append(del, common.ContainerInfo{ItemId: item[0], Num: int32(item[1])})
	}
	priOrder.status = OrderEndedup
	os.orderCount++ // 订单完成计数
	// 等待派送完成
	os.deliveryEndTime = nowTime + 60
	tSysCfg := configtable.SystemTableMgr_GetMe().Get(132)
	if tSysCfg != nil {
		os.deliveryEndTime = nowTime + int64(tSysCfg.ValueA)
	}
	endTime := time.Unix(os.deliveryEndTime, 0)
	logger.Debugf("[订单] 玩家:%v 订单派送 开始时间:%v 结束时间%v", os.owner.playerID, time.Now(), endTime)
	os.isDeliveryEnd = false
	msgSend.OrderCount = os.orderCount
	msgSend.DelieryEndTime = os.deliveryEndTime
	player.sendMsg(command.Command_SCPaymentOrder, msgSend)
	player.triggerTask(common.FinishOrderAct, 0, 1)
	os.sendPriOrderUpdateMsg(priOrder, player)
	// 雷霆统计
	logbus.LoggerTA.OrderFinish(os.owner.playerStrID, priOrder.orderID, priOrder.typeID, player.getLogbusInfo())

	// 完成后立即刷新
	priOrder.waitEndTime = nowTime
	priOrder.status = OrderProcessing
	os.refreshPriOrderByPlayer(priOrder.orderID, true)

	// 道具最后加 加道具可能导致玩家升级,刷新订单列表
	player.delItems(command.ExpStat_ORDERGET, del, strconv.FormatUint(priOrder.orderID, 10))
	// 支付订单获取奖励
	dropItems := configtable.GetDropGroupItems(orderCfg.DropID)
	msgReward := &command.SCRewardBoxNotify{}
	add := make([]common.ContainerInfo, 0, len(dropItems))
	for itemID, itemNum := range dropItems {
		add = append(add, common.ContainerInfo{ItemId: itemID, Num: int32(itemNum)})
		msgReward.RewardList = append(msgReward.RewardList, &command.RewardData{
			Itemid:  itemID,
			Itemnum: itemNum,
		})
	}
	player.addItems(command.ExpStat_ORDERGET, add)
	player.sendMsg(command.Command_SCRewardBoxNotify, msgReward)
	// condition := bson.M{common.PRIMARYKEY: os.owner.playerID}
	// logicMongo.UpdateFields(os.collName, condition, "oc", os.orderCount)
}

func (os *OrderSystem) sendPriOrderUpdateMsg(priOrder *PrivateOrder, player *Player) {
	msgSend := &command.SCUpdatePrivateOrder{Order: &command.PrivateOrderInfo{}}
	priOrder.appendDataToMsg(msgSend.Order)
	player.sendMsg(command.Command_SCUpdatePrivateOrder, msgSend)
}

func (os *OrderSystem) syncOrderSystemData() {
	msgSend := &command.SCSyncOrderSystem{
		PriAccNum:      os.priAccNum,
		PubAccNum:      os.pubAccNum,
		DelieryAccNum:  os.deliveryAccNum,
		DelieryEndTime: os.deliveryEndTime,
		OrderCount:     os.orderCount,
		TotalSeatNum:   os.totalOrderNum,
	}

	for _, priOrder := range os.priOrders {
		orderInfo := &command.PrivateOrderInfo{}
		priOrder.appendDataToMsg(orderInfo)
		msgSend.PriOrders = append(msgSend.PriOrders, orderInfo)
	}
	os.owner.sendMsg(command.Command_SCSyncOrderSystem, msgSend)
}

func (os *OrderSystem) onlineNotify() {
	// 上线时如果派送已经完成通知客户端
	nowTime := time.Now().Unix()
	if nowTime >= os.deliveryEndTime {
		os.isDeliveryEnd = true
		msgSend := &command.SCOrderDelieryFinish{}
		os.owner.sendMsg(command.Command_SCOrderDelieryFinish, msgSend)
	}
}

func (os *OrderSystem) refreshRandOrderByPlayer() {
	for _, order := range os.priOrders {
		os.refreshPriOrderByPlayer(order.orderID, false)
		break
	}
}

func (os *OrderSystem) accelerateRandOrder() {
	for _, order := range os.priOrders {
		if order.status != OrderWaiting {
			continue
		}
		os.acceleratePriOrder(order.orderID, os.owner)
		break
	}
}

func (os *OrderSystem) paymentRandOrder() {
	for _, order := range os.priOrders {

		if order.status != OrderProcessing {
			continue
		}

		smc := os.owner.user.getSmc(command.SMCJobDef_FARMER)
		if smc == nil {
			return
		}

		id := function.CalculateSMCLevelID(command.SMCJobDef(smc.Job), command.SMCTitleDef(smc.Title), uint32(smc.LV))
		smcCfg := configtable.SMCTableMgr_GetMe().Get(id)
		if smcCfg == nil {
			return
		}

		// 订单次数已达上限,不能再完成订单了
		if os.orderCount >= smcCfg.DailyOrderNum {
			continue
		}

		orderCfg := configtable.OrderTableMgr_GetMe().Get(order.typeID)
		if orderCfg == nil {
			continue
		}

		needItems := function.SplitStringTo2dUint32Slice(orderCfg.OrderContent, "|", ",")
		add := make([]common.ContainerInfo, 0, len(needItems))
		// 检测背包中的物品是否满足支付订单需求
		for _, item := range needItems {
			if len(item) < 2 {
				continue
			}
			item[0] = os.owner.checkSpecialID(item[0])
			add = append(add, common.ContainerInfo{ItemId: item[0], Num: int32(item[1])})
		}
		os.owner.addItems(command.ExpStat_ORDERGET, add)
		os.paymentPriOrder(order.orderID, os.owner)
		break
	}
}

func (os *OrderSystem) refreshPriOrderByTask(orderTypeID uint32) {

	orderData := configtable.OrderTableMgr_GetMe().Get(orderTypeID)

	if orderData == nil {
		return
	}

	for orderID, priOrder := range os.priOrders {
		newPriOrder := newPrivateOrder(os, orderData, orderID)
		priOrder.copyFrom(newPriOrder)
		priOrder.waitEndTime = time.Now().Unix()
		priOrder.status = OrderProcessing
		// os.savePrivateOrder(priOrder)
		os.sendPriOrderUpdateMsg(priOrder, os.owner)
		return
	}
}

func (os *OrderSystem) vipUpgrade(lv uint32, write bool) {
	os.priAccNum = 5
	priAccNumCfg := configtable.SystemTableMgr_GetMe().Get(118)
	if priAccNumCfg != nil {
		os.priAccNum = priAccNumCfg.ValueA
	}
	if conf := configtable.SuperLarmTableMgr_GetMe().Get(lv); conf != nil {
		os.priAccNum += conf.AttrNum5
		// if write {
		// 	condition := bson.M{common.PRIMARYKEY: os.owner.playerID}
		// 	logicMongo.UpdateFields(os.collName, condition, "pran", os.priAccNum)
		// }
	}
	os.priAccNum += configtable.DivinationBuffVal(os.owner.divi.Buffid, 5, os.owner.divi.Buffend)
	msgSend := &command.SCAccelerateOrder{
		OrderType:     OrderTypePri,
		Result:        1,
		AccelerateNum: os.priAccNum,
	}
	os.owner.sendMsg(command.Command_SCAccelerateOrder, msgSend)
}

// changeType: true 表示超拉过期  false 表示充值超拉
func (os *OrderSystem) vipChange(changeType, write bool) {
	if conf := configtable.SuperLarmTableMgr_GetMe().Get(1); conf != nil {
		if changeType == true {
			if os.priAccNum >= conf.AttrNum5 {
				os.priAccNum -= conf.AttrNum5
			} else {
				os.priAccNum = 0
			}
		}
		if changeType == false {
			os.priAccNum += conf.AttrNum5
		}
		// if write {
		// 	condition := bson.M{common.PRIMARYKEY: os.owner.playerID}
		// 	logicMongo.UpdateFields(os.collName, condition, "pran", os.priAccNum)
		// }
	}
	msgSend := &command.SCAccelerateOrder{
		OrderType:     OrderTypePri,
		Result:        1,
		AccelerateNum: os.priAccNum,
	}
	os.owner.sendMsg(command.Command_SCAccelerateOrder, msgSend)
}
