package gonsq

import (
	"github.com/nsqio/go-nsq"
)

// MsgHandler 消息处理器
type MsgHandler func(data []byte)

// GameMessageHandler ...
type GameMessageHandler struct {
	handlers map[int]MsgHandler
}

// HandleMessage ...
func (complexs *GameMessageHandler) HandleMessage(msg *nsq.Message) error {
	//requset := command.ComplexMessage{}
	//err := proto.Unmarshal(msg.Body, &requset)
	//if err != nil {
	//	logger.Error("[NSQ] 消息错误:", err)
	//	return err
	//}
	//logger.Debugf("GameMessageHandler %v", requset.Nsqcmd)
	//handler, ok := complexs.handlers[requset.Nsqcmd]
	//if ok {
	//	handler(requset.Data)
	//} else {
	//	logger.Errorf("[NSQ] 未知的消息 cmd:%v", requset.Nsqcmd)
	//	return fmt.Errorf("未知的消息 cmd:%v", requset.Nsqcmd)
	//}
	return nil
}

func (complexs *GameMessageHandler) registerHandler() {
	//complexs.handlers[command.NsqCmd_FRAME] = complexs.chgFrame
	//complexs.handlers[command.NsqCmd_HEAD] = complexs.chgHead
	//complexs.handlers[command.NsqCmd_NICK] = complexs.chgNick
}

var (
	complexsConsumer *nsq.Consumer
	complexsHandler  *GameMessageHandler
)

func nsqProcessMgrInit(srvID string) {
	//gonsq.CreateTopic(gonsq.LogicNSQ, common.GAME_TOPIC)
	//complexsHandler = &GameMessageHandler{
	//	handlers: make(map[command.NsqCmd]MsgHandler),
	//}
	//complexsHandler.registerHandler()
	//complexsConsumer = gonsq.NewConsumer(gonsq.LogicNSQ, common.GAME_TOPIC, "channl-"+srvID, complexsHandler)
}

func nsqProcessMgrEnd() {
	complexsConsumer.Stop()
}

//
//func (complexs *GameMessageHandler) chgNick(data []byte) {
//	chgn := command.ChgNick{}
//	err := proto.Unmarshal(data, &chgn)
//	if err != nil {
//		return
//	}
//	if game, err := gameMgrGetMe().getGame(chgn.PlayerID); err == nil {
//		getPartyMgr().changeNick(game.party.PartyID, chgn.Onick, chgn.Nnick)
//	}
//}
//
//func (complexs *GameMessageHandler) chgHead(data []byte) {
//	chgh := command.ChgHead{}
//	err := proto.Unmarshal(data, &chgh)
//	if err != nil {
//		return
//	}
//	if game, err := gameMgrGetMe().getGame(chgh.PlayerID); err == nil {
//		game.sendToSceneGoroutine(command.SyncCmd_SPChgHead, &chgh)
//	}
//}
//
//func (complexs *GameMessageHandler) chgFrame(data []byte) {
//	chgf := command.ChgFrame{}
//	err := proto.Unmarshal(data, &chgf)
//	if err != nil {
//		return
//	}
//	if game, err := gameMgrGetMe().getGame(chgf.PlayerID); err == nil {
//		game.sendToSceneGoroutine(command.SyncCmd_SPChgFrame, &chgf)
//	}
//}
