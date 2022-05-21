package app

import (
	"easyclash/clash"
	"easyclash/pkg/utils"
	"path/filepath"
	"time"
)

type Version struct {
	NewVersion  int64
	ProdVersion int64
}

func (a *App) GetBaseConf() clash.BaseConf {
	c := clash.BaseConf{}
	_ = a.store.Load(a.ConfDir+"/base_conf.yml", &c)
	c.HomeDir = a.ConfDir
	return c
}

func (a *App) SaveBaseConf(c clash.BaseConf) bool {
	c.Version = time.Now().UnixMicro()
	_ = a.store.Save(a.ConfDir+"/base_conf.yml", c)
	return true
}

func (a *App) GetProxy() *[]clash.Proxy {
	c := new([]clash.Proxy)
	_ = a.store.Load(a.ConfDir+"/proxy.yml", &c)
	return c
}

func (a *App) SaveProxy(c *[]clash.Proxy) bool {
	defer a.updateVersion()
	_ = a.store.Save(a.ConfDir+"/proxy.yml", c)
	return true
}

func (a *App) GetProxyGroup() *[]clash.ProxyGroup {
	c := new([]clash.ProxyGroup)
	_ = a.store.Load(a.ConfDir+"/proxy_group.yml", &c)
	return c
}

func (a *App) SaveProxyGroup(c *[]clash.ProxyGroup) bool {
	defer a.updateVersion()
	_ = a.store.Save(a.ConfDir+"/proxy_group.yml", c)
	return true
}

func (a *App) GetProxyProvider() *clash.ProxyProviders {
	c := new(clash.ProxyProviders)
	_ = a.store.Load(a.ConfDir+"/proxy_provider.yml", &c)
	_c := *c
	for k, v := range _c {
		v.UpdateTime = utils.TimeFormat(utils.GetFileModTime(v.Path))
		_c[k] = v
	}
	return &_c
}

func (a *App) SaveProxyProvider(c *clash.ProxyProviders) bool {
	defer a.updateVersion()
	_c := *c
	for k, v := range _c {
		v.Path = a.ConfDir + "/proxyproviders/" + k + ".yml"
		_c[k] = v
	}
	_ = a.store.Save(a.ConfDir+"/proxy_provider.yml", _c)
	return true
}

func (a *App) GetRule() *[]clash.Rule {
	c := new([]clash.Rule)
	_ = a.store.Load(a.ConfDir+"/rule.yml", &c)
	return c
}

func (a *App) SaveRule(c *[]clash.Rule) bool {
	defer a.updateVersion()
	_ = a.store.Save(a.ConfDir+"/rule.yml", c)
	return true
}

func (a *App) GetRuleSet() *clash.RuleProviders {
	c := new(clash.RuleProviders)
	_ = a.store.Load(a.ConfDir+"/rule_set.yml", &c)
	return c
}

func (a *App) SaveRuleSet(c *clash.RuleProviders) bool {
	defer a.updateVersion()
	_c := *c
	for k, v := range _c {
		v.Path = a.ConfDir + "/ruleset/" + k + ".yml"
		_c[k] = v
	}
	_ = a.store.Save(a.ConfDir+"/rule_set.yml", _c)
	return true
}

func (a *App) HaveNewVersionConf() Version {
	b := a.GetBaseConf()
	c := new(clash.ClashConfig)
	path := filepath.Join(a.ConfDir, "config.yml")
	_ = a.store.Load(path, &c)
	return Version{
		NewVersion:  b.Version,
		ProdVersion: c.Version,
	}
}

func (a *App) ClashIsRunning() bool {
	return a.running
}

func (a *App) updateVersion() {
	b := a.GetBaseConf()
	b.Version = time.Now().UnixMicro()
	a.SaveBaseConf(b)
}

func (a *App) GetConfDir() string {
	return a.ConfDir
}
