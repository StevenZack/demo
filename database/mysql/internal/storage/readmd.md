# MySQL存储结构

## 数据库中的存储结构
记录是按照行来存储的，但是数据库的读取并不以行为单位，在数据库中，不论读一行，还是读多行，都是将这些行所在的页进行加载。也就是说，数据库管理存储空间的基本单位是页（Page）。

一个页中可以存储多个行记录（Row），同时在数据库中，还存在着区（Extent）、段（Segment）和表空间（Tablespace）。行、页、区、段、表空间的关系如下图所示：

![storage](https://img-blog.csdnimg.cn/20191209174007545.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ExODAyMDE2MjM0NA==,size_16,color_FFFFFF,t_70)

# 参考

[Mysql数据结构](https://blog.csdn.net/a18020162344/article/details/103462088)