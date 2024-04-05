# map
1.map是一个无序的key-value对的集合
*每次打印顺序可能都不同*

2.map的格式为
map[keytype]valuetype
（键值唯一，且不能用切片、函数、以及包含切片的结构类型作为键值）

3.map的赋值
如果已存在key值，那就是修改；否则是追加扩容

4.map遍历
for key,value:=range m{

}

5.map判断某个键值是否存在
value,ok:=m[1]
ok==true存在；ok==false不存在

6.删除键值
delete(m,1)//删除键值为1的内容,若一开始没有就是虚删除

7.map做函数参数是引用传递

