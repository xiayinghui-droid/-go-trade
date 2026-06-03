# Go-Trade 校园闲置AI估价桌面应用

<div align="center">

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)
![Vue](https://img.shields.io/badge/Vue-3.4+-4FC08D?style=flat-square&logo=vue.js)
![Wails](https://img.shields.io/badge/Wails-2.6+-FF6B6B?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)

基于 Wails + Go + Vue3 + LLM 的校园闲置 AI 估价桌面项目

</div>

## 📋 项目简介

Go-Trade 是一款面向校园用户的闲置物品交易桌面应用，集成 AI 智能估价功能，帮助用户快速合理地为二手商品定价。采用现代化的技术栈，提供流畅的桌面应用体验。

## ✨ 功能特性

- 🏠 **首页概览** - 商品统计、最新发布、快速入口
- 📦 **商品管理** - 浏览、搜索、筛选闲置商品
- 📝 **发布闲置** - 简单易用的商品发布流程
- 🤖 **AI 智能估价** - 基于 LLM 的智能价格评估
- 💾 **本地存储** - JSON 文件存储，无需数据库
- 🎨 **现代 UI** - Element Plus 组件库，美观易用

## 🏗️ 技术栈

### 前端
- **Vue 3** - 渐进式 JavaScript 框架
- **Vite** - 下一代前端构建工具
- **Vue Router** - 官方路由管理器
- **Pinia** - Vue 状态管理库
- **Element Plus** - Vue 3 组件库

### 后端
- **Go** - 高性能编程语言
- **Wails** - 使用 Go 和 Web 技术构建桌面应用
- **OpenAI API** - LLM 大语言模型集成

## 📁 项目结构

```
go-trade/
├── frontend/              # 前端代码
│   ├── src/
│   │   ├── components/    # 组件
│   │   ├── views/         # 页面视图
│   │   ├── router/        # 路由配置
│   │   ├── assets/        # 静态资源
│   │   ├── App.vue        # 根组件
│   │   └── main.js        # 入口文件
│   ├── package.json
│   └── vite.config.js
├── backend/               # 后端代码
│   ├── app.go             # 应用核心
│   ├── ai_estimator.go    # AI 估价引擎
│   └── llm_service.go     # LLM 服务
├── main.go                # 主入口
├── go.mod                 # Go 模块
├── wails.json             # Wails 配置
└── README.md
```

## 🚀 快速开始

### 环境要求

- Go 1.21+
- Node.js 18+
- Wails CLI 2.0+

### 安装 Wails CLI

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### 安装依赖

```bash
# 克隆项目
git clone https://github.com/yourusername/go-trade.git
cd go-trade

# 安装前端依赖
cd frontend
npm install
cd ..
```

### 开发模式

```bash
wails dev
```

### 构建生产版本

```bash
wails build
```

## 🤖 AI 估价功能

### 配置 LLM API

复制 `.env.example` 为 `.env` 并配置：

```env
OPENAI_API_KEY=your_api_key_here
OPENAI_BASE_URL=https://api.openai.com/v1
OPENAI_MODEL=gpt-3.5-turbo
```

### 估价算法

系统采用混合估价策略：

1. **基础规则引擎** - 基于商品分类、成色、使用时长计算折旧
2. **LLM 增强分析** - 调用大语言模型进行市场分析
3. **价格区间推荐** - 提供合理的定价范围

## 📖 使用说明

1. **浏览商品** - 在首页或商品列表页查看所有闲置商品
2. **发布闲置** - 填写商品信息，支持 AI 估价辅助定价
3. **管理商品** - 在"我的发布"中编辑、删除或标记已售出
4. **AI 估价** - 输入商品信息获取智能估价建议

## 🛠️ 开发指南

### 添加新功能

1. 前端页面：在 `frontend/src/views/` 添加 Vue 组件
2. 路由配置：在 `frontend/src/router/index.js` 注册路由
3. 后端逻辑：在 `backend/` 目录添加 Go 代码
4. 绑定方法：在 `main.go` 的 `Bind` 数组中注册

### 自定义主题

修改 `frontend/src/assets/style/main.scss` 或使用 Element Plus 的主题配置。

## 📝 待办事项

- [ ] 用户登录认证系统
- [ ] 商品图片上传功能
- [ ] 消息通知系统
- [ ] 交易记录管理
- [ ] 数据导出功能
- [ ] 多语言支持
- [ ] 深色模式

## 🤝 贡献指南

欢迎提交 Issue 和 Pull Request！

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🙏 致谢

- [Wails](https://wails.io/) - 桌面应用框架
- [Vue.js](https://vuejs.org/) - 前端框架
- [Element Plus](https://element-plus.org/) - UI 组件库
- [OpenAI](https://openai.com/) - AI 服务

## 📞 联系方式

- 项目地址：[https://github.com/yourusername/go-trade](https://github.com/yourusername/go-trade)
- 问题反馈：[Issues](https://github.com/yourusername/go-trade/issues)

---

<div align="center">
⭐ 如果这个项目对你有帮助，欢迎点个 Star！
</div>
