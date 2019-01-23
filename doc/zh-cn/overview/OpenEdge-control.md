
# OpenEdge可控式设计

OpenEdge可控式设计主要包括3个方面的内容，其本质是OpenEdge联合[智能边缘BIE](https://cloud.baidu.com/product/bie.html)（Baidu-IntelliEdge）云端console管理套件一起实现对OpenEdge配置文件“一键配置、一键下发“，下发配置信息“一键更新、自动热加载”，及在Docker容器模式下“一键更新”容器镜像版本。

其中，“一键配置、一键下发“是指智能边缘BIE云端console支持用户通过在云端进行智能边缘核心设备的建立、身份制定、策略规则制定、函数编写，然后生成配置文件下发至OpenEdge本地运行包，可达到云端管理和应用下发，整体实现“一键配置、一键下发”的能力。另外，智能边缘BIE云端console还可以进行配置版本切换功能，其作用一是可以根据用户需求随时切换应用不同的配置；二是在发现某版本配置文件因使用者不熟练出现配置错误导致生产现场设备间消息收发紊乱的情况下，可以迅速切换稳定版本配置文件，及时止损。

此外，针对从智能边缘BIE下发的配置信息，OpenEdge还能够支持“[热加载](./OpenEdge-design.md)”。目前热加载采用全量更新的方式，既先停止所有老应用模块再启动所有新应用模块，因此模块提供的服务会中断。**需要注意的是**，在停止所有老应用模块时，采取已处于进行中的应用逻辑会持续进行，直到信息处理完成，然后再停止对应模块。

特别地，针对Docker容器模式，OpenEdge还能够支持通过简单修改配置文件达到对应用镜像“一键更新、随时切换”的效果。