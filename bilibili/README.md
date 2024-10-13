## bilibli  播放列表解析

https://api.bilibili.com/x/web-interface/view?bvid=BV1v3411P7kY
bvid 就是bilibili的视频id

非常长的情况下 裁切 json字段下的epsoides下的字段
粘贴到json edit online 里
https://jsoneditoronline.org/#right=local.larefa&left=local.kanaxo

然后把缩小的json 放到 json to go里

https://mholt.github.io/json-to-go/

取消勾选inline type definitions

omitempty 取消勾选 
用在go语言结构体struct标签中，跟在字段名称后面，如果字段值为：0、nil、false，则此字段在转换为json格式时，会没有此字段。
