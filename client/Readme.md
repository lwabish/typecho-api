# typora-typecho-exporter

## 配置

1. 导出
2. 快捷键

## 工作流

1. 新建文章
   1. 创建`240319-文章标题.md`格式的markdown文件
   2. 在yaml front matter中指定分类和标签，例如：
      ```markdown
      ---
      categories:
        - 计算机
      tags:
        - tag1
      ---
      
      这是正文
      ```
   3. 写文章正文
   4. 导出到typecho
   5. 文章将自动发布到typecho，并将主键id回填到yaml front matter中
2. 修改文章
