# Tiny website

## 网站配色方案

[Ant Design 基础色板-明青](https://ant.design/docs/spec/colors-cn#%E5%9F%BA%E7%A1%80%E8%89%B2%E6%9D%BF)

- #e6fffb
- #b5f5ec
- #87e8de
- #5cdbd3
- #36cfc9 font-color: rgba(0, 0, 0, 0.85);
- #13c2c2 font-color: rgb(255, 255, 255);
- #08979c
- #006d75
- #00474f
- #002329

## markdown 编辑器

使用 marked.js 实现，左侧输入右侧预览，参考 [Marked Demo](https://marked.js.org/demo)

增加上传图片按钮，上传至七牛云图床，返回图片图片地址，以 ```![图片名称](图片url)``` 的格式 append 到输入框内

## 样式

尽量不把样式抽离到 css 中，使用 ```{{ define "css-common" }}``` 定义通用样式并在其他模版中使用 ```{{template "css-common" .}}``` 的方式引入，可以将样式按照 style 的方式加载，那么下载下来的 html 文件中已经有样式

## 学习资料

- [跟煎鱼学 Go](https://eddycjy.com/go-categories/)
- [jwt-go](https://github.com/dgrijalva/jwt-go)
- [JWT introduction](https://jwt.io/introduction/)