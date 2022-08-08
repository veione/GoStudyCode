package main

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"runtime"
	"strconv"
	"testpro/component/logger"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	fmt.Println("The program starts ...")

	go func() {
		for {
		}
	}()

	time.Sleep(time.Second)
	fmt.Println("I got scheduled!")

	if mail.Opt == "del_personal_mail" {
		if len(mail.Muid) > 0 && len(mail.MailId) > 0 && len(mail.Muuid) > 0 && len(mail.Mtype) > 0 {
			if mail.Mtype != "Normal" {
				sendGamecenterResponse(-1, "只能删除 普通邮件, Mtype 参数错误.", mail, w)
				return
			}

			uid, e := strconv.ParseUint(mail.Muid, 10, 64)
			if e != nil {
				sendGamecenterResponse(-1, "删除个人邮件参数 Muid 错误.", mail, w)
				return
			}

			Muuid, e := strconv.ParseUint(mail.Muuid, 10, 64)
			if e != nil {
				sendGamecenterResponse(-1, "删除个人邮件参数 Muuid 错误.", mail, w)
				return
			}

			if logicredis.CheckLive(uid) {
				uuid, _ := strconv.Atoi(mail.Muid)
				msg := command.MailMessage{
					Muuid:    uint64(uuid),
					Mtype:    command.EmailType_NORMAL,
					IsGlobal: false,
				}
				msg.Recver = append(msg.Recver, uid)
				data, _ := proto.Marshal(&msg)
				sendMQ(command.NsqCmd_DELMAIN, data)
				sendGamecenterResponse(0, "已经发送删除个人邮件请求到在线服,请验证结果.", mail, w)
				return
			} else {
				condition := bson.M{common.PRIMARYKEY: uid}
				query := bson.M{"$pull": bson.M{"normal": bson.M{"muid": Muuid}}}
				err := logicMongo.Update(common.MailSystemC, condition, query)
				if err == nil {
					sendGamecenterResponse(0, "删除个人邮件OK.", mail, w)
					logger.Debugf("删除个人邮件OK.")
				} else {
					tips := fmt.Sprintf("删除个人邮件失败: %s.", err.Error())
					sendGamecenterResponse(-1, tips, mail, w)
					logger.Debugf(tips)
				}
			}
			return
		}
	}
}
