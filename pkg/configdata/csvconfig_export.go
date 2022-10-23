// this file generated by tools,don't edit it!!!
package configdata

import "trcell/pkg/configdata/csvdef"

// config export interface begin
func (cfg *ConfigData) GetUnionBossCfg(id int32) *csvdef.UnionBoss {
	return cfg.csvModules.unionBossModule.GetData(id)
}
func (cfg *ConfigData) GetUnionBossCfgList() []*csvdef.UnionBoss {
	return cfg.csvModules.unionBossModule.GetDataList()
}

func (cfg *ConfigData) GetUnionBossLvCfg(id int32) *csvdef.UnionBossLv {
	return cfg.csvModules.unionBossLvModule.GetData(id)
}
func (cfg *ConfigData) GetUnionBossLvCfgList() []*csvdef.UnionBossLv {
	return cfg.csvModules.unionBossLvModule.GetDataList()
}

func (cfg *ConfigData) GetChargeCfg(id int32) *csvdef.Charge {
	return cfg.csvModules.chargeModule.GetData(id)
}
func (cfg *ConfigData) GetChargeCfgList() []*csvdef.Charge {
	return cfg.csvModules.chargeModule.GetDataList()
}

func (cfg *ConfigData) GetFormationCfg(id int32) *csvdef.Formation {
	return cfg.csvModules.formationModule.GetData(id)
}
func (cfg *ConfigData) GetFormationCfgList() []*csvdef.Formation {
	return cfg.csvModules.formationModule.GetDataList()
}

func (cfg *ConfigData) GetDrawGuessCfg(id int32) *csvdef.DrawGuess {
	return cfg.csvModules.drawGuessModule.GetData(id)
}
func (cfg *ConfigData) GetDrawGuessCfgList() []*csvdef.DrawGuess {
	return cfg.csvModules.drawGuessModule.GetDataList()
}

func (cfg *ConfigData) GetUndercoverCfg(id int32) *csvdef.Undercover {
	return cfg.csvModules.undercoverModule.GetData(id)
}
func (cfg *ConfigData) GetUndercoverCfgList() []*csvdef.Undercover {
	return cfg.csvModules.undercoverModule.GetDataList()
}

// config export interface end
