# Hot100 题目抽取器

## Features

- 随机选择一道 Leetcode Hot100 中的题目，输出题目相关信息及题目链接
- 记录抽取过的题目，避免重复
- 使用爬虫获取[官网](https://leetcode.cn/problem-list/2cktkvj/)的 Hot100 热题，可以反复爬取来及时更新 Hot100 题目

## Usage

仓库中有已经爬取好的数据，可以选择直接使用

```shell
git clone 
cd leetcode_hot100
python ./Hot100/getQuestion.py
```

如果想要自己爬取官网数据的话，需要安装依赖包，并且可能需要根据不同的操作系统安装相应的软件(google-chrome 以及相关驱动，目前只在 Arch Linux 上测试运行成功)

```shell
git clone 
cd leetcode_hot100
pip install -r requirements.txt
python ./Hot100/getHot100.py
```