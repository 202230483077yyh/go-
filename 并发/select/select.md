# select
可以监听channel的数据流动
格式：
select{
    case <-chan1:   //表明chan1读数据成功

    case chan2<-1:  //表明chan2写入数据成功

    default:    //*若加上default则会轮询，不会阻塞；不写则会阻塞*
}