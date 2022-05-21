package app

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"easyclash/clash"
	"easyclash/pkg/sysproxy"
	"easyclash/pkg/update"
)

var appVersion = "v1.2"
var releaseUrl = "https://api.github.com/repos/daodao97/easyclash/releases/latest"

type Msg struct {
	Code    int
	Message string
}

type App struct {
	ctx     context.Context
	ConfDir string
	store   *clash.Store
	appData string
	msg     []Msg
	running bool
}

func NewApp(path string) *App {
	return &App{
		ConfDir: path + "/.easy_clash",
		store:   clash.NewStore(),
	}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.SaveRuleSet(a.GetRuleSet())
	// if runtime.Environment(ctx).BuildType == "production" {
	go a.StartClash()
	// }
}

func (a App) DomReady(ctx context.Context) {
	go a.checkUpdate(ctx)
}

func (a *App) Shutdown(ctx context.Context) {
	a.UnSetSystemProxy()
}

func (a *App) OnBeforeClose(ctx context.Context) bool {
	return true
}

func (a *App) StartClash() {
	_, err := a.InitConf()
	if err != nil {
		return
	}

	path := filepath.Join(a.ConfDir, "config.yml")
	err = clash.RunClash(true, path)
	if err != nil {
		a.addErr("Clash 核心程序启动失败", err)
		return
	}
	a.SetSystemProxy()

	a.running = true
}

func (a *App) StopClash() {
	a.UnSetSystemProxy()
	runtime.Quit(a.ctx)
}

func (a *App) SetSystemProxy() bool {
	c, _ := clash.NewConf(a.ConfDir)
	err := sysproxy.SetSysHttpProxy(fmt.Sprintf("127.0.0.1:%d", c.Port), "default")
	if err != nil {
		a.addNotify("Clash 设置系统代理失败")
		return false
	}

	err = sysproxy.SetSysSocksProxy(fmt.Sprintf("127.0.0.1:%d", c.SocksPort), "default")
	if err != nil {
		a.addNotify("Clash 设置系统代理失败")
		return false
	}

	return true
}

func (a *App) UnSetSystemProxy() bool {
	err := sysproxy.UnsetSysHttpProxy("default")
	if err != nil {
		a.addNotify("撤销系统代理失败")
		return false
	}
	err = sysproxy.UnsetSysSocksProxy("default")
	if err != nil {
		a.addNotify("撤销系统代理失败")
		return false
	}
	return true
}

func (a *App) InitConf() (*clash.ClashConfig, error) {
	if _, err := os.Stat(a.ConfDir); err != nil {
		err := os.MkdirAll(a.ConfDir, os.ModePerm)
		if err != nil {
			a.addErr("初始化配置文件失败", err)
			return nil, err
		}
	}

	c, err := clash.NewConf(a.ConfDir)
	if err != nil {
		a.addErr("搜集配置 失败", err)
		return nil, err
	}

	if err != nil {
		a.addErr("读取配置失败", err)
		return nil, err
	}
	if len(c.Proxies) == 0 && len(c.ProxyProviders) == 0 {
		a.addNotify("请先添加可用代理")
		return nil, err
	}
	flag := false
	for _, v := range c.ProxyGroups {
		if len(v.Use) > 0 || len(v.Proxies) > 0 {
			flag = true
		}
	}

	if !flag {
		a.addNotify("ProxyGroup 中未添加可用代理")
		return nil, err
	}

	err = clash.CheckClashConf(a.ConfDir, c)
	if err != nil {
		a.addErr("配置有误", err)
		return nil, err
	}

	err = clash.SaveClashConf(a.ConfDir, c)
	if err != nil {
		a.addErr("更新配置 失败", err)
		return nil, err
	}
	return c, nil
}

func (a *App) checkUpdate(ctx context.Context) {
	if runtime.Environment(ctx).BuildType != "production" {
		return
	}
	u := update.New(&update.Options{
		CurrentVersion: appVersion,
		ReleaseUrl:     releaseUrl,
	})
	u.Check(ctx)
}

func (a *App) GetMsg() []Msg {
	defer func() {
		a.msg = []Msg{}
	}()
	return a.msg
}

func (a *App) addErr(msg string, err error) {
	a.msg = append(a.msg, Msg{500, fmt.Sprintf("%v", errors.Wrap(err, msg))})
}

func (a *App) addMsg(m Msg) {
	a.msg = append(a.msg, m)
}

func (a *App) addNotify(m string) {
	a.msg = append(a.msg, Msg{0, m})
}
