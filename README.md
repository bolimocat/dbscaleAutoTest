# dbscaleAutoTest
 GreateDB 4.0 dbscale autotest suit by golang

2022年8月4日 增加内容：
1 生成预期文件的ge工具，增加传递附件功能：
	用法1 ./ge test.sql  > 在file/properties的第一个dbscale上执行test.sql，然后将预期的out文件，放入这个节点的/tmp下，使用时从/tmp下载该文件到本地执行测试对应的模块目录里。
	用法2 ./ge test.sql import.csv > 当执行的sql里需要加载外部文件的，将外部文件做完第二个参数，执行时先把这个附件文件传递到执行sql的节点的/tmp下，注意这里的test.sql中引用
文件的位置需要修改为/tmp/import.csv，同样在实际执行时需要修改sql文件和out文件两个地方的附件地址，在file/模块/ 下。

2 场景文件修改
sceneinfo.csv，增加了一个附件字段，当该条执行场景需要有外部文件配合时，将文件名写入这个字段，实际文件存储在file/模块/下，则执行到该场景时，会自动将附件文件传送到执行端的相应位置。

3 增加了cleanTest一些功能
./cleanTest时，删除本次测试本地testcurrent目录下的内容（实际删除的是场景文件里该条记录状态为TRUE的），然后将testresult中本次的测试结果复制为一个带有当前时间戳的文件，便于保留。

4 增加java运行环境和接口工具
a 增加jre目录，直接使用jre去运行interface中的jdbc接口测试工具和压力测试工具。
b 在interface中存放了2个jar，一个是做接口功能测试，一个是做基于jdbc的压力测试。
c jdbc的相关接口测试，场景同样写入用例sql的位置，具体内容需要继续丰富。