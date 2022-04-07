# PixivAPI
Pixiv API Server for golang (with Auth supported).

## 怎么用

从 [action](https://github.com/DiheChen/PixivAPI/actions) 下载和你对应系统 / 平台的二进制文件, 运行, 登录, 重启。

可以用 `export GIN_MODE=release` (Linux),  `set export GIN_MODE=release` (Windows cmd), 或者 `$Env GIN_MODE="release"` 关闭 debug 模式。

## 实现的 API:

- 获取插画详情
- 获取排行榜
- 获取插画推荐
- 获取用户详情
- 获取用户插画
- 获取用户收藏
- 获取用户正在关注
- 获取用户关注者
- 搜索插画
- 获取趋势标签
- 获取动图元数据
- 获取用户小说
- 获取小说系列
- 获取小说详情
- 获取小说文本
- 获取新的小说
- 收藏插画 (仅做实现)
- 取消收藏 (仅做实现)
- 关注用户 (仅做实现)
- 取关用户 (仅做实现)
