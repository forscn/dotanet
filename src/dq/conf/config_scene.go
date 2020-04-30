// Copyright 2014 mqant Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package conf

import (
	"dq/log"
	"dq/utils"
)

var (
	SceneFileDatas = make(map[interface{}]interface{})
)

//场景配置文件
func LoadSceneFileData() {
	_, SceneFileDatas = utils.ReadXlsxData("bin/conf/scenes.xlsx", (*SceneFileData)(nil))

}
func GetSceneFileData(typeid int32) *SceneFileData {
	//log.Info("find unitfile:%d", typeid)

	re := (SceneFileDatas[int(typeid)])
	if re == nil {
		log.Info("not find Scenefile:%d", typeid)
		return nil
	}
	return (SceneFileDatas[int(typeid)]).(*SceneFileData)
}
func GetAllScene() map[interface{}]interface{} {
	return SceneFileDatas
}

//单位配置文件数据
type SceneFileData struct {
	//配置文件数据
	TypeID          int32  //类型ID
	ScenePath       string //场景路径
	CreateUnit      string //创建单位
	UnitExperience  int32  //击杀单位获得经验
	UnitGold        int32  //击杀单位获得金币
	UnitDiamond     int32  //击杀单位获得的砖石
	StartX          float32
	StartY          float32
	EndX            float32
	EndY            float32
	IsOpen          int32  //1表示开放 2表示关闭
	SceneBuff       string //场景BUFF 进入场景的单位都会添加此BUFF
	ChangeEquipAble int32  //是否可以更换装备 1表示可以 其他表示否
	DeathHuicheng   int32  //死亡后是否回到和平世界 1表示是 其他表示否
	//StartX	StartY	EndX	EndY

	//特殊情况处理
	Exception      int32  //0表示没有特殊情况 1:工会战
	ExceptionParam string //特殊情况处理参数 特殊情况为1的时候:()

}
